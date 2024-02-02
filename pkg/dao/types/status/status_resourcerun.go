package status

const (
	ResourceRunStatusPending ConditionType = "Pending"
	ResourceRunStatusPlan    ConditionType = "Plan"
	ResourceRunStatusApply   ConditionType = "Apply"

	ResourceRunSummaryStatusRunning string = "Running"
	ResourceRunSummaryStatusFailed  string = "Failed"
	ResourceRunSummaryStatusSucceed string = "Succeeded"
)

// resourceRunStatusPaths makes the following decision.
//
//	|  Condition Type  |     Condition Status    | Human Readable Status | Human Sensible Status |
//	| ---------------- | ----------------------- | --------------------- | --------------------- |
//	| Pending          | Unknown                 | Pending               | Transitioning         |
//	| Pending          | False                   | Failed                | Error                 |
//	| Plan             | Unknown                 | Planning              | Transitioning         |
//	| Plan             | False                   | Failed                | Error                 |
//	| Plan             | True                    | Planned               |                       |
//	| Apply            | Unknown                 | Running               | Transitioning         |
//	| Apply            | False                   | Failed                | Error                 |
//	| Apply            | True                    | Succeeded             |                       |
var resourceRunStatusPaths = NewWalker(
	[][]ConditionType{
		{
			ResourceRunStatusPending,
			ResourceRunStatusPlan,
			ResourceRunStatusApply,
		},
	},
)

func WalkResourceRun(st *Status) *Summary {
	return resourceRunStatusPaths.Walk(st)
}
