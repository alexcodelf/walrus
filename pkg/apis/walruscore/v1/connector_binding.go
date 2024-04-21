package v1

import (
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// ConnectorBinding is the schema for the connectorbindings API.
//
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:crd-gen:resource:scope="Namespaced"
type ConnectorBinding struct {
	meta.TypeMeta   `json:",inline"`
	meta.ObjectMeta `json:"metadata,omitempty"`

	Spec ConnectorBindingSpec `json:"spec"`
}

var _ runtime.Object = (*ConnectorBinding)(nil)

// ConnectorBindingSpec defines the desired state of ConnectorBinding.
type ConnectorBindingSpec struct {
	// Connector is the reference to the connector.
	Connector ConnectorReferenceWithType `json:"connector"`
}

// ConnectorBindingList contains a list of ConnectorBinding.
//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ConnectorBindingList struct {
	meta.TypeMeta `json:",inline"`
	meta.ListMeta `json:"metadata,omitempty"`

	Items []ConnectorBinding `json:"items"`
}

var _ runtime.Object = (*ConnectorList)(nil)

// ConnectorReferenceWithType is a reference to a connector with its category and type.
type ConnectorReferenceWithType struct {
	ConnectorReference `json:",inline"`

	// Category is the category of the connector.
	//
	// If the Category is empty,
	// the Category will be set to the category of the connector.
	//
	// +k8s:validation:cel[0]:rule="oldSelf == self"
	// +k8s:validation:cel[0]:message="immutable field"
	Category ConnectorCategory `json:"category,omitempty"`

	// Type is the type of the connector.
	//
	// If the Type is empty,
	// the Type will be set to the type of the connector.
	//
	// +k8s:validation:cel[0]:rule="oldSelf == self"
	// +k8s:validation:cel[0]:message="immutable field"
	Type string `json:"type,omitempty"`
}
