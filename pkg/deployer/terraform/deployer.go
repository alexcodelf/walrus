package terraform

import (
	"context"
	"errors"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"

	revisionbus "github.com/seal-io/walrus/pkg/bus/servicerevision"
	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/model/serviceresource"
	"github.com/seal-io/walrus/pkg/dao/types"
	"github.com/seal-io/walrus/pkg/dao/types/status"
	deptypes "github.com/seal-io/walrus/pkg/deployer/types"
	pkgrevision "github.com/seal-io/walrus/pkg/servicerevision"
	"github.com/seal-io/walrus/pkg/settings"
	"github.com/seal-io/walrus/utils/log"
)

// DeployerType the type of deployer.
const DeployerType = types.DeployerTypeTF

// Deployer terraform deployer to deploy the service.
type Deployer struct {
	logger log.Logger

	modelClient     model.ClientSet
	clientSet       *kubernetes.Clientset
	planner         pkgrevision.IPlan
	revisionManager *pkgrevision.Manager
}

func NewDeployer(_ context.Context, opts deptypes.CreateOptions) (deptypes.Deployer, error) {
	clientSet, err := kubernetes.NewForConfig(opts.KubeConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create kubernetes client set: %w", err)
	}

	if opts.Type != DeployerType {
		return nil, errors.New("deployer type is not terraform")
	}

	return &Deployer{
		logger:          log.WithName("deployer").WithName("tf"),
		modelClient:     opts.ModelClient,
		clientSet:       clientSet,
		planner:         pkgrevision.NewPlan(opts.Type, opts.ModelClient),
		revisionManager: pkgrevision.NewRevisionManager(opts.ModelClient),
	}, nil
}

func (d Deployer) Type() deptypes.Type {
	return DeployerType
}

// Apply creates a new service revision by the given service,
// and drives the Kubernetes Job to create resources of the service.
func (d Deployer) Apply(ctx context.Context, service *model.Service, opts deptypes.ApplyOptions) (err error) {
	revision, err := d.createRevision(ctx, pkgrevision.CreateOptions{
		ServiceID: service.ID,
		JobType:   JobTypeApply,
	})
	if err != nil {
		return err
	}

	defer func() {
		if err == nil {
			return
		}

		// Update a failure status.
		status.ServiceRevisionStatusPending.False(revision, err.Error())

		// Report to service revision.
		_ = d.updateRevisionStatus(ctx, revision)
	}()

	return d.createK8sJob(ctx, createK8sJobOptions{
		Type:            JobTypeApply,
		ServiceRevision: revision,
	})
}

// Destroy creates a new service revision by the given service,
// and drives the Kubernetes Job to clean the resources of the service.
func (d Deployer) Destroy(ctx context.Context, service *model.Service, opts deptypes.DestroyOptions) (err error) {
	revision, err := d.createRevision(ctx, pkgrevision.CreateOptions{
		ServiceID: service.ID,
		JobType:   JobTypeDestroy,
	})
	if err != nil {
		return err
	}

	defer func() {
		if err == nil {
			return
		}

		// Update a failure status.
		status.ServiceRevisionStatusPending.False(revision, err.Error())

		// Report to service revision.
		_ = d.updateRevisionStatus(ctx, revision)
	}()

	// If no resource exists, skip job and set revision status succeed.
	exist, err := d.modelClient.ServiceResources().Query().
		Where(serviceresource.ServiceID(service.ID)).
		Exist(ctx)
	if err != nil {
		return err
	}

	if !exist {
		status.ServiceRevisionStatusPending.True(revision, "")
		return d.updateRevisionStatus(ctx, revision)
	}

	return d.createK8sJob(ctx, createK8sJobOptions{
		Type:            JobTypeDestroy,
		ServiceRevision: revision,
	})
}

type createK8sJobOptions struct {
	// Type indicates the type of the job.
	Type string
	// ServiceRevision indicates the service revision to create the deployment job.
	ServiceRevision *model.ServiceRevision
}

