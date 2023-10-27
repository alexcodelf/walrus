package workflow

import (
	"context"

	"github.com/seal-io/walrus/pkg/auths/session"
	"github.com/seal-io/walrus/pkg/dao/model"
	"k8s.io/client-go/rest"
)

// Apply applies the workflow execution to the argo workflow server.
func Apply(
	ctx context.Context,
	mc model.ClientSet,
	restCfg *rest.Config,
	wf *model.Workflow,
) (*model.WorkflowExecution, error) {
	client, err := NewArgoWorkflowClient(mc, restCfg)
	if err != nil {
		return nil, err
	}

	s := session.MustGetSubject(ctx)

	wfe, err := CreateWorkflowExecution(ctx, mc, wf)
	if err != nil {
		return nil, err
	}

	return wfe, client.Submit(ctx, SubmitOptions{
		WorkflowExecution: wfe,
		SubjectID:         s.ID,
	})
}

// Resubmit resubmits the workflow execution to the argo workflow server.
func Resubmit(
	ctx context.Context,
	mc model.ClientSet,
	restCfg *rest.Config,
	wfe *model.WorkflowExecution,
) error {
	client, err := NewArgoWorkflowClient(mc, restCfg)
	if err != nil {
		return err
	}

	return client.Resubmit(ctx, ResubmitOptions{
		WorkflowExecution: wfe,
	})
}
