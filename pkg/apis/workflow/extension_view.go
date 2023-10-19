package workflow

import (
	"github.com/seal-io/walrus/pkg/dao/model"
)

type RouteGetLatestExecutionRequest struct {
	_ struct{} `route:"GET=/latest-execution"`

	model.WorkflowQueryInput `path:",inline"`
}

type RouteGetLatestExecutionResponse = *model.WorkflowExecutionOutput

type RouteApplyRequest struct {
	_ struct{} `route:"POST=/apply"`

	model.WorkflowQueryInput `path:",inline"`
}

type RouteApplyResponse = *model.WorkflowExecutionOutput

func (r *RouteApplyRequest) Validate() error {
	return r.WorkflowQueryInput.Validate()
}
