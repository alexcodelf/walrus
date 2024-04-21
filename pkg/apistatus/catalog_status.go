package apistatus

import "github.com/seal-io/walrus/pkg/apis/walruscore/v1"

const (
	CatalogConditionFetched         v1.ConditionType = "Fetched"
	CatalogConditionSyncedTemplates v1.ConditionType = "SyncedTemplates"

	CatalogConditionDeleting v1.ConditionType = "Deleting"
)

const (
	CatalogConditionFetchedReasonPartialFailed = "PartialFetchFailed"
	CatalogConditionFetchedReasonAllFailed     = "AllFetchFailed"

	CatalogConditionSyncedTemplatesReasonSyncing = "SyncingTemplates"
	CatalogConditionSyncedTemplatesReasonFailed  = "SyncTemplatesFailed"
)

// catalogStatusPaths makes the following decision.
//
//	|  Condition Type  |     Condition Status    | Human Readable Status | Human Sensible Status |
//	| ---------------- | ----------------------- | --------------------- | --------------------- |
//	| Fetched          | Unknown                 | Fetching              | Transitioning         |
//	| Fetched          | False                   | PartialFetchFailed    | Interrupted           |
//	| Fetched          | False                   | AllFetchFailed        | Interrupted           |
//	| Fetched          | True                    | Fetched               | /                     |
//	| SyncedTemplates  | Unknown                 | SyncingTemplates      | Transitioning         |
//	| SyncedTemplates  | False                   | SyncTemplatesFailed   | Interrupted           |
//	| SyncedTemplates  | True                    | Ready                 | /                     |
//	| Deleting         | True                    | Deleting              | Transitioning         |
var catalogStatusPaths = NewWalker(
	[][]v1.ConditionType{
		{
			CatalogConditionFetched,
			CatalogConditionSyncedTemplates,
		},
		{
			CatalogConditionDeleting,
		},
	},
	func(d Decision[v1.ConditionType]) {
		d.Make(CatalogConditionFetched,
			func(st v1.ConditionStatus, reason string) (string, string, Score) {
				switch st {
				default:
					if reason == "" {
						reason = "Fetching"
					}
					return reason, "", SummaryScoreTransitioning
				case v1.ConditionFalse:
					return reason, "", SummaryScoreInterrupted
				case v1.ConditionTrue:
					return "Fetched", "", SummaryScoreDone
				}
			})
		d.Make(CatalogConditionSyncedTemplates,
			func(st v1.ConditionStatus, reason string) (string, string, Score) {
				switch st {
				default:
					if reason == "" {
						reason = "SyncingTemplates"
					}
					return reason, "", SummaryScoreTransitioning
				case v1.ConditionFalse:
					return reason, "", SummaryScoreInterrupted
				case v1.ConditionTrue:
					return "Ready", "", SummaryScoreDone
				}
			})
		d.Make(CatalogConditionDeleting,
			func(st v1.ConditionStatus, reason string) (string, string, Score) {
				return "Deleting", "", SummaryScoreTransitioning
			})
	},
)

func WalkCatalog(st *v1.StatusDescriptor) *v1.ConditionSummary {
	return catalogStatusPaths.Walk(st)
}
