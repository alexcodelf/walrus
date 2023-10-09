package step

import (
	"context"

	"github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/types"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ConfigService is service to generate service configs.
type ConfigService struct {
	mc model.ClientSet
}

// NewConfigService.
func NewConfigService(mc model.ClientSet) *ConfigService {
	return &ConfigService{
		mc: mc,
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

var testSvcDemo = &model.WorkflowStep{
	Name:        "test-svc-demo",
	Description: "test-svc-demo",
	Type:        types.WorkflowStepTypeService,
	Spec: map[string]any{
		"name":            "bac",
		"attributes":      `{"env": {}, "name": "", "image": "nginx", "ports": [80], "replicas": 1, "limit_cpu": "", "namespace": "", "request_cpu": "0.1", "limit_memory": "", "request_memory": "128Mi"}`,
		"templateName":    "webservice",
		"templateVersion": "v0.0.3",
		"deployerType":    "terraform",
	},
}

func (s *ConfigService) GenerateServiceTemplate(
	ctx context.Context, step *model.WorkflowStep,
) (*v1alpha1.Template, error) {
	t := &v1alpha1.Template{
		Name: step.Name,
		Container: &apiv1.Container{
			ImagePullPolicy: apiv1.PullIfNotPresent,
			Image:           "sealio/terraform-deployer:v0.1.4",
		},
		Script: &v1alpha1.ScriptTemplate{
			Source: `#!/bin/sh
set -e
terraform init -no-color && terraform apply -auto-approve -no-color
`,
		},
	}

	return t, nil
}
