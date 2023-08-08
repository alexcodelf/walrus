// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "seal". DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"github.com/seal-io/seal/pkg/dao/model/distributelock"
)

// DistributeLockCreate is the builder for creating a DistributeLock entity.
type DistributeLockCreate struct {
	config
	mutation   *DistributeLockMutation
	hooks      []Hook
	conflict   []sql.ConflictOption
	object     *DistributeLock
	fromUpsert bool
}

// SetExpireAt sets the "expireAt" field.
func (dlc *DistributeLockCreate) SetExpireAt(i int64) *DistributeLockCreate {
	dlc.mutation.SetExpireAt(i)
	return dlc
}

// SetID sets the "id" field.
func (dlc *DistributeLockCreate) SetID(s string) *DistributeLockCreate {
	dlc.mutation.SetID(s)
	return dlc
}

// Mutation returns the DistributeLockMutation object of the builder.
func (dlc *DistributeLockCreate) Mutation() *DistributeLockMutation {
	return dlc.mutation
}

// Save creates the DistributeLock in the database.
func (dlc *DistributeLockCreate) Save(ctx context.Context) (*DistributeLock, error) {
	return withHooks(ctx, dlc.sqlSave, dlc.mutation, dlc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dlc *DistributeLockCreate) SaveX(ctx context.Context) *DistributeLock {
	v, err := dlc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dlc *DistributeLockCreate) Exec(ctx context.Context) error {
	_, err := dlc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dlc *DistributeLockCreate) ExecX(ctx context.Context) {
	if err := dlc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dlc *DistributeLockCreate) check() error {
	if _, ok := dlc.mutation.ExpireAt(); !ok {
		return &ValidationError{Name: "expireAt", err: errors.New(`model: missing required field "DistributeLock.expireAt"`)}
	}
	if v, ok := dlc.mutation.ID(); ok {
		if err := distributelock.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`model: validator failed for field "DistributeLock.id": %w`, err)}
		}
	}
	return nil
}

func (dlc *DistributeLockCreate) sqlSave(ctx context.Context) (*DistributeLock, error) {
	if err := dlc.check(); err != nil {
		return nil, err
	}
	_node, _spec := dlc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dlc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected DistributeLock.ID type: %T", _spec.ID.Value)
		}
	}
	dlc.mutation.id = &_node.ID
	dlc.mutation.done = true
	return _node, nil
}

func (dlc *DistributeLockCreate) createSpec() (*DistributeLock, *sqlgraph.CreateSpec) {
	var (
		_node = &DistributeLock{config: dlc.config}
		_spec = sqlgraph.NewCreateSpec(distributelock.Table, sqlgraph.NewFieldSpec(distributelock.FieldID, field.TypeString))
	)
	_spec.Schema = dlc.schemaConfig.DistributeLock
	_spec.OnConflict = dlc.conflict
	if id, ok := dlc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := dlc.mutation.ExpireAt(); ok {
		_spec.SetField(distributelock.FieldExpireAt, field.TypeInt64, value)
		_node.ExpireAt = value
	}
	return _node, _spec
}

// Set is different from other Set* methods,
// it sets the value by judging the definition of each field within the entire object.
//
// For required fields, Set calls directly.
//
// For optional fields, Set calls if the value is not zero.
//
// For example:
//
//	## Required
//
//	db.SetX(obj.X)
//
//	## Optional or Default
//
//	if _is_zero_value_(obj.X) {
//	   db.SetX(obj.X)
//	}
func (dlc *DistributeLockCreate) Set(obj *DistributeLock) *DistributeLockCreate {
	// Required.
	dlc.SetID(obj.ID)
	dlc.SetExpireAt(obj.ExpireAt)

	// Optional.

	// Record the given object.
	dlc.object = obj

	return dlc
}

// getClientSet returns the ClientSet for the given builder.
func (dlc *DistributeLockCreate) getClientSet() (mc ClientSet) {
	if _, ok := dlc.config.driver.(*txDriver); ok {
		tx := &Tx{config: dlc.config}
		tx.init()
		mc = tx
	} else {
		cli := &Client{config: dlc.config}
		cli.init()
		mc = cli
	}
	return mc
}

