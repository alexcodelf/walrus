package workflowstepexecution

import (
	"fmt"

	"github.com/seal-io/walrus/pkg/apis/runtime"
	revisionbus "github.com/seal-io/walrus/pkg/bus/servicerevision"
	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/model/service"
	"github.com/seal-io/walrus/pkg/dao/model/servicerevision"
	"github.com/seal-io/walrus/pkg/dao/model/workflowstepexecution"
	"github.com/seal-io/walrus/pkg/dao/types/status"
	"github.com/seal-io/walrus/pkg/datalisten/modelchange"
	"github.com/seal-io/walrus/pkg/workflow/step/types"
	"github.com/seal-io/walrus/utils/topic"
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

	fmt.Println("开始了", req.ID, req.Status)

	entity.Status.SetSummary(status.WalkWorkflowStepExecution(&entity.Status))

	entity, err := h.modelClient.WorkflowStepExecutions().UpdateOne(entity).
		SetRecord(req.Record).
		SetDuration(req.Duration).
		SetStatus(entity.Status).
		Save(req.Context)
	if err != nil {
		return err
	}

	if entity.Type == types.StepTypeService.String() {
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
			status.ServiceRevisionStatusReady.Reset(latestRevision, "")
			status.ServiceRevisionStatusReady.True(latestRevision, "")

		case "Failed", "Error":
			status.ServiceRevisionStatusRunning.False(latestRevision, "")
		default:
			return nil
		}

		latestRevision.Status.SetSummary(status.WalkServiceRevision(&latestRevision.Status))

		latestRevision, err = h.modelClient.ServiceRevisions().UpdateOne(latestRevision).
			SetStatus(latestRevision.Status).
			Save(req.Context)
		if err != nil {
			return err
		}

		return revisionbus.Notify(req.Context, h.modelClient, latestRevision)
	}

	return nil
}

var (
	queryFields = []string{
		workflowstepexecution.FieldID,
		workflowstepexecution.FieldName,
	}
	getFields  = workflowstepexecution.WithoutFields()
	sortFields = []string{
		workflowstepexecution.FieldID,
		workflowstepexecution.FieldName,
	}
)

func (h Handler) CollectionGet(req CollectionGetRequest) (CollectionGetResponse, int, error) {
	query := h.modelClient.WorkflowStepExecutions().Query()

	if queries, ok := req.Querying(queryFields); ok {
		query = query.Where(queries)
	}

	if stream := req.Stream; stream != nil {
		// Handle stream request.
		if fields, ok := req.Extracting(getFields, getFields...); ok {
			query.Select(fields...)
		}

		if orders, ok := req.Sorting(sortFields, model.Desc(workflowstepexecution.FieldCreateTime)); ok {
			query.Order(orders...)
		}

		t, err := topic.Subscribe(modelchange.WorkflowStepExecution)
		if err != nil {
			return nil, 0, err
		}

		defer func() { t.Unsubscribe() }()

		for {
			var event topic.Event

			event, err = t.Receive(stream)
			if err != nil {
				return nil, 0, err
			}

			dm, ok := event.Data.(modelchange.Event)
			if !ok {
				continue
			}

			var items []*model.WorkflowStepExecutionOutput

			switch dm.Type {
			case modelchange.EventTypeCreate, modelchange.EventTypeUpdate:
				entities, err := query.Clone().
					Where(workflowstepexecution.IDIn(dm.IDs...)).
					Unique(false).
					All(stream)
				if err != nil {
					return nil, 0, err
				}

				items = model.ExposeWorkflowStepExecutions(entities)
			case modelchange.EventTypeDelete:
				items = make([]*model.WorkflowStepExecutionOutput, len(dm.IDs))
				for i := range dm.IDs {
					items[i] = &model.WorkflowStepExecutionOutput{
						ID: dm.IDs[i],
					}
				}
			}

			if len(items) == 0 {
				continue
			}

			resp := runtime.TypedResponse(dm.Type.String(), items)
			if err = stream.SendJSON(resp); err != nil {
				return nil, 0, err
			}
		}
	}

	// Handle normal request.

	// Get count.
	count, err := query.Clone().Count(req.Context)
	if err != nil {
		return nil, 0, err
	}

	// Get entities.
	if limit, offset, ok := req.Paging(); ok {
		query.Limit(limit).Offset(offset)
	}

	if fields, ok := req.Extracting(getFields, getFields...); ok {
		query.Select(fields...)
	}

	if orders, ok := req.Sorting(sortFields, model.Desc(workflowstepexecution.FieldCreateTime)); ok {
		query.Order(orders...)
	}

	entities, err := query.
		Unique(false).
		All(req.Context)
	if err != nil {
		return nil, 0, err
	}

	return model.ExposeWorkflowStepExecutions(entities), count, nil
}
