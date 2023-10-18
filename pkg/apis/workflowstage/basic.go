package workflowstage

import (
	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/model/workflowstage"
)

func (h Handler) Get(req GetRequest) (GetResponse, error) {
	entity, err := h.modelClient.WorkflowStages().Query().
		Where(workflowstage.IDEQ(req.ID)).
		Only(req.Context)
	if err != nil {
		return nil, err
	}

	return model.ExposeWorkflowStage(entity), nil
}
