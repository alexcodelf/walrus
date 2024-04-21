package v1

import (
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"

	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
)

// ConnectorConfig is the subresource of the Connector resource for extract configuration.
//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:apireg-gen:resource:scope="Namespaced",categories=["walrus"]
type ConnectorConfig struct {
	meta.TypeMeta   `json:",inline"`
	meta.ObjectMeta `json:"metadata,omitempty"`

	Status ConnectorConfigStatus `json:"status"`
}

var _ runtime.Object = (*ConnectorConfig)(nil)

// ConnectorConfigStatus defines the observed state of ConnectorConfig.
type ConnectorConfigStatus struct {
	// ApplicableEnvironmentType is the environment type that the connector is applicable to.
	ApplicableEnvironmentType walruscore.EnvironmentType `json:"applicableEnvironmentType"`

	// Category is the category of the connector.
	Category walruscore.ConnectorCategory `json:"category"`

	// Type is the type of the connector.
	Type string `json:"type"`

	// Version is the version of the configuration.
	Version string `json:"version"`

	// Data is the configuration of the connector.
	Data map[string][]byte `json:"data"`

	// ConditionSummary is the summary of the conditions.
	walruscore.ConditionSummary `json:",inline"`
}
