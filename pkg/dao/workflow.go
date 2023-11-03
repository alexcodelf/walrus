package dao

import (
	"context"
	"errors"

	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/model/workflowstage"
	"github.com/seal-io/walrus/pkg/dao/model/workflowstep"
	"github.com/seal-io/walrus/pkg/dao/types/object"
)

func WorkflowStagesEdgeSave(ctx context.Context, mc model.ClientSet, entity *model.Workflow) error {
	if len(entity.Edges.Stages) == 0 {
		return nil
	}

	// StageIDs that should be kept.
	// These stage will be updated.
	stageIDs := make([]object.ID, 0, len(entity.Edges.Stages))

	for i := range entity.Edges.Stages {
		stage := entity.Edges.Stages[i]
		if !stage.ID.Valid() {
			continue
		}

		stageIDs = append(stageIDs, stage.ID)
	}

	// Delete stale items.
	_, err := mc.WorkflowStages().Delete().
		Where(
			workflowstage.WorkflowID(entity.ID),
			workflowstage.IDNotIn(stageIDs...)).
		Exec(ctx)
	if err != nil {
		return err
	}

	// Add new items or update existing items.
	newItems := entity.Edges.Stages

	for i := range newItems {
		if newItems[i] == nil {
			return errors.New("invalid input: nil workflow stage")
		}
		newItems[i].WorkflowID = entity.ID
		newItems[i].ProjectID = entity.ProjectID
		newItems[i].Order = i

		if newItems[i].ID.Valid() {
			newItems[i], err = mc.WorkflowStages().UpdateOne(newItems[i]).
				Set(newItems[i]).
				SaveE(ctx, WorkflowStageStepsEdgeSave)
		} else {
			newItems[i], err = mc.WorkflowStages().Create().
				Set(newItems[i]).
				SaveE(ctx, WorkflowStageStepsEdgeSave)
		}

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

	// StepIDs that should be kept.
	// These steps will be updated.
	stepIDs := make([]object.ID, 0, len(entity.Edges.Steps))

	for i := range entity.Edges.Steps {
		step := entity.Edges.Steps[i]
		if !step.ID.Valid() {
			continue
		}

		stepIDs = append(stepIDs, step.ID)
	}

	// Delete stale items.
	_, err := mc.WorkflowSteps().Delete().
		Where(
			workflowstep.WorkflowStageID(entity.ID),
			workflowstep.IDNotIn(stepIDs...),
		).
		Exec(ctx)
	if err != nil {
		return err
	}

	// Add new items or update existing items.
	newItems := entity.Edges.Steps

	for i := range newItems {
		if newItems[i] == nil {
			return errors.New("invalid input: nil workflow step")
		}
		newItems[i].WorkflowStageID = entity.ID
		newItems[i].ProjectID = entity.ProjectID
		newItems[i].WorkflowID = entity.WorkflowID
		newItems[i].Order = i

		if newItems[i].ID.Valid() {
			newItems[i], err = mc.WorkflowSteps().UpdateOne(newItems[i]).
				Set(newItems[i]).
				Save(ctx)
		} else {
			newItems[i], err = mc.WorkflowSteps().Create().
				Set(newItems[i]).
				Save(ctx)
		}

		if err != nil {
			return err
		}
	}

	entity.Edges.Steps = newItems // Feedback.

	return nil
}
