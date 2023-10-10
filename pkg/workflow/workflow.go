package workflow

import (
	"context"

	"github.com/seal-io/walrus/pkg/dao/model"
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

func CreateStage(ctx context.Context, mc model.ClientSet, st *model.WorkflowStage) error {
	return mc.WithTx(ctx, func(tx *model.Tx) error {
		stage, err := tx.WorkflowStages().Create().
			Set(st).
			Save(ctx)
		if err != nil {
			return err
		}

		// Save steps.
		steps := make(model.WorkflowSteps, 0, len(st.Edges.Steps))

		for _, step := range st.Edges.Steps {
			step.StageID = stage.ID

			s, err := tx.WorkflowSteps().Create().
				Set(step).
				Save(ctx)
			if err != nil {
				return err
			}

			steps = append(steps, s)
		}

		// Update stage.
		stage.Edges.Steps = steps

		return nil
	})
}

// TODO Update workflow
// TODO Update stage.
