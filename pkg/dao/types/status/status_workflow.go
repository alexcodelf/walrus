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
