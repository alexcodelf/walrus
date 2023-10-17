package workflow

import (
	"github.com/seal-io/walrus/pkg/apis/runtime"
	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/model/predicate"
	"github.com/seal-io/walrus/pkg/dao/model/workflow"
)

type (
	CreateRequest struct {
		model.WorkflowCreateInput `path:",inline" json:",inline"`
	}

	CreateResponse = *model.WorkflowOutput
)

func (r *CreateRequest) Validate() error {
	if err := r.WorkflowCreateInput.Validate(); err != nil {
		return err
	}

	return nil
}

type (
	GetRequest struct {
		model.WorkflowQueryInput `path:",inline"`
	}

	GetResponse = *model.WorkflowOutput
)

type UpdateRequest struct {
	model.WorkflowUpdateInput `path:",inline" json:",inline"`
}

func (r *UpdateRequest) Validate() error {
	if err := r.WorkflowUpdateInput.Validate(); err != nil {
		return err
	}

	return nil
}

type (
	DeleteRequest struct {
		model.WorkflowQueryInput `path:",inline"`
	}

	DeleteResponse = *model.WorkflowDeleteInput
)

type (
	CollectionGetRequest struct {
		model.WorkflowQueryInputs `path:",inline"`

		runtime.RequestCollection[
			predicate.Workflow, workflow.OrderOption,
		] `query:",inline"`

		Stream *runtime.RequestUnidiStream
	}

	CollectionGetResponse = []*model.WorkflowOutput
)

func (r *CollectionGetRequest) SetStream(stream runtime.RequestUnidiStream) {
	r.Stream = &stream
}

type CollectionDeleteRequest = model.WorkflowDeleteInputs
