package workflowstage

import "github.com/seal-io/walrus/pkg/dao/model"

type GetRequest struct {
	model.WorkflowStageQueryInput `path:",inline"`
}

type GetResponse = *model.WorkflowStageOutput
