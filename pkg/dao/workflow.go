package dao

import (
	"context"
	"errors"

	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/model/workflowstage"
	"github.com/seal-io/walrus/pkg/dao/model/workflowstep"
)

func WorkflowStagesEdgeSave(ctx context.Context, mc model.ClientSet, entity *model.Workflow) error {
	if len(entity.Edges.Stages) == 0 {
		return nil
	}

	// Delete stale items.
	_, err := mc.WorkflowStages().Delete().
		Where(workflowstage.WorkflowID(entity.ID)).
		Exec(ctx)
	if err != nil {
		return err
	}

	// Add new items.
	newItems := entity.Edges.Stages

	for i := range newItems {
		if newItems[i] == nil {
			return errors.New("invalid input: nil relationship")
		}
		newItems[i].WorkflowID = entity.ID
		newItems[i].ProjectID = entity.ProjectID
		newItems[i].Order = i

		newItems[i], err = mc.WorkflowStages().Create().
			Set(newItems[i]).
			SaveE(ctx, WorkflowStageStepsEdgeSave)
		if err != nil {
			return err
		}
	}

	entity.Edges.Stages = newItems // Feedback.

	// Save workflow stage IDs.
	err = mc.Workflows().UpdateOneID(entity.ID).
		Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

// WorkflowStageStepsEdgeSave saves the edge steps of model.WorkflowStage entity.
func WorkflowStageStepsEdgeSave(ctx context.Context, mc model.ClientSet, entity *model.WorkflowStage) error {
	if len(entity.Edges.Steps) == 0 {
		return nil
	}

	// Delete stale items.
	_, err := mc.WorkflowSteps().Delete().
		Where(workflowstep.WorkflowStageID(entity.ID)).
		Exec(ctx)
	if err != nil {
		return err
	}

	// Add new items.
	newItems := entity.Edges.Steps

	for i := range newItems {
		if newItems[i] == nil {
			return errors.New("invalid input: nil relationship")
		}
		newItems[i].WorkflowStageID = entity.ID
		newItems[i].ProjectID = entity.ProjectID
		newItems[i].WorkflowID = entity.WorkflowID
		newItems[i].Order = i
		// TODO avoid save in loop.
		newItems[i], err = mc.WorkflowSteps().Create().
			Set(newItems[i]).
			Save(ctx)
		if err != nil {
			return err
		}
	}

	entity.Edges.Steps = newItems // Feedback.

	return nil
}
