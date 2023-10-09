package step

import "github.com/seal-io/walrus/pkg/dao/types/object"

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
