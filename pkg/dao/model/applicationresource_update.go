// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// GENERATED, DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"github.com/seal-io/seal/pkg/dao/model/applicationresource"
	"github.com/seal-io/seal/pkg/dao/model/internal"
	"github.com/seal-io/seal/pkg/dao/model/predicate"
	"github.com/seal-io/seal/pkg/dao/types"
	"github.com/seal-io/seal/pkg/dao/types/oid"
)

// ApplicationResourceUpdate is the builder for updating ApplicationResource entities.
type ApplicationResourceUpdate struct {
	config
	hooks     []Hook
	mutation  *ApplicationResourceMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the ApplicationResourceUpdate builder.
func (aru *ApplicationResourceUpdate) Where(ps ...predicate.ApplicationResource) *ApplicationResourceUpdate {
	aru.mutation.Where(ps...)
	return aru
}

// SetUpdateTime sets the "updateTime" field.
func (aru *ApplicationResourceUpdate) SetUpdateTime(t time.Time) *ApplicationResourceUpdate {
	aru.mutation.SetUpdateTime(t)
	return aru
}

// SetStatus sets the "status" field.
func (aru *ApplicationResourceUpdate) SetStatus(trs types.ApplicationResourceStatus) *ApplicationResourceUpdate {
	aru.mutation.SetStatus(trs)
	return aru
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (aru *ApplicationResourceUpdate) SetNillableStatus(trs *types.ApplicationResourceStatus) *ApplicationResourceUpdate {
	if trs != nil {
		aru.SetStatus(*trs)
	}
	return aru
}

// ClearStatus clears the value of the "status" field.
func (aru *ApplicationResourceUpdate) ClearStatus() *ApplicationResourceUpdate {
	aru.mutation.ClearStatus()
	return aru
}

// AddComponentIDs adds the "components" edge to the ApplicationResource entity by IDs.
func (aru *ApplicationResourceUpdate) AddComponentIDs(ids ...oid.ID) *ApplicationResourceUpdate {
	aru.mutation.AddComponentIDs(ids...)
	return aru
}

// AddComponents adds the "components" edges to the ApplicationResource entity.
func (aru *ApplicationResourceUpdate) AddComponents(a ...*ApplicationResource) *ApplicationResourceUpdate {
	ids := make([]oid.ID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return aru.AddComponentIDs(ids...)
}

// Mutation returns the ApplicationResourceMutation object of the builder.
func (aru *ApplicationResourceUpdate) Mutation() *ApplicationResourceMutation {
	return aru.mutation
}

// ClearComponents clears all "components" edges to the ApplicationResource entity.
func (aru *ApplicationResourceUpdate) ClearComponents() *ApplicationResourceUpdate {
	aru.mutation.ClearComponents()
	return aru
}

// RemoveComponentIDs removes the "components" edge to ApplicationResource entities by IDs.
func (aru *ApplicationResourceUpdate) RemoveComponentIDs(ids ...oid.ID) *ApplicationResourceUpdate {
	aru.mutation.RemoveComponentIDs(ids...)
	return aru
}

// RemoveComponents removes "components" edges to ApplicationResource entities.
func (aru *ApplicationResourceUpdate) RemoveComponents(a ...*ApplicationResource) *ApplicationResourceUpdate {
	ids := make([]oid.ID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return aru.RemoveComponentIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (aru *ApplicationResourceUpdate) Save(ctx context.Context) (int, error) {
	if err := aru.defaults(); err != nil {
		return 0, err
	}
	return withHooks[int, ApplicationResourceMutation](ctx, aru.sqlSave, aru.mutation, aru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (aru *ApplicationResourceUpdate) SaveX(ctx context.Context) int {
	affected, err := aru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (aru *ApplicationResourceUpdate) Exec(ctx context.Context) error {
	_, err := aru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aru *ApplicationResourceUpdate) ExecX(ctx context.Context) {
	if err := aru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (aru *ApplicationResourceUpdate) defaults() error {
	if _, ok := aru.mutation.UpdateTime(); !ok {
		if applicationresource.UpdateDefaultUpdateTime == nil {
			return fmt.Errorf("model: uninitialized applicationresource.UpdateDefaultUpdateTime (forgotten import model/runtime?)")
		}
		v := applicationresource.UpdateDefaultUpdateTime()
		aru.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (aru *ApplicationResourceUpdate) check() error {
	if _, ok := aru.mutation.InstanceID(); aru.mutation.InstanceCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "ApplicationResource.instance"`)
	}
	if _, ok := aru.mutation.ConnectorID(); aru.mutation.ConnectorCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "ApplicationResource.connector"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (aru *ApplicationResourceUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ApplicationResourceUpdate {
	aru.modifiers = append(aru.modifiers, modifiers...)
	return aru
}

func (aru *ApplicationResourceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := aru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(applicationresource.Table, applicationresource.Columns, sqlgraph.NewFieldSpec(applicationresource.FieldID, field.TypeString))
	if ps := aru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aru.mutation.UpdateTime(); ok {
		_spec.SetField(applicationresource.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := aru.mutation.Status(); ok {
		_spec.SetField(applicationresource.FieldStatus, field.TypeJSON, value)
	}
	if aru.mutation.StatusCleared() {
		_spec.ClearField(applicationresource.FieldStatus, field.TypeJSON)
	}
	if aru.mutation.ComponentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   applicationresource.ComponentsTable,
			Columns: []string{applicationresource.ComponentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: applicationresource.FieldID,
				},
			},
		}
		edge.Schema = aru.schemaConfig.ApplicationResource
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aru.mutation.RemovedComponentsIDs(); len(nodes) > 0 && !aru.mutation.ComponentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   applicationresource.ComponentsTable,
			Columns: []string{applicationresource.ComponentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: applicationresource.FieldID,
				},
			},
		}
		edge.Schema = aru.schemaConfig.ApplicationResource
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aru.mutation.ComponentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   applicationresource.ComponentsTable,
			Columns: []string{applicationresource.ComponentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: applicationresource.FieldID,
				},
			},
		}
		edge.Schema = aru.schemaConfig.ApplicationResource
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = aru.schemaConfig.ApplicationResource
	ctx = internal.NewSchemaConfigContext(ctx, aru.schemaConfig)
	_spec.AddModifiers(aru.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, aru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{applicationresource.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	aru.mutation.done = true
	return n, nil
}

// ApplicationResourceUpdateOne is the builder for updating a single ApplicationResource entity.
type ApplicationResourceUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *ApplicationResourceMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdateTime sets the "updateTime" field.
func (aruo *ApplicationResourceUpdateOne) SetUpdateTime(t time.Time) *ApplicationResourceUpdateOne {
	aruo.mutation.SetUpdateTime(t)
	return aruo
}

// SetStatus sets the "status" field.
func (aruo *ApplicationResourceUpdateOne) SetStatus(trs types.ApplicationResourceStatus) *ApplicationResourceUpdateOne {
	aruo.mutation.SetStatus(trs)
	return aruo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (aruo *ApplicationResourceUpdateOne) SetNillableStatus(trs *types.ApplicationResourceStatus) *ApplicationResourceUpdateOne {
	if trs != nil {
		aruo.SetStatus(*trs)
	}
	return aruo
}

// ClearStatus clears the value of the "status" field.
func (aruo *ApplicationResourceUpdateOne) ClearStatus() *ApplicationResourceUpdateOne {
	aruo.mutation.ClearStatus()
	return aruo
}

// AddComponentIDs adds the "components" edge to the ApplicationResource entity by IDs.
func (aruo *ApplicationResourceUpdateOne) AddComponentIDs(ids ...oid.ID) *ApplicationResourceUpdateOne {
	aruo.mutation.AddComponentIDs(ids...)
	return aruo
}

// AddComponents adds the "components" edges to the ApplicationResource entity.
func (aruo *ApplicationResourceUpdateOne) AddComponents(a ...*ApplicationResource) *ApplicationResourceUpdateOne {
	ids := make([]oid.ID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return aruo.AddComponentIDs(ids...)
}

// Mutation returns the ApplicationResourceMutation object of the builder.
func (aruo *ApplicationResourceUpdateOne) Mutation() *ApplicationResourceMutation {
	return aruo.mutation
}

// ClearComponents clears all "components" edges to the ApplicationResource entity.
func (aruo *ApplicationResourceUpdateOne) ClearComponents() *ApplicationResourceUpdateOne {
	aruo.mutation.ClearComponents()
	return aruo
}

// RemoveComponentIDs removes the "components" edge to ApplicationResource entities by IDs.
func (aruo *ApplicationResourceUpdateOne) RemoveComponentIDs(ids ...oid.ID) *ApplicationResourceUpdateOne {
	aruo.mutation.RemoveComponentIDs(ids...)
	return aruo
}

// RemoveComponents removes "components" edges to ApplicationResource entities.
func (aruo *ApplicationResourceUpdateOne) RemoveComponents(a ...*ApplicationResource) *ApplicationResourceUpdateOne {
	ids := make([]oid.ID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return aruo.RemoveComponentIDs(ids...)
}

// Where appends a list predicates to the ApplicationResourceUpdate builder.
func (aruo *ApplicationResourceUpdateOne) Where(ps ...predicate.ApplicationResource) *ApplicationResourceUpdateOne {
	aruo.mutation.Where(ps...)
	return aruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (aruo *ApplicationResourceUpdateOne) Select(field string, fields ...string) *ApplicationResourceUpdateOne {
	aruo.fields = append([]string{field}, fields...)
	return aruo
}

// Save executes the query and returns the updated ApplicationResource entity.
func (aruo *ApplicationResourceUpdateOne) Save(ctx context.Context) (*ApplicationResource, error) {
	if err := aruo.defaults(); err != nil {
		return nil, err
	}
	return withHooks[*ApplicationResource, ApplicationResourceMutation](ctx, aruo.sqlSave, aruo.mutation, aruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (aruo *ApplicationResourceUpdateOne) SaveX(ctx context.Context) *ApplicationResource {
	node, err := aruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (aruo *ApplicationResourceUpdateOne) Exec(ctx context.Context) error {
	_, err := aruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aruo *ApplicationResourceUpdateOne) ExecX(ctx context.Context) {
	if err := aruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (aruo *ApplicationResourceUpdateOne) defaults() error {
	if _, ok := aruo.mutation.UpdateTime(); !ok {
		if applicationresource.UpdateDefaultUpdateTime == nil {
			return fmt.Errorf("model: uninitialized applicationresource.UpdateDefaultUpdateTime (forgotten import model/runtime?)")
		}
		v := applicationresource.UpdateDefaultUpdateTime()
		aruo.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (aruo *ApplicationResourceUpdateOne) check() error {
	if _, ok := aruo.mutation.InstanceID(); aruo.mutation.InstanceCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "ApplicationResource.instance"`)
	}
	if _, ok := aruo.mutation.ConnectorID(); aruo.mutation.ConnectorCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "ApplicationResource.connector"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (aruo *ApplicationResourceUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ApplicationResourceUpdateOne {
	aruo.modifiers = append(aruo.modifiers, modifiers...)
	return aruo
}

func (aruo *ApplicationResourceUpdateOne) sqlSave(ctx context.Context) (_node *ApplicationResource, err error) {
	if err := aruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(applicationresource.Table, applicationresource.Columns, sqlgraph.NewFieldSpec(applicationresource.FieldID, field.TypeString))
	id, ok := aruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`model: missing "ApplicationResource.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := aruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, applicationresource.FieldID)
		for _, f := range fields {
			if !applicationresource.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
			}
			if f != applicationresource.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := aruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aruo.mutation.UpdateTime(); ok {
		_spec.SetField(applicationresource.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := aruo.mutation.Status(); ok {
		_spec.SetField(applicationresource.FieldStatus, field.TypeJSON, value)
	}
	if aruo.mutation.StatusCleared() {
		_spec.ClearField(applicationresource.FieldStatus, field.TypeJSON)
	}
	if aruo.mutation.ComponentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   applicationresource.ComponentsTable,
			Columns: []string{applicationresource.ComponentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: applicationresource.FieldID,
				},
			},
		}
		edge.Schema = aruo.schemaConfig.ApplicationResource
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aruo.mutation.RemovedComponentsIDs(); len(nodes) > 0 && !aruo.mutation.ComponentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   applicationresource.ComponentsTable,
			Columns: []string{applicationresource.ComponentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: applicationresource.FieldID,
				},
			},
		}
		edge.Schema = aruo.schemaConfig.ApplicationResource
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aruo.mutation.ComponentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   applicationresource.ComponentsTable,
			Columns: []string{applicationresource.ComponentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: applicationresource.FieldID,
				},
			},
		}
		edge.Schema = aruo.schemaConfig.ApplicationResource
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = aruo.schemaConfig.ApplicationResource
	ctx = internal.NewSchemaConfigContext(ctx, aruo.schemaConfig)
	_spec.AddModifiers(aruo.modifiers...)
	_node = &ApplicationResource{config: aruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, aruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{applicationresource.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	aruo.mutation.done = true
	return _node, nil
}
