package resourcerun

import (
	"context"
	"errors"

	"github.com/seal-io/walrus/pkg/dao"
	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/types"
	"github.com/seal-io/walrus/pkg/dao/types/object"
	"github.com/seal-io/walrus/pkg/dao/types/status"
	pkgenv "github.com/seal-io/walrus/pkg/environment"
)

// ConfigData is the config file data.
type ConfigData = []byte

// InputLoader is the interface to construct input configs and dependency connectors for the run.
type InputLoader interface {
	InputConfigLoader
	InputProviderLoader
}

// InputConfigLoader is the interface to construct input configs for the run.
type InputConfigLoader interface {
	// LoadMain loads the main config file of the config options.
	LoadMain(context.Context, model.ClientSet, *ConfigLoaderOptions) (ConfigData, error)
	// LoadAll loads the configs files of the config options.
	LoadAll(context.Context, model.ClientSet, *ConfigLoaderOptions) (map[string]ConfigData, error)
}

// InputProviderLoader is the interface to construct dependency connectors files for the run.
type InputProviderLoader interface {
	// LoadProviders loads the providers of the run required,
	// Some connectors may be required to deploy the service.
	LoadProviders(model.Connectors) (map[string]ConfigData, error)
}

// ConfigLoaderOptions are the options for load a run config files.
type ConfigLoaderOptions struct {
	// SecretMountPath of the deploy job.
	SecretMountPath string

	ResourceRun *model.ResourceRun
	Connectors  model.Connectors
	SubjectID   object.ID
	// Walrus Context.
	Context Context
}

// NewPlan creates a new plan with the plan type.
func NewInputLoader(deployerType string) InputLoader {
	switch deployerType {
	case types.DeployerTypeTF:
		return NewTerraformInput()
	default:
		return nil
	}
}

// GetConfigLoaderOptions sets the config loader options.
// It will fetch the resource run, environment, project, resource and subject.
func GetConfigLoaderOptions(
	ctx context.Context,
	mc model.ClientSet,
	run *model.ResourceRun,
	secretMountPath string,
) (*ConfigLoaderOptions, error) {
	opts := &ConfigLoaderOptions{
		ResourceRun:     run,
		SecretMountPath: secretMountPath,
	}

	if !status.ResourceRunStatusReady.IsUnknown(run) {
		return nil, errors.New("resource run is not running")
	}

	connectors, err := dao.GetConnectors(ctx, mc, run.EnvironmentID)
	if err != nil {
		return nil, err
	}

	proj, err := mc.Projects().Get(ctx, run.ProjectID)
	if err != nil {
		return nil, err
	}

	env, err := dao.GetEnvironmentByID(ctx, mc, run.EnvironmentID)
	if err != nil {
		return nil, err
	}

	res, err := mc.Resources().Get(ctx, run.ResourceID)
	if err != nil {
		return nil, err
	}

	sj, err := getSubject(ctx, mc, res)
	if err != nil {
		return nil, err
	}

	opts.Connectors = connectors
	opts.SubjectID = sj.ID

	// Walrus Context.
	opts.Context = *NewContext().
		SetProject(proj.ID, proj.Name).
		SetEnvironment(env.ID, env.Name, pkgenv.GetManagedNamespaceName(env)).
		SetResource(res.ID, res.Name)

	return opts, nil
}
