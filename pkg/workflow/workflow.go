package workflow

import (
	"context"

	"k8s.io/client-go/rest"

	"github.com/seal-io/walrus/pkg/auths/session"
	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/types/status"
	"github.com/seal-io/walrus/utils/log"
)

// Apply applies the workflow execution to the argo workflow server.
func Apply(
	ctx context.Context,
	restCfg *rest.Config,
	wf *model.Workflow,
	opts ExecuteOptions,
) (*model.WorkflowExecution, error) {
	logger := log.WithName("workflow")

	client, err := NewArgoWorkflowClient(opts.ModelClient, restCfg)
	if err != nil {
		return nil, err
	}

	s := session.MustGetSubject(ctx)

	wfe, err := CreateWorkflowExecution(ctx, CreateWorkflowExecutionOptions{
		ExecuteOptions: opts,
		Workflow:       wf,
	})
	if err != nil {
		return nil, err
	}

	err = client.Submit(ctx, SubmitOptions{
		WorkflowExecution: wfe,
		SubjectID:         s.ID,
	})

	if err != nil {
		{
			ctx := context.Background()

			status.WorkflowExecutionStatusPending.False(wfe, err.Error())

			err := opts.ModelClient.WorkflowExecutions().UpdateOne(wfe).
				SetStatus(wfe.Status).
				Exec(ctx)
			if err != nil {
				logger.Errorf("failed to update workflow execution status: %v", err)
			}
		}
	}

	return wfe, err
}

// Resubmit resubmits the workflow execution to the argo workflow server.
func Resubmit(
	ctx context.Context,
	mc model.ClientSet,
	restCfg *rest.Config,
	wfe *model.WorkflowExecution,
) error {
	logger := log.WithName("workflow")

	client, err := NewArgoWorkflowClient(mc, restCfg)
	if err != nil {
		return err
	}

	err = client.Resubmit(ctx, ResubmitOptions{
		WorkflowExecution: wfe,
	})

	if err != nil {
		{
			// If the workflow execution is not found, reset the status to pending.
			ctx := context.Background()

			status.WorkflowExecutionStatusPending.False(wfe, err.Error())

			err := mc.WorkflowExecutions().UpdateOne(wfe).
				SetStatus(wfe.Status).
				Exec(ctx)
			if err != nil {
				logger.Errorf("failed to update workflow execution status: %v", err)
			}
		}
	}

	return err
}
