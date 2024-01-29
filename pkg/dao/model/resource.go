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

	"github.com/seal-io/walrus/pkg/dao/model/environment"
	"github.com/seal-io/walrus/pkg/dao/model/project"
	"github.com/seal-io/walrus/pkg/dao/model/resource"
	"github.com/seal-io/walrus/pkg/dao/model/resourcedefinition"
	"github.com/seal-io/walrus/pkg/dao/model/resourcedefinitionmatchingrule"
	"github.com/seal-io/walrus/pkg/dao/model/resourcestate"
	"github.com/seal-io/walrus/pkg/dao/model/templateversion"
	"github.com/seal-io/walrus/pkg/dao/types"
	"github.com/seal-io/walrus/pkg/dao/types/object"
	"github.com/seal-io/walrus/pkg/dao/types/property"
	"github.com/seal-io/walrus/pkg/dao/types/status"
	"github.com/seal-io/walrus/utils/json"
)

// Resource is the model entity for the Resource schema.
type Resource struct {
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
	// ID of the environment to which the resource deploys.
	EnvironmentID object.ID `json:"environment_id,omitempty"`
	// ID of the template version to which the resource belong.
	TemplateID *object.ID `json:"template_id,omitempty"`
	// Type of the resource referring to a resource definition type.
	Type string `json:"type,omitempty"`
	// ID of the resource definition to which the resource use.
	ResourceDefinitionID *object.ID `json:"resource_definition_id,omitempty"`
	// ID of the resource definition matching rule to which the resource use.
	ResourceDefinitionMatchingRuleID *object.ID `json:"resource_definition_matching_rule_id,omitempty"`
	// Attributes to configure the template.
	Attributes property.Values `json:"attributes,omitempty"`
	// Computed attributes generated from attributes and schemas.
	ComputedAttributes property.Values `json:"computed_attributes,omitempty"`
	// Endpoints of the resource.
	Endpoints types.ResourceEndpoints `json:"endpoints,omitempty,cli-table-column"`
	// Change comment of the resource.
	ChangeComment string `json:"change_comment,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ResourceQuery when eager-loading is set.
	Edges        ResourceEdges `json:"edges,omitempty"`
	selectValues sql.SelectValues
}

// ResourceEdges holds the relations/edges for other nodes in the graph.
type ResourceEdges struct {
	// Project to which the resource belongs.
	Project *Project `json:"project,omitempty"`
	// Environment to which the resource belongs.
	Environment *Environment `json:"environment,omitempty"`
	// Template to which the resource belongs.
	Template *TemplateVersion `json:"template,omitempty"`
	// Definition of the resource.
	ResourceDefinition *ResourceDefinition `json:"-"`
	// Resource definition matching rule which the resource matches.
	ResourceDefinitionMatchingRule *ResourceDefinitionMatchingRule `json:"resource_definition_matching_rule,omitempty"`
	// Runs that belong to the resource.
	Runs []*ResourceRun `json:"runs,omitempty"`
	// Components that belong to the resource.
	Components []*ResourceComponent `json:"components,omitempty"`
	// Dependencies holds the value of the dependencies edge.
	Dependencies []*ResourceRelationship `json:"dependencies,omitempty"`
	// State of the resource.
	State *ResourceState `json:"state,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [9]bool
}

// ProjectOrErr returns the Project value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ResourceEdges) ProjectOrErr() (*Project, error) {
	if e.loadedTypes[0] {
		if e.Project == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: project.Label}
		}
		return e.Project, nil
	}
	return nil, &NotLoadedError{edge: "project"}
}

// EnvironmentOrErr returns the Environment value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ResourceEdges) EnvironmentOrErr() (*Environment, error) {
	if e.loadedTypes[1] {
		if e.Environment == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: environment.Label}
		}
		return e.Environment, nil
	}
	return nil, &NotLoadedError{edge: "environment"}
}

// TemplateOrErr returns the Template value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ResourceEdges) TemplateOrErr() (*TemplateVersion, error) {
	if e.loadedTypes[2] {
		if e.Template == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: templateversion.Label}
		}
		return e.Template, nil
	}
	return nil, &NotLoadedError{edge: "template"}
}

// ResourceDefinitionOrErr returns the ResourceDefinition value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ResourceEdges) ResourceDefinitionOrErr() (*ResourceDefinition, error) {
	if e.loadedTypes[3] {
		if e.ResourceDefinition == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: resourcedefinition.Label}
		}
		return e.ResourceDefinition, nil
	}
	return nil, &NotLoadedError{edge: "resource_definition"}
}

// ResourceDefinitionMatchingRuleOrErr returns the ResourceDefinitionMatchingRule value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ResourceEdges) ResourceDefinitionMatchingRuleOrErr() (*ResourceDefinitionMatchingRule, error) {
	if e.loadedTypes[4] {
		if e.ResourceDefinitionMatchingRule == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: resourcedefinitionmatchingrule.Label}
		}
		return e.ResourceDefinitionMatchingRule, nil
	}
	return nil, &NotLoadedError{edge: "resource_definition_matching_rule"}
}

