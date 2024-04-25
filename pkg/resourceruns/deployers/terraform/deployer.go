package terraform

import (
	"context"

	wf "github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
	"github.com/go-logr/logr"
	core "k8s.io/api/core/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"

	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/clients/clientset"
	runconfig "github.com/seal-io/walrus/pkg/resourceruns/config"
	runconfigs "github.com/seal-io/walrus/pkg/resourceruns/configs"
	"github.com/seal-io/walrus/pkg/resourceruns/deployer"
	runstatus "github.com/seal-io/walrus/pkg/resourceruns/status"
	"github.com/seal-io/walrus/pkg/system"
	"github.com/seal-io/walrus/pkg/systemsetting"
)

// TerraformDeployer terraform deployer to deploy the resource.
type TerraformDeployer struct {
	logger          logr.Logger
	clientSet       clientset.Interface
	runConfigurator runconfig.Configurator
}

func NewDeployer(opts deployer.CreateOptions) (deployer.Deployer, error) {
	clientSet := system.LoopbackKubeClient.Get()

	return &TerraformDeployer{
		logger:          klog.Background().WithName("resource-run").WithName("deployer-tf"),
		clientSet:       clientSet,
		runConfigurator: runconfigs.NewConfigurator(runconfig.ConfigTypeTerraform),
	}, nil
}

func (d TerraformDeployer) Type() deployer.Type {
	return deployer.TypeTerraform
}

// Apply creates a new resource run by the given resource,
// and drives the Kubernetes Job to create components of the resource.
func (d TerraformDeployer) Apply(
	ctx context.Context,
	run *walruscore.ResourceRun,
	opts deployer.ApplyOptions,
) (tmpl *wf.Template, err error) {
	defer d.errorHandle(run, err)

	run, err = runstatus.UpdateStatus(ctx, run)
	if err != nil {
		return
	}

	return d.computeTemplate(ctx, run, computeTemplateOptions{
		ResourceRunStepType: walruscore.ResourceRunStepTypeApply,
	})
}

func (d TerraformDeployer) Plan(
	ctx context.Context,
	run *walruscore.ResourceRun,
	opts deployer.PlanOptions,
) (tmpl *wf.Template, err error) {
	defer d.errorHandle(run, err)

	run, err = runstatus.UpdateStatus(ctx, run)
	if err != nil {
		return
	}
	return d.computeTemplate(ctx, run, computeTemplateOptions{
		ResourceRunStepType: walruscore.ResourceRunStepTypePlan,
	})
}

// Destroy creates a new resource run by the given resource,
// and drives the Kubernetes Job to clean the components of the resource.
func (d TerraformDeployer) Destroy(
	ctx context.Context,
	run *walruscore.ResourceRun,
	opts deployer.DestroyOptions,
) (tmpl *wf.Template, err error) {
	defer d.errorHandle(run, err)

	return d.computeTemplate(ctx, run, computeTemplateOptions{
		ResourceRunStepType: walruscore.ResourceRunStepTypeApply,
	})
}

// errorHandle handles the error of the deployer operation.
func (d TerraformDeployer) errorHandle(run *walruscore.ResourceRun, err error) {
	if err == nil {
		return
	}

	// Update a failure status.
	runstatus.SetStatusFalse(run, err.Error())

	// Report to resource run.
	_, updateErr := runstatus.UpdateStatus(context.Background(), run)
	if updateErr != nil {
		d.logger.Error(updateErr, "failed to update the status of the resource run")
	}
}

type computeTemplateOptions struct {
	// Type indicates the type of the job.
	ResourceRunStepType walruscore.ResourceRunStepType
}

