package workflowstepexecution

import (
	"github.com/seal-io/walrus/pkg/apis/runtime"
	"github.com/seal-io/walrus/pkg/dao/model"
)

type RouteLogRequest struct {
	_ struct{} `route:"GET=/log"`

	model.WorkflowStepExecutionQueryInput `path:",inline"`

	Stream *runtime.RequestUnidiStream
}

func (r *RouteLogRequest) Validate() error {
	if err := r.WorkflowStepExecutionQueryInput.Validate(); err != nil {
		return err
	}

	return nil
}

func (r *RouteLogRequest) SetStream(stream runtime.RequestUnidiStream) {
	r.Stream = &stream
}
