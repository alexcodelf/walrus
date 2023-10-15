package stage

import (
	"github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
	"github.com/seal-io/walrus/pkg/dao/model"
)

// CreateStageTemplate creates a DAG template for a stage.
// Edge WorkflowStep is required.
func CreateStageTemplate(stage *model.WorkflowStage) *v1alpha1.Template {
	panic("implement me")
}
