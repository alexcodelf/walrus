package workflowexecution

import (
	"github.com/seal-io/walrus/pkg/apis/runtime"
	"github.com/seal-io/walrus/pkg/apis/workflowstageexecution"
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
	return "WorkflowExecution"
}

func (h Handler) SubResourceHandlers() []runtime.IResourceHandler {
	return []runtime.IResourceHandler{
		workflowstageexecution.Handle(h.modelClient, h.k8sConfig),
	}
}
