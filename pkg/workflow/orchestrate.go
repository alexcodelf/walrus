package workflow

import (
	"context"
	"fmt"

	"github.com/seal-io/walrus/pkg/auths/session"
	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/model/subject"
	"github.com/seal-io/walrus/pkg/dao/types"
	"github.com/seal-io/walrus/pkg/dao/types/object"
	"github.com/seal-io/walrus/pkg/dao/types/status"
)

func CreateWorkflowExecution(
	ctx context.Context,
	mc model.ClientSet,
	wf *model.Workflow,
) (*model.WorkflowExecution, error) {
	s := session.MustGetSubject(ctx)

	var trigger types.WorkflowExecutionTrigger

	switch wf.Type {
	case types.WorkflowTypeDefault:
		userSubject, err := mc.Subjects().Query().
			Where(subject.ID(s.ID)).
			Only(ctx)
		if err != nil {
			return nil, err
		}
		trigger = types.WorkflowExecutionTrigger{
			Type: types.WorkflowExecutionTriggerTypeManual,
			User: userSubject.Name,
		}
	default:
		return nil, fmt.Errorf("invalid workflow type: %s", wf.Type)
	}

	workflowExecution := &model.WorkflowExecution{
		Name:       wf.Name,
		ProjectID:  wf.ProjectID,
		WorkflowID: wf.ID,
		SubjectID:  s.ID,
		Trigger:    trigger,
		Version:    wf.Version + 1,
	}

	status.WorkflowExecutionStatusPending.Unknown(workflowExecution, "")
	workflowExecution.Status.SetSummary(status.WalkWorkflowExecution(&workflowExecution.Status))

	stageMap := make(map[object.ID]*model.WorkflowStage)
	for i := range wf.Edges.Stages {
		stageMap[wf.Edges.Stages[i].ID] = wf.Edges.Stages[i]
	}

	entity, err := mc.WorkflowExecutions().Create().
		Set(workflowExecution).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	stageExecutions := make(model.WorkflowStageExecutions, len(wf.Edges.Stages))

	for i, stage := range wf.Edges.Stages {
		// Create workflow stage execution.
		stageExecution, err := CreateWorkflowStageExecution(ctx, mc, stage, entity)
		if err != nil {
			return nil, err
		}

		stageExecutions[i] = stageExecution
	}

	entity.Edges.Stages = stageExecutions

	return entity, nil
}

func CreateWorkflowStageExecution(
	ctx context.Context,
	mc model.ClientSet,
	stage *model.WorkflowStage,
	we *model.WorkflowExecution,
) (*model.WorkflowStageExecution, error) {
	stageExec := &model.WorkflowStageExecution{
		Name:                stage.Name,
		ProjectID:           stage.ProjectID,
		WorkflowID:          stage.WorkflowID,
		WorkflowStageID:     stage.ID,
		WorkflowExecutionID: we.ID,
		Order:               stage.Order,
	}

	status.WorkflowStageStatusPending.Unknown(stageExec, "")
	stageExec.Status.SetSummary(status.WalkWorkflowStageExecution(&stageExec.Status))

	entity, err := mc.WorkflowStageExecutions().Create().
		Set(stageExec).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	stepMap := make(map[object.ID]*model.WorkflowStep)
	for i := range stage.Edges.Steps {
		stepMap[stage.Edges.Steps[i].ID] = stage.Edges.Steps[i]
	}

	stepExecutions := make(model.WorkflowStepExecutions, len(stage.Edges.Steps))

	for i, step := range stage.Edges.Steps {
		// Create workflow step execution.
		stepExecution, err := CreateWorkflowStepExecution(ctx, mc, step, entity)
		if err != nil {
			return nil, err
		}

		stepExecutions[i] = stepExecution
	}

	entity.Edges.Steps = stepExecutions

	return entity, nil
}

func CreateWorkflowStepExecution(
	ctx context.Context,
	mc model.ClientSet,
	step *model.WorkflowStep,
	wse *model.WorkflowStageExecution,
) (*model.WorkflowStepExecution, error) {
	stepExec := &model.WorkflowStepExecution{
		Name:                     step.Name,
		Type:                     step.Type,
		Order:                    step.Order,
		ProjectID:                step.ProjectID,
		WorkflowID:               step.WorkflowID,
		WorkflowStepID:           step.ID,
		WorkflowExecutionID:      wse.WorkflowExecutionID,
		WorkflowStageExecutionID: wse.ID,
		Spec:                     step.Spec,
		RetryStrategy:            step.RetryStrategy,
	}

	status.WorkflowStepExecutionStatusPending.Unknown(stepExec, "")
	stepExec.Status.SetSummary(status.WalkWorkflowStepExecution(&stepExec.Status))

	return mc.WorkflowStepExecutions().Create().
		Set(stepExec).
		Save(ctx)
}