// SaveE calls the given function after created the DistributeLock entity,
// which is always good for cascading create operations.
func (dlc *DistributeLockCreate) SaveE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *DistributeLock) error) (*DistributeLock, error) {
	obj, err := dlc.Save(ctx)
	if err != nil {
		return nil, err
	}

	if len(cbs) == 0 {
		return obj, err
	}

	mc := dlc.getClientSet()

	if x := dlc.object; x != nil {
		if _, set := dlc.mutation.Field(distributelock.FieldExpireAt); set {
			obj.ExpireAt = x.ExpireAt
		}
	}

	for i := range cbs {
		if err = cbs[i](ctx, mc, obj); err != nil {
			return nil, err
		}
	}

	return obj, nil
}

// SaveEX is like SaveE, but panics if an error occurs.
func (dlc *DistributeLockCreate) SaveEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *DistributeLock) error) *DistributeLock {
	obj, err := dlc.SaveE(ctx, cbs...)
	if err != nil {
		panic(err)
	}
	return obj
}

// ExecE calls the given function after executed the query,
// which is always good for cascading create operations.
func (dlc *DistributeLockCreate) ExecE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *DistributeLock) error) error {
	_, err := dlc.SaveE(ctx, cbs...)
	return err
}

// ExecEX is like ExecE, but panics if an error occurs.
func (dlc *DistributeLockCreate) ExecEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *DistributeLock) error) {
	if err := dlc.ExecE(ctx, cbs...); err != nil {
		panic(err)
	}
}

// Set leverages the DistributeLockCreate Set method,
// it sets the value by judging the definition of each field within the entire item of the given list.
//
// For required fields, Set calls directly.
//
// For optional fields, Set calls if the value is not zero.
//
// For example:
//
//	## Required
//
//	db.SetX(obj.X)
//
//	## Optional or Default
//
//	if _is_zero_value_(obj.X) {
//	   db.SetX(obj.X)
//	}
func (dlcb *DistributeLockCreateBulk) Set(objs ...*DistributeLock) *DistributeLockCreateBulk {
	if len(objs) != 0 {
		client := NewDistributeLockClient(dlcb.config)

		dlcb.builders = make([]*DistributeLockCreate, len(objs))
		for i := range objs {
			dlcb.builders[i] = client.Create().Set(objs[i])
		}

		// Record the given objects.
		dlcb.objects = objs
	}

	return dlcb
}

// getClientSet returns the ClientSet for the given builder.
func (dlcb *DistributeLockCreateBulk) getClientSet() (mc ClientSet) {
	if _, ok := dlcb.config.driver.(*txDriver); ok {
		tx := &Tx{config: dlcb.config}
		tx.init()
		mc = tx
	} else {
		cli := &Client{config: dlcb.config}
		cli.init()
		mc = cli
	}
	return mc
}

// SaveE calls the given function after created the DistributeLock entities,
// which is always good for cascading create operations.
func (dlcb *DistributeLockCreateBulk) SaveE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *DistributeLock) error) ([]*DistributeLock, error) {
	objs, err := dlcb.Save(ctx)
	if err != nil {
		return nil, err
	}

	if len(cbs) == 0 {
		return objs, err
	}

	mc := dlcb.getClientSet()

	if x := dlcb.objects; x != nil {
		for i := range x {
			if _, set := dlcb.builders[i].mutation.Field(distributelock.FieldExpireAt); set {
				objs[i].ExpireAt = x[i].ExpireAt
			}
		}
	}

	for i := range objs {
		for j := range cbs {
			if err = cbs[j](ctx, mc, objs[i]); err != nil {
				return nil, err
			}
		}
	}

	return objs, nil
}

// SaveEX is like SaveE, but panics if an error occurs.
func (dlcb *DistributeLockCreateBulk) SaveEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *DistributeLock) error) []*DistributeLock {
	objs, err := dlcb.SaveE(ctx, cbs...)
	if err != nil {
		panic(err)
	}
	return objs
}

// ExecE calls the given function after executed the query,
// which is always good for cascading create operations.
func (dlcb *DistributeLockCreateBulk) ExecE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *DistributeLock) error) error {
	_, err := dlcb.SaveE(ctx, cbs...)
	return err
}

