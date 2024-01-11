package terraform

import (
	"context"
	"fmt"
	"io"
	"time"

	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	coreclient "k8s.io/client-go/kubernetes/typed/core/v1"

	"github.com/seal-io/walrus/pkg/dao/types"
	"github.com/seal-io/walrus/pkg/dao/types/object"
	opk8s "github.com/seal-io/walrus/pkg/operator/k8s"
	"github.com/seal-io/walrus/pkg/operator/k8s/kube"
	pkgrun "github.com/seal-io/walrus/pkg/resourcerun"
	"github.com/seal-io/walrus/utils/log"
	"github.com/seal-io/walrus/utils/pointer"
)

const (
	JobTypeApply   = "apply"
	JobTypeDestroy = "destroy"
)

type JobCreateOptions struct {
	// Type is the deployment type of job, apply or destroy or other.
	Type          string
	ResourceRunID string
	Image         string
	Env           []corev1.EnvVar
	DockerMode    bool
}

type StreamJobLogsOptions struct {
	Cli     *coreclient.CoreV1Client
	RunID   object.ID
	JobType string
	Out     io.Writer
}

const (
	// _podName the name of the pod.
	_podName = "deployer"

	// _jobNameFormat the format of job name.
	_jobNameFormat = "tf-job-%s-%s"
	// _jobSecretPrefix the prefix of secret name.
	_jobSecretPrefix = "tf-secret-"
	// SecretMountPath the path to mount the secret.
	SecretMountPath = "/var/terraform/secrets"
	// _workdir the working directory of the job.
	_workdir = "/var/terraform/workspace"
)

const (
	// _applyCommands the commands to apply deployment of the resource.
	_applyCommands = "terraform init -no-color && terraform apply -auto-approve -no-color"
	// _destroyCommands the commands to destroy deployment of the resource.
	_destroyCommands = "terraform init -no-color && terraform destroy -auto-approve -no-color"
)

// CreateJob create a job to run terraform deployment.
func CreateJob(ctx context.Context, clientSet *kubernetes.Clientset, opts JobCreateOptions) error {
	var (
		logger = log.WithName("deployer").WithName("tf")

		backoffLimit            int32 = 0
		ttlSecondsAfterFinished int32 = 3600
		name                          = getK8sJobName(_jobNameFormat, opts.Type, opts.ResourceRunID)
		configName                    = _jobSecretPrefix + opts.ResourceRunID
	)

	secret, err := clientSet.CoreV1().Secrets(types.WalrusSystemNamespace).Get(ctx, configName, metav1.GetOptions{})
	if err != nil {
		return err
	}

	podTemplate := getPodTemplate(opts.ResourceRunID, configName, opts)
	job := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec: batchv1.JobSpec{
			Template:                podTemplate,
			BackoffLimit:            &backoffLimit,
			TTLSecondsAfterFinished: &ttlSecondsAfterFinished,
		},
	}

	job, err = clientSet.BatchV1().Jobs(types.WalrusSystemNamespace).Create(ctx, job, metav1.CreateOptions{})
	if err != nil {
		if kerrors.IsAlreadyExists(err) {
			logger.Warnf("k8s job %s already exists", name)
		} else {
			return err
		}
	}

	// Set ownerReferences to secret with the job name.
	secret.ObjectMeta = metav1.ObjectMeta{
		Name: configName,
		OwnerReferences: []metav1.OwnerReference{
			{
				APIVersion: batchv1.SchemeGroupVersion.String(),
				Kind:       "Job",
				Name:       name,
				UID:        job.UID,
			},
		},
	}

	_, err = clientSet.CoreV1().Secrets(types.WalrusSystemNamespace).Update(ctx, secret, metav1.UpdateOptions{})
	if err != nil {
		return err
	}

	logger.Debugf("k8s job %s created", name)

	return nil
}

// CreateSecret create a secret to store terraform config.
func CreateSecret(ctx context.Context, clientSet *kubernetes.Clientset, name string, data map[string][]byte) error {
	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{Name: name},
		Data:       data,
	}

	_, err := clientSet.CoreV1().Secrets(types.WalrusSystemNamespace).Create(ctx, secret, metav1.CreateOptions{})
	if err != nil && !kerrors.IsAlreadyExists(err) {
		return err
	}

	return nil
}

