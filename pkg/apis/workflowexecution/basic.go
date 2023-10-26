package workflowexecution

import (
	"github.com/seal-io/walrus/pkg/apis/runtime"
	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/model/workflowexecution"
	"github.com/seal-io/walrus/pkg/dao/model/workflowstageexecution"
	"github.com/seal-io/walrus/pkg/dao/model/workflowstepexecution"
	"github.com/seal-io/walrus/pkg/dao/types/status"
	"github.com/seal-io/walrus/pkg/datalisten/modelchange"
	"github.com/seal-io/walrus/pkg/workflow"
	"github.com/seal-io/walrus/utils/topic"
)

func (h Handler) Get(req GetRequest) (GetResponse, error) {
	entity, err := h.modelClient.WorkflowExecutions().Query().
		Where(workflowexecution.ID(req.ID)).
		WithStages(func(wsgq *model.WorkflowStageExecutionQuery) {
			wsgq.WithSteps(func(wseq *model.WorkflowStepExecutionQuery) {
				wseq.Order(model.Asc(workflowstepexecution.FieldCreateTime))
			}).
				Order(model.Asc(workflowstageexecution.FieldCreateTime))
		}).
		Only(req.Context)
	if err != nil {
		return nil, err
	}

	return model.ExposeWorkflowExecution(entity), nil
}

func (h Handler) Update(req UpdateRequest) error {
	entity := req.Model()

	switch req.Status {
	case "Succeeded":
		status.WorkflowExecutionStatusRunning.True(entity, "")
		status.WorkflowExecutionStatusReady.True(entity, "")
	case "Error", "Failed":
		status.WorkflowExecutionStatusRunning.False(entity, "execute failed")
	case "Running":
		status.WorkflowExecutionStatusRunning.Unknown(entity, "")
	}

	entity.Status.SetSummary(status.WalkWorkflowExecution(&entity.Status))

	statusManager := workflow.NewStatusManager(h.modelClient)
	if err := statusManager.HandleWorkflowExecutionFailed(req.Context, entity); err != nil {
		return err
	}

	return h.modelClient.WorkflowExecutions().UpdateOne(entity).
		SetDescription(req.Description).
		SetDuration(req.Duration).
		SetRecord(req.Record).
		SetStatus(entity.Status).
		Exec(req.Context)
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
	query := h.modelClient.WorkflowExecutions().Query().
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
							wseq.Order(model.Asc(workflowstepexecution.FieldCreateTime))
						}).
							Order(model.Asc(workflowstageexecution.FieldCreateTime))
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
				wseq.Order(model.Asc(workflowstepexecution.FieldCreateTime))
			}).
				Order(model.Asc(workflowstageexecution.FieldCreateTime))
		}).
		All(req.Context)
	if err != nil {
		return nil, 0, err
	}

	return model.ExposeWorkflowExecutions(entities), count, nil
}
