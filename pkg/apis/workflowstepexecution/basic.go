package workflowstepexecution

import (
	revisionbus "github.com/seal-io/walrus/pkg/bus/servicerevision"
	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/model/service"
	"github.com/seal-io/walrus/pkg/dao/model/servicerevision"
	"github.com/seal-io/walrus/pkg/dao/types/status"
	"github.com/seal-io/walrus/pkg/workflow/step/types"
)

func (h Handler) Update(req UpdateRequest) error {
	entity := req.Model()

	switch req.Status {
	case "Succeeded":
		status.WorkflowStepExecutionStatusReady.Reset(entity, "")
		status.WorkflowStepExecutionStatusReady.True(entity, "")
	case "Error", "Failed":
		status.WorkflowStepExecutionStatusRunning.Reset(entity, "")
		status.WorkflowStepExecutionStatusRunning.False(entity, "execute failed")
	case "Running":
		status.WorkflowStepExecutionStatusRunning.Reset(entity, "")
	}

	entity.Status.SetSummary(status.WalkWorkflowStepExecution(&entity.Status))

	entity, err := h.modelClient.WorkflowStepExecutions().UpdateOne(entity).
		SetRecord(req.Record).
		SetDuration(req.Duration).
		SetStatus(entity.Status).
		Save(req.Context)
	if err != nil {
		return err
	}

	switch entity.Type {
	case types.StepTypeService.String():
		service, err := h.modelClient.Services().Query().
			Where(service.WorkflowStepID(entity.WorkflowStepID)).
			Only(req.Context)
		if err != nil {
			return err
		}

		latestRevision, err := h.modelClient.ServiceRevisions().Query().
			Where(servicerevision.ServiceID(service.ID)).
			Order(model.Desc(servicerevision.FieldCreateTime)).
			First(req.Context)
		if err != nil {
			return err
		}

		switch req.Status {
		case "Succeeded":
			status.ServiceRevisionStatusRunning.Reset(latestRevision, "")
			status.ServiceRevisionStatusReady.True(latestRevision, "")

		case "Failed", "Error":
			status.ServiceRevisionStatusRunning.False(latestRevision, "")
		}

		latestRevision, err = h.modelClient.ServiceRevisions().UpdateOne(latestRevision).
			SetStatus(latestRevision.Status).
			Save(req.Context)
		if err != nil {
			return err
		}

		latestRevision.Status.SetSummary(status.WalkServiceRevision(&latestRevision.Status))

		return revisionbus.Notify(req.Context, h.modelClient, latestRevision)

	case types.StepTypeApproval.String():
		// TODO.
	}

	return nil
}
