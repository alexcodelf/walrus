package workflowexecution

import "github.com/seal-io/walrus/pkg/dao/model"

type RouteResubmitRequest struct {
	_ struct{} `route:"PUT=/resubmit"`

	model.WorkflowExecutionQueryInput `path:",inline"`
}

func (r *RouteResubmitRequest) Validate() error {
	return r.WorkflowExecutionQueryInput.Validate()
}
