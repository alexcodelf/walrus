// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// GENERATED, DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"github.com/seal-io/seal/pkg/dao/model/role"
	"github.com/seal-io/seal/pkg/dao/schema"
	"github.com/seal-io/seal/pkg/dao/types/oid"
)

// RoleCreate is the builder for creating a Role entity.
type RoleCreate struct {
	config
	mutation *RoleMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreateTime sets the "createTime" field.
func (rc *RoleCreate) SetCreateTime(t time.Time) *RoleCreate {
	rc.mutation.SetCreateTime(t)
	return rc
}

// SetNillableCreateTime sets the "createTime" field if the given value is not nil.
func (rc *RoleCreate) SetNillableCreateTime(t *time.Time) *RoleCreate {
	if t != nil {
		rc.SetCreateTime(*t)
	}
	return rc
}

// SetUpdateTime sets the "updateTime" field.
func (rc *RoleCreate) SetUpdateTime(t time.Time) *RoleCreate {
	rc.mutation.SetUpdateTime(t)
	return rc
}

// SetNillableUpdateTime sets the "updateTime" field if the given value is not nil.
func (rc *RoleCreate) SetNillableUpdateTime(t *time.Time) *RoleCreate {
	if t != nil {
		rc.SetUpdateTime(*t)
	}
	return rc
}

// SetDomain sets the "domain" field.
func (rc *RoleCreate) SetDomain(s string) *RoleCreate {
	rc.mutation.SetDomain(s)
	return rc
}

// SetNillableDomain sets the "domain" field if the given value is not nil.
func (rc *RoleCreate) SetNillableDomain(s *string) *RoleCreate {
	if s != nil {
		rc.SetDomain(*s)
	}
	return rc
}

// SetName sets the "name" field.
func (rc *RoleCreate) SetName(s string) *RoleCreate {
	rc.mutation.SetName(s)
	return rc
}

// SetDescription sets the "description" field.
func (rc *RoleCreate) SetDescription(s string) *RoleCreate {
	rc.mutation.SetDescription(s)
	return rc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (rc *RoleCreate) SetNillableDescription(s *string) *RoleCreate {
	if s != nil {
		rc.SetDescription(*s)
	}
	return rc
}

// SetPolicies sets the "policies" field.
func (rc *RoleCreate) SetPolicies(sp schema.RolePolicies) *RoleCreate {
	rc.mutation.SetPolicies(sp)
	return rc
}

// SetBuiltin sets the "builtin" field.
func (rc *RoleCreate) SetBuiltin(b bool) *RoleCreate {
	rc.mutation.SetBuiltin(b)
	return rc
}

// SetNillableBuiltin sets the "builtin" field if the given value is not nil.
func (rc *RoleCreate) SetNillableBuiltin(b *bool) *RoleCreate {
	if b != nil {
		rc.SetBuiltin(*b)
	}
	return rc
}

// SetSession sets the "session" field.
func (rc *RoleCreate) SetSession(b bool) *RoleCreate {
	rc.mutation.SetSession(b)
	return rc
}

// SetNillableSession sets the "session" field if the given value is not nil.
func (rc *RoleCreate) SetNillableSession(b *bool) *RoleCreate {
	if b != nil {
		rc.SetSession(*b)
	}
	return rc
}

// SetID sets the "id" field.
func (rc *RoleCreate) SetID(o oid.ID) *RoleCreate {
	rc.mutation.SetID(o)
	return rc
}

// Mutation returns the RoleMutation object of the builder.
func (rc *RoleCreate) Mutation() *RoleMutation {
	return rc.mutation
}

// Save creates the Role in the database.
func (rc *RoleCreate) Save(ctx context.Context) (*Role, error) {
	if err := rc.defaults(); err != nil {
		return nil, err
	}
	return withHooks[*Role, RoleMutation](ctx, rc.sqlSave, rc.mutation, rc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RoleCreate) SaveX(ctx context.Context) *Role {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *RoleCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *RoleCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rc *RoleCreate) defaults() error {
	if _, ok := rc.mutation.CreateTime(); !ok {
		if role.DefaultCreateTime == nil {
			return fmt.Errorf("model: uninitialized role.DefaultCreateTime (forgotten import model/runtime?)")
		}
		v := role.DefaultCreateTime()
		rc.mutation.SetCreateTime(v)
	}
	if _, ok := rc.mutation.UpdateTime(); !ok {
		if role.DefaultUpdateTime == nil {
			return fmt.Errorf("model: uninitialized role.DefaultUpdateTime (forgotten import model/runtime?)")
		}
		v := role.DefaultUpdateTime()
		rc.mutation.SetUpdateTime(v)
	}
	if _, ok := rc.mutation.Domain(); !ok {
		v := role.DefaultDomain
		rc.mutation.SetDomain(v)
	}
	if _, ok := rc.mutation.Policies(); !ok {
		v := role.DefaultPolicies
		rc.mutation.SetPolicies(v)
	}
	if _, ok := rc.mutation.Builtin(); !ok {
		v := role.DefaultBuiltin
		rc.mutation.SetBuiltin(v)
	}
	if _, ok := rc.mutation.Session(); !ok {
		v := role.DefaultSession
		rc.mutation.SetSession(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (rc *RoleCreate) check() error {
	if _, ok := rc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "createTime", err: errors.New(`model: missing required field "Role.createTime"`)}
	}
	if _, ok := rc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "updateTime", err: errors.New(`model: missing required field "Role.updateTime"`)}
	}
	if _, ok := rc.mutation.Domain(); !ok {
		return &ValidationError{Name: "domain", err: errors.New(`model: missing required field "Role.domain"`)}
	}
	if _, ok := rc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`model: missing required field "Role.name"`)}
	}
	if v, ok := rc.mutation.Name(); ok {
		if err := role.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`model: validator failed for field "Role.name": %w`, err)}
		}
	}
	if _, ok := rc.mutation.Policies(); !ok {
		return &ValidationError{Name: "policies", err: errors.New(`model: missing required field "Role.policies"`)}
	}
	if _, ok := rc.mutation.Builtin(); !ok {
		return &ValidationError{Name: "builtin", err: errors.New(`model: missing required field "Role.builtin"`)}
	}
	if _, ok := rc.mutation.Session(); !ok {
		return &ValidationError{Name: "session", err: errors.New(`model: missing required field "Role.session"`)}
	}
	return nil
}

