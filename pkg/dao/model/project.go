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
	"github.com/seal-io/walrus/pkg/dao/types/object"
	"github.com/seal-io/walrus/utils/json"
)

// Project is the model entity for the Project schema.
type Project struct {
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
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProjectQuery when eager-loading is set.
	Edges        ProjectEdges `json:"edges,omitempty"`
	selectValues sql.SelectValues
}

// ProjectEdges holds the relations/edges for other nodes in the graph.
type ProjectEdges struct {
	// Environments that belong to the project.
	Environments []*Environment `json:"environments,omitempty"`
	// Connectors that belong to the project.
	Connectors []*Connector `json:"connectors,omitempty"`
	// Roles of a subject that belong to the project.
	SubjectRoles []*SubjectRoleRelationship `json:"subject_roles,omitempty"`
	// Services that belong to the project.
	Services []*Service `json:"services,omitempty"`
	// ServiceResources that belong to the project.
	ServiceResources []*ServiceResource `json:"service_resources,omitempty"`
	// ServiceRevisions that belong to the project.
	ServiceRevisions []*ServiceRevision `json:"service_revisions,omitempty"`
	// Variables that belong to the project.
	Variables []*Variable `json:"variables,omitempty"`
	// Workflows that belong to the project.
	Workflows []*Workflow `json:"workflows,omitempty"`
	// WorkflowStages that belong to the project.
	WorkflowStages []*WorkflowStage `json:"workflow_stages,omitempty"`
	// WorkflowSteps that belong to the project.
	WorkflowSteps []*WorkflowStep `json:"workflow_steps,omitempty"`
	// WorkflowExecutions that belong to the project.
	WorkflowExecutions []*WorkflowExecution `json:"workflow_executions,omitempty"`
	// WorkflowStageExecutions that belong to the project.
	WorkflowStageExecutions []*WorkflowStageExecution `json:"workflow_stage_executions,omitempty"`
	// WorkflowStepExecutions that belong to the project.
	WorkflowStepExecutions []*WorkflowStepExecution `json:"workflow_step_executions,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [13]bool
}

// EnvironmentsOrErr returns the Environments value or an error if the edge
// was not loaded in eager-loading.
func (e ProjectEdges) EnvironmentsOrErr() ([]*Environment, error) {
	if e.loadedTypes[0] {
		return e.Environments, nil
	}
	return nil, &NotLoadedError{edge: "environments"}
}

// ConnectorsOrErr returns the Connectors value or an error if the edge
// was not loaded in eager-loading.
func (e ProjectEdges) ConnectorsOrErr() ([]*Connector, error) {
	if e.loadedTypes[1] {
		return e.Connectors, nil
	}
	return nil, &NotLoadedError{edge: "connectors"}
}

// SubjectRolesOrErr returns the SubjectRoles value or an error if the edge
// was not loaded in eager-loading.
func (e ProjectEdges) SubjectRolesOrErr() ([]*SubjectRoleRelationship, error) {
	if e.loadedTypes[2] {
		return e.SubjectRoles, nil
	}
	return nil, &NotLoadedError{edge: "subject_roles"}
}

// ServicesOrErr returns the Services value or an error if the edge
// was not loaded in eager-loading.
func (e ProjectEdges) ServicesOrErr() ([]*Service, error) {
	if e.loadedTypes[3] {
		return e.Services, nil
	}
	return nil, &NotLoadedError{edge: "services"}
}

// ServiceResourcesOrErr returns the ServiceResources value or an error if the edge
// was not loaded in eager-loading.
func (e ProjectEdges) ServiceResourcesOrErr() ([]*ServiceResource, error) {
	if e.loadedTypes[4] {
		return e.ServiceResources, nil
	}
	return nil, &NotLoadedError{edge: "service_resources"}
}

// ServiceRevisionsOrErr returns the ServiceRevisions value or an error if the edge
// was not loaded in eager-loading.
func (e ProjectEdges) ServiceRevisionsOrErr() ([]*ServiceRevision, error) {
	if e.loadedTypes[5] {
		return e.ServiceRevisions, nil
	}
	return nil, &NotLoadedError{edge: "service_revisions"}
}

// VariablesOrErr returns the Variables value or an error if the edge
// was not loaded in eager-loading.
func (e ProjectEdges) VariablesOrErr() ([]*Variable, error) {
	if e.loadedTypes[6] {
		return e.Variables, nil
	}
	return nil, &NotLoadedError{edge: "variables"}
}

// WorkflowsOrErr returns the Workflows value or an error if the edge
// was not loaded in eager-loading.
func (e ProjectEdges) WorkflowsOrErr() ([]*Workflow, error) {
	if e.loadedTypes[7] {
		return e.Workflows, nil
	}
	return nil, &NotLoadedError{edge: "workflows"}
}

// WorkflowStagesOrErr returns the WorkflowStages value or an error if the edge
// was not loaded in eager-loading.
func (e ProjectEdges) WorkflowStagesOrErr() ([]*WorkflowStage, error) {
	if e.loadedTypes[8] {
		return e.WorkflowStages, nil
	}
	return nil, &NotLoadedError{edge: "workflow_stages"}
}

// WorkflowStepsOrErr returns the WorkflowSteps value or an error if the edge
// was not loaded in eager-loading.
func (e ProjectEdges) WorkflowStepsOrErr() ([]*WorkflowStep, error) {
	if e.loadedTypes[9] {
		return e.WorkflowSteps, nil
	}
	return nil, &NotLoadedError{edge: "workflow_steps"}
}

// WorkflowExecutionsOrErr returns the WorkflowExecutions value or an error if the edge
// was not loaded in eager-loading.
func (e ProjectEdges) WorkflowExecutionsOrErr() ([]*WorkflowExecution, error) {
	if e.loadedTypes[10] {
		return e.WorkflowExecutions, nil
	}
	return nil, &NotLoadedError{edge: "workflow_executions"}
}

// WorkflowStageExecutionsOrErr returns the WorkflowStageExecutions value or an error if the edge
// was not loaded in eager-loading.
func (e ProjectEdges) WorkflowStageExecutionsOrErr() ([]*WorkflowStageExecution, error) {
	if e.loadedTypes[11] {
		return e.WorkflowStageExecutions, nil
	}
	return nil, &NotLoadedError{edge: "workflow_stage_executions"}
}

// WorkflowStepExecutionsOrErr returns the WorkflowStepExecutions value or an error if the edge
// was not loaded in eager-loading.
func (e ProjectEdges) WorkflowStepExecutionsOrErr() ([]*WorkflowStepExecution, error) {
	if e.loadedTypes[12] {
		return e.WorkflowStepExecutions, nil
	}
	return nil, &NotLoadedError{edge: "workflow_step_executions"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Project) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case project.FieldLabels, project.FieldAnnotations:
			values[i] = new([]byte)
		case project.FieldID:
			values[i] = new(object.ID)
		case project.FieldName, project.FieldDescription:
			values[i] = new(sql.NullString)
		case project.FieldCreateTime, project.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Project fields.
func (pr *Project) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case project.FieldID:
			if value, ok := values[i].(*object.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				pr.ID = *value
			}
		case project.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pr.Name = value.String
			}
		case project.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				pr.Description = value.String
			}
		case project.FieldLabels:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field labels", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pr.Labels); err != nil {
					return fmt.Errorf("unmarshal field labels: %w", err)
				}
			}
		case project.FieldAnnotations:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field annotations", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pr.Annotations); err != nil {
					return fmt.Errorf("unmarshal field annotations: %w", err)
				}
			}
		case project.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				pr.CreateTime = new(time.Time)
				*pr.CreateTime = value.Time
			}
		case project.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				pr.UpdateTime = new(time.Time)
				*pr.UpdateTime = value.Time
			}
		default:
			pr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Project.
// This includes values selected through modifiers, order, etc.
func (pr *Project) Value(name string) (ent.Value, error) {
	return pr.selectValues.Get(name)
}

// QueryEnvironments queries the "environments" edge of the Project entity.
func (pr *Project) QueryEnvironments() *EnvironmentQuery {
	return NewProjectClient(pr.config).QueryEnvironments(pr)
}

// QueryConnectors queries the "connectors" edge of the Project entity.
func (pr *Project) QueryConnectors() *ConnectorQuery {
	return NewProjectClient(pr.config).QueryConnectors(pr)
}

// QuerySubjectRoles queries the "subject_roles" edge of the Project entity.
func (pr *Project) QuerySubjectRoles() *SubjectRoleRelationshipQuery {
	return NewProjectClient(pr.config).QuerySubjectRoles(pr)
}

// QueryServices queries the "services" edge of the Project entity.
func (pr *Project) QueryServices() *ServiceQuery {
	return NewProjectClient(pr.config).QueryServices(pr)
}

// QueryServiceResources queries the "service_resources" edge of the Project entity.
func (pr *Project) QueryServiceResources() *ServiceResourceQuery {
	return NewProjectClient(pr.config).QueryServiceResources(pr)
}

// QueryServiceRevisions queries the "service_revisions" edge of the Project entity.
func (pr *Project) QueryServiceRevisions() *ServiceRevisionQuery {
	return NewProjectClient(pr.config).QueryServiceRevisions(pr)
}

// QueryVariables queries the "variables" edge of the Project entity.
func (pr *Project) QueryVariables() *VariableQuery {
	return NewProjectClient(pr.config).QueryVariables(pr)
}

// QueryWorkflows queries the "workflows" edge of the Project entity.
func (pr *Project) QueryWorkflows() *WorkflowQuery {
	return NewProjectClient(pr.config).QueryWorkflows(pr)
}

// QueryWorkflowStages queries the "workflow_stages" edge of the Project entity.
func (pr *Project) QueryWorkflowStages() *WorkflowStageQuery {
	return NewProjectClient(pr.config).QueryWorkflowStages(pr)
}

// QueryWorkflowSteps queries the "workflow_steps" edge of the Project entity.
func (pr *Project) QueryWorkflowSteps() *WorkflowStepQuery {
	return NewProjectClient(pr.config).QueryWorkflowSteps(pr)
}

// QueryWorkflowExecutions queries the "workflow_executions" edge of the Project entity.
func (pr *Project) QueryWorkflowExecutions() *WorkflowExecutionQuery {
	return NewProjectClient(pr.config).QueryWorkflowExecutions(pr)
}

// QueryWorkflowStageExecutions queries the "workflow_stage_executions" edge of the Project entity.
func (pr *Project) QueryWorkflowStageExecutions() *WorkflowStageExecutionQuery {
	return NewProjectClient(pr.config).QueryWorkflowStageExecutions(pr)
}

// QueryWorkflowStepExecutions queries the "workflow_step_executions" edge of the Project entity.
func (pr *Project) QueryWorkflowStepExecutions() *WorkflowStepExecutionQuery {
	return NewProjectClient(pr.config).QueryWorkflowStepExecutions(pr)
}

// Update returns a builder for updating this Project.
// Note that you need to call Project.Unwrap() before calling this method if this Project
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Project) Update() *ProjectUpdateOne {
	return NewProjectClient(pr.config).UpdateOne(pr)
}

// Unwrap unwraps the Project entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Project) Unwrap() *Project {
	_tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("model: Project is not a transactional entity")
	}
	pr.config.driver = _tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Project) String() string {
	var builder strings.Builder
	builder.WriteString("Project(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pr.ID))
	builder.WriteString("name=")
	builder.WriteString(pr.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(pr.Description)
	builder.WriteString(", ")
	builder.WriteString("labels=")
	builder.WriteString(fmt.Sprintf("%v", pr.Labels))
	builder.WriteString(", ")
	builder.WriteString("annotations=")
	builder.WriteString(fmt.Sprintf("%v", pr.Annotations))
	builder.WriteString(", ")
	if v := pr.CreateTime; v != nil {
		builder.WriteString("create_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := pr.UpdateTime; v != nil {
		builder.WriteString("update_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Projects is a parsable slice of Project.
type Projects []*Project
