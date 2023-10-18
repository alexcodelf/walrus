package workflowstep

import "github.com/seal-io/walrus/pkg/dao/model"

type GetRequest struct {
	model.WorkflowStepQueryInput `path:",inline"`
}

type GetResponse = *model.WorkflowStepOutput
