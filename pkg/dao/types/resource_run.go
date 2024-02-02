package types

import (
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
