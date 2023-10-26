package workflow

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqljson"

	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/model/workflowstageexecution"
	"github.com/seal-io/walrus/pkg/dao/model/workflowstepexecution"
	"github.com/seal-io/walrus/pkg/dao/types/status"
)

const (
	summaryStatusPending = "Pending"
	summaryStatusRunning = "Running"
)

type StatusManager struct {
	mc model.ClientSet
}

func NewStatusManager(mc model.ClientSet) *StatusManager {
	return &StatusManager{
		mc: mc,
	}
}

// HandleWorkflowExecutionFailed When workflowExecution is failed,
// mark all pending stages and steps status as failed.
func (m *StatusManager) HandleWorkflowExecutionFailed(
	ctx context.Context,
	workflowExecution *model.WorkflowExecution,
) error {
	if !status.WorkflowExecutionStatusRunning.IsFalse(workflowExecution) &&
		!status.WorkflowExecutionStatusReady.IsFalse(workflowExecution) {
		return nil
	}

	stages, err := m.mc.WorkflowStageExecutions().Query().
		Select(
			workflowstageexecution.FieldID,
			workflowstageexecution.FieldStatus,
			workflowstageexecution.FieldWorkflowExecutionID,
		).
		Where(
			workflowstageexecution.WorkflowExecutionID(workflowExecution.ID),
			func(s *sql.Selector) {
				s.Where(sqljson.ValueIn(
					workflowstageexecution.FieldStatus,
					[]any{
						summaryStatusRunning,
						summaryStatusPending,
					},
					sqljson.Path("summaryStatus"),
				))
			},
			func(s *sql.Selector) {
				s.Where(sqljson.ValueNEQ(
					workflowstageexecution.FieldStatus,
					false,
					sqljson.Path("error"),
				))
			},
		).
		All(ctx)
	if err != nil {
		return err
	}

	for i := range stages {
		stage := stages[i]

		err = m.HandleWorkflowStageExecutionFailed(ctx, stage)
		if err != nil {
			return err
		}

		status.WorkflowStageExecutionStatusRunning.False(stage, "")
		stage.Status.SetSummary(status.WalkWorkflowStageExecution(&stage.Status))

		// TODO (alex) update status in batch.
		err = m.mc.WorkflowStageExecutions().UpdateOne(stage).
			SetStatus(stage.Status).
			Exec(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

// HandleWorkflowStageExecutionFailed When workflowStageExecution is failed,
// mark all pending and running steps status as failed.
func (m *StatusManager) HandleWorkflowStageExecutionFailed(
	ctx context.Context,
	workflowStageExecution *model.WorkflowStageExecution,
) error {
	if !status.WorkflowStageExecutionStatusRunning.IsFalse(workflowStageExecution) &&
		!status.WorkflowExecutionStatusPending.IsFalse(workflowStageExecution) {
		return nil
	}

	steps, err := m.mc.WorkflowStepExecutions().Query().
		Select(
			workflowstageexecution.FieldID,
			workflowstageexecution.FieldStatus,
			workflowstageexecution.FieldWorkflowExecutionID,
		).
		Where(
			workflowstepexecution.WorkflowStageExecutionID(workflowStageExecution.ID),
			func(s *sql.Selector) {
				s.Where(sqljson.ValueIn(
					workflowstepexecution.FieldStatus,
					[]any{
						summaryStatusRunning,
						summaryStatusPending,
					},
					sqljson.Path("summaryStatus"),
				))
			},
			func(s *sql.Selector) {
				s.Where(sqljson.ValueNEQ(
					workflowstepexecution.FieldStatus,
					false,
					sqljson.Path("error"),
				))
			},
		).All(ctx)
	if err != nil {
		return err
	}

	for i := range steps {
		step := steps[i]
		status.WorkflowStepExecutionStatusRunning.False(step, "")
		step.Status.SetSummary(status.WalkWorkflowStepExecution(&step.Status))

		// TODO (alex) update status in batch.
		err = m.mc.WorkflowStepExecutions().UpdateOne(step).
			SetStatus(step.Status).
			Exec(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}
