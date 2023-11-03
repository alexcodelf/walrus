package step

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"

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
-H "Authorization: Bearer {{inputs.parameters.token}}" \
-d '{{inputs.parameters.executionSpec}}' $tlsVerify -s --fail --show-error

if tar -xzf config.tar.gz; then
	# run terraform
	terraform {{inputs.parameters.tfCommand}}
else
    echo "Extraction Config failed"
    cat config.tar.gz
fi
`

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
	stepExecution *model.WorkflowStepExecution,
) (*v1alpha1.Template, error) {
	deployerImage := settings.DeployerImage.ShouldValue(ctx, s.mc)

	environment, ok := stepExecution.Spec["environment"].(map[string]any)
	if !ok {
		return nil, errors.New("environment is not found")
	}

	environmentID, ok := environment["id"].(string)
	if !ok {
		return nil, errors.New("environment id is not found")
	}

	// Inject workflow step execution id to request.
	stepSpec := stepExecution.Spec
	stepSpec["workflowStepExecutionID"] = stepExecution.ID.String()

	execSpec, err := json.Marshal(stepExecution.Spec)
	if err != nil {
		return nil, err
	}

	st := &v1alpha1.Template{
		Name: "step-execution-" + stepExecution.ID.String(),
		Metadata: v1alpha1.Metadata{
			Labels: map[string]string{
				"step-execution-id": stepExecution.ID.String(),
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
					Value: v1alpha1.AnyStringPtr(stepExecution.WorkflowID.String()),
				},
				{
					Name:  "workflowStepID",
					Value: v1alpha1.AnyStringPtr(stepExecution.WorkflowStepID.String()),
				},
				{
					Name:  "executionSpec",
					Value: v1alpha1.AnyStringPtr(string(execSpec)),
				},
				{
					Name:  "tfCommand",
					Value: v1alpha1.AnyStringPtr("init -no-color && terraform apply -auto-approve -no-color"),
				},
				{
					Name: "token",
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

	if stepExecution.RetryStrategy != nil {
		st.RetryStrategy = stepExecution.RetryStrategy
	}

	if stepExecution.Timeout > 0 {
		st.Timeout = strconv.Itoa(stepExecution.Timeout)
	}

	return st, nil
}
