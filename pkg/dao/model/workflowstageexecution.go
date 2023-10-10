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

	"github.com/seal-io/walrus/pkg/dao/model/workflowexecution"
	"github.com/seal-io/walrus/pkg/dao/model/workflowstage"
	"github.com/seal-io/walrus/pkg/dao/model/workflowstageexecution"
	"github.com/seal-io/walrus/pkg/dao/types/object"
	"github.com/seal-io/walrus/pkg/dao/types/status"
	"github.com/seal-io/walrus/utils/json"
)

// WorkflowStageExecution is the model entity for the WorkflowStageExecution schema.
type WorkflowStageExecution struct {
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
	// ID of the project to belong.
	ProjectID object.ID `json:"project_id,omitempty"`
	// Duration of the workflow stage execution.
	Duration int `json:"duration,omitempty"`
	// ID of the workflow stage that this workflow stage execution belongs to.
	StageID object.ID `json:"stage_id,omitempty"`
	// ID of the workflow execution that this workflow stage execution belongs to.
	WorkflowExecutionID object.ID `json:"workflow_execution_id,omitempty"`
	// ID list of the workflow step executions that belong to this workflow stage execution.
	StepExecutionIds []object.ID `json:"step_execution_ids,omitempty"`
	// Log record of the workflow stage execution.
	Record string `json:"record,omitempty"`
	// Input of the workflow stage execution. It's the yaml file that defines the workflow stage execution.
	Input string `json:"input,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the WorkflowStageExecutionQuery when eager-loading is set.
	Edges        WorkflowStageExecutionEdges `json:"edges,omitempty"`
	selectValues sql.SelectValues
}

// WorkflowStageExecutionEdges holds the relations/edges for other nodes in the graph.
type WorkflowStageExecutionEdges struct {
	// Workflow step executions that belong to this workflow stage execution.
	StepExecutions []*WorkflowStepExecution `json:"step_executions,omitempty"`
	// Workflow stage that this workflow stage execution belongs to.
	Stage *WorkflowStage `json:"stage,omitempty"`
	// Workflow execution that this workflow stage execution belongs to.
	WorkflowExecution *WorkflowExecution `json:"workflow_execution,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// StepExecutionsOrErr returns the StepExecutions value or an error if the edge
// was not loaded in eager-loading.
func (e WorkflowStageExecutionEdges) StepExecutionsOrErr() ([]*WorkflowStepExecution, error) {
	if e.loadedTypes[0] {
		return e.StepExecutions, nil
	}
	return nil, &NotLoadedError{edge: "step_executions"}
}

// StageOrErr returns the Stage value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WorkflowStageExecutionEdges) StageOrErr() (*WorkflowStage, error) {
	if e.loadedTypes[1] {
		if e.Stage == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: workflowstage.Label}
		}
		return e.Stage, nil
	}
	return nil, &NotLoadedError{edge: "stage"}
}

// WorkflowExecutionOrErr returns the WorkflowExecution value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WorkflowStageExecutionEdges) WorkflowExecutionOrErr() (*WorkflowExecution, error) {
	if e.loadedTypes[2] {
		if e.WorkflowExecution == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: workflowexecution.Label}
		}
		return e.WorkflowExecution, nil
	}
	return nil, &NotLoadedError{edge: "workflow_execution"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*WorkflowStageExecution) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case workflowstageexecution.FieldLabels, workflowstageexecution.FieldAnnotations, workflowstageexecution.FieldStatus, workflowstageexecution.FieldStepExecutionIds:
			values[i] = new([]byte)
		case workflowstageexecution.FieldID, workflowstageexecution.FieldProjectID, workflowstageexecution.FieldStageID, workflowstageexecution.FieldWorkflowExecutionID:
			values[i] = new(object.ID)
		case workflowstageexecution.FieldDuration:
			values[i] = new(sql.NullInt64)
		case workflowstageexecution.FieldName, workflowstageexecution.FieldDescription, workflowstageexecution.FieldRecord, workflowstageexecution.FieldInput:
			values[i] = new(sql.NullString)
		case workflowstageexecution.FieldCreateTime, workflowstageexecution.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the WorkflowStageExecution fields.
