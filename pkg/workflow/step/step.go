package step

import (
	"context"
	"fmt"
	"net/http"

	"github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
	"k8s.io/apimachinery/pkg/util/intstr"

	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/settings"
	"github.com/seal-io/walrus/pkg/workflow/step/types"
)

const notificationPath = "/v1/projects/%s/workflows/%s/workflow-steps/%s/notification"

var managerCreators map[types.Type]func(types.CreateOptions) types.StepManager

func init() {
	managerCreators = map[types.Type]func(types.CreateOptions) types.StepManager{
		types.StepTypeService:  NewServiceStepManager,
		types.StepTypeApproval: NewApprovalStepManager,
	}
}

func GetStepManager(opts types.CreateOptions) (types.StepManager, error) {
	constructor, ok := managerCreators[opts.Type]
	if !ok {
		return nil, fmt.Errorf("unknown step type: %s", opts.Type)
	}

	return constructor(opts), nil
}

func GetStepNotificationStep(
	ctx context.Context,
	mc model.ClientSet,
	stepExec *model.WorkflowStepExecution,
) (*v1alpha1.Template, error) {
	limit := intstr.FromString("3")

	serverAddress, err := settings.ServeUrl.Value(ctx, mc)
	if err != nil {
		return nil, err
	}

	notifyAddr := fmt.Sprintf(serverAddress+notificationPath, stepExec.ProjectID, stepExec.WorkflowID, stepExec.ID)

	stepNotification := &v1alpha1.Template{
		Name:    "notification" + stepExec.ID.String(),
		Timeout: "30s",
		RetryStrategy: &v1alpha1.RetryStrategy{
			Limit: &limit,
		},
		Inputs: v1alpha1.Inputs{
			Parameters: []v1alpha1.Parameter{
				{
					Name:  "addr",
					Value: v1alpha1.AnyStringPtr(notifyAddr),
				},
				{
					Name:  "project-id",
					Value: v1alpha1.AnyStringPtr(stepExec.ProjectID.String()),
				},
				{
					Name: "workflow-id",
				},
			},
		},
		HTTP: &v1alpha1.HTTP{
			URL:    "{{inputs.parameters.addr}}",
			Method: http.MethodPost,
			Body: `{
				"projectID": "{{inputs.parameters.project-id}}",
				"project": {
					"id": "{{inputs.parameters.project-id}}"
				},
				"workflow": {
					"id": "{{inputs.parameters.workflow-id}}"
				},
			}`,
		},
	}

	return stepNotification, nil
}
