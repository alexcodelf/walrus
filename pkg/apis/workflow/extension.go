package workflow

import (
	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/model/workflow"
	"github.com/seal-io/walrus/pkg/dao/model/workflowexecution"
	"github.com/seal-io/walrus/pkg/dao/model/workflowstage"
	"github.com/seal-io/walrus/pkg/dao/model/workflowstageexecution"
	"github.com/seal-io/walrus/pkg/dao/model/workflowstep"
	"github.com/seal-io/walrus/pkg/dao/model/workflowstepexecution"
	"github.com/seal-io/walrus/pkg/k8s"
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
		WithStages(func(wsq *model.WorkflowStageExecutionQuery) {
			wsq.WithSteps(func(weeq *model.WorkflowStepExecutionQuery) {
				weeq.Order(model.Asc(workflowstepexecution.FieldCreateTime))
			}).
				Order(model.Asc(workflowstageexecution.FieldCreateTime))
		}).
		First(req.Context)
	if err != nil {
		return nil, err
	}

	return model.ExposeWorkflowExecution(wf), nil
}

func (h Handler) RouteApplyRequest(req RouteApplyRequest) (RouteApplyResponse, error) {
	wf, err := h.modelClient.Workflows().Query().
		Where(workflow.ID(req.ID)).
		WithStages(func(wsq *model.WorkflowStageQuery) {
			wsq.WithSteps(func(wsq *model.WorkflowStepQuery) {
				wsq.Order(model.Asc(workflowstep.FieldCreateTime))
			}).Order(model.Asc(workflowstage.FieldCreateTime))
		}).
		Only(req.Context)
	if err != nil {
		return nil, err
	}

	apiConfig := k8s.ToClientCmdApiConfig(h.k8sConfig)
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
