package workflowstep

import (
	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/model/workflowstep"
)

func (h Handler) Get(req GetRequest) (GetResponse, error) {
	entity, err := h.modelClient.WorkflowSteps().Query().
		Where(workflowstep.IDEQ(req.ID)).
		Only(req.Context)
	if err != nil {
		return nil, err
	}

	return model.ExposeWorkflowStep(entity), nil
}
