package step

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
	apiv1 "k8s.io/api/core/v1"

	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/settings"
	"github.com/seal-io/walrus/pkg/workflow/step/types"
)

//nolint:lll
const stepSource = `#!/bin/sh
set -e
set -o pipefail

# if skip tls verify
tlsVerify="-k"
if [ "{{workflow.parameters.tlsVerify}}" == "true" ]; then
	tlsVerify=""
fi

# get config
curl -o config.tar.gz -X POST \
{{workflow.parameters.server}}/v1/projects/{{workflow.parameters.projectID}}/environments/{{inputs.parameters.environmentID}}/services/_/workflow \
-H 'Content-Type: application/json' \
-H "Authorization: Bearer {{workflow.parameters.token}}" \
-d '{{inputs.parameters.executionSpec}}' $tlsVerify -s

tar -xzf config.tar.gz

# run terraform
terraform {{inputs.parameters.tfCommand}}`

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

func (s *ServiceStepManager) GenerateTemplate(
	ctx context.Context,
	stepExec *model.WorkflowStepExecution,
) (*v1alpha1.Template, error) {
	deployerImage := settings.DeployerImage.ShouldValue(ctx, s.mc)

	environment, ok := stepExec.Spec["environment"].(map[string]interface{})
	if !ok {
		return nil, errors.New("environment is not found")
	}
	environmentID, ok := environment["id"].(string)
	if !ok {
		return nil, errors.New("environment id is not found")
	}

	// Inject workflow step execution id to request.
	stepSpec := stepExec.Spec
	stepSpec["workflowStepExecutionID"] = stepExec.ID.String()

	execSpec, err := json.Marshal(stepExec.Spec)
	if err != nil {
		return nil, err
	}

	st := &v1alpha1.Template{
		Name: "step-execution-" + stepExec.ID.String(),
		Metadata: v1alpha1.Metadata{
			Labels: map[string]string{
				"step-execution-id": stepExec.ID.String(),
			},
		},
		Inputs: v1alpha1.Inputs{
			Parameters: []v1alpha1.Parameter{
				{
					Name:  "environmentID",
					Value: v1alpha1.AnyStringPtr(environmentID),
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
		Script: &v1alpha1.ScriptTemplate{
			Container: apiv1.Container{
				Image:           deployerImage,
				ImagePullPolicy: apiv1.PullIfNotPresent,
				Command:         []string{"sh"},
			},
			Source: stepSource,
		},
	}

	return st, nil
}
