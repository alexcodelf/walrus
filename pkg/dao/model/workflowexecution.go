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

	"github.com/seal-io/walrus/pkg/dao/model/project"
	"github.com/seal-io/walrus/pkg/dao/model/workflow"
	"github.com/seal-io/walrus/pkg/dao/model/workflowexecution"
	"github.com/seal-io/walrus/pkg/dao/types"
	"github.com/seal-io/walrus/pkg/dao/types/object"
	"github.com/seal-io/walrus/pkg/dao/types/status"
	"github.com/seal-io/walrus/utils/json"
)

// WorkflowExecution is the model entity for the WorkflowExecution schema.
type WorkflowExecution struct {
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
	// Version of the workflow execution.
	Version int `json:"version,omitempty"`
	// Type of the workflow execution.
	Type string `json:"type,omitempty"`
	// ID of the workflow that this workflow execution belongs to.
	WorkflowID object.ID `json:"workflow_id,omitempty"`
	// ID of the subject that create workflow execution.
	SubjectID object.ID `json:"subject_id,omitempty"`
	// Time of the workflow execution started.
	ExecuteTime time.Time `json:"execute_time,omitempty"`
	// Number of times that this workflow execution has been executed.
	Times int `json:"times,omitempty"`
	// Duration seconds of the workflow execution.
	Duration int `json:"duration,omitempty"`
	// Number of task pods that can be executed in parallel of workflow.
	Parallelism int `json:"parallelism,omitempty"`
	// Timeout of the workflow execution.
	Timeout int `json:"timeout,omitempty"`
	// Trigger of the workflow execution.
	Trigger types.WorkflowExecutionTrigger `json:"trigger,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the WorkflowExecutionQuery when eager-loading is set.
	Edges        WorkflowExecutionEdges `json:"edges,omitempty"`
	selectValues sql.SelectValues
}

// WorkflowExecutionEdges holds the relations/edges for other nodes in the graph.
type WorkflowExecutionEdges struct {
	// Project to which the workflow execution belongs.
	Project *Project `json:"project,omitempty"`
	// Workflow stage executions that belong to this workflow execution.
	Stages []*WorkflowStageExecution `json:"stages,omitempty"`
	// Workflow that this workflow execution belongs to.
	Workflow *Workflow `json:"workflow,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// ProjectOrErr returns the Project value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WorkflowExecutionEdges) ProjectOrErr() (*Project, error) {
	if e.loadedTypes[0] {
		if e.Project == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: project.Label}
		}
		return e.Project, nil
	}
	return nil, &NotLoadedError{edge: "project"}
}

// StagesOrErr returns the Stages value or an error if the edge
// was not loaded in eager-loading.
func (e WorkflowExecutionEdges) StagesOrErr() ([]*WorkflowStageExecution, error) {
	if e.loadedTypes[1] {
		return e.Stages, nil
	}
	return nil, &NotLoadedError{edge: "stages"}
}

// WorkflowOrErr returns the Workflow value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WorkflowExecutionEdges) WorkflowOrErr() (*Workflow, error) {
	if e.loadedTypes[2] {
		if e.Workflow == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: workflow.Label}
		}
		return e.Workflow, nil
	}
	return nil, &NotLoadedError{edge: "workflow"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*WorkflowExecution) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case workflowexecution.FieldLabels, workflowexecution.FieldAnnotations, workflowexecution.FieldStatus, workflowexecution.FieldTrigger:
			values[i] = new([]byte)
		case workflowexecution.FieldID, workflowexecution.FieldProjectID, workflowexecution.FieldWorkflowID, workflowexecution.FieldSubjectID:
			values[i] = new(object.ID)
		case workflowexecution.FieldVersion, workflowexecution.FieldTimes, workflowexecution.FieldDuration, workflowexecution.FieldParallelism, workflowexecution.FieldTimeout:
			values[i] = new(sql.NullInt64)
		case workflowexecution.FieldName, workflowexecution.FieldDescription, workflowexecution.FieldType:
			values[i] = new(sql.NullString)
		case workflowexecution.FieldCreateTime, workflowexecution.FieldUpdateTime, workflowexecution.FieldExecuteTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the WorkflowExecution fields.
