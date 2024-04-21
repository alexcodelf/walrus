package apistatus

import "github.com/seal-io/walrus/pkg/apis/walruscore/v1"

const (
	TemplateConditionSynced v1.ConditionType = "Synced"

	TemplateConditionDeleting v1.ConditionType = "Deleting"
)

const (
	TemplateConditionSyncedReasonFailed = "SyncFailed"
)

// templateStatusPaths makes the following decision.
//
//	|  Condition Type  |     Condition Status    | Human Readable Status | Human Sensible Status |
//	| ---------------- | ----------------------- | --------------------- | --------------------- |
//	| Synced           | Unknown                 | Syncing               | Transitioning         |
//	| Synced           | False                   | SyncFailed            | Interrupted           |
//	| Synced           | True                    | Ready                 | /                     |
//	| Deleting         | True                    | Deleting              | Transitioning         |
var templateStatusPaths = NewWalker(
	[][]v1.ConditionType{
		{
			TemplateConditionSynced,
		},
		{
			TemplateConditionDeleting,
		},
	},
	func(d Decision[v1.ConditionType]) {
		d.Make(TemplateConditionSynced,
			func(st v1.ConditionStatus, reason string) (string, string, Score) {
				switch st {
				default:
					if reason == "" {
						reason = "Syncing"
					}
					return reason, "", SummaryScoreTransitioning
				case v1.ConditionFalse:
					return reason, "", SummaryScoreInterrupted
				case v1.ConditionTrue:
					return "Ready", "", SummaryScoreDone
				}
			})
		d.Make(TemplateConditionDeleting,
			func(st v1.ConditionStatus, reason string) (string, string, Score) {
				return "Deleting", "", SummaryScoreTransitioning
			})
	},
)

func WalkTemplate(st *v1.StatusDescriptor) *v1.ConditionSummary {
	return templateStatusPaths.Walk(st)
}
