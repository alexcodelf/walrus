package types

import (
	"context"

	"github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"

	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/types"
	"github.com/seal-io/walrus/pkg/dao/types/object"
)

// Type is a type of step.
type Type string

func (t Type) String() string {
	return string(t)
}

const (
	StepTypeService  Type = types.WorkflowStepTypeService
	StepTypeApproval Type = types.WorkflowStepTypeApproval
)

type StepManager interface {
	GenerateTemplate(context.Context, *model.WorkflowStepExecution) (*v1alpha1.Template, error)
}

// WorkflowStepServiceSepc is the spec of WorkflowStepService.
type WorkflowStepServiceSepc struct {
	Name          string    `json:"name"`
	ProjectID     object.ID `json:"projectID"`
	EnvironmentID object.ID `json:"environmentID"`

	TemplateName    string `json:"templateName"`
	TemplateVersion string `json:"templateVersion"`
	DeployerType    string `json:"deployerType"`
	Attributes      string `json:"attributes"`
}
