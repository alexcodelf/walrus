package types

import (
	"encoding/json"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/zclconf/go-cty/cty"

	"github.com/seal-io/walrus/pkg/dao/types/property"
)

type OutputValue struct {
	Name      string          `json:"name,omitempty"`
	Value     property.Value  `json:"value,omitempty,cli-table-column"`
	Type      cty.Type        `json:"type,omitempty"`
	Sensitive bool            `json:"sensitive,omitempty"`
	Schema    openapi3.Schema `json:"schema,omitempty"`
}

// Task type defines the type of the task to be performed with deployer.
const (
	RunTaskTypeApply   RunJobType = "apply"
	RunTaskTypePlan    RunJobType = "plan"
	RunTaskTypeDestroy RunJobType = "destroy"
)

const (
	RunTypeCreate   RunType = "create"
	RunTypeUpgrade  RunType = "upgrade"
	RunTypeDelete   RunType = "delete"
	RunTypeStart    RunType = "start"
	RunTypeStop     RunType = "stop"
	RunTypeRollback RunType = "rollback"
)

type RunType string

func (t RunType) String() string {
	return string(t)
}

type RunJobType string

func (t RunJobType) String() string {
	return string(t)
}

type ResourceRunConfigData = []byte

type TerraformPlan struct {
	FormatVersion            string                    `json:"format_version"`
	TerraformVersion         string                    `json:"terraform_version"`
	Variables                json.RawMessage           `json:"variables"`
	PlannedValues            json.RawMessage           `json:"planned_values"`
	ResourceComponentChanges []ResourceComponentChange `json:"resource_changes"`
	OutputChanges            json.RawMessage           `json:"output_changes"`
	PriorState               json.RawMessage           `json:"prior_state"`
	Configuration            json.RawMessage           `json:"configuration"`
	RelevantAttributes       []RelevantAttribute       `json:"relevant_attributes"`
	Timestamp                time.Time                 `json:"timestamp"`
}

// ResourceComponentChangeSummary is the summary of the resource component changes.
type ResourceComponentChangeSummary struct {
	Created int `json:"created"`
	Updated int `json:"updated"`
	Deleted int `json:"deleted"`
}

type RelevantAttribute struct {
	Resource  string `json:"resource"`
	Attribute []any  `json:"attribute"`
}

type ResourceComponentChange struct {
	Address       string `json:"address"`
	Mode          string `json:"mode"`
	Type          string `json:"type"`
	Name          string `json:"name"`
	ProviderName  string `json:"provider_name"`
	Change        Change `json:"change"`
	ModuleAddress string `json:"module_address,omitempty"`
	Index         string `json:"index,omitempty"`
}

const (
	ResourceComponentChangeTypeCreate = "create"
	ResourceComponentChangeTypeUpdate = "update"
	ResourceComponentChangeTypeDelete = "delete"
)

type Change struct {
	Actions         []string        `json:"actions"`
	Type            string          `json:"type"`
	Before          json.RawMessage `json:"before"`
	After           json.RawMessage `json:"after"`
	AfterUnknown    json.RawMessage `json:"after_unknown"`
	BeforeSensitive json.RawMessage `json:"before_sensitive"`
	AfterSensitive  json.RawMessage `json:"after_sensitive"`
}
