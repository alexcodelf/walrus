package types

type WorkflowStepTemplateSchema = map[string]any

const (
	WorkflowTypeBasic       = "Basic"
	WorkflowTypeDAG         = "DAG"
	WorkflowTypeParentChild = "ParentChild"
)

type RetryStrategy map[string]any
