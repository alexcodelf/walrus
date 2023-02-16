// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// GENERATED, DO NOT EDIT.

package model

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"

	"github.com/seal-io/seal/pkg/dao/model/project"
	"github.com/seal-io/seal/pkg/dao/types"
)

// Project is the model entity for the Project schema.
type Project struct {
	config `json:"-"`
	// ID of the ent.
	ID types.ID `json:"id,omitempty"`
	// Name of the resource.
	Name string `json:"name"`
	// Description of the resource.
	Description string `json:"description,omitempty"`
	// Labels of the resource.
	Labels map[string]string `json:"labels,omitempty"`
	// Describe creation time.
	CreateTime *time.Time `json:"createTime,omitempty"`
	// Describe modification time.
	UpdateTime *time.Time `json:"updateTime,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProjectQuery when eager-loading is set.
	Edges ProjectEdges `json:"edges,omitempty"`
}

// ProjectEdges holds the relations/edges for other nodes in the graph.
type ProjectEdges struct {
	// Applications that belong to the project.
	Applications []*Application `json:"applications,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes       [1]bool
	namedApplications map[string][]*Application
}

// ApplicationsOrErr returns the Applications value or an error if the edge
// was not loaded in eager-loading.
func (e ProjectEdges) ApplicationsOrErr() ([]*Application, error) {
	if e.loadedTypes[0] {
		return e.Applications, nil
	}
	return nil, &NotLoadedError{edge: "applications"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Project) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case project.FieldLabels:
			values[i] = new([]byte)
		case project.FieldName, project.FieldDescription:
			values[i] = new(sql.NullString)
		case project.FieldCreateTime, project.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case project.FieldID:
			values[i] = new(types.ID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Project", columns[i])
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
			if value, ok := values[i].(*types.ID); !ok {
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
		case project.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createTime", values[i])
			} else if value.Valid {
				pr.CreateTime = new(time.Time)
				*pr.CreateTime = value.Time
			}
		case project.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updateTime", values[i])
			} else if value.Valid {
				pr.UpdateTime = new(time.Time)
				*pr.UpdateTime = value.Time
			}
		}
	}
	return nil
}

// QueryApplications queries the "applications" edge of the Project entity.
func (pr *Project) QueryApplications() *ApplicationQuery {
	return NewProjectClient(pr.config).QueryApplications(pr)
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
	if v := pr.CreateTime; v != nil {
		builder.WriteString("createTime=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := pr.UpdateTime; v != nil {
		builder.WriteString("updateTime=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// NamedApplications returns the Applications named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pr *Project) NamedApplications(name string) ([]*Application, error) {
	if pr.Edges.namedApplications == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pr.Edges.namedApplications[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pr *Project) appendNamedApplications(name string, edges ...*Application) {
	if pr.Edges.namedApplications == nil {
		pr.Edges.namedApplications = make(map[string][]*Application)
	}
	if len(edges) == 0 {
		pr.Edges.namedApplications[name] = []*Application{}
	} else {
		pr.Edges.namedApplications[name] = append(pr.Edges.namedApplications[name], edges...)
	}
}

// Projects is a parsable slice of Project.
type Projects []*Project

func (pr Projects) config(cfg config) {
	for _i := range pr {
		pr[_i].config = cfg
	}
}
