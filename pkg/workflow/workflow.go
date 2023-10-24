package workflow

import (
	"context"
	"fmt"

	"github.com/seal-io/walrus/pkg/auths/session"
	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/types/object"
	"github.com/seal-io/walrus/pkg/dao/types/status"
	"k8s.io/client-go/tools/clientcmd"
)

func Create(ctx context.Context, mc model.ClientSet, wf *model.Workflow) error {
	return mc.WithTx(ctx, func(tx *model.Tx) error {
		workflow, err := tx.Workflows().Create().
			Set(wf).
			Save(ctx)
		if err != nil {
			return err
		}

		stages := make(model.WorkflowStages, 0, len(wf.Edges.Stages))

		for _, stage := range wf.Edges.Stages {
			stage.WorkflowID = workflow.ID

			st, err := tx.WorkflowStages().Create().
				Set(stage).
				Save(ctx)
			if err != nil {
				return err
			}

			stages = append(stages, st)
		}

		// Update workflow stages.
		workflow.Edges.Stages = stages
		wf = workflow

		return nil
	})
}

// Apply applies the workflow execution to the argo workflow server.
func Apply(
	ctx context.Context,
	mc model.ClientSet,
	clientConfig clientcmd.ClientConfig,
	wf *model.Workflow,
) (*model.WorkflowExecution, error) {
	client, err := NewArgoWorkflowClient(mc, clientConfig)
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

func CreateWorkflowExecution(
	ctx context.Context,
	mc model.ClientSet,
	wf *model.Workflow,
) (*model.WorkflowExecution, error) {
	s := session.MustGetSubject(ctx)

	// Create workflow execution.
	progress := fmt.Sprintf("%d/%d", 0, len(wf.Edges.Stages))

	workflowExecution := &model.WorkflowExecution{
		Name:       wf.Name,
		ProjectID:  wf.ProjectID,
		WorkflowID: wf.ID,
		Progress:   progress,
		SubjectID:  s.ID,
	}

	status.WorkflowExecutionStatusPending.Unknown(workflowExecution, "")
	workflowExecution.Status.SetSummary(status.WalkWorkflowExecution(&workflowExecution.Status))

	entity, err := mc.WorkflowExecutions().Create().
		Set(workflowExecution).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	stageMap := make(map[object.ID]*model.WorkflowStage)
	for i := range wf.Edges.Stages {
		stageMap[wf.Edges.Stages[i].ID] = wf.Edges.Stages[i]
	}

	ordered := make(model.WorkflowStages, len(wf.Edges.Stages))
	for i := range wf.StageIds {
		ordered[i] = stageMap[wf.StageIds[i]]
	}

	wf.Edges.Stages = ordered

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
		StageID:             stage.ID,
		WorkflowExecutionID: we.ID,
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

	ordered := make(model.WorkflowSteps, len(stage.Edges.Steps))
	for i := range stage.StepIds {
		ordered[i] = stepMap[stage.StepIds[i]]
	}

	stage.Edges.Steps = ordered

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
		ProjectID:                step.ProjectID,
		WorkflowID:               step.WorkflowID,
		Type:                     step.Type,
		WorkflowStepID:           step.ID,
		WorkflowExecutionID:      wse.WorkflowExecutionID,
		WorkflowStageExecutionID: wse.ID,
		Spec:                     step.Spec,
	}

	status.WorkflowStepExecutionStatusPending.Unknown(stepExec, "")
	stepExec.Status.SetSummary(status.WalkWorkflowStepExecution(&stepExec.Status))

	return mc.WorkflowStepExecutions().Create().
		Set(stepExec).
		Save(ctx)
}