// createK8sJob creates a k8s job to deploy, destroy or rollback the service.
func (d Deployer) createK8sJob(ctx context.Context, opts createK8sJobOptions) error {
	// Prepare tfConfig for deployment.
	secretOpts := pkgrevision.PlanOptions{
		SecretMountPath: SecretMountPath,
		ServiceRevision: opts.ServiceRevision,
	}

	if err := pkgrevision.SetPlanOptions(ctx, d.modelClient, &secretOpts); err != nil {
		return err
	}

	if err := d.createK8sSecrets(ctx, secretOpts); err != nil {
		return err
	}

	jobImage, err := settings.DeployerImage.Value(ctx, d.modelClient)
	if err != nil {
		return err
	}

	jobEnv, err := d.getProxyEnv(ctx)
	if err != nil {
		return err
	}

	// Create deployment job.
	jobOpts := JobCreateOptions{
		Type:              opts.Type,
		ServiceRevisionID: opts.ServiceRevision.ID.String(),
		Image:             jobImage,
		Env:               jobEnv,
	}

	return CreateJob(ctx, d.clientSet, jobOpts)
}

func (d Deployer) getProxyEnv(ctx context.Context) ([]corev1.EnvVar, error) {
	var env []corev1.EnvVar

	allProxy, err := settings.DeployerAllProxy.Value(ctx, d.modelClient)
	if err != nil {
		return nil, err
	}

	if allProxy != "" {
		env = append(env, corev1.EnvVar{
			Name:  "ALL_PROXY",
			Value: allProxy,
		})
	}

	httpProxy, err := settings.DeployerHttpProxy.Value(ctx, d.modelClient)
	if err != nil {
		return nil, err
	}

	if httpProxy != "" {
		env = append(env, corev1.EnvVar{
			Name:  "HTTP_PROXY",
			Value: httpProxy,
		})
	}

	httpsProxy, err := settings.DeployerHttpsProxy.Value(ctx, d.modelClient)
	if err != nil {
		return nil, err
	}

	if httpsProxy != "" {
		env = append(env, corev1.EnvVar{
			Name:  "HTTPS_PROXY",
			Value: httpsProxy,
		})
	}

	noProxy, err := settings.DeployerNoProxy.Value(ctx, d.modelClient)
	if err != nil {
		return nil, err
	}

	if noProxy != "" {
		env = append(env, corev1.EnvVar{
			Name:  "NO_PROXY",
			Value: noProxy,
		})
	}

	return env, nil
}

func (d Deployer) updateRevisionStatus(ctx context.Context, ar *model.ServiceRevision) error {
	// Report to service revision.
	ar.Status.SetSummary(status.WalkServiceRevision(&ar.Status))

	ar, err := d.modelClient.ServiceRevisions().UpdateOne(ar).
		SetStatus(ar.Status).
		Save(ctx)
	if err != nil {
		return err
	}

	if err = revisionbus.Notify(ctx, d.modelClient, ar); err != nil {
		d.logger.Error(err)
		return err
	}

	return nil
}

// createK8sSecrets creates the k8s secrets for deployment.
func (d Deployer) createK8sSecrets(ctx context.Context, opts pkgrevision.PlanOptions) error {
	secretData := make(map[string][]byte)
	// SecretName terraform tfConfig name.
	secretName := _jobSecretPrefix + string(opts.ServiceRevision.ID)

	// Prepare terraform config files bytes for deployment.
	tfConfigs, err := d.planner.LoadConfigs(ctx, opts)
	if err != nil {
		return err
	}

	for k, v := range tfConfigs {
		secretData[k] = v
	}

	// Mount the provider configs(e.g. kubeconfig) to secret.
	providerData, err := d.planner.LoadConnectorConfigs(opts.Connectors)
	if err != nil {
		return err
	}

	for k, v := range providerData {
		secretData[k] = v
	}

	// Create deployment secret.
	if err = CreateSecret(ctx, d.clientSet, secretName, secretData); err != nil {
		return err
	}

	return nil
}

// createRevision creates a new service revision.
// Get the latest revision, and check it if it is running.
// If not running, then apply the latest revision.
// If running, then wait for the latest revision to be applied.
func (d Deployer) createRevision(
	ctx context.Context,
	opts pkgrevision.CreateOptions,
) (*model.ServiceRevision, error) {
	return d.revisionManager.Create(ctx, opts)
}