// RunsOrErr returns the Runs value or an error if the edge
// was not loaded in eager-loading.
func (e ResourceEdges) RunsOrErr() ([]*ResourceRun, error) {
	if e.loadedTypes[5] {
		return e.Runs, nil
	}
	return nil, &NotLoadedError{edge: "runs"}
}

// ComponentsOrErr returns the Components value or an error if the edge
// was not loaded in eager-loading.
func (e ResourceEdges) ComponentsOrErr() ([]*ResourceComponent, error) {
	if e.loadedTypes[6] {
		return e.Components, nil
	}
	return nil, &NotLoadedError{edge: "components"}
}

// DependenciesOrErr returns the Dependencies value or an error if the edge
// was not loaded in eager-loading.
func (e ResourceEdges) DependenciesOrErr() ([]*ResourceRelationship, error) {
	if e.loadedTypes[7] {
		return e.Dependencies, nil
	}
	return nil, &NotLoadedError{edge: "dependencies"}
}

// StateOrErr returns the State value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ResourceEdges) StateOrErr() (*ResourceState, error) {
	if e.loadedTypes[8] {
		if e.State == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: resourcestate.Label}
		}
		return e.State, nil
	}
	return nil, &NotLoadedError{edge: "state"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Resource) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case resource.FieldTemplateID, resource.FieldResourceDefinitionID, resource.FieldResourceDefinitionMatchingRuleID:
			values[i] = &sql.NullScanner{S: new(object.ID)}
		case resource.FieldLabels, resource.FieldAnnotations, resource.FieldStatus, resource.FieldEndpoints:
			values[i] = new([]byte)
		case resource.FieldID, resource.FieldProjectID, resource.FieldEnvironmentID:
			values[i] = new(object.ID)
		case resource.FieldAttributes, resource.FieldComputedAttributes:
			values[i] = new(property.Values)
		case resource.FieldName, resource.FieldDescription, resource.FieldType, resource.FieldChangeComment:
			values[i] = new(sql.NullString)
		case resource.FieldCreateTime, resource.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Resource fields.
func (r *Resource) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case resource.FieldID:
			if value, ok := values[i].(*object.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				r.ID = *value
			}
		case resource.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				r.Name = value.String
			}
		case resource.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				r.Description = value.String
			}
		case resource.FieldLabels:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field labels", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &r.Labels); err != nil {
					return fmt.Errorf("unmarshal field labels: %w", err)
				}
			}
		case resource.FieldAnnotations:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field annotations", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &r.Annotations); err != nil {
					return fmt.Errorf("unmarshal field annotations: %w", err)
				}
			}
		case resource.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				r.CreateTime = new(time.Time)
				*r.CreateTime = value.Time
			}
		case resource.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				r.UpdateTime = new(time.Time)
				*r.UpdateTime = value.Time
			}
		case resource.FieldStatus:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &r.Status); err != nil {
					return fmt.Errorf("unmarshal field status: %w", err)
				}
			}
		case resource.FieldProjectID:
			if value, ok := values[i].(*object.ID); !ok {
				return fmt.Errorf("unexpected type %T for field project_id", values[i])
			} else if value != nil {
				r.ProjectID = *value
			}
		case resource.FieldEnvironmentID:
			if value, ok := values[i].(*object.ID); !ok {
				return fmt.Errorf("unexpected type %T for field environment_id", values[i])
			} else if value != nil {
				r.EnvironmentID = *value
			}
		case resource.FieldTemplateID:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field template_id", values[i])
			} else if value.Valid {
				r.TemplateID = new(object.ID)
				*r.TemplateID = *value.S.(*object.ID)
			}
		case resource.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				r.Type = value.String
			}
		case resource.FieldResourceDefinitionID:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field resource_definition_id", values[i])
			} else if value.Valid {
				r.ResourceDefinitionID = new(object.ID)
				*r.ResourceDefinitionID = *value.S.(*object.ID)
			}
		case resource.FieldResourceDefinitionMatchingRuleID:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field resource_definition_matching_rule_id", values[i])
			} else if value.Valid {
				r.ResourceDefinitionMatchingRuleID = new(object.ID)
				*r.ResourceDefinitionMatchingRuleID = *value.S.(*object.ID)
			}
		case resource.FieldAttributes:
			if value, ok := values[i].(*property.Values); !ok {
				return fmt.Errorf("unexpected type %T for field attributes", values[i])
			} else if value != nil {
				r.Attributes = *value
			}
		case resource.FieldComputedAttributes:
			if value, ok := values[i].(*property.Values); !ok {
				return fmt.Errorf("unexpected type %T for field computed_attributes", values[i])
			} else if value != nil {
				r.ComputedAttributes = *value
			}
		case resource.FieldEndpoints:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field endpoints", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &r.Endpoints); err != nil {
					return fmt.Errorf("unmarshal field endpoints: %w", err)
				}
			}
		case resource.FieldChangeComment:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field change_comment", values[i])
			} else if value.Valid {
				r.ChangeComment = value.String
			}
		default:
			r.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Resource.
