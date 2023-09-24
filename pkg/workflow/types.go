package workflow

import (
	"github.com/seal-io/walrus/pkg/dao/types/object"
	"github.com/seal-io/walrus/pkg/dao/types/property"
)

type Workflow struct {
	ID          object.ID        `json:"id,omitempty"`
	Name        string           `json:"name"`
	DisplayName string           `json:"displayName"`
	Type        string           `json:"type"`
	Stages      []*WorkflowStage `json:"stages"`
}

type WorkflowStage struct {
	ID         object.ID `json:"id"`
	WorkflowID object.ID `json:"workflowID"`

	Steps []*WorkflowStep `json:"steps"`
}

type WorkflowStep struct {
	ID          object.ID         `json:"id"`
	Name        string            `json:"name"`
	DisplayName string            `json:"displayName"`
	Type        string            `json:"type"`
	Annotation  map[string]string `json:"annotation,omitempty"`
	Timeout     int               `json:"timeout,omitempty"`
	Spec        map[string]any    `json:"spec"`
	Input       map[string]any    `json:"input"`
	Output      map[string]any    `json:"output"`

	TemplateID object.ID `json:"templateID"`
	StageID    object.ID `json:"stageID"`

	Dependency []string `json:"dependency"`
}

type WorkflowTemplate struct {
	ID   object.ID `json:"id"`
	Name string    `json:"name"`

	Schema property.Schemas `json:"schema"`
}
