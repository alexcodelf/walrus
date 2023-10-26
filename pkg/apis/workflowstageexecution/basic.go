package workflowstageexecution

import (
	"github.com/seal-io/walrus/pkg/dao/types/status"
	"github.com/seal-io/walrus/pkg/workflow"
)

func (h Handler) Update(req UpdateRequest) error {
	entity := req.Model()

	switch req.Status {
	case "Succeeded":
		status.WorkflowStageExecutionStatusRunning.True(entity, "")
		status.WorkflowStageExecutionStatusReady.True(entity, "")
	case "Error", "Failed":
		status.WorkflowStageExecutionStatusRunning.False(entity, "execute failed")
	case "Running":
		status.WorkflowStageExecutionStatusRunning.Unknown(entity, "")
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
