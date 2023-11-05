package step

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
	apiv1 "k8s.io/api/core/v1"

	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/model/service"
	"github.com/seal-io/walrus/pkg/dao/types/object"
	"github.com/seal-io/walrus/pkg/settings"
	"github.com/seal-io/walrus/pkg/workflow/step/types"
)

//nolint:lll
const stepSource = `#!/bin/sh
set -e
set -o pipefail

serverURL="{{workflow.parameters.server}}"
projectID="{{workflow.parameters.projectID}}"
environmentID="{{inputs.parameters.environmentID}}"
token="{{inputs.parameters.token}}"
jobType="{{inputs.parameters.jobType}}"
spec='{{inputs.parameters.spec}}'
serviceName="{{inputs.parameters.serviceName}}"
commomPath="$serverURL/v1/projects/$projectID/environments/$environmentID"

# if skip tls verify
tlsVerify="-k"
if [ "{{workflow.parameters.tlsVerify}}" == "true" ]; then
	tlsVerify=""
fi

# If jobType create service.
if [ "$jobType" == "create" ]; then
	response=$(curl -s "$commomPath/services" -X "POST" -H "content-type: application/json" -H "Authorization: Bearer $token" -d $spec $tlsVerify)

	serviceName=$(echo $response | jq -r '.name')
	if [ "$serviceName" == "null" ]; then
		echo "service name is null"
		echo "create response: $response"
		exit 1
	fi
fi

# If jobType upgrade service.
if [ "$jobType" == "upgrade" ]; then
	response=$(curl -s $commomPath/services/$serviceName/upgrade -X "PUT" -H "content-type: application/json" -H "Authorization: Bearer $token" -d $spec $tlsVerify)
fi

# Get latest revision id
revisionResponse=$(curl -s "$commomPath/services/$serviceName/revisions?page=1&perPage=1&sort=-createTime" -X GET -H "Authorization: Bearer $token" $tlsVerify)
revisionID=$(echo $revisionResponse | jq -r '.items[0].id')

# Watch service logs until the service finished.
curl -o - -s "$commomPath/services/$serviceName/revisions/$revisionID/log?jobType=$watchType&watch=true" -X GET -H "Authorization: Bearer $token" $tlsVerify --compressed
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

func (s *ServiceStepManager) GenerateTemplates(
	ctx context.Context,
	stepExecution *model.WorkflowStepExecution,
) (main *v1alpha1.Template, subs []*v1alpha1.Template, err error) {
	deployerImage := settings.WorkflowStepServiceImage.ShouldValue(ctx, s.mc)

	environment, ok := stepExecution.Spec["environment"].(map[string]any)
	if !ok {
		return nil, nil, errors.New("environment is not found")
	}

	environmentID, ok := environment["id"].(string)
	if !ok {
		return nil, nil, errors.New("environment id is not found")
	}

	serviceName, ok := stepExecution.Spec["name"].(string)
	if !ok {
		return nil, nil, errors.New("service name is not found")
	}

	// If service exist in environment, job type is upgrade.
	// Otherwise, job type is create.
	svc, err := s.mc.Services().Query().
		Select(
			service.FieldID,
			service.FieldName,
			service.FieldEnvironmentID,
		).
		Where(
			service.EnvironmentID(object.ID(environmentID)),
			service.Name(serviceName),
		).
		Only(ctx)
	if err != nil && !model.IsNotFound(err) {
		return nil, nil, fmt.Errorf("failed to get service: %w", err)
	}

	jobType := "create"
	if svc != nil {
		jobType = "upgrade"
	}

	// Inject workflow step execution id to request.
	stepSpec := stepExecution.Spec
	stepSpec["workflowStepExecutionID"] = stepExecution.ID.String()

	spec, err := json.Marshal(stepExecution.Spec)
	if err != nil {
		return nil, nil, err
	}

	// An service type workflow template
	// Interact with walrus server to create or update service.
	// Watch service logs until the service finished.
	main = &v1alpha1.Template{
		Name: fmt.Sprintf("step-execution-%s", stepExecution.ID.String()),
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
					Name:  "spec",
					Value: v1alpha1.AnyStringPtr(string(spec)),
				},
				{
					Name:  "jobType",
					Value: v1alpha1.AnyStringPtr(jobType),
				},
				{
					Name: "token",
				},
				{
					Name:  "serviceName",
					Value: v1alpha1.AnyStringPtr(serviceName),
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
		main.RetryStrategy = stepExecution.RetryStrategy
	}

	if stepExecution.Timeout > 0 {
		main.Timeout = strconv.Itoa(stepExecution.Timeout)
	}

	return main, nil, nil
}