// getPodTemplate returns a pod template for deployment.
func getPodTemplate(runID, configName string, opts JobCreateOptions) corev1.PodTemplateSpec {
	var (
		command       = []string{"/bin/sh", "-c"}
		deployCommand = fmt.Sprintf("cp %s/main.tf main.tf && ", SecretMountPath)
		varfile       = fmt.Sprintf(" -var-file=%s/terraform.tfvars", SecretMountPath)
	)

	switch opts.Type {
	case JobTypeApply:
		deployCommand += _applyCommands + varfile
	case JobTypeDestroy:
		deployCommand += _destroyCommands + varfile
	}

	command = append(command, deployCommand)

	volumeMounts := []corev1.VolumeMount{
		{
			Name:      configName,
			MountPath: SecretMountPath,
			ReadOnly:  false,
		},
	}

	volumes := []corev1.Volume{
		{
			Name: configName,
			VolumeSource: corev1.VolumeSource{
				Secret: &corev1.SecretVolumeSource{
					SecretName: configName,
				},
			},
		},
	}

	securityContext := &corev1.PodSecurityContext{}

	if opts.DockerMode {
		volumeMounts = append(volumeMounts, corev1.VolumeMount{
			Name:      "docker-sock",
			MountPath: "/var/run/docker.sock",
		})

		volumes = append(volumes, corev1.Volume{
			Name: "docker-sock",
			VolumeSource: corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: "/var/run/docker.sock",
				},
			},
		})
		securityContext.RunAsUser = pointer.Int64(0)
	}

	return corev1.PodTemplateSpec{
		ObjectMeta: metav1.ObjectMeta{
			Name: _podName,
			Labels: map[string]string{
				pkgrun.ResourceRunIDLabel: runID,
			},
		},
		Spec: corev1.PodSpec{
			HostNetwork:        true,
			DNSPolicy:          corev1.DNSClusterFirstWithHostNet,
			ServiceAccountName: types.DeployerServiceAccountName,
			RestartPolicy:      corev1.RestartPolicyNever,
			Containers: []corev1.Container{
				{
					Name:            "deployment",
					Image:           opts.Image,
					WorkingDir:      _workdir,
					Command:         command,
					ImagePullPolicy: corev1.PullIfNotPresent,
					VolumeMounts:    volumeMounts,
					Env:             opts.Env,
				},
			},
			Volumes:         volumes,
			SecurityContext: securityContext,
		},
	}
}

// getK8sJobName returns the kubernetes job name for the given resource run id.
func getK8sJobName(format, jobType, runID string) string {
	return fmt.Sprintf(format, jobType, runID)
}

// StreamJobLogs streams the logs of a job.
func StreamJobLogs(ctx context.Context, opts StreamJobLogsOptions) error {
	var (
		jobName       = getK8sJobName(_jobNameFormat, opts.JobType, opts.RunID.String())
		labelSelector = "job-name=" + jobName
	)

	podList, err := opts.Cli.Pods(types.WalrusSystemNamespace).
		List(ctx, metav1.ListOptions{LabelSelector: labelSelector})
	if err != nil {
		return err
	}

	if len(podList.Items) == 0 {
		return nil
	}

	jobPod := &podList.Items[0]

	err = wait.PollUntilContextTimeout(ctx, 1*time.Second, 1*time.Minute, true,
		func(ctx context.Context) (bool, error) {
			var getErr error

			jobPod, getErr = opts.Cli.Pods(types.WalrusSystemNamespace).Get(ctx, jobPod.Name, metav1.GetOptions{
				ResourceVersion: "0",
			})
			if getErr != nil {
				return false, getErr
			}

			return kube.IsPodReady(jobPod), nil
		})
	if err != nil {
		return err
	}

	states := kube.GetContainerStates(jobPod)
	if len(states) == 0 {
		return nil
	}

	var (
		containerName, containerType = states[0].Name, states[0].Type
		follow                       = kube.IsContainerRunning(jobPod, kube.Container{
			Type: containerType,
			Name: containerName,
		})
		podLogOpts = &corev1.PodLogOptions{
			Container: containerName,
			Follow:    follow,
		}
	)

	return opk8s.GetPodLogs(ctx, opts.Cli, types.WalrusSystemNamespace, jobPod.Name, opts.Out, podLogOpts)
}
