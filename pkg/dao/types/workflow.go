package types

import "github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"

type WorkflowStepTemplateSchema = map[string]any

const (
	WorkflowTypeDefault = "default"
	WorkflowTypeCron    = "cron"
)

type RetryStrategy = v1alpha1.RetryStrategy

const (
	WorkflowStepTypeService  = "service"
	WorkflowStepTypeApproval = "approval"
)

const (
	WorkflowExecutionTriggerTypeManual = "manual"
)

// WorkflowExecutionTrigger is the trigger of a workflow execution.
type WorkflowExecutionTrigger struct {
	Type string `json:"type"`
	User string `json:"user"`
}

const (
	ExecutionStatusRunning   = "Running"
	ExecutionStatusSucceeded = "Succeeded"
	ExecutionStatusFailed    = "Failed"
	ExecutionStatusError     = "Error"
)