// ExecEX is like ExecE, but panics if an error occurs.
func (dlcb *DistributeLockCreateBulk) ExecEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *DistributeLock) error) {
	if err := dlcb.ExecE(ctx, cbs...); err != nil {
		panic(err)
	}
}

// ExecE calls the given function after executed the query,
// which is always good for cascading create operations.
func (u *DistributeLockUpsertOne) ExecE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *DistributeLock) error) error {
	if len(u.create.conflict) == 0 {
		return errors.New("model: missing options for DistributeLockUpsertOne.OnConflict")
	}
	u.create.fromUpsert = true
	return u.create.ExecE(ctx, cbs...)
}

// ExecEX is like ExecE, but panics if an error occurs.
func (u *DistributeLockUpsertOne) ExecEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *DistributeLock) error) {
	if err := u.ExecE(ctx, cbs...); err != nil {
		panic(err)
	}
}

// ExecE calls the given function after executed the query,
// which is always good for cascading create operations.
func (u *DistributeLockUpsertBulk) ExecE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *DistributeLock) error) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("model: OnConflict was set for builder %d. Set it on the DistributeLockUpsertBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("model: missing options for DistributeLockUpsertBulk.OnConflict")
	}
	u.create.fromUpsert = true
	return u.create.ExecE(ctx, cbs...)
}

// ExecEX is like ExecE, but panics if an error occurs.
func (u *DistributeLockUpsertBulk) ExecEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *DistributeLock) error) {
	if err := u.ExecE(ctx, cbs...); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.DistributeLock.Create().
//		SetExpireAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DistributeLockUpsert) {
//			SetExpireAt(v+v).
//		}).
//		Exec(ctx)
func (dlc *DistributeLockCreate) OnConflict(opts ...sql.ConflictOption) *DistributeLockUpsertOne {
	dlc.conflict = opts
	return &DistributeLockUpsertOne{
		create: dlc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DistributeLock.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (dlc *DistributeLockCreate) OnConflictColumns(columns ...string) *DistributeLockUpsertOne {
	dlc.conflict = append(dlc.conflict, sql.ConflictColumns(columns...))
	return &DistributeLockUpsertOne{
		create: dlc,
	}
}

type (
	// DistributeLockUpsertOne is the builder for "upsert"-ing
	//  one DistributeLock node.
	DistributeLockUpsertOne struct {
		create *DistributeLockCreate
	}

	// DistributeLockUpsert is the "OnConflict" setter.
	DistributeLockUpsert struct {
		*sql.UpdateSet
	}
)

// SetExpireAt sets the "expireAt" field.
func (u *DistributeLockUpsert) SetExpireAt(v int64) *DistributeLockUpsert {
	u.Set(distributelock.FieldExpireAt, v)
	return u
}

// UpdateExpireAt sets the "expireAt" field to the value that was provided on create.
func (u *DistributeLockUpsert) UpdateExpireAt() *DistributeLockUpsert {
	u.SetExcluded(distributelock.FieldExpireAt)
	return u
}

// AddExpireAt adds v to the "expireAt" field.
func (u *DistributeLockUpsert) AddExpireAt(v int64) *DistributeLockUpsert {
	u.Add(distributelock.FieldExpireAt, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.DistributeLock.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(distributelock.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *DistributeLockUpsertOne) UpdateNewValues() *DistributeLockUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(distributelock.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.DistributeLock.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *DistributeLockUpsertOne) Ignore() *DistributeLockUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DistributeLockUpsertOne) DoNothing() *DistributeLockUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DistributeLockCreate.OnConflict
// documentation for more info.
func (u *DistributeLockUpsertOne) Update(set func(*DistributeLockUpsert)) *DistributeLockUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DistributeLockUpsert{UpdateSet: update})
	}))
	return u
}

// SetExpireAt sets the "expireAt" field.
func (u *DistributeLockUpsertOne) SetExpireAt(v int64) *DistributeLockUpsertOne {
	return u.Update(func(s *DistributeLockUpsert) {
		s.SetExpireAt(v)
	})
}

// AddExpireAt adds v to the "expireAt" field.
func (u *DistributeLockUpsertOne) AddExpireAt(v int64) *DistributeLockUpsertOne {
	return u.Update(func(s *DistributeLockUpsert) {
		s.AddExpireAt(v)
	})
}

