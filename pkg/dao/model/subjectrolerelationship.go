// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "ent". DO NOT EDIT.

package model

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"

	"github.com/seal-io/seal/pkg/dao/model/project"
	"github.com/seal-io/seal/pkg/dao/model/role"
	"github.com/seal-io/seal/pkg/dao/model/subject"
	"github.com/seal-io/seal/pkg/dao/model/subjectrolerelationship"
	"github.com/seal-io/seal/pkg/dao/types/oid"
)

// SubjectRoleRelationship is the model entity for the SubjectRoleRelationship schema.
type SubjectRoleRelationship struct {
	config `json:"-"`
	// ID of the ent.
	ID oid.ID `json:"id,omitempty" sql:"id"`
	// ID of the project to which the resource belongs, empty means using for global level.
	ProjectID oid.ID `json:"projectID,omitempty" sql:"projectID"`
	// Describe creation time.
	CreateTime *time.Time `json:"createTime,omitempty" sql:"createTime"`
	// ID of the subject to which the relationship connects.
	SubjectID oid.ID `json:"subjectID" sql:"subjectID"`
	// ID of the role to which the relationship connects.
	RoleID string `json:"roleID" sql:"roleID"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SubjectRoleRelationshipQuery when eager-loading is set.
	Edges        SubjectRoleRelationshipEdges `json:"edges,omitempty"`
	selectValues sql.SelectValues
}

// SubjectRoleRelationshipEdges holds the relations/edges for other nodes in the graph.
type SubjectRoleRelationshipEdges struct {
	// Project to which the subject role belongs.
	Project *Project `json:"project,omitempty" sql:"project"`
	// Subject that connect to the relationship.
	Subject *Subject `json:"subject,omitempty" sql:"subject"`
	// Role that connect to the relationship.
	Role *Role `json:"role,omitempty" sql:"role"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// ProjectOrErr returns the Project value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SubjectRoleRelationshipEdges) ProjectOrErr() (*Project, error) {
	if e.loadedTypes[0] {
		if e.Project == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: project.Label}
		}
		return e.Project, nil
	}
	return nil, &NotLoadedError{edge: "project"}
}

// SubjectOrErr returns the Subject value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SubjectRoleRelationshipEdges) SubjectOrErr() (*Subject, error) {
	if e.loadedTypes[1] {
		if e.Subject == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: subject.Label}
		}
		return e.Subject, nil
	}
	return nil, &NotLoadedError{edge: "subject"}
}

// RoleOrErr returns the Role value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SubjectRoleRelationshipEdges) RoleOrErr() (*Role, error) {
	if e.loadedTypes[2] {
		if e.Role == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: role.Label}
		}
		return e.Role, nil
	}
	return nil, &NotLoadedError{edge: "role"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SubjectRoleRelationship) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case subjectrolerelationship.FieldID, subjectrolerelationship.FieldProjectID, subjectrolerelationship.FieldSubjectID:
			values[i] = new(oid.ID)
		case subjectrolerelationship.FieldRoleID:
			values[i] = new(sql.NullString)
		case subjectrolerelationship.FieldCreateTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SubjectRoleRelationship fields.
func (srr *SubjectRoleRelationship) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case subjectrolerelationship.FieldID:
			if value, ok := values[i].(*oid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				srr.ID = *value
			}
		case subjectrolerelationship.FieldProjectID:
			if value, ok := values[i].(*oid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field projectID", values[i])
			} else if value != nil {
				srr.ProjectID = *value
			}
		case subjectrolerelationship.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createTime", values[i])
			} else if value.Valid {
				srr.CreateTime = new(time.Time)
				*srr.CreateTime = value.Time
			}
		case subjectrolerelationship.FieldSubjectID:
			if value, ok := values[i].(*oid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field subject_id", values[i])
			} else if value != nil {
				srr.SubjectID = *value
			}
		case subjectrolerelationship.FieldRoleID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field role_id", values[i])
			} else if value.Valid {
				srr.RoleID = value.String
			}
		default:
			srr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the SubjectRoleRelationship.
// This includes values selected through modifiers, order, etc.
func (srr *SubjectRoleRelationship) Value(name string) (ent.Value, error) {
	return srr.selectValues.Get(name)
}

// QueryProject queries the "project" edge of the SubjectRoleRelationship entity.
func (srr *SubjectRoleRelationship) QueryProject() *ProjectQuery {
	return NewSubjectRoleRelationshipClient(srr.config).QueryProject(srr)
}

// QuerySubject queries the "subject" edge of the SubjectRoleRelationship entity.
func (srr *SubjectRoleRelationship) QuerySubject() *SubjectQuery {
	return NewSubjectRoleRelationshipClient(srr.config).QuerySubject(srr)
}

// QueryRole queries the "role" edge of the SubjectRoleRelationship entity.
func (srr *SubjectRoleRelationship) QueryRole() *RoleQuery {
	return NewSubjectRoleRelationshipClient(srr.config).QueryRole(srr)
}

// Update returns a builder for updating this SubjectRoleRelationship.
// Note that you need to call SubjectRoleRelationship.Unwrap() before calling this method if this SubjectRoleRelationship
// was returned from a transaction, and the transaction was committed or rolled back.
func (srr *SubjectRoleRelationship) Update() *SubjectRoleRelationshipUpdateOne {
	return NewSubjectRoleRelationshipClient(srr.config).UpdateOne(srr)
}

// Unwrap unwraps the SubjectRoleRelationship entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (srr *SubjectRoleRelationship) Unwrap() *SubjectRoleRelationship {
	_tx, ok := srr.config.driver.(*txDriver)
	if !ok {
		panic("model: SubjectRoleRelationship is not a transactional entity")
	}
	srr.config.driver = _tx.drv
	return srr
}

// String implements the fmt.Stringer.
func (srr *SubjectRoleRelationship) String() string {
	var builder strings.Builder
	builder.WriteString("SubjectRoleRelationship(")
	builder.WriteString(fmt.Sprintf("id=%v, ", srr.ID))
	builder.WriteString("projectID=")
	builder.WriteString(fmt.Sprintf("%v", srr.ProjectID))
	builder.WriteString(", ")
	if v := srr.CreateTime; v != nil {
		builder.WriteString("createTime=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("subject_id=")
	builder.WriteString(fmt.Sprintf("%v", srr.SubjectID))
	builder.WriteString(", ")
	builder.WriteString("role_id=")
	builder.WriteString(srr.RoleID)
	builder.WriteByte(')')
	return builder.String()
}

// SubjectRoleRelationships is a parsable slice of SubjectRoleRelationship.
type SubjectRoleRelationships []*SubjectRoleRelationship
