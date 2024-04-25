package apistatus

import v1 "github.com/seal-io/walrus/pkg/apis/walruscore/v1"

const (
	ResourceConditionUnDeployed v1.ConditionType = "Undeployed"
	ResourceConditionDeployed   v1.ConditionType = "Deployed"

	ResourceConditionStopped v1.ConditionType = "Stopped"
	ResourceConditionDeleted v1.ConditionType = "Deleted"

	ResourceConditionProgressing v1.ConditionType = "Progressing"

	ResourceConditionReady v1.ConditionType = "Ready"
)

const (
	ResourceConditionReasonProgressing = "Progressing"
	ResourceConditionReasonProgressed  = "Progressed"
	ResourceConditionReasonError       = "Error"

	ResourceConditionReasonUnDeployed = "UnDeployed"

	ResourceConditionReasonDeploying    = "Deploying"
	ResourceConditionReasonDeployed     = "Deployed"
	ResourceConditionReasonDeployFailed = "DeployFailed"

	ResourceConditionReasonDeleteFailed = "DeleteFailed"
	ResourceConditionReasonDeleting     = "Deleting"

	ResourceConditionReasonPreparing = "Preparing"
	ResourceConditionReasonReady     = "Ready"
	ResourceConditionReasonNotReady  = "NotReady"

	ResourceConditionReasonStopping   = "Stopping"
	ResourceConditionReasonStopped    = "Stopped"
	ResourceConditionReasonStopFailed = "StopFailed"
)

// resourceStatusPaths makes the following decision.
//
//	|  Condition Type  |     Condition Status    | Human Readable Status | Human Sensible Status |
//	| ---------------- | ----------------------- | --------------------- | --------------------- |
//	| Progressing      | Unknown                 | Progressing           | Transitioning         |
//	| Progressing      | False                   | Progressing           | Error                 |
//	| Progressing      | True                    | Progressed            |                       |
//	| Deployed         | Unknown                 | Deploying             | Transitioning         |
//	| Deployed         | False                   | DeployFailed          | Error                 |
//	| Deployed         | True                    | Deployed              |                       |
//	| UnDeployed       | Unknown                 | Transitioning         | Transitioning         |
//	| UnDeployed       | False                   | Error                 | Error                 |
//	| UnDeployed       | True                    | Undeployed            |                       |
//	| Stopped          | Unknown                 | Stopping              | Transitioning         |
//	| Stopped          | False                   | StopFailed            | Error                 |
//	| Stopped          | True                    | Stopped               |                       |
//	| Ready            | Unknown                 | Preparing             | Transitioning         |
//	| Ready            | False                   | NotReady              | Error                 |
//	| Ready            | True                    | Ready                 |                       |
//	| Deleted          | Unknown                 | Deleting              | Transitioning         |
//	| Deleted          | False                   | DeleteFailed          | Error                 |
//	| Deleted          | True                    | Deleted               |                       |
var resourceStatusPaths = NewWalker(
	[][]v1.ConditionType{
		{
			ResourceConditionDeleted,
			ResourceConditionProgressing,
			ResourceConditionDeployed,
			ResourceConditionUnDeployed,
			ResourceConditionStopped,
			ResourceConditionReady,
		},
	},
	func(d Decision[v1.ConditionType]) {
		d.Make(ResourceConditionDeleted,
			func(st v1.ConditionStatus, reason string) (string, string, Score) {
				switch st {
				case v1.ConditionUnknown:
					return "Deleting", "", SummaryScoreTransitioning
				case v1.ConditionFalse:
					return "DeleteFailed", "", SummaryScoreInterrupted
				}
				return "Deleted", "", SummaryScoreDone
			})
		d.Make(ResourceConditionUnDeployed,
			func(st v1.ConditionStatus, reason string) (string, string, Score) {
				switch st {
				case v1.ConditionUnknown:
					return "Transitioning", "", SummaryScoreTransitioning
				case v1.ConditionFalse:
					return "Error", "", SummaryScoreInterrupted
				}
				return "Undeployed", "", SummaryScoreDone
			})
		d.Make(ResourceConditionStopped,
			func(st v1.ConditionStatus, reason string) (string, string, Score) {
				switch st {
				case v1.ConditionUnknown:
					return "Stopping", "", SummaryScoreTransitioning
				case v1.ConditionFalse:
					return "StopFailed", "", SummaryScoreInterrupted
				}
				return "Stopped", "", SummaryScoreDone
			})
	},
)

// WalkResource walks the given status by resource flow.
func WalkResource(st *v1.StatusDescriptor) *v1.ConditionSummary {
	return resourceStatusPaths.Walk(st)
}
