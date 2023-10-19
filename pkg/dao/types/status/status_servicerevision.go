package status

const (
	ServiceRevisionStatusPending ConditionType = "Pending"
	ServiceRevisionStatusRunning ConditionType = "Running"
	ServiceRevisionStatusReady   ConditionType = "Ready"

	ServiceRevisionSummaryStatusRunning string = "Running"
	ServiceRevisionSummaryStatusFailed  string = "Failed"
	ServiceRevisionSummaryStatusSucceed string = "Succeed"
)

// serviceRevisionStatusPaths makes the following decision.
//
//	|  Condition Type  |     Condition Status    | Human Readable Status | Human Sensible Status |
//	| ---------------- | ----------------------- | --------------------- | --------------------- |
//	| Pending          | Unknown                 | Pending               | Transitioning         |
//	| Pending          | False                   | Failed                | Error                 |
//	| Pending          | True                    | Pended                |                       |
//	| Running          | Unknown                 | Running               | Transitioning         |
//	| Running          | False                   | Failed                | Error                 |
//	| Running          | True                    | Deployed              |                       |
//	| Ready            | Unknown                 | Preparing             | Transitioning         |
//	| Ready            | False                   | Failed                | Error                 |
//	| Ready            | True                    | Ready                 |                       |
var serviceRevisionStatusPaths = NewWalker(
	[][]ConditionType{
		{
			ServiceRevisionStatusPending,
			ServiceRevisionStatusRunning,
			ServiceRevisionStatusReady,
		},
	},
)

func WalkServiceRevision(st *Status) *Summary {
	return serviceRevisionStatusPaths.Walk(st)
}