// UpdateExpireAt sets the "expireAt" field to the value that was provided on create.
func (u *DistributeLockUpsertOne) UpdateExpireAt() *DistributeLockUpsertOne {
	return u.Update(func(s *DistributeLockUpsert) {
		s.UpdateExpireAt()
	})
}

// Exec executes the query.
func (u *DistributeLockUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("model: missing options for DistributeLockCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DistributeLockUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *DistributeLockUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("model: DistributeLockUpsertOne.ID is not supported by MySQL driver. Use DistributeLockUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *DistributeLockUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// DistributeLockCreateBulk is the builder for creating many DistributeLock entities in bulk.
type DistributeLockCreateBulk struct {
	config
	builders   []*DistributeLockCreate
	conflict   []sql.ConflictOption
	objects    []*DistributeLock
	fromUpsert bool
}

// Save creates the DistributeLock entities in the database.
func (dlcb *DistributeLockCreateBulk) Save(ctx context.Context) ([]*DistributeLock, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dlcb.builders))
	nodes := make([]*DistributeLock, len(dlcb.builders))
	mutators := make([]Mutator, len(dlcb.builders))
	for i := range dlcb.builders {
		func(i int, root context.Context) {
			builder := dlcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DistributeLockMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, dlcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = dlcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dlcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, dlcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dlcb *DistributeLockCreateBulk) SaveX(ctx context.Context) []*DistributeLock {
	v, err := dlcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dlcb *DistributeLockCreateBulk) Exec(ctx context.Context) error {
	_, err := dlcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dlcb *DistributeLockCreateBulk) ExecX(ctx context.Context) {
	if err := dlcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.DistributeLock.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DistributeLockUpsert) {
//			SetExpireAt(v+v).
//		}).
//		Exec(ctx)
func (dlcb *DistributeLockCreateBulk) OnConflict(opts ...sql.ConflictOption) *DistributeLockUpsertBulk {
	dlcb.conflict = opts
	return &DistributeLockUpsertBulk{
		create: dlcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DistributeLock.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (dlcb *DistributeLockCreateBulk) OnConflictColumns(columns ...string) *DistributeLockUpsertBulk {
	dlcb.conflict = append(dlcb.conflict, sql.ConflictColumns(columns...))
	return &DistributeLockUpsertBulk{
		create: dlcb,
	}
}

// DistributeLockUpsertBulk is the builder for "upsert"-ing
// a bulk of DistributeLock nodes.
type DistributeLockUpsertBulk struct {
	create *DistributeLockCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.DistributeLock.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(distributelock.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *DistributeLockUpsertBulk) UpdateNewValues() *DistributeLockUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(distributelock.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.DistributeLock.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *DistributeLockUpsertBulk) Ignore() *DistributeLockUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DistributeLockUpsertBulk) DoNothing() *DistributeLockUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DistributeLockCreateBulk.OnConflict
// documentation for more info.
func (u *DistributeLockUpsertBulk) Update(set func(*DistributeLockUpsert)) *DistributeLockUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DistributeLockUpsert{UpdateSet: update})
	}))
	return u
}

// SetExpireAt sets the "expireAt" field.
func (u *DistributeLockUpsertBulk) SetExpireAt(v int64) *DistributeLockUpsertBulk {
	return u.Update(func(s *DistributeLockUpsert) {
		s.SetExpireAt(v)
	})
}

// AddExpireAt adds v to the "expireAt" field.
func (u *DistributeLockUpsertBulk) AddExpireAt(v int64) *DistributeLockUpsertBulk {
	return u.Update(func(s *DistributeLockUpsert) {
		s.AddExpireAt(v)
	})
}

// UpdateExpireAt sets the "expireAt" field to the value that was provided on create.
func (u *DistributeLockUpsertBulk) UpdateExpireAt() *DistributeLockUpsertBulk {
	return u.Update(func(s *DistributeLockUpsert) {
		s.UpdateExpireAt()
	})
}

// Exec executes the query.
func (u *DistributeLockUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("model: OnConflict was set for builder %d. Set it on the DistributeLockCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("model: missing options for DistributeLockCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DistributeLockUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}