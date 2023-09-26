package servicerevision

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-config-inspect/tfconfig"
	"k8s.io/apimachinery/pkg/util/sets"

	apiconfig "github.com/seal-io/walrus/pkg/apis/config"
	"github.com/seal-io/walrus/pkg/auths"
	revisionbus "github.com/seal-io/walrus/pkg/bus/servicerevision"
	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/model/predicate"
	"github.com/seal-io/walrus/pkg/dao/model/templateversion"
	"github.com/seal-io/walrus/pkg/dao/types"
	"github.com/seal-io/walrus/pkg/dao/types/crypto"
	"github.com/seal-io/walrus/pkg/dao/types/property"
	opk8s "github.com/seal-io/walrus/pkg/operator/k8s"
	pkgservice "github.com/seal-io/walrus/pkg/service"
	"github.com/seal-io/walrus/pkg/settings"
	"github.com/seal-io/walrus/pkg/terraform/config"
	"github.com/seal-io/walrus/pkg/terraform/parser"
	"github.com/seal-io/walrus/pkg/terraform/util"
	"github.com/seal-io/walrus/utils/log"
	"github.com/seal-io/walrus/utils/pointer"
)

// TerraformPlan handler the revision Plan of terraform deployer.
type TerraformPlan struct {
	logger log.Logger

	modelClient model.ClientSet
}

// NewTerraformPlan creates a new terraform plan of service revision.
func NewTerraformPlan(mc model.ClientSet) *TerraformPlan {
	return &TerraformPlan{
		logger:      log.WithName("service.revision-plan"),
		modelClient: mc,
	}
}

// LoadPlan will get the revision input plan of terraform.
// It will generate the main.tf for the revision.
func (t TerraformPlan) LoadPlan(ctx context.Context, opts PlanOptions) ([]byte, error) {
	configBytes, err := t.LoadConfigs(ctx, opts)
	if err != nil {
		return nil, err
	}

	return configBytes[config.FileMain], nil
}

func (t TerraformPlan) LoadConfigs(ctx context.Context, opts PlanOptions) (map[string][]byte, error) {
	if opts.ServiceRevision.InputPlanConfigs != nil {
		return opts.ServiceRevision.InputPlanConfigs, nil
	}

	// Prepare terraform tfConfig.
	//  get module configs from service revision.
	moduleConfig, providerRequirements, err := t.getModuleConfig(ctx, opts)
	if err != nil {
		return nil, err
	}

	// Merge current and previous required providers.
	providerRequirements = append(providerRequirements,
		opts.ServiceRevision.PreviousRequiredProviders...)

	requiredProviders := make(map[string]*tfconfig.ProviderRequirement, 0)
	for _, p := range providerRequirements {
		if _, ok := requiredProviders[p.Name]; !ok {
			requiredProviders[p.Name] = p.ProviderRequirement
		} else {
			t.logger.Warnf("duplicate provider requirement: %s", p.Name)
		}
	}

	serviceOpts := pkgservice.ParseAttributesOptions{
		ServiceRevision: opts.ServiceRevision,
		ServiceName:     opts.ServiceName,
		ProjectID:       opts.ProjectID,
		EnvironmentID:   opts.EnvironmentID,
	}
	// Parse module attributes.
	variables, dependencyOutputs, err := pkgservice.ParseModuleAttributes(
		ctx,
		t.modelClient,
		moduleConfig.Attributes,
		false,
		serviceOpts,
	)
	if err != nil {
		return nil, err
	}

	// Update output sensitive with variables.
	wrapVariables, err := setOutputSensitiveWithVariables(variables, moduleConfig)
	if err != nil {
		return nil, err
	}

	// Prepare terraform config files to be mounted to secret.
	requiredProviderNames := sets.NewString()
	for _, p := range providerRequirements {
		requiredProviderNames = requiredProviderNames.Insert(p.Name)
	}

	address, token, err := t.getBackendConfig(ctx, opts)
	if err != nil {
		return nil, err
	}

	// Options for create terraform config.
	planConfigOptions := map[string]config.CreateOptions{
		config.FileMain: {
			TerraformOptions: &config.TerraformOptions{
				Token:                token,
				Address:              address,
				SkipTLSVerify:        !apiconfig.TlsCertified.Get(),
				ProviderRequirements: requiredProviders,
			},
			ProviderOptions: &config.ProviderOptions{
				RequiredProviderNames: requiredProviderNames.List(),
				Connectors:            opts.Connectors,
				SecretMonthPath:       opts.SecretMountPath,
				ConnectorSeparator:    parser.ConnectorSeparator,
			},
			ModuleOptions: &config.ModuleOptions{
				ModuleConfigs: []*config.ModuleConfig{moduleConfig},
			},
			VariableOptions: &config.VariableOptions{
				VariablePrefix:    pkgservice.VariablePrefix,
				ServicePrefix:     pkgservice.ServicePrefix,
				Variables:         wrapVariables,
				DependencyOutputs: dependencyOutputs,
			},
			OutputOptions: moduleConfig.Outputs,
		},
		config.FileVars: getVarConfigOptions(variables, dependencyOutputs),
	}
	planConfigs := make(map[string][]byte, 0)

	for k, v := range planConfigOptions {
		planConfigs[k], err = config.CreateConfigToBytes(v)
		if err != nil {
			return nil, err
		}
	}

	// Save input plan to service revision.
	opts.ServiceRevision.InputPlanConfigs = planConfigs
	// If service revision does not inherit variables from cloned revision,
	// then save the parsed variables to service revision.
	if len(opts.ServiceRevision.Variables) == 0 {
		variableMap := make(crypto.Map[string, string], len(variables))
		for _, s := range variables {
			variableMap[s.Name] = string(s.Value)
		}
		opts.ServiceRevision.Variables = variableMap
	}

	revision, err := t.modelClient.ServiceRevisions().UpdateOne(opts.ServiceRevision).
		Set(opts.ServiceRevision).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	if err = revisionbus.Notify(ctx, t.modelClient, revision); err != nil {
		return nil, err
	}

	return planConfigs, nil
}

