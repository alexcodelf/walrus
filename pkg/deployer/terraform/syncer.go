package terraform

import (
	"context"

	busrun "github.com/seal-io/walrus/pkg/bus/resourcerun"
	"github.com/seal-io/walrus/pkg/dao/model/resource"
	"github.com/seal-io/walrus/pkg/dao/types/status"
)

// SyncResourceRunStatus updates the status of the run according to its recent finished resource run.
func SyncResourceRunStatus(ctx context.Context, bm busrun.BusMessage) (err error) {
	var (
		mc  = bm.TransactionalModelClient
		run = bm.Refer
	)

	// Report to resource.
	entity, err := mc.Resources().Query().
		Where(resource.ID(run.ResourceID)).
		Select(
			resource.FieldID,
			resource.FieldStatus).
		Only(ctx)
	if err != nil {
		return err
	}

	if status.ResourceRunStatusReady.IsTrue(run) {
		switch {
		case status.ResourceStatusDeleted.IsUnknown(entity):
			return mc.Resources().DeleteOne(entity).
				Exec(ctx)
		case status.ResourceStatusStopped.IsUnknown(entity):
			// Stopping -> Stopped.
			status.ResourceStatusStopped.True(entity, "")
		default:
			// Deployed.
			status.ResourceStatusDeployed.True(entity, "")
			status.ResourceStatusReady.Unknown(entity, "")
		}
	} else if status.ResourceRunStatusReady.IsFalse(run) {
		switch {
		case status.ResourceStatusDeleted.IsUnknown(entity):
			status.ResourceStatusDeleted.False(entity, "")
		default:
			status.ResourceStatusDeployed.False(entity, "")
		}

		entity.Status.SummaryStatusMessage = run.Status.SummaryStatusMessage
	}

	entity.Status.SetSummary(status.WalkResource(&entity.Status))

	return mc.Resources().UpdateOne(entity).
		SetStatus(entity.Status).
		Exec(ctx)
}
