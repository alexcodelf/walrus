package workflow

import (
	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/model/workflow"
	"github.com/seal-io/walrus/pkg/dao/model/workflowexecution"
	pkgworkflow "github.com/seal-io/walrus/pkg/workflow"
	"k8s.io/client-go/tools/clientcmd"
)

func (h Handler) RouteGetLatestExecutionRequest(req RouteGetLatestExecutionRequest) (
	RouteGetLatestExecutionResponse,
	error,
) {
	wf, err := h.modelClient.WorkflowExecutions().Query().
		Where(workflowexecution.WorkflowID(req.ID)).
		Order(model.Desc(workflowexecution.FieldCreateTime)).
		Only(req.Context)
	if err != nil {
		return nil, err
	}

	return model.ExposeWorkflowExecution(wf), nil
}

func (h Handler) RouteApplyRequest(req RouteApplyRequest) (RouteApplyResponse, error) {
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

	var wfe *model.WorkflowExecution
	err = h.modelClient.WithTx(req.Context, func(tx *model.Tx) error {
		wfe, err = pkgworkflow.Apply(req.Context, h.modelClient, clientConfig, wf)
		if err != nil {
			return err
		}

		return nil
	})

	return model.ExposeWorkflowExecution(wfe), err
}
