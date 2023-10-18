package workflow

import (
	"github.com/seal-io/walrus/pkg/apis/runtime"
	"github.com/seal-io/walrus/pkg/dao"
	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/model/catalog"
	"github.com/seal-io/walrus/pkg/dao/model/workflow"
	"github.com/seal-io/walrus/pkg/dao/types/status"
	"github.com/seal-io/walrus/pkg/datalisten/modelchange"
	"github.com/seal-io/walrus/utils/topic"
)

func (h Handler) Create(req CreateRequest) (CreateResponse, error) {
	entity := req.Model()

	status.WorkflowStatusInitialized.Unknown(entity, "Workflow is initialized.")
	entity.Status.SetSummary(status.WalkWorkflow(&entity.Status))

	var err error
	err = h.modelClient.WithTx(req.Context, func(tx *model.Tx) error {
		entity, err = tx.Workflows().Create().
			Set(entity).
			SaveE(req.Context, dao.WorkflowStagesEdgeSave)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return model.ExposeWorkflow(entity), nil
}

func (h Handler) Update(req UpdateRequest) error {
	entity := req.Model()

	status.WorkflowStatusInitialized.Unknown(entity, "Workflow is initialized.")
	entity.Status.SetSummary(status.WalkWorkflow(&entity.Status))

	var err error
	return h.modelClient.WithTx(req.Context, func(tx *model.Tx) error {
		entity, err = tx.Workflows().UpdateOne(entity).
			Set(entity).
			SaveE(req.Context, dao.WorkflowStagesEdgeSave)
		if err != nil {
			return err
		}

		return nil
	})
}

func (h Handler) Get(req GetRequest) (GetResponse, error) {
	entity, err := h.modelClient.Workflows().Query().
		Where(workflow.ID(req.ID)).
		WithStages(func(sgq *model.WorkflowStageQuery) {
			sgq.WithSteps()
		}).
		Only(req.Context)
	if err != nil {
		return nil, err
	}

	return model.ExposeWorkflow(entity), nil
}

func (h Handler) Delete(req DeleteRequest) (err error) {
	return h.modelClient.Workflows().DeleteOneID(req.ID).
		Exec(req.Context)
}

var (
	queryFields = []string{
		workflow.FieldID,
		workflow.FieldName,
	}
	getFields  = workflow.WithoutFields()
	sortFields = []string{
		workflow.FieldID,
		workflow.FieldName,
	}
)

func (h Handler) CollectionGet(req CollectionGetRequest) (CollectionGetResponse, int, error) {
	query := h.modelClient.Workflows().Query()

	if queries, ok := req.Querying(queryFields); ok {
		query = query.Where(queries)
	}

	if stream := req.Stream; stream != nil {
		// Handle stream request.
		if fields, ok := req.Extracting(getFields, getFields...); ok {
			query.Select(fields...)
		}

		if orders, ok := req.Sorting(sortFields, model.Desc(catalog.FieldCreateTime)); ok {
			query.Order(orders...)
		}

		t, err := topic.Subscribe(modelchange.Catalog)
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

			var items []*model.WorkflowOutput

			switch dm.Type {
			case modelchange.EventTypeCreate, modelchange.EventTypeUpdate:
				entities, err := query.Clone().
					Where(workflow.IDIn(dm.IDs...)).
					Unique(false).
					All(stream)
				if err != nil {
					return nil, 0, err
				}

				items = model.ExposeWorkflows(entities)
			case modelchange.EventTypeDelete:
				items = make([]*model.WorkflowOutput, len(dm.IDs))
				for i := range dm.IDs {
					items[i] = &model.WorkflowOutput{
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

	if orders, ok := req.Sorting(sortFields, model.Desc(catalog.FieldCreateTime)); ok {
		query.Order(orders...)
	}

	entities, err := query.
		Unique(false).
		All(req.Context)
	if err != nil {
		return nil, 0, err
	}

	return model.ExposeWorkflows(entities), count, nil
}