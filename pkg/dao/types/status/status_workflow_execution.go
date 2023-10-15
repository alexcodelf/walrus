package status

const (
	WorkflowExecutionStatusPending ConditionType = "Pending"
	WorkflowExecutionStatusRunning ConditionType = "Running"
	WorkflowExecutionStatusReady   ConditionType = "Ready"
)

// workflowExecutionStatusPaths makes the following decision.
//
//	|  Condition Type  |     Condition Status    | Human Readable Status | Human Sensible Status |
//	| ---------------- | ----------------------- | --------------------- | --------------------- |
//	| Pending          | Unknown                 | Pending               | Transitioning         |
//	| Pending          | False                   | PendingFailed         | Error                 |
//	| Pending          | True                    | Pending               |                       |
//	| Running          | Unknown                 | Running               | Transitioning         |
//	| Running          | False                   | RunFailed             | Error                 |
//	| Running          | True                    | Running               |                       |
//	| Ready            | Unknown                 | Preparing             | Transitioning         |
//	| Ready            | False                   | NotReady              | Error                 |
//	| Ready            | True                    | Ready                 |                       |
var workflowExecutionStatusPaths = NewWalker(
	[][]ConditionType{
		{
			WorkflowExecutionStatusPending,
			WorkflowExecutionStatusRunning,
			WorkflowExecutionStatusReady,
		},
	},
)

func WalkWorkflowExecution(st *Status) *Summary {
	return workflowExecutionStatusPaths.Walk(st)
}

const (
	WorkflowStageExecutionStatusPending ConditionType = "Pending"
	WorkflowStageExecutionStatusRunning ConditionType = "Running"
	WorkflowStageExecutionStatusReady   ConditionType = "Ready"
)

// workflowStageExecutionStatusPaths makes the following decision.
//
//	|  Condition Type  |     Condition Status    | Human Readable Status | Human Sensible Status |
//	| ---------------- | ----------------------- | --------------------- | --------------------- |
//	| Pending          | Unknown                 | Pending               | Transitioning         |
//	| Pending          | False                   | PendingFailed         | Error                 |
//	| Pending          | True                    | Pending               |                       |
//	| Running          | Unknown                 | Running               | Transitioning         |
//	| Running          | False                   | RunFailed             | Error                 |
//	| Running          | True                    | Running               |                       |
//	| Ready            | Unknown                 | Preparing             | Transitioning         |
//	| Ready            | False                   | NotReady              | Error                 |
//	| Ready            | True                    | Ready                 |                       |
var workflowStageExecutionStatusPaths = NewWalker(
	[][]ConditionType{
		{
			WorkflowStageExecutionStatusPending,
			WorkflowStageExecutionStatusRunning,
			WorkflowStageExecutionStatusReady,
		},
	},
)

func WalkWorkflowStageExecution(st *Status) *Summary {
	return workflowStageExecutionStatusPaths.Walk(st)
}

const (
	WorkflowStepExecutionStatusPending ConditionType = "Pending"
	WorkflowStepExecutionStatusRunning ConditionType = "Running"
	WorkflowStepExecutionStatusReady   ConditionType = "Ready"
)

// workflowStepExecutionStatusPaths makes the following decision.
//
//	|  Condition Type  |     Condition Status    | Human Readable Status | Human Sensible Status |
//	| ---------------- | ----------------------- | --------------------- | --------------------- |
//	| Pending          | Unknown                 | Pending               | Transitioning         |
//	| Pending          | False                   | PendingFailed         | Error                 |
//	| Pending          | True                    | Pending               |                       |
//	| Running          | Unknown                 | Running               | Transitioning         |
//	| Running          | False                   | RunFailed             | Error                 |
//	| Running          | True                    | Running               |                       |
//	| Ready            | Unknown                 | Preparing             | Transitioning         |
//	| Ready            | False                   | NotReady              | Error                 |
//	| Ready            | True                    | Ready                 |                       |
var workflowStepExecutionStatusPaths = NewWalker(
	[][]ConditionType{
		{
			WorkflowStepExecutionStatusPending,
			WorkflowStepExecutionStatusRunning,
			WorkflowStepExecutionStatusReady,
		},
	},
)

func WalkWorkflowStepExecution(st *Status) *Summary {
	return workflowStepExecutionStatusPaths.Walk(st)
}
