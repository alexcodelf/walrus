package resourcerun

import (
	"context"
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-config-inspect/tfconfig"
	apiconfig "github.com/seal-io/walrus/pkg/apis/config"
	"github.com/seal-io/walrus/pkg/auths"
	busrun "github.com/seal-io/walrus/pkg/bus/resourcerun"
	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/model/predicate"
	"github.com/seal-io/walrus/pkg/dao/model/templateversion"
	"github.com/seal-io/walrus/pkg/dao/types"
	"github.com/seal-io/walrus/pkg/dao/types/crypto"
	opk8s "github.com/seal-io/walrus/pkg/operator/k8s"
	"github.com/seal-io/walrus/pkg/settings"
	"github.com/seal-io/walrus/pkg/terraform/config"
	"github.com/seal-io/walrus/pkg/terraform/parser"
	"github.com/seal-io/walrus/pkg/terraform/util"
	"github.com/seal-io/walrus/utils/log"
	"github.com/seal-io/walrus/utils/pointer"
	"k8s.io/apimachinery/pkg/util/sets"
)

const (
	// _backendAPI the API path to terraform deploy backend.
	// Terraform will get and update deployment states from this API.
	_backendAPI = "/v1/projects/%s/environments/%s/resources/%s/runs/%s/terraform-states"
)

func NewTerraformInput() InputLoader {
	return &TerraformInput{}
}

// TerraformInput constructs the terraform config files for the run.
type TerraformInput struct {
	logger log.Logger
}

func (c TerraformInput) LoadMain(
	ctx context.Context,
	mc model.ClientSet,
	opts *ConfigLoaderOptions,
) (ConfigData, error) {
	planConfig, err := c.LoadAll(ctx, mc, opts)
	if err != nil {
		return nil, err
	}

	return planConfig[config.FileMain], nil
}

func (c TerraformInput) LoadAll(
	ctx context.Context,
	mc model.ClientSet,
	opts *ConfigLoaderOptions,
) (map[string]ConfigData, error) {
	// Prepare terraform tfConfig.
	//  get module configs from resource run.
	moduleConfig, providerRequirements, err := c.getModuleConfig(ctx, mc, opts)
	if err != nil {
		return nil, err
	}

	// Merge current and previous required providers.
	providerRequirements = append(providerRequirements,
		opts.ResourceRun.PreviousRequiredProviders...)

	requiredProviders := make(map[string]*tfconfig.ProviderRequirement, 0)
	for _, p := range providerRequirements {
		if _, ok := requiredProviders[p.Name]; !ok {
			requiredProviders[p.Name] = p.ProviderRequirement
		} else {
			c.logger.Warnf("duplicate provider requirement: %s", p.Name)
		}
	}

	runOpts := ParseOpts{
		ResourceRun:   opts.ResourceRun,
		ResourceName:  opts.Context.Resource.Name,
		ProjectID:     opts.Context.Project.ID,
		EnvironmentID: opts.Context.Environment.ID,
	}
	// Parse module attributes.
	attrs, variables, dependencyOutputs, err := ParseModuleAttributes(
		ctx,
		mc,
		moduleConfig.Attributes,
		false,
		runOpts,
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

	address, token, err := c.getBackendConfig(ctx, mc, opts)
	if err != nil {
		return nil, err
	}

	tfCreateOpts := map[string]config.CreateOptions{
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
				VariablePrefix:    _variablePrefix,
				ResourcePrefix:    _resourcePrefix,
				Variables:         wrapVariables,
				DependencyOutputs: dependencyOutputs,
			},
			OutputOptions: moduleConfig.Outputs,
		},
		config.FileVars: getVarConfigOptions(variables, dependencyOutputs),
	}

	configFiles := make(map[string][]byte, len(tfCreateOpts))

	for k, v := range tfCreateOpts {
		configFiles[k], err = config.CreateConfigToBytes(v)
		if err != nil {
			return nil, err
		}
	}

	// Save input plan to resource run.
	opts.ResourceRun.InputPlan = string(configFiles[config.FileMain])
	// If resource run does not inherit variables from cloned run,
	// then save the parsed variables to resource run.
	if len(opts.ResourceRun.Variables) == 0 {
		variableMap := make(crypto.Map[string, string], len(variables))
		for _, s := range variables {
			variableMap[s.Name] = string(s.Value)
		}
		opts.ResourceRun.Variables = variableMap
	}

	run, err := mc.ResourceRuns().UpdateOne(opts.ResourceRun).
		Set(opts.ResourceRun).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	if err = busrun.Notify(ctx, mc, run); err != nil {
		return nil, err
	}

	return configFiles, nil
}

func (c TerraformInput) LoadProviders(connectors model.Connectors) (map[string]ConfigData, error) {
	secretData := make(map[string]ConfigData)

	for _, c := range connectors {
		if c.Type != types.ConnectorTypeKubernetes {
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

func (c TerraformInput) getBackendConfig(
	ctx context.Context,
	mc model.ClientSet,
	opts *ConfigLoaderOptions,
) (address, token string, err error) {
	// Prepare address for terraform backend.
	serverAddress, err := settings.ServeUrl.Value(ctx, mc)
	if err != nil {
		return "", "", err
	}

	if serverAddress == "" {
		return "", "", errors.New("server address is empty")
	}
	address = fmt.Sprintf("%s%s", serverAddress,
		fmt.Sprintf(_backendAPI,
			opts.Context.Project.ID,
			opts.Context.Environment.ID,
			opts.Context.Resource.ID,
			opts.ResourceRun.ID))

	// Prepare API token for terraform backend.
	const _1Day = 60 * 60 * 24

	at, err := auths.CreateAccessToken(ctx,
		mc, opts.SubjectID, types.TokenKindDeployment, string(opts.ResourceRun.ID), pointer.Int(_1Day))
	if err != nil {
		return "", "", err
	}

	token = at.AccessToken

	return
}

// getModuleConfig returns module configs and required connectors to
// get terraform module config block from resource run.
func (c TerraformInput) getModuleConfig(
	ctx context.Context,
	mc model.ClientSet,
	opts *ConfigLoaderOptions,
) (*config.ModuleConfig, []types.ProviderRequirement, error) {
	var (
		requiredProviders = make([]types.ProviderRequirement, 0)
		predicates        = make([]predicate.TemplateVersion, 0)
	)

	predicates = append(predicates, templateversion.And(
		templateversion.Version(opts.ResourceRun.TemplateVersion),
		templateversion.TemplateID(opts.ResourceRun.TemplateID),
	))

	templateVersion, err := mc.TemplateVersions().
		Query().
		Select(
			templateversion.FieldID,
			templateversion.FieldTemplateID,
			templateversion.FieldName,
			templateversion.FieldVersion,
			templateversion.FieldSource,
			templateversion.FieldSchema,
			templateversion.FieldUiSchema,
		).
		Where(templateversion.Or(predicates...)).
		Only(ctx)
	if err != nil {
		return nil, nil, err
	}

	if len(templateVersion.Schema.RequiredProviders) != 0 {
		requiredProviders = append(requiredProviders, templateVersion.Schema.RequiredProviders...)
	}

	moduleConfig, err := getModuleConfig(opts.ResourceRun, templateVersion, opts)
	if err != nil {
		return nil, nil, err
	}

	return moduleConfig, requiredProviders, err
}
