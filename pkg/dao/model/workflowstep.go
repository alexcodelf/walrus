// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus". DO NOT EDIT.

package model

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"

	"github.com/seal-io/walrus/pkg/dao/model/workflowstage"
	"github.com/seal-io/walrus/pkg/dao/model/workflowstep"
	"github.com/seal-io/walrus/pkg/dao/types/object"
	"github.com/seal-io/walrus/pkg/dao/types/status"
	"github.com/seal-io/walrus/utils/json"
)

// WorkflowStep is the model entity for the WorkflowStep schema.
type WorkflowStep struct {
	config `json:"-"`
	// ID of the ent.
	ID object.ID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Labels holds the value of the "labels" field.
	Labels map[string]string `json:"labels,omitempty"`
	// Annotations holds the value of the "annotations" field.
	Annotations map[string]string `json:"annotations,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime *time.Time `json:"update_time,omitempty"`
	// Status holds the value of the "status" field.
	Status status.Status `json:"status,omitempty"`
	// Type of the workflow step.
	Type string `json:"type,omitempty"`
	// ID of the workflow that this workflow step belongs to.
	WorkflowID object.ID `json:"workflow_id,omitempty"`
	// ID of the stage that this workflow step belongs to.
	StageID object.ID `json:"stage_id,omitempty"`
	// Spec of the workflow step.
	Spec map[string]interface{} `json:"spec,omitempty"`
	// Input of the workflow step.
	Input map[string]interface{} `json:"input,omitempty"`
	// Output of the workflow step.
	Output map[string]interface{} `json:"output,omitempty"`
	// ID list of the workflow steps that this workflow step depends on.
	Dependencies []object.ID `json:"dependencies,omitempty"`
	// Retry policy of the workflow step.
	RetryStrategy v1alpha1.RetryStrategy `json:"retryStrategy,omitempty"`
	// Timeout seconds of the workflow step, 0 means no timeout.
	Timeout int `json:"timeout,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the WorkflowStepQuery when eager-loading is set.
	Edges        WorkflowStepEdges `json:"edges,omitempty"`
	selectValues sql.SelectValues
}

