package terraform

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/go-logr/logr"
	inspectconfig "github.com/hashicorp/terraform-config-inspect/tfconfig"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"

	walrus "github.com/seal-io/walrus/pkg/apis/walrus/v1"
	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/resourcehandler"
	"github.com/seal-io/walrus/pkg/resourcehandlers/k8s"
	runapi "github.com/seal-io/walrus/pkg/resourceruns/api"
	"github.com/seal-io/walrus/pkg/resourceruns/config"
	"github.com/seal-io/walrus/pkg/resourceruns/parse"
	resapi "github.com/seal-io/walrus/pkg/resources/api"
	servervars "github.com/seal-io/walrus/pkg/servers/vars"
	"github.com/seal-io/walrus/pkg/system"
	"github.com/seal-io/walrus/pkg/templates/api"
	templateapi "github.com/seal-io/walrus/pkg/templates/api"
	"github.com/seal-io/walrus/pkg/templates/openapi"
	"github.com/seal-io/walrus/pkg/templates/translator"
	tfconfig "github.com/seal-io/walrus/pkg/terraform/config"
	"github.com/seal-io/walrus/pkg/terraform/parser"
	"github.com/seal-io/walrus/pkg/terraform/util"
)

const (
	// _backendAPI the API path to terraform deploy backend.
	// Terraform will get and update deployment states from this API.
	_backendAPI = "/v1/projects/%s/environments/%s/resources/%s/runs/%s/state"
)

// TerraformConfigurator constructs the terraform config files for the run.
type TerraformConfigurator struct {
	logger logr.Logger
}

func NewConfigurator() config.Configurator {
	return &TerraformConfigurator{
		logger: klog.Background().WithName("resource-run").WithName("tf"),
	}
}

func (c *TerraformConfigurator) Load(
	ctx context.Context,
	run *walruscore.ResourceRun,
	opts *config.Options,
) (map[string][]byte, error) {
	// Prepare terraform tfConfig.
	//  get module configs from resource run.
	moduleConfig, providerRequirements, err := c.getModuleConfig(ctx, run)
	if err != nil {
		return nil, err
	}

	requiredProviders := make(map[string]*inspectconfig.ProviderRequirement)
	for _, p := range providerRequirements {
		if _, ok := requiredProviders[p.Name]; !ok {
			requiredProviders[p.Name] = p.ProviderRequirement
		} else {
			c.logger.Infof("duplicate provider requirement: %s", p.Name)
		}
	}

	// Parse module attributes.
	attrs, variables, dependencyOutputs, err := parse.ParseModuleAttributes(
		ctx,
		run,
		moduleConfig.Attributes,
		false,
	)
	if err != nil {
		return nil, err
	}

	moduleConfig.Attributes = attrs

	// Update output sensitive with variables.
	wrapVariables, err := updateOutputWithVariables(variables, moduleConfig)
	if err != nil {
		return nil, err
	}

	// Prepare terraform config files to be mounted to secret.
	requiredProviderNames := sets.NewString()
	for _, p := range providerRequirements {
		requiredProviderNames = requiredProviderNames.Insert(p.Name)
	}

	address := fmt.Sprintf("%s%s", opts.ServerURL,
		fmt.Sprintf(_backendAPI,
			run.Spec.Project,
			run.Namespace,
			run.Spec.ResourceName,
			run.Name))

	tfCreateOpts := map[string]tfconfig.CreateOptions{
		tfconfig.FileMain: {
			TerraformOptions: &tfconfig.TerraformOptions{
				Token:                opts.Token,
				Address:              address,
				SkipTLSVerify:        !servervars.IsTlsCertified(),
				ProviderRequirements: requiredProviders,
			},
			ProviderOptions: &tfconfig.ProviderOptions{
				RequiredProviderNames: requiredProviderNames.List(),
				Connectors:            opts.Connectors,
				SecretMonthPath:       opts.SecretMountPath,
				ConnectorSeparator:    parser.ConnectorSeparator,
			},
			ModuleOptions: &tfconfig.ModuleOptions{
				ModuleConfigs: []*tfconfig.ModuleConfig{moduleConfig},
			},
			VariableOptions: &tfconfig.VariableOptions{
				VariablePrefix:    config.VariablePrefix,
				ResourcePrefix:    config.ResourcePrefix,
				Variables:         wrapVariables,
				DependencyOutputs: dependencyOutputs,
			},
			OutputOptions: moduleConfig.Outputs,
		},
		tfconfig.FileVars: getVarConfigOptions(variables, dependencyOutputs),
	}

	inputConfigs := make(map[string][]byte, len(tfCreateOpts))
	for k, v := range tfCreateOpts {
		inputConfigs[k], err = tfconfig.CreateConfigToBytes(ctx, v)
		if err != nil {
			return nil, err
		}
	}

	// Save input plan to resource run.
	run.Status.ConfigSecretName = walruscore.ResourceRunConfigSecretName(run)

	return inputConfigs, nil
}