// LoadConnectorConfigs loads the connector for terraform provider.
func (t TerraformPlan) LoadConnectorConfigs(connectors model.Connectors) (map[string][]byte, error) {
	secretData := make(map[string][]byte)

	for _, c := range connectors {
		if c.Type != types.ConnectorTypeK8s {
			continue
		}

		_, s, err := opk8s.LoadApiConfig(*c)
		if err != nil {
			return nil, err
		}

		// NB(alex) the secret file name must be config + connector id to
		// match with terraform provider in config convert.
		secretFileName := util.GetK8sSecretName(c.ID.String())
		secretData[secretFileName] = []byte(s)
	}

	return secretData, nil
}

func (t TerraformPlan) getBackendConfig(ctx context.Context, opts PlanOptions) (address, token string, err error) {
	// Prepare address for terraform backend.
	serverAddress, err := settings.ServeUrl.Value(ctx, t.modelClient)
	if err != nil {
		return "", "", err
	}

	if serverAddress == "" {
		return "", "", errors.New("server address is empty")
	}
	address = fmt.Sprintf("%s%s", serverAddress,
		fmt.Sprintf(_backendAPI,
			opts.ProjectID,
			opts.EnvironmentID,
			opts.ServiceID,
			opts.ServiceRevision.ID))

	// Prepare API token for terraform backend.
	const _1Day = 60 * 60 * 24

	at, err := auths.CreateAccessToken(ctx,
		t.modelClient, opts.SubjectID, types.TokenKindDeployment, string(opts.ServiceRevision.ID), pointer.Int(_1Day))
	if err != nil {
		return "", "", err
	}

	token = at.AccessToken

	return
}

// getModuleConfig returns module configs and required connectors to
// get terraform module config block from service revision.
func (t TerraformPlan) getModuleConfig(
	ctx context.Context,
	opts PlanOptions,
) (*config.ModuleConfig, []types.ProviderRequirement, error) {
	var (
		requiredProviders = make([]types.ProviderRequirement, 0)
		predicates        = make([]predicate.TemplateVersion, 0)
	)

	predicates = append(predicates, templateversion.And(
		templateversion.Name(opts.ServiceRevision.TemplateName),
		templateversion.Version(opts.ServiceRevision.TemplateVersion),
	))

	templateVersion, err := t.modelClient.TemplateVersions().
		Query().
		Select(
			templateversion.FieldID,
			templateversion.FieldTemplateID,
			templateversion.FieldName,
			templateversion.FieldVersion,
			templateversion.FieldSource,
			templateversion.FieldSchema,
		).
		Where(templateversion.Or(predicates...)).
		Only(ctx)
	if err != nil {
		return nil, nil, err
	}

	if templateVersion.Schema != nil {
		requiredProviders = append(requiredProviders, templateVersion.Schema.RequiredProviders...)
	}

	mc, err := getModuleConfig(opts.ServiceRevision, templateVersion, opts)
	if err != nil {
		return nil, nil, err
	}

	return mc, requiredProviders, err
}

