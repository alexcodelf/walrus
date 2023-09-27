package workflow

import (
	"context"

	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/model/workflowstage"
	"github.com/seal-io/walrus/utils/log"
)

type WorkflowExecutor interface {
	Execute(context.Context, *model.Workflow) error
}

// ArgoWorkflowExecutor is the executor for Argo workflow.
type ArgoWorkflowExecutor struct {
	WorkflowExecutor

	logger log.Logger
	mc     *model.Client
}

func NewArgoWorkflowExecutor(mc *model.Client) *ArgoWorkflowExecutor {
	return &ArgoWorkflowExecutor{
		logger: log.WithName("workflow-executor"),
		mc:     mc,
	}
}

func (e *ArgoWorkflowExecutor) Execute(ctx context.Context, wf *model.Workflow) error {
	// 1. Get workflow stages.
	stages, err := e.mc.WorkflowStages().Query().
		Where(workflowstage.WorkflowID(wf.ID)).
		All(ctx)
	if err != nil {
		return err
	}

	e.logger.Info("workflow stages", "stages", stages)

	// 2. Get workflow steps.
	// 3. Get workflow.

	// Transition.
	// 1. Create workflow Execution.
	// 2. Create workflow stages Execution.
	// 3. Create workflow steps Execution.

	panic("not implemented")
}
