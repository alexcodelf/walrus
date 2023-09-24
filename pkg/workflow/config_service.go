package workflow

import (
	"github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
	"github.com/seal-io/walrus/pkg/dao/model"
)

// ConfigService is the interface that defines the operations of workflow config.
// It will convert walrus model.Workflow to argo workflow config.
type ConfigService interface {
	// Convert converts a walrus model.Workflow to argo workflow config.
	Convert(*model.Workflow) (*v1alpha1.Workflow, error)
}