// getModuleConfig returns module configs and required connectors to
// get terraform module config block from resource run.
func (c *TerraformConfigurator) getModuleConfig(
	ctx context.Context,
	run *walruscore.ResourceRun,
) (*tfconfig.ModuleConfig, []templateapi.ProviderRequirement, error) {
	// Get template version and template version generated schema.
	templateVersion, schema, err := templateapi.GetTemplateVersionReferenceSchema(ctx, run.Spec.Template, templateapi.TemplateVesionSchemaTypeOrigin)
	if err != nil {
		return nil, nil, err
	}

	var templateSchema *api.TemplateSchema
	if err := json.Unmarshal(schema.Status.Value.Raw, &templateSchema); err != nil {
		return nil, nil, err
	}

	mc := &tfconfig.ModuleConfig{
		Name:   run.Spec.ResourceName,
		Source: templateVersion.URL,
	}

	if templateSchema.IsEmpty() {
		return mc, nil, nil
	}

	mc.SchemaData = templateSchema.Data

	// Variables.
	var (
		variableSchema     = templateSchema.VariableSchema()
		outputSchema       = templateSchema.OutputSchema()
		sensitiveVariables = sets.Set[string]{}
	)

	if variableSchema != nil {
		// Convert variable schema to terraform module config.
		values := make(map[string]json.RawMessage, len(variableSchema.Properties))

		if err := json.Unmarshal(run.Status.ComputedAttributes.Raw, &values); err != nil {
			return nil, nil, err
		}

		attrs, err := translator.ToGoTypeValues(values, *variableSchema)
		if err != nil {
			return nil, nil, err
		}

		mc.Attributes = attrs

		for n, v := range variableSchema.Properties {
			// Add sensitive from schema variable.
			if v.Value.WriteOnly {
				sensitiveVariables.Insert(fmt.Sprintf(`var\.%s`, n))
			}

			if n == openapi.WalrusContextVariableName {
				mc.Attributes[n] = resapi.NewContext().
					SetEnvironment(run.Namespace).
					SetProject(run.Spec.Project).
					SetResource(run.Spec.ResourceName)
			}
		}
	}

	// Outputs.
	if outputSchema != nil {
		sps := outputSchema.Properties
		mc.Outputs = make([]tfconfig.Output, 0, len(sps))

		sensitiveVariableRegex, err := matchAnyRegex(sensitiveVariables.UnsortedList())
		if err != nil {
			return nil, nil, err
		}

		for k, v := range sps {
			origin := openapi.GetExtOriginal(v.Value.Extensions)
			co := tfconfig.Output{
				Sensitive:    v.Value.WriteOnly,
				Name:         k,
				ResourceName: run.Spec.ResourceName,
				Value:        origin.ValueExpression,
			}

			if !v.Value.WriteOnly {
				// Update sensitive while output is from sensitive data, like secret.
				if sensitiveVariables.Len() != 0 && sensitiveVariableRegex.Match(origin.ValueExpression) {
					co.Sensitive = true
				}
			}

			mc.Outputs = append(mc.Outputs, co)
		}
	}

	return mc, templateSchema.Data.Terraform.RequiredProviders, nil
}

