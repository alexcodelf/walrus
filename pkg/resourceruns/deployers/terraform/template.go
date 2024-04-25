package terraform

import (
	"fmt"

	core "k8s.io/api/core/v1"
	"k8s.io/utils/pointer"

	wf "github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/resourceruns/deployer"
)

const (
	// _podName the name of the pod.
	_podName = "terraform"

	// _secretMountPath the path to mount the secret.
	_secretMountPath = "/var/terraform/secrets"
	// _workdir the working directory of the job.
	_workdir = "/var/terraform/workspace"

	// _accessTokenkey the key of token in secret.
	_accessTokenkey = "access-token"
)

// generateTemplate generate a argo workflow template for deployment.
func generateTemplate(run *walruscore.ResourceRun, opts deployer.CreateTemplateOptions) (*wf.Template, error) {
	var (
		command       = []string{"/bin/sh", "-c"}
		deployCommand = fmt.Sprintf("cp %s/main.tf main.tf &&"+
			" export TF_HTTP_PASSWORD=$(cat /var/run/secrets/kubernetes.io/serviceaccount/token) && ", _secretMountPath)
	)

	if opts.ResourceRunStepType == walruscore.ResourceRunStepTypePlan {
		deployCommand += getPlanCommands(run, opts)
	} else if opts.ResourceRunStepType == walruscore.ResourceRunStepTypeApply {
		switch run.Spec.Type {
		case walruscore.ResourceRunTypeCreate,
			walruscore.ResourceRunTypeUpdate,
			walruscore.ResourceRunTypeRollback,
			walruscore.ResourceRunTypeStart:

			deployCommand += getApplyCommands(run, opts)
		case walruscore.ResourceRunTypeStop,
			walruscore.ResourceRunTypeDelete:

			deployCommand += getDestroyCommands(run, opts)
		}
	}

	command = append(command, deployCommand)

	volumeMounts := []core.VolumeMount{
		{
			Name:      run.Status.ConfigSecretName,
			MountPath: _secretMountPath,
			ReadOnly:  false,
		},
	}

	volumes := []core.Volume{
		{
			Name: run.Status.ConfigSecretName,
			VolumeSource: core.VolumeSource{
				Secret: &core.SecretVolumeSource{
					SecretName: run.Status.ConfigSecretName,
				},
			},
		},
	}

	securityContext := &core.PodSecurityContext{}

	if opts.DockerMode {
		volumeMounts = append(volumeMounts, core.VolumeMount{
			Name:      "docker-sock",
			MountPath: "/var/run/docker.sock",
		})

		volumes = append(volumes, core.Volume{
			Name: "docker-sock",
			VolumeSource: core.VolumeSource{
				HostPath: &core.HostPathVolumeSource{
					Path: "/var/run/docker.sock",
				},
			},
		})
		securityContext.RunAsUser = pointer.Int64(0)
	}

	return &wf.Template{
		Name: _podName,
		Metadata: wf.Metadata{
			Annotations: map[string]string{
				deployer.LabelWalrusResourceRunName:     run.Name,
				deployer.LabelWalrusResourceRunStepType: opts.ResourceRunStepType.String(),
			},
		},
		Container: &core.Container{
			Name:            "deployment",
			Image:           opts.Image,
			WorkingDir:      _workdir,
			Command:         command,
			ImagePullPolicy: core.PullIfNotPresent,
			VolumeMounts:    volumeMounts,
			Env:             opts.Env,
		},
		Volumes:         volumes,
		SecurityContext: securityContext,
	}, nil
}
