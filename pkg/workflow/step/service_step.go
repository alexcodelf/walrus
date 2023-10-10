package step

import (
	"context"
	"encoding/json"

	"github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/settings"
	"github.com/seal-io/walrus/pkg/workflow/step/types"

	apiconfig "github.com/seal-io/walrus/pkg/apis/config"
)

// ServiceStepManager is service to generate service configs.
type ServiceStepManager struct {
	mc model.ClientSet
}

// NewServiceStepManager.
func NewServiceStepManager(opts types.CreateOptions) types.StepManager {
	return &ServiceStepManager{
		mc: opts.ModelClient,
	}
}

var serviceParameter = v1alpha1.Parameter{
	Name:  "deployer-image",
	Value: v1alpha1.AnyStringPtr("sealio/terraform-deployer:v0.1.4"),
}

var svcTpl = &v1alpha1.WorkflowTemplate{
	ObjectMeta: metav1.ObjectMeta{
		Name: "svc-tpl",
	},
	Spec: v1alpha1.WorkflowSpec{
		Entrypoint: "main",
		Arguments: v1alpha1.Arguments{
			Parameters: []v1alpha1.Parameter{
				serviceParameter,
			},
		},
		Templates: []v1alpha1.Template{
			{
				Name: "main",
				DAG: &v1alpha1.DAGTemplate{
					Tasks: []v1alpha1.DAGTask{
						{
							Name:     "webservice",
							Template: "webservice",
						},
					},
				},
			},
			{
				Name: "webservice",
				Container: &apiv1.Container{
					Image: "{{inputs.parameters.deployer-image}",
				},
				Script: &v1alpha1.ScriptTemplate{
					Source: `#!/bin/sh
set -e
terraform init -no-color && terraform apply -auto-approve -no-color
`,
				},
			},
		},
	},
}

var testSvcDemo = &model.WorkflowStepExecution{
	Name:        "test-svc-demo",
	Description: "test-svc-demo",
	Type:        types.StepTypeService.String(),
	Spec: map[string]any{
		"name": "bac",
		"attributes": map[string]any{
			"env":            nil,
			"name":           "",
			"image":          "nginx",
			"ports":          []int{80},
			"replicas":       1,
			"limit_cpu":      "",
			"namespace":      "",
			"request_cpu":    "0.1",
			"limit_memory":   "",
			"request_memory": "128Mi",
		},
		"templateName":    "webservice",
		"templateVersion": "v0.0.3",
		"deployerType":    "terraform",
	},
}

func (s *ServiceStepManager) GenerateTemplate(
	ctx context.Context,
	stepExec *model.WorkflowStepExecution,
) (*v1alpha1.Template, error) {
	tlsVerify := apiconfig.TlsCertified.Get()

	deployerImage := settings.DeployerImage.ShouldValue(ctx, s.mc)

	execSpec, err := json.Marshal(stepExec.Spec)
	if err != nil {
		return nil, err
	}

	st := &v1alpha1.Template{
		Name: stepExec.Name, // TODO use stepExec.ID.String() as name.
		Inputs: v1alpha1.Inputs{
			Parameters: []v1alpha1.Parameter{
				{
					Name:  "tlsVerify",
					Value: v1alpha1.AnyStringPtr(tlsVerify),
				},
				{
					Name:  "projectID",
					Value: v1alpha1.AnyStringPtr(stepExec.ProjectID.String()),
				},
				{
					Name:  "workflowID",
					Value: v1alpha1.AnyStringPtr(stepExec.WorkflowID.String()),
				},
				{
					Name:  "workflowStepID",
					Value: v1alpha1.AnyStringPtr(stepExec.WorkflowStepID.String()),
				},
				{
					Name:  "executionSpec",
					Value: v1alpha1.AnyStringPtr(string(execSpec)),
				},
				{
					Name:  "tfCommand",
					Value: v1alpha1.AnyStringPtr("init -no-color && terraform apply -auto-approve -no-color"),
				},
			},
		},
		//nolint:lll
		Script: &v1alpha1.ScriptTemplate{
			Container: apiv1.Container{
				Image:           deployerImage,
				ImagePullPolicy: apiv1.PullIfNotPresent,
				Command:         []string{"sh"},
			},
			Source: `#!/bin/sh
set -e
set -o pipefail

# if skip tls verify
tlsVerify="-k"
if [ "{{inputs.parameters.tls-verify}}" == "false" ]; then
	tlsVerify=""
fi

# get config
curl -o config.tar.gz -X POST \
{{inputs.parameters.server}}/v1/projects/{{inputs.parameters.project-id}}/environments/{{inputs.parameters.environment-id}}/services/_/workflow \
-H 'Content-Type: application/json' \
-H "Authorization: Bearer {{inputs.parameters.token}}" \
-d '{{inputs.parameters.executionSpec}}' $tlsVerify -s

tar -xzf config.tar.gz

# run terraform
terraform {{inputs.parameters.tfCommand}} `,
		},
	}

	return st, nil
}
