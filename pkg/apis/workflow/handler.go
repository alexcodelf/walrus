package workflow

import (
	"github.com/seal-io/walrus/pkg/dao/model"
	"k8s.io/client-go/rest"
)

func Handle(mc model.ClientSet, k8sConfig *rest.Config) Handler {
	return Handler{
		modelClient: mc,
		k8sConfig:   k8sConfig,
	}
}

type Handler struct {
	modelClient model.ClientSet
	k8sConfig   *rest.Config
}

func (Handler) Kind() string {
	return "Workflow"
}