func (rc *RoleCreate) sqlSave(ctx context.Context) (*Role, error) {
	if err := rc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*oid.ID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	rc.mutation.id = &_node.ID
	rc.mutation.done = true
	return _node, nil
}

func (rc *RoleCreate) createSpec() (*Role, *sqlgraph.CreateSpec) {
	var (
		_node = &Role{config: rc.config}
		_spec = sqlgraph.NewCreateSpec(role.Table, sqlgraph.NewFieldSpec(role.FieldID, field.TypeString))
	)
	_spec.Schema = rc.schemaConfig.Role
	_spec.OnConflict = rc.conflict
	if id, ok := rc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := rc.mutation.CreateTime(); ok {
		_spec.SetField(role.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = &value
	}
	if value, ok := rc.mutation.UpdateTime(); ok {
		_spec.SetField(role.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = &value
	}
	if value, ok := rc.mutation.Domain(); ok {
		_spec.SetField(role.FieldDomain, field.TypeString, value)
		_node.Domain = value
	}
	if value, ok := rc.mutation.Name(); ok {
		_spec.SetField(role.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := rc.mutation.Description(); ok {
		_spec.SetField(role.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := rc.mutation.Policies(); ok {
		_spec.SetField(role.FieldPolicies, field.TypeJSON, value)
		_node.Policies = value
	}
	if value, ok := rc.mutation.Builtin(); ok {
		_spec.SetField(role.FieldBuiltin, field.TypeBool, value)
		_node.Builtin = value
	}
	if value, ok := rc.mutation.Session(); ok {
		_spec.SetField(role.FieldSession, field.TypeBool, value)
		_node.Session = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Role.Create().
//		SetCreateTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.RoleUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (rc *RoleCreate) OnConflict(opts ...sql.ConflictOption) *RoleUpsertOne {
	rc.conflict = opts
	return &RoleUpsertOne{
		create: rc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Role.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (rc *RoleCreate) OnConflictColumns(columns ...string) *RoleUpsertOne {
	rc.conflict = append(rc.conflict, sql.ConflictColumns(columns...))
	return &RoleUpsertOne{
		create: rc,
	}
}

type (
	// RoleUpsertOne is the builder for "upsert"-ing
	//  one Role node.
	RoleUpsertOne struct {
		create *RoleCreate
	}

	// RoleUpsert is the "OnConflict" setter.
	RoleUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdateTime sets the "updateTime" field.
func (u *RoleUpsert) SetUpdateTime(v time.Time) *RoleUpsert {
	u.Set(role.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "updateTime" field to the value that was provided on create.
func (u *RoleUpsert) UpdateUpdateTime() *RoleUpsert {
	u.SetExcluded(role.FieldUpdateTime)
	return u
}

// SetDescription sets the "description" field.
func (u *RoleUpsert) SetDescription(v string) *RoleUpsert {
	u.Set(role.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *RoleUpsert) UpdateDescription() *RoleUpsert {
	u.SetExcluded(role.FieldDescription)
	return u
}

// ClearDescription clears the value of the "description" field.
func (u *RoleUpsert) ClearDescription() *RoleUpsert {
	u.SetNull(role.FieldDescription)
	return u
}

// SetPolicies sets the "policies" field.
func (u *RoleUpsert) SetPolicies(v schema.RolePolicies) *RoleUpsert {
	u.Set(role.FieldPolicies, v)
	return u
}

// UpdatePolicies sets the "policies" field to the value that was provided on create.
func (u *RoleUpsert) UpdatePolicies() *RoleUpsert {
	u.SetExcluded(role.FieldPolicies)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Role.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(role.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *RoleUpsertOne) UpdateNewValues() *RoleUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(role.FieldID)
		}
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(role.FieldCreateTime)
		}
		if _, exists := u.create.mutation.Domain(); exists {
			s.SetIgnore(role.FieldDomain)
		}
		if _, exists := u.create.mutation.Name(); exists {
			s.SetIgnore(role.FieldName)
		}
		if _, exists := u.create.mutation.Builtin(); exists {
			s.SetIgnore(role.FieldBuiltin)
		}
		if _, exists := u.create.mutation.Session(); exists {
			s.SetIgnore(role.FieldSession)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Role.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *RoleUpsertOne) Ignore() *RoleUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *RoleUpsertOne) DoNothing() *RoleUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the RoleCreate.OnConflict
// documentation for more info.
func (u *RoleUpsertOne) Update(set func(*RoleUpsert)) *RoleUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&RoleUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "updateTime" field.
func (u *RoleUpsertOne) SetUpdateTime(v time.Time) *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "updateTime" field to the value that was provided on create.
func (u *RoleUpsertOne) UpdateUpdateTime() *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetDescription sets the "description" field.
func (u *RoleUpsertOne) SetDescription(v string) *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *RoleUpsertOne) UpdateDescription() *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *RoleUpsertOne) ClearDescription() *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.ClearDescription()
	})
}

// SetPolicies sets the "policies" field.
func (u *RoleUpsertOne) SetPolicies(v schema.RolePolicies) *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.SetPolicies(v)
	})
}

// UpdatePolicies sets the "policies" field to the value that was provided on create.
func (u *RoleUpsertOne) UpdatePolicies() *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.UpdatePolicies()
	})
}

