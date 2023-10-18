package workflowstepexecution

import (
	"github.com/seal-io/walrus/pkg/dao/model"
	"k8s.io/client-go/rest"
)

func Handle(mc model.ClientSet, kc *rest.Config) Handler {
	return Handler{
		modelClient: mc,
		k8sConfig:   kc,
	}
}

type Handler struct {
	modelClient model.ClientSet
	k8sConfig   *rest.Config
}

func (Handler) Kind() string {
	return "WorkflowStepExecution"
}
