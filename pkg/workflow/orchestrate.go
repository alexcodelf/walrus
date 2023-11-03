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

type (
	ExecuteOptions struct {
		ModelClient model.ClientSet
		Params      map[string]string
		// Description is the description of the workflow execution.
		Description string
	}
	CreateWorkflowExecutionOptions struct {
		ExecuteOptions `json:",inline"`

		Workflow *model.Workflow
	}

	CreateWorkflowStageExecutionOptions struct {
		ExecuteOptions `json:",inline"`

		WorkflowExecution *model.WorkflowExecution
		Stage             *model.WorkflowStage
	}

	CreateWorkflowStepExecutionOptions struct {
		ExecuteOptions `json:",inline"`

		StageExecution *model.WorkflowStageExecution
		Step           *model.WorkflowStep
	}
)

func CreateWorkflowExecution(
	ctx context.Context,
	opts CreateWorkflowExecutionOptions,
) (*model.WorkflowExecution, error) {
	s := session.MustGetSubject(ctx)

	var trigger types.WorkflowExecutionTrigger

	wf := opts.Workflow

	switch wf.Type {
	case types.WorkflowTypeDefault:
		userSubject, err := opts.ModelClient.Subjects().Query().
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
		Name:        wf.Name,
		Description: opts.Description,
		Type:        wf.Type,
		ProjectID:   wf.ProjectID,
		WorkflowID:  wf.ID,
		SubjectID:   s.ID,
		Parallelism: wf.Parallelism,
		// When creating a workflow execution, the execute times is always 1.
		Times:   1,
		Timeout: wf.Timeout,
		Trigger: trigger,
		Version: wf.Version + 1,
	}

	status.WorkflowExecutionStatusPending.Unknown(workflowExecution, "")
	workflowExecution.Status.SetSummary(status.WalkWorkflowExecution(&workflowExecution.Status))

	stageMap := make(map[object.ID]*model.WorkflowStage)
	for i := range wf.Edges.Stages {
		stageMap[wf.Edges.Stages[i].ID] = wf.Edges.Stages[i]
	}

	entity, err := opts.ModelClient.WorkflowExecutions().Create().
		Set(workflowExecution).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	stageExecutions := make(model.WorkflowStageExecutions, len(wf.Edges.Stages))

	for i, stage := range wf.Edges.Stages {
		// Create workflow stage execution.
		stageExecution, err := CreateWorkflowStageExecution(ctx, CreateWorkflowStageExecutionOptions{
			ExecuteOptions:    opts.ExecuteOptions,
			WorkflowExecution: entity,
			Stage:             stage,
		})
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
	opts CreateWorkflowStageExecutionOptions,
) (*model.WorkflowStageExecution, error) {
	stage := opts.Stage
	stageExec := &model.WorkflowStageExecution{
		Name:                stage.Name,
		ProjectID:           stage.ProjectID,
		WorkflowID:          stage.WorkflowID,
		WorkflowStageID:     stage.ID,
		WorkflowExecutionID: opts.WorkflowExecution.ID,
		Order:               stage.Order,
	}

	status.WorkflowStageExecutionStatusPending.Unknown(stageExec, "")
	stageExec.Status.SetSummary(status.WalkWorkflowStageExecution(&stageExec.Status))

	entity, err := opts.ModelClient.WorkflowStageExecutions().Create().
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
		stepExecution, err := CreateWorkflowStepExecution(ctx, CreateWorkflowStepExecutionOptions{
			ExecuteOptions: opts.ExecuteOptions,
			StageExecution: entity,
			Step:           step,
		})
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
	opts CreateWorkflowStepExecutionOptions,
) (*model.WorkflowStepExecution, error) {
	var (
		step   = opts.Step
		wse    = opts.StageExecution
		params = opts.Params
	)

	spec, err := parseParams(step.Spec, params)
	if err != nil {
		return nil, err
	}

	stepExec := &model.WorkflowStepExecution{
		Name:                     step.Name,
		Type:                     step.Type,
		Order:                    step.Order,
		ProjectID:                step.ProjectID,
		WorkflowID:               step.WorkflowID,
		WorkflowStepID:           step.ID,
		WorkflowExecutionID:      wse.WorkflowExecutionID,
		WorkflowStageExecutionID: wse.ID,
		Spec:                     spec,
		Times:                    1,
		Timeout:                  step.Timeout,
		RetryStrategy:            step.RetryStrategy,
	}

	status.WorkflowStepExecutionStatusPending.Unknown(stepExec, "")
	stepExec.Status.SetSummary(status.WalkWorkflowStepExecution(&stepExec.Status))

	return opts.ModelClient.WorkflowStepExecutions().Create().
		Set(stepExec).
		Save(ctx)
}