// This includes values selected through modifiers, order, etc.
func (r *Resource) Value(name string) (ent.Value, error) {
	return r.selectValues.Get(name)
}

// QueryProject queries the "project" edge of the Resource entity.
func (r *Resource) QueryProject() *ProjectQuery {
	return NewResourceClient(r.config).QueryProject(r)
}

// QueryEnvironment queries the "environment" edge of the Resource entity.
func (r *Resource) QueryEnvironment() *EnvironmentQuery {
	return NewResourceClient(r.config).QueryEnvironment(r)
}

// QueryTemplate queries the "template" edge of the Resource entity.
func (r *Resource) QueryTemplate() *TemplateVersionQuery {
	return NewResourceClient(r.config).QueryTemplate(r)
}

// QueryResourceDefinition queries the "resource_definition" edge of the Resource entity.
func (r *Resource) QueryResourceDefinition() *ResourceDefinitionQuery {
	return NewResourceClient(r.config).QueryResourceDefinition(r)
}

// QueryResourceDefinitionMatchingRule queries the "resource_definition_matching_rule" edge of the Resource entity.
func (r *Resource) QueryResourceDefinitionMatchingRule() *ResourceDefinitionMatchingRuleQuery {
	return NewResourceClient(r.config).QueryResourceDefinitionMatchingRule(r)
}

// QueryRuns queries the "runs" edge of the Resource entity.
func (r *Resource) QueryRuns() *ResourceRunQuery {
	return NewResourceClient(r.config).QueryRuns(r)
}

// QueryComponents queries the "components" edge of the Resource entity.
func (r *Resource) QueryComponents() *ResourceComponentQuery {
	return NewResourceClient(r.config).QueryComponents(r)
}

// QueryDependencies queries the "dependencies" edge of the Resource entity.
func (r *Resource) QueryDependencies() *ResourceRelationshipQuery {
	return NewResourceClient(r.config).QueryDependencies(r)
}

// QueryState queries the "state" edge of the Resource entity.
func (r *Resource) QueryState() *ResourceStateQuery {
	return NewResourceClient(r.config).QueryState(r)
}

// Update returns a builder for updating this Resource.
// Note that you need to call Resource.Unwrap() before calling this method if this Resource
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Resource) Update() *ResourceUpdateOne {
	return NewResourceClient(r.config).UpdateOne(r)
}

// Unwrap unwraps the Resource entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Resource) Unwrap() *Resource {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("model: Resource is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Resource) String() string {
	var builder strings.Builder
	builder.WriteString("Resource(")
	builder.WriteString(fmt.Sprintf("id=%v, ", r.ID))
	builder.WriteString("name=")
	builder.WriteString(r.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(r.Description)
	builder.WriteString(", ")
	builder.WriteString("labels=")
	builder.WriteString(fmt.Sprintf("%v", r.Labels))
	builder.WriteString(", ")
	builder.WriteString("annotations=")
	builder.WriteString(fmt.Sprintf("%v", r.Annotations))
	builder.WriteString(", ")
	if v := r.CreateTime; v != nil {
		builder.WriteString("create_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := r.UpdateTime; v != nil {
		builder.WriteString("update_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", r.Status))
	builder.WriteString(", ")
	builder.WriteString("project_id=")
	builder.WriteString(fmt.Sprintf("%v", r.ProjectID))
	builder.WriteString(", ")
	builder.WriteString("environment_id=")
	builder.WriteString(fmt.Sprintf("%v", r.EnvironmentID))
	builder.WriteString(", ")
	if v := r.TemplateID; v != nil {
		builder.WriteString("template_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(r.Type)
	builder.WriteString(", ")
	if v := r.ResourceDefinitionID; v != nil {
		builder.WriteString("resource_definition_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := r.ResourceDefinitionMatchingRuleID; v != nil {
		builder.WriteString("resource_definition_matching_rule_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("attributes=")
	builder.WriteString(fmt.Sprintf("%v", r.Attributes))
	builder.WriteString(", ")
	builder.WriteString("computed_attributes=")
	builder.WriteString(fmt.Sprintf("%v", r.ComputedAttributes))
	builder.WriteString(", ")
	builder.WriteString("endpoints=")
	builder.WriteString(fmt.Sprintf("%v", r.Endpoints))
	builder.WriteString(", ")
	builder.WriteString("change_comment=")
	builder.WriteString(r.ChangeComment)
	builder.WriteByte(')')
	return builder.String()
}

// Resources is a parsable slice of Resource.
type Resources []*Resource
