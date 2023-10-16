package workflow

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
	"github.com/seal-io/walrus/pkg/auths/session"
	"github.com/seal-io/walrus/pkg/dao/model"
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

func Apply(ctx context.Context, mc model.ClientSet, clientConfig clientcmd.ClientConfig, wf *model.Workflow) error {
	client, err := NewArgoWorkflowClient(mc, clientConfig)
	if err != nil {
		return err
	}

	s := session.MustGetSubject(ctx)

	wfe, err := CreateWorkflowExecution(ctx, mc, wf)
	if err != nil {
		return err
	}

	return client.Submit(ctx, SubmitOptions{
		WorkflowExecution: wfe,
		SubjectID:         s.ID,
	})
}

func getExitTemplate(mc model.ClientSet, wf *model.WorkflowExecution) *v1alpha1.Template {
	return &v1alpha1.Template{
		Name: "notify",
		HTTP: &v1alpha1.HTTP{
			//nolint:lll
			URL:    "{{workflow.parameters.server}}" + "v1/projects/" + wf.ProjectID.String() + "workflow-executions/" + wf.ID.String() + "/done",
			Method: http.MethodPost,
			Headers: v1alpha1.HTTPHeaders{
				{
					Name:  "Content-Type",
					Value: "application/json",
				},
				{
					Name:  "Authorization",
					Value: "Bearer {{workflow.parameters.token}}",
				},
			},
			Body: `{"project":{"id": "{{workflow.parameters.projectID}}"}}`,
		},
	}
}

func CreateWorkflowExecution(
	ctx context.Context,
	mc model.ClientSet,
	wf *model.Workflow,
) (*model.WorkflowExecution, error) {
	s := session.MustGetSubject(ctx)

	// TODO check if the workflow is already running.

	// create workflow execution
	progress := fmt.Sprintf("%d/%d", 0, len(wf.Edges.Stages))

	workflowExecution := &model.WorkflowExecution{
		Name:       fmt.Sprintf("%s-%d", wf.Name, time.Now().Unix()),
		ProjectID:  wf.ProjectID,
		WorkflowID: wf.ID,
		Progress:   progress,
		SubjectID:  s.ID,
	}

	entity, err := mc.WorkflowExecutions().Create().
		Set(workflowExecution).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	stageExecutions := make(model.WorkflowStageExecutions, len(wf.Edges.Stages))

	for i, stage := range wf.Edges.Stages {
		// create workflow stage execution
		stageExecution, err := CreateWorkflowStageExecution(ctx, mc, stage, entity)
		if err != nil {
			return nil, err
		}

		stageExecutions[i] = stageExecution
	}

	entity.Edges.WorkflowStageExecutions = stageExecutions

	return entity, nil
}

func CreateWorkflowStageExecution(
	ctx context.Context,
	mc model.ClientSet,
	stage *model.WorkflowStage,
	we *model.WorkflowExecution,
) (*model.WorkflowStageExecution, error) {
	stageExec := &model.WorkflowStageExecution{
		Name:                fmt.Sprintf("%s-%d", stage.Name, time.Now().Unix()),
		ProjectID:           stage.ProjectID,
		StageID:             stage.ID,
		WorkflowExecutionID: we.ID,
	}

	entity, err := mc.WorkflowStageExecutions().Create().
		Set(stageExec).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	stepExecutions := make(model.WorkflowStepExecutions, len(stage.Edges.Steps))

	for i, step := range stage.Edges.Steps {
		// create workflow step execution
		stepExecution, err := CreateWorkflowStepExecution(ctx, mc, step, entity)
		if err != nil {
			return nil, err
		}

		stepExecutions[i] = stepExecution
	}

	entity.Edges.StepExecutions = stepExecutions

	return entity, nil
}

func CreateWorkflowStepExecution(
	ctx context.Context,
	mc model.ClientSet,
	step *model.WorkflowStep,
	wse *model.WorkflowStageExecution,
) (*model.WorkflowStepExecution, error) {
	stepExec := &model.WorkflowStepExecution{
		Name:                     fmt.Sprintf("%s-%d", step.Name, time.Now().Unix()),
		ProjectID:                step.ProjectID,
		WorkflowID:               step.WorkflowID,
		Type:                     step.Type,
		WorkflowStepID:           step.ID,
		WorkflowExecutionID:      wse.WorkflowExecutionID,
		WorkflowStageExecutionID: wse.ID,
		Spec:                     step.Spec,
	}

	return mc.WorkflowStepExecutions().Create().
		Set(stepExec).
		Save(ctx)
}
