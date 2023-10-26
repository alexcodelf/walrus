package workflowexecution

import (
	"fmt"

	"github.com/seal-io/walrus/pkg/apis/runtime"
	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/model/workflowexecution"
	"github.com/seal-io/walrus/pkg/dao/model/workflowstageexecution"
	"github.com/seal-io/walrus/pkg/dao/model/workflowstepexecution"
	"github.com/seal-io/walrus/pkg/dao/types/object"
	"github.com/seal-io/walrus/pkg/dao/types/status"
	"github.com/seal-io/walrus/pkg/datalisten/modelchange"
	"github.com/seal-io/walrus/utils/topic"
)

func (h Handler) Get(req GetRequest) (GetResponse, error) {
	entity, err := h.modelClient.WorkflowExecutions().Query().
		Where(workflowexecution.ID(req.ID)).
		WithStages(func(wsgq *model.WorkflowStageExecutionQuery) {
			wsgq.WithSteps(func(wseq *model.WorkflowStepExecutionQuery) {
				wseq.Order(model.Asc(workflowstepexecution.FieldOrder))
			}).
				Order(model.Asc(workflowstageexecution.FieldOrder))
		}).
		Only(req.Context)
	if err != nil {
		return nil, err
	}

	return model.ExposeWorkflowExecution(entity), nil
}

func (h Handler) Update(req UpdateRequest) error {
	entity := req.Model()

	fmt.Println("workflow execution update", entity.ID, req.Status)

	switch req.Status {
	case "Succeeded":
		status.WorkflowExecutionStatusRunning.True(entity, "")
		status.WorkflowExecutionStatusReady.True(entity, "")
	case "Error", "Failed":
		status.WorkflowExecutionStatusRunning.False(entity, "")
	case "Running":
		status.WorkflowExecutionStatusPending.True(entity, "")
		status.WorkflowExecutionStatusRunning.Unknown(entity, "")
	default:
		return nil
	}

	entity.Status.SetSummary(status.WalkWorkflowExecution(&entity.Status))

	update := h.modelClient.WorkflowExecutions().UpdateOne(entity).
		SetStatus(entity.Status)

	if req.Record != "" {
		update = update.SetRecord(req.Record)
	}

	if req.Duration > 0 {
		update = update.SetDuration(req.Duration)
	}

	entity, err := update.Save(req.Context)
	if err != nil {
		return err
	}

	// Publish workflow execution topic.
	return topic.Publish(req.Context, modelchange.Workflow, modelchange.Event{
		Type: modelchange.EventTypeUpdate,
		IDs:  []object.ID{entity.WorkflowID},
	})
}

var (
	queryFields = []string{
		workflowexecution.FieldID,
		workflowexecution.FieldName,
		workflowexecution.FieldWorkflowID,
	}
	getFields  = workflowexecution.WithoutFields()
	sortFields = []string{
		workflowexecution.FieldID,
		workflowexecution.FieldName,
	}
)

func (h Handler) CollectionGet(req CollectionGetRequest) (CollectionGetResponse, int, error) {
	query := h.modelClient.WithDebug().WorkflowExecutions().Query().
		Where(workflowexecution.WorkflowID(req.Workflow.ID))

	if queries, ok := req.Querying(queryFields); ok {
		query = query.Where(queries)
	}

	if stream := req.Stream; stream != nil {
		// Handle stream request.
		if fields, ok := req.Extracting(getFields, getFields...); ok {
			query.Select(fields...)
		}

		if orders, ok := req.Sorting(sortFields, model.Desc(workflowexecution.FieldCreateTime)); ok {
			query.Order(orders...)
		}

		t, err := topic.Subscribe(modelchange.WorkflowExecution)
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

			var items []*model.WorkflowExecutionOutput

			switch dm.Type {
			case modelchange.EventTypeCreate, modelchange.EventTypeUpdate:
				entities, err := query.Clone().
					Where(workflowexecution.IDIn(dm.IDs...)).
					WithStages(func(wsgq *model.WorkflowStageExecutionQuery) {
						wsgq.WithSteps(func(wseq *model.WorkflowStepExecutionQuery) {
							wseq.Order(model.Asc(workflowstepexecution.FieldOrder))
						}).Order(model.Asc(workflowstageexecution.FieldOrder))
					}).
					Unique(false).
					All(stream)
				if err != nil {
					return nil, 0, err
				}

				items = model.ExposeWorkflowExecutions(entities)
			case modelchange.EventTypeDelete:
				items = make([]*model.WorkflowExecutionOutput, len(dm.IDs))
				for i := range dm.IDs {
					items[i] = &model.WorkflowExecutionOutput{
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

	if orders, ok := req.Sorting(sortFields, model.Desc(workflowexecution.FieldCreateTime)); ok {
		query.Order(orders...)
	}

	entities, err := query.
		Unique(false).
		WithStages(func(wsgq *model.WorkflowStageExecutionQuery) {
			wsgq.WithSteps(func(wseq *model.WorkflowStepExecutionQuery) {
				wseq.Order(model.Asc(workflowstepexecution.FieldOrder))
			}).Order(model.Asc(workflowstageexecution.FieldOrder))
		}).
		All(req.Context)
	if err != nil {
		return nil, 0, err
	}

	return model.ExposeWorkflowExecutions(entities), count, nil
}