func (we *WorkflowExecution) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case workflowexecution.FieldID:
			if value, ok := values[i].(*object.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				we.ID = *value
			}
		case workflowexecution.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				we.Name = value.String
			}
		case workflowexecution.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				we.Description = value.String
			}
		case workflowexecution.FieldLabels:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field labels", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &we.Labels); err != nil {
					return fmt.Errorf("unmarshal field labels: %w", err)
				}
			}
		case workflowexecution.FieldAnnotations:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field annotations", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &we.Annotations); err != nil {
					return fmt.Errorf("unmarshal field annotations: %w", err)
				}
			}
		case workflowexecution.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				we.CreateTime = new(time.Time)
				*we.CreateTime = value.Time
			}
		case workflowexecution.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				we.UpdateTime = new(time.Time)
				*we.UpdateTime = value.Time
			}
		case workflowexecution.FieldStatus:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &we.Status); err != nil {
					return fmt.Errorf("unmarshal field status: %w", err)
				}
			}
		case workflowexecution.FieldProjectID:
			if value, ok := values[i].(*object.ID); !ok {
				return fmt.Errorf("unexpected type %T for field project_id", values[i])
			} else if value != nil {
				we.ProjectID = *value
			}
		case workflowexecution.FieldVersion:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				we.Version = int(value.Int64)
			}
		case workflowexecution.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				we.Type = value.String
			}
		case workflowexecution.FieldWorkflowID:
			if value, ok := values[i].(*object.ID); !ok {
				return fmt.Errorf("unexpected type %T for field workflow_id", values[i])
			} else if value != nil {
				we.WorkflowID = *value
			}
		case workflowexecution.FieldSubjectID:
			if value, ok := values[i].(*object.ID); !ok {
				return fmt.Errorf("unexpected type %T for field subject_id", values[i])
			} else if value != nil {
				we.SubjectID = *value
			}
		case workflowexecution.FieldExecuteTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field execute_time", values[i])
			} else if value.Valid {
				we.ExecuteTime = value.Time
			}
		case workflowexecution.FieldTimes:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field times", values[i])
			} else if value.Valid {
				we.Times = int(value.Int64)
			}
		case workflowexecution.FieldDuration:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field duration", values[i])
			} else if value.Valid {
				we.Duration = int(value.Int64)
			}
		case workflowexecution.FieldParallelism:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field parallelism", values[i])
			} else if value.Valid {
				we.Parallelism = int(value.Int64)
			}
		case workflowexecution.FieldTimeout:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field timeout", values[i])
			} else if value.Valid {
				we.Timeout = int(value.Int64)
			}
		case workflowexecution.FieldTrigger:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field trigger", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &we.Trigger); err != nil {
					return fmt.Errorf("unmarshal field trigger: %w", err)
				}
			}
		default:
			we.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the WorkflowExecution.
// This includes values selected through modifiers, order, etc.
func (we *WorkflowExecution) Value(name string) (ent.Value, error) {
	return we.selectValues.Get(name)
}

// QueryProject queries the "project" edge of the WorkflowExecution entity.
func (we *WorkflowExecution) QueryProject() *ProjectQuery {
	return NewWorkflowExecutionClient(we.config).QueryProject(we)
}

// QueryStages queries the "stages" edge of the WorkflowExecution entity.
func (we *WorkflowExecution) QueryStages() *WorkflowStageExecutionQuery {
	return NewWorkflowExecutionClient(we.config).QueryStages(we)
}

// QueryWorkflow queries the "workflow" edge of the WorkflowExecution entity.
func (we *WorkflowExecution) QueryWorkflow() *WorkflowQuery {
	return NewWorkflowExecutionClient(we.config).QueryWorkflow(we)
}

// Update returns a builder for updating this WorkflowExecution.
// Note that you need to call WorkflowExecution.Unwrap() before calling this method if this WorkflowExecution
// was returned from a transaction, and the transaction was committed or rolled back.
func (we *WorkflowExecution) Update() *WorkflowExecutionUpdateOne {
	return NewWorkflowExecutionClient(we.config).UpdateOne(we)
}

// Unwrap unwraps the WorkflowExecution entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (we *WorkflowExecution) Unwrap() *WorkflowExecution {
	_tx, ok := we.config.driver.(*txDriver)
	if !ok {
		panic("model: WorkflowExecution is not a transactional entity")
	}
	we.config.driver = _tx.drv
	return we
}

// String implements the fmt.Stringer.
func (we *WorkflowExecution) String() string {
	var builder strings.Builder
	builder.WriteString("WorkflowExecution(")
	builder.WriteString(fmt.Sprintf("id=%v, ", we.ID))
	builder.WriteString("name=")
	builder.WriteString(we.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(we.Description)
	builder.WriteString(", ")
	builder.WriteString("labels=")
	builder.WriteString(fmt.Sprintf("%v", we.Labels))
	builder.WriteString(", ")
	builder.WriteString("annotations=")
	builder.WriteString(fmt.Sprintf("%v", we.Annotations))
	builder.WriteString(", ")
	if v := we.CreateTime; v != nil {
		builder.WriteString("create_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := we.UpdateTime; v != nil {
		builder.WriteString("update_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", we.Status))
	builder.WriteString(", ")
	builder.WriteString("project_id=")
	builder.WriteString(fmt.Sprintf("%v", we.ProjectID))
	builder.WriteString(", ")
	builder.WriteString("version=")
	builder.WriteString(fmt.Sprintf("%v", we.Version))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(we.Type)
	builder.WriteString(", ")
	builder.WriteString("workflow_id=")
	builder.WriteString(fmt.Sprintf("%v", we.WorkflowID))
	builder.WriteString(", ")
	builder.WriteString("subject_id=")
	builder.WriteString(fmt.Sprintf("%v", we.SubjectID))
	builder.WriteString(", ")
	builder.WriteString("execute_time=")
	builder.WriteString(we.ExecuteTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("times=")
	builder.WriteString(fmt.Sprintf("%v", we.Times))
	builder.WriteString(", ")
	builder.WriteString("duration=")
	builder.WriteString(fmt.Sprintf("%v", we.Duration))
	builder.WriteString(", ")
	builder.WriteString("parallelism=")
	builder.WriteString(fmt.Sprintf("%v", we.Parallelism))
	builder.WriteString(", ")
	builder.WriteString("timeout=")
	builder.WriteString(fmt.Sprintf("%v", we.Timeout))
	builder.WriteString(", ")
	builder.WriteString("trigger=")
	builder.WriteString(fmt.Sprintf("%v", we.Trigger))
	builder.WriteByte(')')
	return builder.String()
}

// WorkflowExecutions is a parsable slice of WorkflowExecution.
type WorkflowExecutions []*WorkflowExecution
