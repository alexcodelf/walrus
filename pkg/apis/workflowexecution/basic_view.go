package workflowexecution

import (
	"github.com/seal-io/walrus/pkg/apis/runtime"
	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/model/predicate"
	"github.com/seal-io/walrus/pkg/dao/model/workflowexecution"
)

type (
	GetRequest struct {
		model.WorkflowExecutionQueryInput `path:",inline"`
	}

	GetResponse = *model.WorkflowExecutionOutput
)

type UpdateRequest struct {
	model.WorkflowExecutionUpdateInput `path:",inline" json:",inline"`

	Status string `json:"status"`
}

func (r *UpdateRequest) Validate() error {
	if err := r.WorkflowExecutionUpdateInput.Validate(); err != nil {
		return err
	}

	return nil
}

type (
	CollectionGetRequest struct {
		model.WorkflowExecutionQueryInputs `path:",inline"`

		runtime.RequestCollection[
			predicate.WorkflowExecution, workflowexecution.OrderOption,
		] `query:",inline"`

		Stream *runtime.RequestUnidiStream
	}

	CollectionGetResponse []*model.WorkflowExecutionOutput
)

func (r *CollectionGetRequest) SetStream(stream runtime.RequestUnidiStream) {
	r.Stream = &stream
}
