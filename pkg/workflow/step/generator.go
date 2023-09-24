package step

import (
	wfv1 "github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
	"github.com/seal-io/walrus/pkg/dao/model"
)

type StepGenerator interface {
	Generate(*model.WorkflowStep) (*wfv1.Template, error)
}
