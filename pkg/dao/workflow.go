package dao

import (
	"context"
	"errors"

	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/model/workflowstage"
	"github.com/seal-io/walrus/pkg/dao/model/workflowstep"
	"github.com/seal-io/walrus/pkg/dao/types/object"
	"github.com/seal-io/walrus/pkg/dao/types/status"
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

	steps := make([][]*model.WorkflowStep, len(entity.Edges.Stages))
	// Add new items.
	newItems := entity.Edges.Stages
	for i := range newItems {
		if newItems[i] == nil {
			return errors.New("invalid input: nil relationship")
		}
		newItems[i].WorkflowID = entity.ID
		newItems[i].ProjectID = entity.ProjectID

		status.WorkflowStageStatusInitialized.Unknown(newItems[i], "Workflow stage is initialized.")
		newItems[i].Status.SetSummary(status.WalkWorkflowStage(&newItems[i].Status))

		steps[i] = newItems[i].Edges.Steps
	}

	newItems, err = mc.WorkflowStages().CreateBulk().
		Set(newItems...).
		Save(ctx)
	if err != nil {
		return err
	}

	// Save steps after stage is saved.
	for i := range newItems {
		newItems[i].Edges.Steps = steps[i]
		// TODO avoid save steps in loop.
		err := WorkflowStageStepsEdgeSave(ctx, mc, newItems[i])
		if err != nil {
			return err
		}
	}

	entity.Edges.Stages = newItems // Feedback.

	stageIDs := make([]object.ID, len(newItems))
	for i := range newItems {
		stageIDs[i] = newItems[i].ID
	}

	// Save workflow stage IDs.
	err = mc.Workflows().UpdateOneID(entity.ID).
		SetStageIds(stageIDs).
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
		Where(workflowstep.StageID(entity.ID)).
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
		newItems[i].StageID = entity.ID
		newItems[i].ProjectID = entity.ProjectID
		newItems[i].WorkflowID = entity.WorkflowID

		status.WorkflowStageStatusInitialized.Unknown(newItems[i], "Workflow step is initialized.")
		newItems[i].Status.SetSummary(status.WalkWorkflowStep(&newItems[i].Status))
	}

	newItems, err = mc.WorkflowSteps().CreateBulk().
		Set(newItems...).
		Save(ctx)
	if err != nil {
		return err
	}

	entity.Edges.Steps = newItems // Feedback.

	stepIDs := make([]object.ID, len(newItems))
	for i := range newItems {
		stepIDs[i] = newItems[i].ID
	}

	// Save workflow step IDs.
	err = mc.WorkflowStages().UpdateOneID(entity.ID).
		SetStepIds(stepIDs).
		Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