// computeTemplate creates a argo workflow template to deploy, destroy or rollback the resource.
func (d TerraformDeployer) computeTemplate(ctx context.Context, run *walruscore.ResourceRun, opts computeTemplateOptions) (*wf.Template, error) {
	// Prepare tfConfig for deployment.
	secretOpts, err := runconfig.GetConfigOptions(ctx, run, _secretMountPath)
	if err != nil {
		return nil, err
	}

	if err = d.createK8sSecrets(ctx, run, secretOpts); err != nil {
		return nil, err
	}

	image, err := systemsetting.TerraformDeployerImage.Value(ctx)
	if err != nil {
		return nil, err
	}

	env := d.getEnv(ctx)

	localEnvironmentMode, err := systemsetting.DefaultEnvironmentMode.Value(ctx)
	if err != nil {
		return nil, err
	}

	// Create a argo workflow template.
	templateOpts := deployer.CreateTemplateOptions{
		ResourceRunStepType: opts.ResourceRunStepType,
		Image:               image,
		Env:                 env,
		DockerMode:          localEnvironmentMode == "docker",
		ServerURL:           secretOpts.ServerURL,
		Token:               secretOpts.Token,
	}

	return generateTemplate(run, templateOpts)
}

func (d TerraformDeployer) getEnv(ctx context.Context) (env []core.EnvVar) {
	if v, err := systemsetting.DeployerAllProxy.Value(ctx); err == nil && v != "" {
		env = append(env, core.EnvVar{
			Name:  "ALL_PROXY",
			Value: v,
		})
	}

	if v, err := systemsetting.DeployerHttpProxy.Value(ctx); err == nil && v != "" {
		env = append(env, core.EnvVar{
			Name:  "HTTP_PROXY",
			Value: v,
		})
	}

	if v, err := systemsetting.DeployerHttpsProxy.Value(ctx); err == nil && v != "" {
		env = append(env, core.EnvVar{
			Name:  "HTTPS_PROXY",
			Value: v,
		})
	}

	if v, err := systemsetting.DeployerNoProxy.Value(ctx); err == nil && v != "" {
		env = append(env, core.EnvVar{
			Name:  "NO_PROXY",
			Value: v,
		})
	}

	if enable, err := systemsetting.EnableRemoteTlsVerify.ValueBool(ctx); !enable && err != nil {
		env = append(env, core.EnvVar{
			Name:  "GIT_SSL_NO_VERIFY",
			Value: "true",
		})
	}

	if v, err := systemsetting.TerraformDeployerNetworkMirrorUrl.Value(ctx); err == nil && v != "" {
		env = append(env,
			core.EnvVar{
				Name:  "TF_CLI_NETWORK_MIRROR_URL",
				Value: v,
			},
			core.EnvVar{
				Name:  "TF_CLI_NETWORK_MIRROR_INSECURE_SKIP_VERIFY",
				Value: "true",
			})
	}

	return env
}

// createK8sSecrets creates the k8s secrets for deployer template.
func (d TerraformDeployer) createK8sSecrets(ctx context.Context, run *walruscore.ResourceRun, opts *runconfig.Options) error {
	secretData := make(map[string][]byte)

	// Prepare terraform config files bytes for deployment.
	inputConfigs, err := d.runConfigurator.Load(ctx, run, opts)
	if err != nil {
		return err
	}

	for k, v := range inputConfigs {
		secretData[k] = v
	}

	// Mount the provider configs(e.g. kubeconfig) to secret.
	providerConfigs, err := d.runConfigurator.LoadProviders(ctx, opts.Connectors)
	if err != nil {
		return err
	}

	for k, v := range providerConfigs {
		secretData[k] = v
	}

	// Mount deploy access token to secret.
	secretData[_accessTokenkey] = []byte(opts.Token)

	// Create deployment secret.
	secret := &core.Secret{
		ObjectMeta: meta.ObjectMeta{
			Namespace: run.Namespace,
			Name:      run.Status.ConfigSecretName,
			OwnerReferences: []meta.OwnerReference{
				{
					APIVersion: walruscore.SchemeGroupVersion.String(),
					Kind:       "ResourceRun",
					Name:       run.Name,
					UID:        run.UID,
				},
			},
		},
		Data: secretData,
	}

	_, err = d.clientSet.CoreV1().Secrets(run.Namespace).Create(ctx, secret, meta.CreateOptions{})
	if err != nil {
		return err
	}

	return nil
}
