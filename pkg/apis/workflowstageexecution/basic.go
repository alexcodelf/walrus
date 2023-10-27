package workflowstageexecution

import (
	"fmt"
	"time"

	"github.com/seal-io/walrus/pkg/dao/model/workflowstageexecution"
	"github.com/seal-io/walrus/pkg/dao/types"
	"github.com/seal-io/walrus/pkg/dao/types/object"
	"github.com/seal-io/walrus/pkg/dao/types/status"
	"github.com/seal-io/walrus/pkg/datalisten/modelchange"
	"github.com/seal-io/walrus/utils/topic"
)

func (h Handler) Update(req UpdateRequest) error {
	entity, err := h.modelClient.WorkflowStageExecutions().Query().
		Where(workflowstageexecution.ID(req.ID)).
		Only(req.Context)
	if err != nil {
		return err
	}

	fmt.Println("stage execution update", entity.ID, req.Status)

	switch req.Status {
	case types.ExecutionStatusSucceeded:
		status.WorkflowStageExecutionStatusRunning.True(entity, "")
		status.WorkflowStageExecutionStatusReady.True(entity, "")
	case types.ExecutionStatusFailed, types.ExecutionStatusError:
		status.WorkflowStageExecutionStatusRunning.False(entity, "execute failed")
	case types.ExecutionStatusRunning:
		status.WorkflowExecutionStatusPending.True(entity, "")
		status.WorkflowStageExecutionStatusRunning.Unknown(entity, "")
	default:
		return nil
	}

	entity.Status.SetSummary(status.WalkWorkflowStageExecution(&entity.Status))

	update := h.modelClient.WorkflowStageExecutions().UpdateOne(entity).
		SetStatus(entity.Status)

	if req.Status == types.ExecutionStatusSucceeded {
		update.SetDuration(int(time.Since(*entity.CreateTime).Seconds()))
	}

	entity, err = update.Save(req.Context)
	if err != nil {
		return err
	}

	// Publish workflow execution topic,
	// stage execution update will trigger workflow execution update.
	err = topic.Publish(req.Context, modelchange.WorkflowExecution, modelchange.Event{
		Type: modelchange.EventTypeUpdate,
		IDs:  []object.ID{entity.WorkflowExecutionID},
	})
	if err != nil {
		return err
	}

	return topic.Publish(req.Context, modelchange.Workflow, modelchange.Event{
		Type: modelchange.EventTypeUpdate,
		IDs:  []object.ID{entity.WorkflowID},
	})
}
