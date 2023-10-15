package workflowstageexecution

import "github.com/seal-io/walrus/pkg/dao/model"

type Handler struct {
	modelClient model.ClientSet
}

func Handle(mc model.ClientSet) Handler {
	return Handler{
		modelClient: mc,
	}
}

func (Handler) Kind() string {
	return "WorkflowStageExecution"
}
