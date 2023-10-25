package types

type WorkflowStepTemplateSchema = map[string]any

const (
	WorkflowTypeBasic       = "Basic"
	WorkflowTypeDAG         = "DAG"
	WorkflowTypeParentChild = "ParentChild"
)

type RetryStrategy map[string]any

const (
	WorkflowStepTypeService = "Service"
)

const (
	WorkflowExecutionTriggerTypeManual = "Manual"
)

// WorkflowExecutionTrigger is the trigger of a workflow execution.
type WorkflowExecutionTrigger struct {
	Type string `json:"type"`
	User string `json:"user"`
}
