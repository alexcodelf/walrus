package workflow

import (
	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/workflow"
)

type RouteGetLatestExecutionRequest struct {
	_ struct{} `route:"GET=/latest-execution"`

	model.WorkflowQueryInput `path:",inline"`
}

type RouteGetLatestExecutionResponse = *model.WorkflowExecutionOutput

type RouteApplyRequest struct {
	_ struct{} `route:"POST=/apply"`

	model.WorkflowQueryInput `path:",inline" json:",inline"`

	Params      map[string]string `json:"params"`
	Description string            `json:"description,omitempty"`
}

type RouteApplyResponse = *model.WorkflowExecutionOutput

func (r *RouteApplyRequest) Validate() error {
	err := r.WorkflowQueryInput.Validate()
	if err != nil {
		return err
	}

	return workflow.CheckParams(r.Params)
}
