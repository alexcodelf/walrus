// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus". DO NOT EDIT.

package model

import (
	"context"
	stdsql "database/sql"
	"errors"
	"fmt"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"github.com/seal-io/walrus/pkg/dao/model/internal"
	"github.com/seal-io/walrus/pkg/dao/model/predicate"
	"github.com/seal-io/walrus/pkg/dao/model/resourcestate"
)

// ResourceStateUpdate is the builder for updating ResourceState entities.
type ResourceStateUpdate struct {
	config
	hooks     []Hook
	mutation  *ResourceStateMutation
	modifiers []func(*sql.UpdateBuilder)
	object    *ResourceState
}

// Where appends a list predicates to the ResourceStateUpdate builder.
func (rsu *ResourceStateUpdate) Where(ps ...predicate.ResourceState) *ResourceStateUpdate {
	rsu.mutation.Where(ps...)
	return rsu
}

// SetData sets the "data" field.
func (rsu *ResourceStateUpdate) SetData(s string) *ResourceStateUpdate {
	rsu.mutation.SetData(s)
	return rsu
}

// SetNillableData sets the "data" field if the given value is not nil.
func (rsu *ResourceStateUpdate) SetNillableData(s *string) *ResourceStateUpdate {
	if s != nil {
		rsu.SetData(*s)
	}
	return rsu
}

// Mutation returns the ResourceStateMutation object of the builder.
func (rsu *ResourceStateUpdate) Mutation() *ResourceStateMutation {
	return rsu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (rsu *ResourceStateUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, rsu.sqlSave, rsu.mutation, rsu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rsu *ResourceStateUpdate) SaveX(ctx context.Context) int {
	affected, err := rsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (rsu *ResourceStateUpdate) Exec(ctx context.Context) error {
	_, err := rsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rsu *ResourceStateUpdate) ExecX(ctx context.Context) {
	if err := rsu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rsu *ResourceStateUpdate) check() error {
	if _, ok := rsu.mutation.ResourceID(); rsu.mutation.ResourceCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "ResourceState.resource"`)
	}
	return nil
}

// Set is different from other Set* methods,
// it sets the value by judging the definition of each field within the entire object.
//
// For default fields, Set calls if the value is not zero.
//
// For no default but required fields, Set calls directly.
//
// For no default but optional fields, Set calls if the value is not zero,
// or clears if the value is zero.
//
// For example:
//
//	## Without Default
//
//	### Required
//
//	db.SetX(obj.X)
//
//	### Optional or Default
//
//	if _is_zero_value_(obj.X) {
//	   db.SetX(obj.X)
//	} else {
//	   db.ClearX()
//	}
//
//	## With Default
//
//	if _is_zero_value_(obj.X) {
//	   db.SetX(obj.X)
//	}
func (rsu *ResourceStateUpdate) Set(obj *ResourceState) *ResourceStateUpdate {
	// Without Default.
	rsu.SetData(obj.Data)

	// With Default.

	// Record the given object.
	rsu.object = obj

	return rsu
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (rsu *ResourceStateUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ResourceStateUpdate {
	rsu.modifiers = append(rsu.modifiers, modifiers...)
	return rsu
}

func (rsu *ResourceStateUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := rsu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(resourcestate.Table, resourcestate.Columns, sqlgraph.NewFieldSpec(resourcestate.FieldID, field.TypeString))
	if ps := rsu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rsu.mutation.Data(); ok {
		_spec.SetField(resourcestate.FieldData, field.TypeString, value)
	}
	_spec.Node.Schema = rsu.schemaConfig.ResourceState
	ctx = internal.NewSchemaConfigContext(ctx, rsu.schemaConfig)
	_spec.AddModifiers(rsu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, rsu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{resourcestate.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	rsu.mutation.done = true
	return n, nil
}

// ResourceStateUpdateOne is the builder for updating a single ResourceState entity.
type ResourceStateUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *ResourceStateMutation
	modifiers []func(*sql.UpdateBuilder)
	object    *ResourceState
}

// SetData sets the "data" field.
func (rsuo *ResourceStateUpdateOne) SetData(s string) *ResourceStateUpdateOne {
	rsuo.mutation.SetData(s)
	return rsuo
}

// SetNillableData sets the "data" field if the given value is not nil.
func (rsuo *ResourceStateUpdateOne) SetNillableData(s *string) *ResourceStateUpdateOne {
	if s != nil {
		rsuo.SetData(*s)
	}
	return rsuo
}

// Mutation returns the ResourceStateMutation object of the builder.
func (rsuo *ResourceStateUpdateOne) Mutation() *ResourceStateMutation {
	return rsuo.mutation
}

// Where appends a list predicates to the ResourceStateUpdate builder.
func (rsuo *ResourceStateUpdateOne) Where(ps ...predicate.ResourceState) *ResourceStateUpdateOne {
	rsuo.mutation.Where(ps...)
	return rsuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (rsuo *ResourceStateUpdateOne) Select(field string, fields ...string) *ResourceStateUpdateOne {
	rsuo.fields = append([]string{field}, fields...)
	return rsuo
}

// Save executes the query and returns the updated ResourceState entity.
func (rsuo *ResourceStateUpdateOne) Save(ctx context.Context) (*ResourceState, error) {
	return withHooks(ctx, rsuo.sqlSave, rsuo.mutation, rsuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rsuo *ResourceStateUpdateOne) SaveX(ctx context.Context) *ResourceState {
	node, err := rsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (rsuo *ResourceStateUpdateOne) Exec(ctx context.Context) error {
	_, err := rsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rsuo *ResourceStateUpdateOne) ExecX(ctx context.Context) {
	if err := rsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rsuo *ResourceStateUpdateOne) check() error {
	if _, ok := rsuo.mutation.ResourceID(); rsuo.mutation.ResourceCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "ResourceState.resource"`)
	}
	return nil
}