func (c *TerraformConfigurator) LoadProviders(
	ctx context.Context,
	connectors walruscore.Connectors,
) (map[string][]byte, error) {
	providerConfigs := make(map[string][]byte, len(connectors))
	loopbackKubeClient := system.LoopbackKubeClient.Get()

	for _, c := range connectors {
		if c.Spec.Type != resourcehandler.ConnectorTypeKubernetes {
			continue
		}

		// Get secret.
		sec, err := loopbackKubeClient.CoreV1().Secrets(c.Namespace).Get(ctx, c.Spec.SecretName, meta.GetOptions{})
		if err != nil {
			return nil, err
		}

		connCfg := walrus.ConnectorConfig{
			ObjectMeta: meta.ObjectMeta{
				Namespace: c.Namespace,
				Name:      c.Name,
			},
			Status: walrus.ConnectorConfigStatus{
				ApplicableEnvironmentType: c.Spec.ApplicableEnvironmentType,
				Category:                  c.Spec.Category,
				Type:                      c.Spec.Type,
				Version:                   c.Spec.Config.Version,
				Data:                      sec.Data,
				ConditionSummary:          c.Status.ConditionSummary,
			},
		}

		_, s, err := k8s.LoadApiConfig(connCfg)
		if err != nil {
			return nil, err
		}

		// NB(alex) the secret file name must be config + connector id to
		// match with terraform provider in config convert.
		secretFileName := util.GetK8sSecretName(c.Name)
		providerConfigs[secretFileName] = []byte(s)
	}

	return providerConfigs, nil
}

func matchAnyRegex(list []string) (*regexp.Regexp, error) {
	var sb strings.Builder

	sb.WriteString("(")

	for i, v := range list {
		sb.WriteString(v)

		if i < len(list)-1 {
			sb.WriteString("|")
		}
	}

	sb.WriteString(")")

	return regexp.Compile(sb.String())
}

func updateOutputWithVariables(variables walrus.Variables, moduleConfig *tfconfig.ModuleConfig) (map[string]bool, error) {
	var (
		variableOpts         = make(map[string]bool)
		encryptVariableNames = sets.NewString()
	)

	for _, s := range variables {
		variableOpts[s.Name] = s.Spec.Sensitive

		if s.Spec.Sensitive {
			encryptVariableNames.Insert(config.VariablePrefix + s.Name)
		}
	}

	if encryptVariableNames.Len() == 0 {
		return variableOpts, nil
	}

	reg, err := matchAnyRegex(encryptVariableNames.UnsortedList())
	if err != nil {
		return nil, err
	}

	var shouldEncryptAttr []string

	for k, v := range moduleConfig.Attributes {
		b, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}

		matches := reg.FindAllString(string(b), -1)
		if len(matches) != 0 {
			shouldEncryptAttr = append(shouldEncryptAttr, fmt.Sprintf(`var\.%s`, k))
		}
	}

	// Outputs use encrypted variable should set to sensitive.
	for i, v := range moduleConfig.Outputs {
		if v.Sensitive {
			continue
		}

		reg, err := matchAnyRegex(shouldEncryptAttr)
		if err != nil {
			return nil, err
		}

		if reg.MatchString(string(v.Value)) {
			moduleConfig.Outputs[i].Sensitive = true
		}
	}

	return variableOpts, nil
}

func getVarConfigOptions(
	variables walrus.Variables,
	resourceOutputs map[string]runapi.OutputValue,
) tfconfig.CreateOptions {
	varsConfigOpts := tfconfig.CreateOptions{
		Attributes: map[string]any{},
	}

	for _, v := range variables {
		varsConfigOpts.Attributes[config.VariablePrefix+v.Name] = v.Spec.Value
	}

	// Setup resource outputs.
	for n, v := range resourceOutputs {
		varsConfigOpts.Attributes[config.ResourcePrefix+n] = v.Value
	}

	return varsConfigOpts
}