// getModuleConfig get module config of terraform.
func getModuleConfig(
	revision *model.ServiceRevision,
	template *model.TemplateVersion,
	opts PlanOptions,
) (*config.ModuleConfig, error) {
	var (
		props              = make(property.Properties, len(revision.Attributes))
		typesWith          = revision.Attributes.TypesWith(template.Schema.Variables)
		sensitiveVariables = sets.Set[string]{}
	)

	for k, v := range revision.Attributes {
		props[k] = property.Property{
			Type:  typesWith[k],
			Value: v,
		}
	}

	attrs, err := props.TypedValues()
	if err != nil {
		return nil, err
	}

	mc := &config.ModuleConfig{
		Name:       opts.ServiceName,
		Source:     template.Source,
		Schema:     template.Schema,
		Attributes: attrs,
	}

	if template.Schema == nil {
		return mc, nil
	}

	for _, v := range template.Schema.Variables {
		// Add sensitive from schema variable.
		if v.Sensitive {
			sensitiveVariables.Insert(fmt.Sprintf(`var\.%s`, v.Name))
		}

		// Add seal metadata.
		var attrValue string

		switch v.Name {
		case pkgservice.WalrusMetadataProjectName:
			attrValue = opts.ProjectName
		case pkgservice.WalrusMetadataEnvironmentName:
			attrValue = opts.EnvironmentName
		case pkgservice.WalrusMetadataServiceName:
			attrValue = opts.ServiceName
		case pkgservice.WalrusMetadataProjectID:
			attrValue = opts.ProjectID.String()
		case pkgservice.WalrusMetadataEnvironmentID:
			attrValue = opts.EnvironmentID.String()
		case pkgservice.WalrusMetadataServiceID:
			attrValue = opts.ServiceID.String()
		case pkgservice.WalrusMetadataNamespaceName:
			attrValue = opts.ManagedNamespaceName
		}

		if attrValue != "" {
			mc.Attributes[v.Name] = attrValue
		}
	}

	sensitiveVariableRegex, err := matchAnyRegex(sensitiveVariables.UnsortedList())
	if err != nil {
		return nil, err
	}

	mc.Outputs = make([]config.Output, len(template.Schema.Outputs))
	for i, v := range template.Schema.Outputs {
		mc.Outputs[i].Sensitive = v.Sensitive
		mc.Outputs[i].Name = v.Name
		mc.Outputs[i].ServiceName = opts.ServiceName
		mc.Outputs[i].Value = v.Value

		if v.Sensitive {
			continue
		}

		// Update sensitive while output is from sensitive data, like secret.
		if sensitiveVariables.Len() != 0 && sensitiveVariableRegex.Match(v.Value) {
			mc.Outputs[i].Sensitive = true
		}
	}

	return mc, nil
}

// matchAnyRegex get regex of match any list string.
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

// setOutputSensitiveWithVariables update output with variables.
// Sensitive output should not show the value.
func setOutputSensitiveWithVariables(
	variables model.Variables,
	moduleConfig *config.ModuleConfig,
) (map[string]bool, error) {
	var (
		variableOpts         = make(map[string]bool)
		encryptVariableNames = sets.NewString()
	)

	for _, s := range variables {
		variableOpts[s.Name] = s.Sensitive

		if s.Sensitive {
			encryptVariableNames.Insert(pkgservice.VariablePrefix + s.Name)
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

// getVarConfigOptions get terraform tf.vars config.
func getVarConfigOptions(variables model.Variables, serviceOutputs map[string]parser.OutputState) config.CreateOptions {
	varsConfigOpts := config.CreateOptions{
		Attributes: map[string]any{},
	}

	for _, v := range variables {
		varsConfigOpts.Attributes[pkgservice.VariablePrefix+v.Name] = v.Value
	}

	// Setup service outputs.
	for n, v := range serviceOutputs {
		varsConfigOpts.Attributes[pkgservice.ServicePrefix+n] = v.Value
	}

	return varsConfigOpts
}