// Set is different from other Set* methods,
// it sets the value by judging the definition of each field within the entire object.
//
// For default fields, Set calls if the value changes from the original.
//
// For no default but required fields, Set calls if the value changes from the original.
//
// For no default but optional fields, Set calls if the value changes from the original,
// or clears if changes to zero.
//
// For example:
//
//	## Without Default
//
//	### Required
//
//	db.SetX(obj.X)
//
//	### Optional or Default
//
//	if _is_zero_value_(obj.X) {
//	   if _is_not_equal_(db.X, obj.X) {
//	      db.SetX(obj.X)
//	   }
//	} else {
//	   db.ClearX()
//	}
//
//	## With Default
//
//	if _is_zero_value_(obj.X) && _is_not_equal_(db.X, obj.X) {
//	   db.SetX(obj.X)
//	}
func (rsuo *ResourceStateUpdateOne) Set(obj *ResourceState) *ResourceStateUpdateOne {
	h := func(n ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			mt := m.(*ResourceStateMutation)
			db, err := mt.Client().ResourceState.Get(ctx, *mt.id)
			if err != nil {
				return nil, fmt.Errorf("failed getting ResourceState with id: %v", *mt.id)
			}

			// Without Default.
			if db.Data != obj.Data {
				rsuo.SetData(obj.Data)
			}

			// With Default.

			// Record the given object.
			rsuo.object = obj

			return n.Mutate(ctx, m)
		})
	}

	rsuo.hooks = append(rsuo.hooks, h)

	return rsuo
}

// getClientSet returns the ClientSet for the given builder.
func (rsuo *ResourceStateUpdateOne) getClientSet() (mc ClientSet) {
	if _, ok := rsuo.config.driver.(*txDriver); ok {
		tx := &Tx{config: rsuo.config}
		tx.init()
		mc = tx
	} else {
		cli := &Client{config: rsuo.config}
		cli.init()
		mc = cli
	}
	return mc
}

// SaveE calls the given function after updated the ResourceState entity,
// which is always good for cascading update operations.
func (rsuo *ResourceStateUpdateOne) SaveE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *ResourceState) error) (*ResourceState, error) {
	obj, err := rsuo.Save(ctx)
	if err != nil &&
		(rsuo.object == nil || !errors.Is(err, stdsql.ErrNoRows)) {
		return nil, err
	}

	if len(cbs) == 0 {
		return obj, err
	}

	mc := rsuo.getClientSet()

	if obj == nil {
		obj = rsuo.object
	} else if x := rsuo.object; x != nil {
		if _, set := rsuo.mutation.Field(resourcestate.FieldData); set {
			obj.Data = x.Data
		}
		obj.Edges = x.Edges
	}

	for i := range cbs {
		if err = cbs[i](ctx, mc, obj); err != nil {
			return nil, err
		}
	}

	return obj, nil
}

// SaveEX is like SaveE, but panics if an error occurs.
func (rsuo *ResourceStateUpdateOne) SaveEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *ResourceState) error) *ResourceState {
	obj, err := rsuo.SaveE(ctx, cbs...)
	if err != nil {
		panic(err)
	}
	return obj
}

// ExecE calls the given function after executed the query,
// which is always good for cascading update operations.
func (rsuo *ResourceStateUpdateOne) ExecE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *ResourceState) error) error {
	_, err := rsuo.SaveE(ctx, cbs...)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rsuo *ResourceStateUpdateOne) ExecEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *ResourceState) error) {
	if err := rsuo.ExecE(ctx, cbs...); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (rsuo *ResourceStateUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ResourceStateUpdateOne {
	rsuo.modifiers = append(rsuo.modifiers, modifiers...)
	return rsuo
}

func (rsuo *ResourceStateUpdateOne) sqlSave(ctx context.Context) (_node *ResourceState, err error) {
	if err := rsuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(resourcestate.Table, resourcestate.Columns, sqlgraph.NewFieldSpec(resourcestate.FieldID, field.TypeString))
	id, ok := rsuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`model: missing "ResourceState.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := rsuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, resourcestate.FieldID)
		for _, f := range fields {
			if !resourcestate.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
			}
			if f != resourcestate.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := rsuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rsuo.mutation.Data(); ok {
		_spec.SetField(resourcestate.FieldData, field.TypeString, value)
	}
	_spec.Node.Schema = rsuo.schemaConfig.ResourceState
	ctx = internal.NewSchemaConfigContext(ctx, rsuo.schemaConfig)
	_spec.AddModifiers(rsuo.modifiers...)
	_node = &ResourceState{config: rsuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, rsuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{resourcestate.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	rsuo.mutation.done = true
	return _node, nil
}