package workflowstageexecution

import (
	"github.com/seal-io/walrus/pkg/dao/types/status"
	"github.com/seal-io/walrus/pkg/workflow"
)

func (h Handler) Update(req UpdateRequest) error {
	entity := req.Model()

	switch req.Status {
	case "Succeeded":
		status.WorkflowStageExecutionStatusReady.Reset(entity, "")
		status.WorkflowStageExecutionStatusReady.True(entity, "")
	case "Error", "Failed":
		status.WorkflowStageExecutionStatusPending.Reset(entity, "")
		status.WorkflowStageExecutionStatusPending.False(entity, "execute failed")
	case "Running":
		status.WorkflowStageExecutionStatusPending.Reset(entity, "")
	}

	entity.Status.SetSummary(status.WalkWorkflowStageExecution(&entity.Status))

	statusManager := workflow.NewStatusManager(h.modelClient)
	if err := statusManager.HandleWorkflowStageExecutionFailed(req.Context, entity); err != nil {
		return err
	}

	return h.modelClient.WorkflowStageExecutions().UpdateOne(entity).
		SetRecord(req.Record).
		SetStatus(entity.Status).
		Exec(req.Context)
}
