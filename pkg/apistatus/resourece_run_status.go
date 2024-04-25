package apistatus

import v1 "github.com/seal-io/walrus/pkg/apis/walruscore/v1"

const (
	ResourceRunConditionPending  v1.ConditionType = "Pending"
	ResourceRunConditionRunning  v1.ConditionType = "Running"
	ResourceRunConditionCanceled v1.ConditionType = "Canceled"
)

const (
	ResourceRunConditionReasonPending   string = "Pending"
	ResourceRunConditionReasonRunning   string = "Running"
	ResourceRunConditionReasonFailed    string = "Failed"
	ResourceRunConditionReasonCanceled  string = "Canceled"
	ResourceRunConditionReasonSucceeded string = "Succeeded"
)

// resourceRunStatusPaths makes the following decision.
//
//	|  Condition Type  |     Condition Status    | Human Readable Status | Human Sensible Status |
//	| ---------------- | ----------------------- | --------------------- | --------------------- |
//	| Pending          | Unknown                 | Pending               | Transitioning         |
//	| Pending          | False                   | Failed                | Error                 |
//	| Pending          | True                    |                       |                       |
//	| Running          | Unknown                 | Running               | Transitioning         |
//	| Running          | False                   | Failed                | Error                 |
//	| Running          | True                    |                       |                       |
//	| Canceled         | True                    | Canceled              | Canceled              |
var resourceRunStatusPaths = NewWalker(
	[][]v1.ConditionType{
		{
			ResourceRunConditionPending,
			ResourceRunConditionRunning,
			ResourceRunConditionCanceled,
		},
	},
	func(d Decision[v1.ConditionType]) {
		d.Make(ResourceRunConditionCanceled,
			func(st v1.ConditionStatus, reason string) (string, string, Score) {
				switch st {
				case v1.ConditionUnknown:
					return "Canceling", "", SummaryScoreTransitioning
				case v1.ConditionFalse:
					return "CancelFailed", "", SummaryScoreInterrupted
				}
				return "Canceled", "", SummaryScoreDone
			})
	},
)

func WalkResourceRun(st *v1.StatusDescriptor) *v1.ConditionSummary {
	return resourceRunStatusPaths.Walk(st)
}

const (
	ResourceRunStepConditionPending v1.ConditionType = "Pending"
	ResourceRunStepConditionRunning v1.ConditionType = "Running"
)

const (
	ResourceRunStepConditionReasonPending string = "Pending"
	ResourceRunStepConditionReasonRunning string = "Running"
	ResourceRunStepConditionReasonFailed  string = "Failed"
	ResourceRunStepConditionReasonSucceed string = "Succeeded"
)
