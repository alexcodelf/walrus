package workflow

import (
	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/model/workflow"
	pkgworkflow "github.com/seal-io/walrus/pkg/workflow"
)

func (h Handler) CollectionRouteGetTestRequest(req CollectionRouteGetTestRequest) (any, error) {
	wf := &model.Workflow{
		ID:        "480121971162554616",
		Name:      "test",
		ProjectID: "480121971162504616",
	}

	err := pkgworkflow.Apply(req.Context, h.modelClient, h.k8sConfig, wf)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (h Handler) RouteApplyRequest(req RouteApplyRequest) (any, error) {
	wf, err := h.modelClient.Workflows().Query().
		Where(workflow.ID(req.ID)).
		WithStages(func(wsq *model.WorkflowStageQuery) {
			wsq.WithSteps()
		}).
		Only(req.Context)
	if err != nil {
		return nil, err
	}

	err = h.modelClient.WithTx(req.Context, func(tx *model.Tx) error {
		return pkgworkflow.Apply(req.Context, h.modelClient, h.k8sConfig, wf)
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}
