package v1

import (
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

// Connector is the schema for the connectors API.
//
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:crd-gen:resource:scope="Namespaced",subResources=["status"]
type Connector struct {
	meta.TypeMeta   `json:",inline"`
	meta.ObjectMeta `json:"metadata,omitempty"`

	Spec   ConnectorSpec   `json:"spec"`
	Status ConnectorStatus `json:"status,omitempty"`
}

var _ runtime.Object = (*Connector)(nil)

// ConnectorReference is a reference to a connector.
type ConnectorReference struct {
	// Name is the name of the connector.
	Name string `json:"name"`
	// Namespace is the namespace of the connector.
	Namespace string `json:"namespace"`
}

type Connectors = []Connector

func (in *ConnectorReference) ToNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Namespace: in.Namespace,
		Name:      in.Name,
	}
}

// ConnectorCategory is the category of the connector.
//
// +enum
type ConnectorCategory string

const (
	ConnectorCategoryDocker        ConnectorCategory = "Docker"
	ConnectorCategoryKubernetes    ConnectorCategory = "Kubernetes"
	ConnectorCategoryCustom        ConnectorCategory = "Custom"
	ConnectorCategoryCloudProvider ConnectorCategory = "CloudProvider"
)

// ConnectorSpec defines the desired state of Connector.
type ConnectorSpec struct {
	// ApplicableEnvironmentType is the environment type that the connector is applicable to.
	//
	// +k8s:validation:cel[0]:rule="oldSelf == self"
	// +k8s:validation:cel[0]:message="immutable field"
	// +k8s:validation:enum=["Development","Staging","Production"]
	ApplicableEnvironmentType EnvironmentType `json:"applicableEnvironmentType,omitempty"`

	// Category is the category of the connector.
	//
	// +k8s:validation:cel[0]:rule="oldSelf == self"
	// +k8s:validation:cel[0]:message="immutable field"
	// +k8s:validation:enum=["Docker","Kubernetes","Custom","CloudProvider"]
	Category ConnectorCategory `json:"category"`

	// Type is the type of the connector.
	//
	// +k8s:validation:cel[0]:rule="oldSelf == self"
	// +k8s:validation:cel[0]:message="immutable field"
	Type string `json:"type"`

	// Config is the configuration of the connector.
	//
	// Any sensitive configuration entry will be erased before storing.
	Config ConnectorConfig `json:"config"`

	// SecretName is the name of the secret that stores the Config.
	//
	// If the secret name is not provided, a secret will be created to store the Config,
	// otherwise, the Config will be stored in the secret with the provided name.
	//
	// +k8s:validation:cel[0]:rule="oldSelf == self"
	// +k8s:validation:cel[0]:message="immutable field"
	SecretName string `json:"secretName,omitempty"`

	// Description is the description of the connector.
	Description string `json:"description,omitempty"`
}

type (
	// ConnectorConfig defines the configuration of the Connector.
	ConnectorConfig struct {
		// Version is the version of the configuration.
		Version string `json:"version"`
		// Data holds the configuration entries.
		//
		// +mapType=atomic
		Data map[string]ConnectorConfigEntry `json:"data"`
	}

	// ConnectorConfigEntry defines the configuration entry of the Connector.
	ConnectorConfigEntry struct {
		// Sensitive indicates whether the entry is sensitive.
		Sensitive bool `json:"sensitive"`
		// Value is the value of the configuration entry.
		//
		// When Sensitive is true,
		// it is provided as a write-only input field,
		// and returns "(sensitive)".
		Value string `json:"value"`
	}
)

// ConnectorStatus defines the observed state of Connector.
type ConnectorStatus struct {
	// StatusDescriptor defines the status of the Connector.
	StatusDescriptor `json:",inline"`

	// Project is the project that the connector belongs to.
	Project string `json:"project,omitempty"`
}

// ConnectorList holds the list of Connector.
//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ConnectorList struct {
	meta.TypeMeta `json:",inline"`
	meta.ListMeta `json:"metadata,omitempty"`

	Items []Connector `json:"items"`
}

var _ runtime.Object = (*ConnectorList)(nil)

type Property struct {
	// Value indicates the value of property.
	Value string `json:"value"`
	// Visible indicates to show the value of this property or not.
	Visible bool `json:"visible"`
}