// Exec executes the query.
func (u *RoleUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("model: missing options for RoleCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *RoleUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *RoleUpsertOne) ID(ctx context.Context) (id oid.ID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("model: RoleUpsertOne.ID is not supported by MySQL driver. Use RoleUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *RoleUpsertOne) IDX(ctx context.Context) oid.ID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// RoleCreateBulk is the builder for creating many Role entities in bulk.
type RoleCreateBulk struct {
	config
	builders []*RoleCreate
	conflict []sql.ConflictOption
}

// Save creates the Role entities in the database.
func (rcb *RoleCreateBulk) Save(ctx context.Context) ([]*Role, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Role, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RoleMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = rcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *RoleCreateBulk) SaveX(ctx context.Context) []*Role {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *RoleCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *RoleCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Role.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.RoleUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (rcb *RoleCreateBulk) OnConflict(opts ...sql.ConflictOption) *RoleUpsertBulk {
	rcb.conflict = opts
	return &RoleUpsertBulk{
		create: rcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Role.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (rcb *RoleCreateBulk) OnConflictColumns(columns ...string) *RoleUpsertBulk {
	rcb.conflict = append(rcb.conflict, sql.ConflictColumns(columns...))
	return &RoleUpsertBulk{
		create: rcb,
	}
}

// RoleUpsertBulk is the builder for "upsert"-ing
// a bulk of Role nodes.
type RoleUpsertBulk struct {
	create *RoleCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Role.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(role.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *RoleUpsertBulk) UpdateNewValues() *RoleUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(role.FieldID)
			}
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(role.FieldCreateTime)
			}
			if _, exists := b.mutation.Domain(); exists {
				s.SetIgnore(role.FieldDomain)
			}
			if _, exists := b.mutation.Name(); exists {
				s.SetIgnore(role.FieldName)
			}
			if _, exists := b.mutation.Builtin(); exists {
				s.SetIgnore(role.FieldBuiltin)
			}
			if _, exists := b.mutation.Session(); exists {
				s.SetIgnore(role.FieldSession)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Role.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *RoleUpsertBulk) Ignore() *RoleUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *RoleUpsertBulk) DoNothing() *RoleUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the RoleCreateBulk.OnConflict
// documentation for more info.
func (u *RoleUpsertBulk) Update(set func(*RoleUpsert)) *RoleUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&RoleUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "updateTime" field.
func (u *RoleUpsertBulk) SetUpdateTime(v time.Time) *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "updateTime" field to the value that was provided on create.
func (u *RoleUpsertBulk) UpdateUpdateTime() *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetDescription sets the "description" field.
func (u *RoleUpsertBulk) SetDescription(v string) *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *RoleUpsertBulk) UpdateDescription() *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *RoleUpsertBulk) ClearDescription() *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.ClearDescription()
	})
}

// SetPolicies sets the "policies" field.
func (u *RoleUpsertBulk) SetPolicies(v schema.RolePolicies) *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.SetPolicies(v)
	})
}

// UpdatePolicies sets the "policies" field to the value that was provided on create.
func (u *RoleUpsertBulk) UpdatePolicies() *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.UpdatePolicies()
	})
}

// Exec executes the query.
func (u *RoleUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("model: OnConflict was set for builder %d. Set it on the RoleCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("model: missing options for RoleCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *RoleUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
