package workflowstage

import (
	"github.com/seal-io/walrus/pkg/apis/runtime"
	"github.com/seal-io/walrus/pkg/apis/workflowstep"
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
	return "WorkflowStage"
}

func (h Handler) SubResourceHandlers() []runtime.IResourceHandler {
	return []runtime.IResourceHandler{
		workflowstep.Handle(h.modelClient),
	}
}
