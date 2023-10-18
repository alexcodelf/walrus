package workflow

import (
	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/model/workflow"
	pkgworkflow "github.com/seal-io/walrus/pkg/workflow"
	"k8s.io/client-go/tools/clientcmd"
)

func (h Handler) CollectionRouteGetTestRequest(req CollectionRouteGetTestRequest) (any, error) {
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

	apiConfig := pkgworkflow.CreateKubeconfigFileForRestConfig(h.k8sConfig)
	clientConfig := clientcmd.NewDefaultClientConfig(apiConfig, nil)

	err = h.modelClient.WithTx(req.Context, func(tx *model.Tx) error {
		return pkgworkflow.Apply(req.Context, h.modelClient, clientConfig, wf)
	})

	return nil, err
}