// WorkflowStepEdges holds the relations/edges for other nodes in the graph.
type WorkflowStepEdges struct {
	// Workflow step executions that belong to this workflow step.
	Executions []*WorkflowStepExecution `json:"executions,omitempty"`
	// Workflow stage that this workflow step belongs to.
	Stage *WorkflowStage `json:"stage,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ExecutionsOrErr returns the Executions value or an error if the edge
// was not loaded in eager-loading.
func (e WorkflowStepEdges) ExecutionsOrErr() ([]*WorkflowStepExecution, error) {
	if e.loadedTypes[0] {
		return e.Executions, nil
	}
	return nil, &NotLoadedError{edge: "executions"}
}

// StageOrErr returns the Stage value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WorkflowStepEdges) StageOrErr() (*WorkflowStage, error) {
	if e.loadedTypes[1] {
		if e.Stage == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: workflowstage.Label}
		}
		return e.Stage, nil
	}
	return nil, &NotLoadedError{edge: "stage"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*WorkflowStep) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case workflowstep.FieldLabels, workflowstep.FieldAnnotations, workflowstep.FieldStatus, workflowstep.FieldSpec, workflowstep.FieldInput, workflowstep.FieldOutput, workflowstep.FieldDependencies, workflowstep.FieldRetryStrategy:
			values[i] = new([]byte)
		case workflowstep.FieldID, workflowstep.FieldWorkflowID, workflowstep.FieldStageID:
			values[i] = new(object.ID)
		case workflowstep.FieldTimeout:
			values[i] = new(sql.NullInt64)
		case workflowstep.FieldName, workflowstep.FieldDescription, workflowstep.FieldType:
			values[i] = new(sql.NullString)
		case workflowstep.FieldCreateTime, workflowstep.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the WorkflowStep fields.
func (ws *WorkflowStep) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case workflowstep.FieldID:
			if value, ok := values[i].(*object.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ws.ID = *value
			}
		case workflowstep.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ws.Name = value.String
			}
		case workflowstep.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				ws.Description = value.String
			}
		case workflowstep.FieldLabels:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field labels", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &ws.Labels); err != nil {
					return fmt.Errorf("unmarshal field labels: %w", err)
				}
			}
		case workflowstep.FieldAnnotations:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field annotations", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &ws.Annotations); err != nil {
					return fmt.Errorf("unmarshal field annotations: %w", err)
				}
			}
		case workflowstep.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				ws.CreateTime = new(time.Time)
				*ws.CreateTime = value.Time
			}
		case workflowstep.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				ws.UpdateTime = new(time.Time)
				*ws.UpdateTime = value.Time
			}
		case workflowstep.FieldStatus:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &ws.Status); err != nil {
					return fmt.Errorf("unmarshal field status: %w", err)
				}
			}
		case workflowstep.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				ws.Type = value.String
			}
		case workflowstep.FieldWorkflowID:
			if value, ok := values[i].(*object.ID); !ok {
				return fmt.Errorf("unexpected type %T for field workflow_id", values[i])
			} else if value != nil {
				ws.WorkflowID = *value
			}
		case workflowstep.FieldStageID:
			if value, ok := values[i].(*object.ID); !ok {
				return fmt.Errorf("unexpected type %T for field stage_id", values[i])
			} else if value != nil {
				ws.StageID = *value
			}
		case workflowstep.FieldSpec:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field spec", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &ws.Spec); err != nil {
					return fmt.Errorf("unmarshal field spec: %w", err)
				}
			}
		case workflowstep.FieldInput:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field input", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &ws.Input); err != nil {
					return fmt.Errorf("unmarshal field input: %w", err)
				}
			}
		case workflowstep.FieldOutput:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field output", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &ws.Output); err != nil {
					return fmt.Errorf("unmarshal field output: %w", err)
				}
			}
		case workflowstep.FieldDependencies:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field dependencies", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &ws.Dependencies); err != nil {
					return fmt.Errorf("unmarshal field dependencies: %w", err)
				}
			}
		case workflowstep.FieldRetryStrategy:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field retryStrategy", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &ws.RetryStrategy); err != nil {
					return fmt.Errorf("unmarshal field retryStrategy: %w", err)
				}
			}
		case workflowstep.FieldTimeout:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field timeout", values[i])
			} else if value.Valid {
				ws.Timeout = int(value.Int64)
			}
		default:
			ws.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the WorkflowStep.
// This includes values selected through modifiers, order, etc.
func (ws *WorkflowStep) Value(name string) (ent.Value, error) {
	return ws.selectValues.Get(name)
}

// QueryExecutions queries the "executions" edge of the WorkflowStep entity.
func (ws *WorkflowStep) QueryExecutions() *WorkflowStepExecutionQuery {
	return NewWorkflowStepClient(ws.config).QueryExecutions(ws)
}

// QueryStage queries the "stage" edge of the WorkflowStep entity.
func (ws *WorkflowStep) QueryStage() *WorkflowStageQuery {
	return NewWorkflowStepClient(ws.config).QueryStage(ws)
}

// Update returns a builder for updating this WorkflowStep.
// Note that you need to call WorkflowStep.Unwrap() before calling this method if this WorkflowStep
// was returned from a transaction, and the transaction was committed or rolled back.
func (ws *WorkflowStep) Update() *WorkflowStepUpdateOne {
	return NewWorkflowStepClient(ws.config).UpdateOne(ws)
}

// Unwrap unwraps the WorkflowStep entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ws *WorkflowStep) Unwrap() *WorkflowStep {
	_tx, ok := ws.config.driver.(*txDriver)
	if !ok {
		panic("model: WorkflowStep is not a transactional entity")
	}
	ws.config.driver = _tx.drv
	return ws
}

// String implements the fmt.Stringer.
func (ws *WorkflowStep) String() string {
	var builder strings.Builder
	builder.WriteString("WorkflowStep(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ws.ID))
	builder.WriteString("name=")
	builder.WriteString(ws.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(ws.Description)
	builder.WriteString(", ")
	builder.WriteString("labels=")
	builder.WriteString(fmt.Sprintf("%v", ws.Labels))
	builder.WriteString(", ")
	builder.WriteString("annotations=")
	builder.WriteString(fmt.Sprintf("%v", ws.Annotations))
	builder.WriteString(", ")
	if v := ws.CreateTime; v != nil {
		builder.WriteString("create_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := ws.UpdateTime; v != nil {
		builder.WriteString("update_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", ws.Status))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(ws.Type)
	builder.WriteString(", ")
	builder.WriteString("workflow_id=")
	builder.WriteString(fmt.Sprintf("%v", ws.WorkflowID))
	builder.WriteString(", ")
	builder.WriteString("stage_id=")
	builder.WriteString(fmt.Sprintf("%v", ws.StageID))
	builder.WriteString(", ")
	builder.WriteString("spec=")
	builder.WriteString(fmt.Sprintf("%v", ws.Spec))
	builder.WriteString(", ")
	builder.WriteString("input=")
	builder.WriteString(fmt.Sprintf("%v", ws.Input))
	builder.WriteString(", ")
	builder.WriteString("output=")
	builder.WriteString(fmt.Sprintf("%v", ws.Output))
	builder.WriteString(", ")
	builder.WriteString("dependencies=")
	builder.WriteString(fmt.Sprintf("%v", ws.Dependencies))
	builder.WriteString(", ")
	builder.WriteString("retryStrategy=")
	builder.WriteString(fmt.Sprintf("%v", ws.RetryStrategy))
	builder.WriteString(", ")
	builder.WriteString("timeout=")
	builder.WriteString(fmt.Sprintf("%v", ws.Timeout))
	builder.WriteByte(')')
	return builder.String()
}

// WorkflowSteps is a parsable slice of WorkflowStep.
type WorkflowSteps []*WorkflowStep
