package workflow

import (
	"github.com/seal-io/walrus/pkg/dao"
	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/model/workflow"
	"github.com/seal-io/walrus/pkg/dao/model/workflowexecution"
	"github.com/seal-io/walrus/pkg/dao/model/workflowstage"
	"github.com/seal-io/walrus/pkg/dao/model/workflowstageexecution"
	"github.com/seal-io/walrus/pkg/dao/model/workflowstep"
	"github.com/seal-io/walrus/pkg/dao/model/workflowstepexecution"
	pkgworkflow "github.com/seal-io/walrus/pkg/workflow"
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
				weeq.Order(model.Asc(workflowstepexecution.FieldOrder))
			}).
				Order(model.Asc(workflowstageexecution.FieldOrder))
		}).
		First(req.Context)
	if err != nil && !model.IsNotFound(err) {
		return nil, err
	}

	if wf == nil {
		return nil, nil
	}

	return model.ExposeWorkflowExecution(wf), nil
}

func (h Handler) RouteApplyRequest(req RouteRunRequest) (RouteRunResponse, error) {
	wf, err := h.modelClient.Workflows().Query().
		Where(workflow.ID(req.ID)).
		WithStages(func(wsq *model.WorkflowStageQuery) {
			wsq.WithSteps(func(wsq *model.WorkflowStepQuery) {
				wsq.Order(model.Asc(workflowstep.FieldOrder))
			}).Order(model.Asc(workflowstage.FieldOrder))
		}).
		Only(req.Context)
	if err != nil {
		return nil, err
	}

	var wfe *model.WorkflowExecution

	err = h.modelClient.WithTx(req.Context, func(tx *model.Tx) error {
		wfe, err = pkgworkflow.Run(req.Context, h.modelClient, wf, dao.ExecuteOptions{
			RestCfg:     h.k8sConfig,
			Params:      req.Params,
			Description: req.Description,
		})
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	// Update workflow version.
	wf, err = h.modelClient.Workflows().UpdateOne(wf).
		AddVersion(1).
		Save(req.Context)
	if err != nil {
		return nil, err
	}

	return model.ExposeWorkflowExecution(wfe), nil
}
