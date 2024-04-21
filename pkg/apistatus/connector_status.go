package apistatus

import v1 "github.com/seal-io/walrus/pkg/apis/walruscore/v1"

const (
	ConnectorConditionConnected v1.ConditionType = "Connected"

	ConnectorConditionDeleting v1.ConditionType = "Deleting"
)

const (
	ConnectorConditionConnectedReasonFailed = "Disconnected"
)

// connectorStatusPaths makes the following decision.
//
//	|  Condition Type  |     Condition Status    | Human Readable Status | Human Sensible Status |
//	| ---------------- | ----------------------- | --------------------- | --------------------- |
//	| Connected        | Unknown                 | Connecting            | Transitioning         |
//	| Connected        | False                   | Disconnected          | Interrupted           |
//	| Connected        | True                    | Connected             | /                     |
//	| Deleting         | True                    | Deleting              | Transitioning         |
var connectorStatusPaths = NewWalker(
	[][]v1.ConditionType{
		{
			ConnectorConditionConnected,
		},
		{
			ConnectorConditionDeleting,
		},
	},
	func(d Decision[v1.ConditionType]) {
		d.Make(TemplateConditionDeleting,
			func(st v1.ConditionStatus, reason string) (string, string, Score) {
				return "Deleting", "", SummaryScoreTransitioning
			})
	},
)

// WalkConnector walks the given status by connector flow.
func WalkConnector(st *v1.StatusDescriptor) *v1.ConditionSummary {
	return connectorStatusPaths.Walk(st)
}
