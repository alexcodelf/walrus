package workflow

import (
	"github.com/gin-gonic/gin"
	"github.com/seal-io/walrus/pkg/dao/model"
)

type CollectionRouteGetTestRequest struct {
	_ struct{} `route:"GET=/test"`

	Context *gin.Context
}

type CollectionRouteGetTestResponse = any

func (r *CollectionRouteGetTestRequest) SetGinContext(ctx *gin.Context) {
	r.Context = ctx
}

type RouteApplyRequest struct {
	_ struct{} `route:"POST=/apply"`

	model.WorkflowQueryInput `path:",inline"`
}

func (r *RouteApplyRequest) Validate() error {
	return r.WorkflowQueryInput.Validate()
}
