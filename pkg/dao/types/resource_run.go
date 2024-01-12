package types

import (
	"encoding/json"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/zclconf/go-cty/cty"

	"github.com/seal-io/walrus/pkg/dao/types/property"
)

type OutputValue struct {
	Name      string          `json:"name,omitempty"`
	Value     property.Value  `json:"value,omitempty"`
	Type      cty.Type        `json:"type,omitempty"`
	Sensitive bool            `json:"sensitive,omitempty"`
	Schema    openapi3.Schema `json:"schema,omitempty"`
}

type ResourceChangeCount struct {
	// Create is the number of resources to be created.
	Create int `json:"create,omitempty"`
	// Update is the number of resources to be updated.
	Update int `json:"update,omitempty"`
	// Delete is the number of resources to be deleted.
	Delete int `json:"delete,omitempty"`
}

type ResourceComponentChange struct {
	Address       string  `json:"address"`
	ModuleAddress string  `json:"module_address"`
	Mode          string  `json:"mode"`
	Type          string  `json:"type"`
	Name          string  `json:"name"`
	ProviderName  string  `json:"provider_name"`
	Change        *Change `json:"change"`
}

type Change struct {
	Actions         []string        `json:"actions"`
	Before          json.RawMessage `json:"before"`
	After           json.RawMessage `json:"after"`
	AfterUnknown    json.RawMessage `json:"after_unknown"`
	BeforeSensitive json.RawMessage `json:"before_sensitive"`
	AfterSensitive  json.RawMessage `json:"after_sensitive"`
}