func (wse *WorkflowStageExecution) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case workflowstageexecution.FieldID:
			if value, ok := values[i].(*object.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				wse.ID = *value
			}
		case workflowstageexecution.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				wse.Name = value.String
			}
		case workflowstageexecution.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				wse.Description = value.String
			}
		case workflowstageexecution.FieldLabels:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field labels", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &wse.Labels); err != nil {
					return fmt.Errorf("unmarshal field labels: %w", err)
				}
			}
		case workflowstageexecution.FieldAnnotations:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field annotations", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &wse.Annotations); err != nil {
					return fmt.Errorf("unmarshal field annotations: %w", err)
				}
			}
		case workflowstageexecution.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				wse.CreateTime = new(time.Time)
				*wse.CreateTime = value.Time
			}
		case workflowstageexecution.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				wse.UpdateTime = new(time.Time)
				*wse.UpdateTime = value.Time
			}
		case workflowstageexecution.FieldStatus:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &wse.Status); err != nil {
					return fmt.Errorf("unmarshal field status: %w", err)
				}
			}
		case workflowstageexecution.FieldProjectID:
			if value, ok := values[i].(*object.ID); !ok {
				return fmt.Errorf("unexpected type %T for field project_id", values[i])
			} else if value != nil {
				wse.ProjectID = *value
			}
		case workflowstageexecution.FieldDuration:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field duration", values[i])
			} else if value.Valid {
				wse.Duration = int(value.Int64)
			}
		case workflowstageexecution.FieldStageID:
			if value, ok := values[i].(*object.ID); !ok {
				return fmt.Errorf("unexpected type %T for field stage_id", values[i])
			} else if value != nil {
				wse.StageID = *value
			}
		case workflowstageexecution.FieldWorkflowExecutionID:
			if value, ok := values[i].(*object.ID); !ok {
				return fmt.Errorf("unexpected type %T for field workflow_execution_id", values[i])
			} else if value != nil {
				wse.WorkflowExecutionID = *value
			}
		case workflowstageexecution.FieldStepExecutionIds:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field step_execution_ids", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &wse.StepExecutionIds); err != nil {
					return fmt.Errorf("unmarshal field step_execution_ids: %w", err)
				}
			}
		case workflowstageexecution.FieldRecord:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field record", values[i])
			} else if value.Valid {
				wse.Record = value.String
			}
		case workflowstageexecution.FieldInput:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field input", values[i])
			} else if value.Valid {
				wse.Input = value.String
			}
		default:
			wse.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the WorkflowStageExecution.
// This includes values selected through modifiers, order, etc.
func (wse *WorkflowStageExecution) Value(name string) (ent.Value, error) {
	return wse.selectValues.Get(name)
}

// QueryStepExecutions queries the "step_executions" edge of the WorkflowStageExecution entity.
func (wse *WorkflowStageExecution) QueryStepExecutions() *WorkflowStepExecutionQuery {
	return NewWorkflowStageExecutionClient(wse.config).QueryStepExecutions(wse)
}

// QueryStage queries the "stage" edge of the WorkflowStageExecution entity.
func (wse *WorkflowStageExecution) QueryStage() *WorkflowStageQuery {
	return NewWorkflowStageExecutionClient(wse.config).QueryStage(wse)
}

// QueryWorkflowExecution queries the "workflow_execution" edge of the WorkflowStageExecution entity.
func (wse *WorkflowStageExecution) QueryWorkflowExecution() *WorkflowExecutionQuery {
	return NewWorkflowStageExecutionClient(wse.config).QueryWorkflowExecution(wse)
}

// Update returns a builder for updating this WorkflowStageExecution.
// Note that you need to call WorkflowStageExecution.Unwrap() before calling this method if this WorkflowStageExecution
// was returned from a transaction, and the transaction was committed or rolled back.
func (wse *WorkflowStageExecution) Update() *WorkflowStageExecutionUpdateOne {
	return NewWorkflowStageExecutionClient(wse.config).UpdateOne(wse)
}

// Unwrap unwraps the WorkflowStageExecution entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (wse *WorkflowStageExecution) Unwrap() *WorkflowStageExecution {
	_tx, ok := wse.config.driver.(*txDriver)
	if !ok {
		panic("model: WorkflowStageExecution is not a transactional entity")
	}
	wse.config.driver = _tx.drv
	return wse
}

// String implements the fmt.Stringer.
func (wse *WorkflowStageExecution) String() string {
	var builder strings.Builder
	builder.WriteString("WorkflowStageExecution(")
	builder.WriteString(fmt.Sprintf("id=%v, ", wse.ID))
	builder.WriteString("name=")
	builder.WriteString(wse.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(wse.Description)
	builder.WriteString(", ")
	builder.WriteString("labels=")
	builder.WriteString(fmt.Sprintf("%v", wse.Labels))
	builder.WriteString(", ")
	builder.WriteString("annotations=")
	builder.WriteString(fmt.Sprintf("%v", wse.Annotations))
	builder.WriteString(", ")
	if v := wse.CreateTime; v != nil {
		builder.WriteString("create_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := wse.UpdateTime; v != nil {
		builder.WriteString("update_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", wse.Status))
	builder.WriteString(", ")
	builder.WriteString("project_id=")
	builder.WriteString(fmt.Sprintf("%v", wse.ProjectID))
	builder.WriteString(", ")
	builder.WriteString("duration=")
	builder.WriteString(fmt.Sprintf("%v", wse.Duration))
	builder.WriteString(", ")
	builder.WriteString("stage_id=")
	builder.WriteString(fmt.Sprintf("%v", wse.StageID))
	builder.WriteString(", ")
	builder.WriteString("workflow_execution_id=")
	builder.WriteString(fmt.Sprintf("%v", wse.WorkflowExecutionID))
	builder.WriteString(", ")
	builder.WriteString("step_execution_ids=")
	builder.WriteString(fmt.Sprintf("%v", wse.StepExecutionIds))
	builder.WriteString(", ")
	builder.WriteString("record=")
	builder.WriteString(wse.Record)
	builder.WriteString(", ")
	builder.WriteString("input=")
	builder.WriteString(wse.Input)
	builder.WriteByte(')')
	return builder.String()
}

// WorkflowStageExecutions is a parsable slice of WorkflowStageExecution.
type WorkflowStageExecutions []*WorkflowStageExecution
