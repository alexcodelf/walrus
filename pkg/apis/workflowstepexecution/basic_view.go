package workflowstepexecution

import (
	"github.com/seal-io/walrus/pkg/apis/runtime"
	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/model/predicate"
	"github.com/seal-io/walrus/pkg/dao/model/workflowstepexecution"
)

type UpdateRequest struct {
	model.WorkflowStepExecutionUpdateInput `path:",inline" json:",inline"`

	Status string `json:"status"`
}

func (r *UpdateRequest) Validate() error {
	if err := r.WorkflowStepExecutionUpdateInput.Validate(); err != nil {
		return err
	}

	return nil
}

type (
	CollectionGetRequest struct {
		model.WorkflowQueryInputs `path:",inline"`

		runtime.RequestCollection[
			predicate.WorkflowStepExecution, workflowstepexecution.OrderOption,
		] `query:",inline"`

		Stream *runtime.RequestUnidiStream
	}

	CollectionGetResponse = []*model.WorkflowStepExecutionOutput
)

func (r *CollectionGetRequest) SetStream(stream runtime.RequestUnidiStream) {
	r.Stream = &stream
}
