package workflowexecution

import "github.com/seal-io/walrus/pkg/dao/model"

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
