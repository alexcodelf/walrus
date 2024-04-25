package config

import (
	"context"
	"errors"
	"regexp"

	meta "k8s.io/apimachinery/pkg/apis/meta/v1"

	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/system"
	"github.com/seal-io/walrus/pkg/systemsetting"
)

const (
	ConfigTypeTerraform = "terraform"
)

// Configurator is the interface to construct input configs and dependency connectors for the run.
type Configurator interface {
	InputLoader
	ProviderLoader
}

// InputLoader is the interface to construct input configs for the run.
type InputLoader interface {
	// LoadAll loads the configs files of the config options.
	Load(context.Context, *walruscore.ResourceRun, *Options) (map[string][]byte, error)
}

// ProviderLoader is the interface to construct dependency connectors files for the run.
type ProviderLoader interface {
	// LoadProviders loads the providers of the run required,
	// Some connectors may be required to deploy the service.
	LoadProviders(context.Context, walruscore.Connectors) (map[string][]byte, error)
}

// Options are the options for load a run config files.
type Options struct {
	// SecretMountPath of the deployment job.
	SecretMountPath string

	Connectors walruscore.Connectors

	ServerURL string
	Token     string
}

// GetConfigOptions sets the config loader options.
// It will fetch the resource run, environment, project, resource and subject.
func GetConfigOptions(
	ctx context.Context,
	run *walruscore.ResourceRun,
	secretMountPath string,
) (*Options, error) {
	loopbackKubeClient := system.LoopbackKubeClient.Get()

	opts := &Options{
		SecretMountPath: secretMountPath,
	}

	connectorBindingList, err := loopbackKubeClient.WalruscoreV1().ConnectorBindings(run.Namespace).List(ctx, meta.ListOptions{})
	if err != nil {
		return nil, err
	}

	for _, connectorBinding := range connectorBindingList.Items {
		connector, err := loopbackKubeClient.WalruscoreV1().Connectors(connectorBinding.Spec.Connector.Namespace).Get(
			ctx, connectorBinding.Spec.Connector.Name, meta.GetOptions{})
		if err != nil {
			return nil, err
		}

		opts.Connectors = append(opts.Connectors, *connector)
	}

	// Prepare address for terraform backend.
	opts.ServerURL, err = systemsetting.ServeUrl.Value(ctx)
	if err != nil {
		return nil, err
	}

	if opts.ServerURL == "" {
		return nil, errors.New("server address is empty")
	}

	return opts, nil
}

const (
	// VariablePrefix the prefix of the variable name.
	VariablePrefix = "_walrus_var_"

	// ResourcePrefix the prefix of the resource output name.
	ResourcePrefix = "_walrus_res_"
)

// InterpolationReg is the regular expression for matching non-reference or non-variable expressions.
// Reference: https://developer.hashicorp.com/terraform/language/expressions/strings#escape-sequences-1
// To handle escape sequences, ${xxx} is converted to $${xxx}.
// If there are more than two consecutive $ symbols, like $${xxx}, they are further converted to $$${xxx}.
// During Terraform processing, $${} is ultimately transformed back to ${};
// this interpolation is used to ensure a WYSIWYG user experience.
var InterpolationReg = regexp.MustCompile(`\$\{((var\.)?([^.}]+)(?:\.([^.}]+))?)[^\}]*\}`)

type RunOpts struct {
	ResourceRun *walruscore.ResourceRun

	ResourceName string

	ProjectName     string
	EnvironmentName string
}
