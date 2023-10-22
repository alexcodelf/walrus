package status

const (
	WorkflowStatusInitialized ConditionType = "Initialized"
	WorkflowStatusRunning     ConditionType = "Running"
	WorkflowStatusReady       ConditionType = "Ready"
)

// workflowStatusPaths makes the following decision.
//
//	|  Condition Type  |     Condition Status    | Human Readable Status | Human Sensible Status |
//	| ---------------- | ----------------------- | --------------------- | --------------------- |
//	| Initialized      | Unknown                 | Initializing          | Transitioning         |
//	| Initialized      | False                   | InitializeFailed      | Error                 |
//	| Initialized      | True                    | Initialized           |                       |
//	| Running          | Unknown                 | Running               | Transitioning         |
//	| Running          | False                   | RunFailed             | Error                 |
//	| Running          | True                    | Running               |                       |
//	| Ready            | Unknown                 | Preparing             | Transitioning         |
//	| Ready            | False                   | NotReady              | Error                 |
//	| Ready            | True                    | Ready                 |                       |
var workflowStatusPaths = NewWalker(
	[][]ConditionType{
		{
			WorkflowStatusInitialized,
			WorkflowStatusRunning,
			WorkflowStatusReady,
		},
	},
)

func WalkWorkflow(st *Status) *Summary {
	return workflowStatusPaths.Walk(st)
}

const (
	WorkflowStageStatusPending ConditionType = "Pending"
	WorkflowStageStatusRunning ConditionType = "Running"
	WorkflowStageStatusReady   ConditionType = "Ready"
)

// workflowStageStatusPaths makes the following decision.
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
var workflowStageStatusPaths = NewWalker(
	[][]ConditionType{
		{
			WorkflowStageStatusPending,
			WorkflowStageStatusRunning,
			WorkflowStageStatusReady,
		},
	},
)

func WalkWorkflowStage(st *Status) *Summary {
	return workflowStageStatusPaths.Walk(st)
}

const (
	WorkflowStepStatusPending ConditionType = "Pending"
	WorkflowStepStatusRunning ConditionType = "Running"
	WorkflowStepStatusReady   ConditionType = "Ready"
)

// workflowStepStatusPaths makes the following decision.
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
var workflowStepStatusPaths = NewWalker(
	[][]ConditionType{
		{
			WorkflowStepStatusPending,
			WorkflowStepStatusRunning,
			WorkflowStepStatusReady,
		},
	},
)

func WalkWorkflowStep(st *Status) *Summary {
	return workflowStepStatusPaths.Walk(st)
}
