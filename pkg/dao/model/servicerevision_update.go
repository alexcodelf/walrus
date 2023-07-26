// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus". DO NOT EDIT.

package model

import (
	"context"
	stdsql "database/sql"
	"errors"
	"fmt"
	"reflect"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"

	"github.com/seal-io/walrus/pkg/dao/model/internal"
	"github.com/seal-io/walrus/pkg/dao/model/predicate"
	"github.com/seal-io/walrus/pkg/dao/model/servicerevision"
	"github.com/seal-io/walrus/pkg/dao/types"
	"github.com/seal-io/walrus/pkg/dao/types/crypto"
	"github.com/seal-io/walrus/pkg/dao/types/property"
)

// ServiceRevisionUpdate is the builder for updating ServiceRevision entities.
type ServiceRevisionUpdate struct {
	config
	hooks     []Hook
	mutation  *ServiceRevisionMutation
	modifiers []func(*sql.UpdateBuilder)
	object    *ServiceRevision
}

// Where appends a list predicates to the ServiceRevisionUpdate builder.
func (sru *ServiceRevisionUpdate) Where(ps ...predicate.ServiceRevision) *ServiceRevisionUpdate {
	sru.mutation.Where(ps...)
	return sru
}

// SetStatus sets the "status" field.
func (sru *ServiceRevisionUpdate) SetStatus(s string) *ServiceRevisionUpdate {
	sru.mutation.SetStatus(s)
	return sru
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (sru *ServiceRevisionUpdate) SetNillableStatus(s *string) *ServiceRevisionUpdate {
	if s != nil {
		sru.SetStatus(*s)
	}
	return sru
}

// ClearStatus clears the value of the "status" field.
func (sru *ServiceRevisionUpdate) ClearStatus() *ServiceRevisionUpdate {
	sru.mutation.ClearStatus()
	return sru
}

// SetStatusMessage sets the "status_message" field.
func (sru *ServiceRevisionUpdate) SetStatusMessage(s string) *ServiceRevisionUpdate {
	sru.mutation.SetStatusMessage(s)
	return sru
}

// SetNillableStatusMessage sets the "status_message" field if the given value is not nil.
func (sru *ServiceRevisionUpdate) SetNillableStatusMessage(s *string) *ServiceRevisionUpdate {
	if s != nil {
		sru.SetStatusMessage(*s)
	}
	return sru
}

// ClearStatusMessage clears the value of the "status_message" field.
func (sru *ServiceRevisionUpdate) ClearStatusMessage() *ServiceRevisionUpdate {
	sru.mutation.ClearStatusMessage()
	return sru
}

// SetType sets the "type" field.
func (sru *ServiceRevisionUpdate) SetType(s string) *ServiceRevisionUpdate {
	sru.mutation.SetType(s)
	return sru
}

// SetTemplateVersion sets the "template_version" field.
func (sru *ServiceRevisionUpdate) SetTemplateVersion(s string) *ServiceRevisionUpdate {
	sru.mutation.SetTemplateVersion(s)
	return sru
}

// SetAttributes sets the "attributes" field.
func (sru *ServiceRevisionUpdate) SetAttributes(pr property.Values) *ServiceRevisionUpdate {
	sru.mutation.SetAttributes(pr)
	return sru
}

// ClearAttributes clears the value of the "attributes" field.
func (sru *ServiceRevisionUpdate) ClearAttributes() *ServiceRevisionUpdate {
	sru.mutation.ClearAttributes()
	return sru
}

// SetVariables sets the "variables" field.
func (sru *ServiceRevisionUpdate) SetVariables(c crypto.Map[string, string]) *ServiceRevisionUpdate {
	sru.mutation.SetVariables(c)
	return sru
}

// SetInputPlan sets the "input_plan" field.
func (sru *ServiceRevisionUpdate) SetInputPlan(s string) *ServiceRevisionUpdate {
	sru.mutation.SetInputPlan(s)
	return sru
}

// SetOutput sets the "output" field.
func (sru *ServiceRevisionUpdate) SetOutput(s string) *ServiceRevisionUpdate {
	sru.mutation.SetOutput(s)
	return sru
}

// SetDeployerType sets the "deployer_type" field.
func (sru *ServiceRevisionUpdate) SetDeployerType(s string) *ServiceRevisionUpdate {
	sru.mutation.SetDeployerType(s)
	return sru
}

// SetNillableDeployerType sets the "deployer_type" field if the given value is not nil.
func (sru *ServiceRevisionUpdate) SetNillableDeployerType(s *string) *ServiceRevisionUpdate {
	if s != nil {
		sru.SetDeployerType(*s)
	}
	return sru
}

// SetDuration sets the "duration" field.
func (sru *ServiceRevisionUpdate) SetDuration(i int) *ServiceRevisionUpdate {
	sru.mutation.ResetDuration()
	sru.mutation.SetDuration(i)
	return sru
}

// SetNillableDuration sets the "duration" field if the given value is not nil.
func (sru *ServiceRevisionUpdate) SetNillableDuration(i *int) *ServiceRevisionUpdate {
	if i != nil {
		sru.SetDuration(*i)
	}
	return sru
}

// AddDuration adds i to the "duration" field.
func (sru *ServiceRevisionUpdate) AddDuration(i int) *ServiceRevisionUpdate {
	sru.mutation.AddDuration(i)
	return sru
}

// SetPreviousRequiredProviders sets the "previous_required_providers" field.
func (sru *ServiceRevisionUpdate) SetPreviousRequiredProviders(tr []types.ProviderRequirement) *ServiceRevisionUpdate {
	sru.mutation.SetPreviousRequiredProviders(tr)
	return sru
}

// AppendPreviousRequiredProviders appends tr to the "previous_required_providers" field.
func (sru *ServiceRevisionUpdate) AppendPreviousRequiredProviders(tr []types.ProviderRequirement) *ServiceRevisionUpdate {
	sru.mutation.AppendPreviousRequiredProviders(tr)
	return sru
}

// SetTags sets the "tags" field.
func (sru *ServiceRevisionUpdate) SetTags(s []string) *ServiceRevisionUpdate {
	sru.mutation.SetTags(s)
	return sru
}

// AppendTags appends s to the "tags" field.
func (sru *ServiceRevisionUpdate) AppendTags(s []string) *ServiceRevisionUpdate {
	sru.mutation.AppendTags(s)
	return sru
}

// Mutation returns the ServiceRevisionMutation object of the builder.
func (sru *ServiceRevisionUpdate) Mutation() *ServiceRevisionMutation {
	return sru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (sru *ServiceRevisionUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, sru.sqlSave, sru.mutation, sru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (sru *ServiceRevisionUpdate) SaveX(ctx context.Context) int {
	affected, err := sru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (sru *ServiceRevisionUpdate) Exec(ctx context.Context) error {
	_, err := sru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sru *ServiceRevisionUpdate) ExecX(ctx context.Context) {
	if err := sru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sru *ServiceRevisionUpdate) check() error {
	if v, ok := sru.mutation.TemplateVersion(); ok {
		if err := servicerevision.TemplateVersionValidator(v); err != nil {
			return &ValidationError{Name: "template_version", err: fmt.Errorf(`model: validator failed for field "ServiceRevision.template_version": %w`, err)}
		}
	}
	if _, ok := sru.mutation.ProjectID(); sru.mutation.ProjectCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "ServiceRevision.project"`)
	}
	if _, ok := sru.mutation.EnvironmentID(); sru.mutation.EnvironmentCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "ServiceRevision.environment"`)
	}
	if _, ok := sru.mutation.ServiceID(); sru.mutation.ServiceCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "ServiceRevision.service"`)
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
func (sru *ServiceRevisionUpdate) Set(obj *ServiceRevision) *ServiceRevisionUpdate {
	// Without Default.
	if obj.Status != "" {
		sru.SetStatus(obj.Status)
	}
	if obj.StatusMessage != "" {
		sru.SetStatusMessage(obj.StatusMessage)
	}
	sru.SetType(obj.Type)
	sru.SetTemplateVersion(obj.TemplateVersion)
	if !reflect.ValueOf(obj.Attributes).IsZero() {
		sru.SetAttributes(obj.Attributes)
	} else {
		sru.ClearAttributes()
	}
	sru.SetVariables(obj.Variables)
	sru.SetInputPlan(obj.InputPlan)
	sru.SetOutput(obj.Output)
	sru.SetDeployerType(obj.DeployerType)
	sru.SetDuration(obj.Duration)
	sru.SetPreviousRequiredProviders(obj.PreviousRequiredProviders)
	sru.SetTags(obj.Tags)

	// With Default.

	// Record the given object.
	sru.object = obj

	return sru
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (sru *ServiceRevisionUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ServiceRevisionUpdate {
	sru.modifiers = append(sru.modifiers, modifiers...)
	return sru
}

func (sru *ServiceRevisionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := sru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(servicerevision.Table, servicerevision.Columns, sqlgraph.NewFieldSpec(servicerevision.FieldID, field.TypeString))
	if ps := sru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sru.mutation.Status(); ok {
		_spec.SetField(servicerevision.FieldStatus, field.TypeString, value)
	}
	if sru.mutation.StatusCleared() {
		_spec.ClearField(servicerevision.FieldStatus, field.TypeString)
	}
	if value, ok := sru.mutation.StatusMessage(); ok {
		_spec.SetField(servicerevision.FieldStatusMessage, field.TypeString, value)
	}
	if sru.mutation.StatusMessageCleared() {
		_spec.ClearField(servicerevision.FieldStatusMessage, field.TypeString)
	}
	if value, ok := sru.mutation.GetType(); ok {
		_spec.SetField(servicerevision.FieldType, field.TypeString, value)
	}
	if value, ok := sru.mutation.TemplateVersion(); ok {
		_spec.SetField(servicerevision.FieldTemplateVersion, field.TypeString, value)
	}
	if value, ok := sru.mutation.Attributes(); ok {
		_spec.SetField(servicerevision.FieldAttributes, field.TypeOther, value)
	}
	if sru.mutation.AttributesCleared() {
		_spec.ClearField(servicerevision.FieldAttributes, field.TypeOther)
	}
	if value, ok := sru.mutation.Variables(); ok {
		_spec.SetField(servicerevision.FieldVariables, field.TypeOther, value)
	}
	if value, ok := sru.mutation.InputPlan(); ok {
		_spec.SetField(servicerevision.FieldInputPlan, field.TypeString, value)
	}
	if value, ok := sru.mutation.Output(); ok {
		_spec.SetField(servicerevision.FieldOutput, field.TypeString, value)
	}
	if value, ok := sru.mutation.DeployerType(); ok {
		_spec.SetField(servicerevision.FieldDeployerType, field.TypeString, value)
	}
	if value, ok := sru.mutation.Duration(); ok {
		_spec.SetField(servicerevision.FieldDuration, field.TypeInt, value)
	}
	if value, ok := sru.mutation.AddedDuration(); ok {
		_spec.AddField(servicerevision.FieldDuration, field.TypeInt, value)
	}
	if value, ok := sru.mutation.PreviousRequiredProviders(); ok {
		_spec.SetField(servicerevision.FieldPreviousRequiredProviders, field.TypeJSON, value)
	}
	if value, ok := sru.mutation.AppendedPreviousRequiredProviders(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, servicerevision.FieldPreviousRequiredProviders, value)
		})
	}
	if value, ok := sru.mutation.Tags(); ok {
		_spec.SetField(servicerevision.FieldTags, field.TypeJSON, value)
	}
	if value, ok := sru.mutation.AppendedTags(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, servicerevision.FieldTags, value)
		})
	}
	_spec.Node.Schema = sru.schemaConfig.ServiceRevision
	ctx = internal.NewSchemaConfigContext(ctx, sru.schemaConfig)
	_spec.AddModifiers(sru.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, sru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{servicerevision.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	sru.mutation.done = true
	return n, nil
}

// ServiceRevisionUpdateOne is the builder for updating a single ServiceRevision entity.
type ServiceRevisionUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *ServiceRevisionMutation
	modifiers []func(*sql.UpdateBuilder)
	object    *ServiceRevision
}

// SetStatus sets the "status" field.
func (sruo *ServiceRevisionUpdateOne) SetStatus(s string) *ServiceRevisionUpdateOne {
	sruo.mutation.SetStatus(s)
	return sruo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (sruo *ServiceRevisionUpdateOne) SetNillableStatus(s *string) *ServiceRevisionUpdateOne {
	if s != nil {
		sruo.SetStatus(*s)
	}
	return sruo
}

// ClearStatus clears the value of the "status" field.
func (sruo *ServiceRevisionUpdateOne) ClearStatus() *ServiceRevisionUpdateOne {
	sruo.mutation.ClearStatus()
	return sruo
}

// SetStatusMessage sets the "status_message" field.
func (sruo *ServiceRevisionUpdateOne) SetStatusMessage(s string) *ServiceRevisionUpdateOne {
	sruo.mutation.SetStatusMessage(s)
	return sruo
}

// SetNillableStatusMessage sets the "status_message" field if the given value is not nil.
func (sruo *ServiceRevisionUpdateOne) SetNillableStatusMessage(s *string) *ServiceRevisionUpdateOne {
	if s != nil {
		sruo.SetStatusMessage(*s)
	}
	return sruo
}

// ClearStatusMessage clears the value of the "status_message" field.
func (sruo *ServiceRevisionUpdateOne) ClearStatusMessage() *ServiceRevisionUpdateOne {
	sruo.mutation.ClearStatusMessage()
	return sruo
}

// SetType sets the "type" field.
func (sruo *ServiceRevisionUpdateOne) SetType(s string) *ServiceRevisionUpdateOne {
	sruo.mutation.SetType(s)
	return sruo
}

// SetTemplateVersion sets the "template_version" field.
func (sruo *ServiceRevisionUpdateOne) SetTemplateVersion(s string) *ServiceRevisionUpdateOne {
	sruo.mutation.SetTemplateVersion(s)
	return sruo
}

// SetAttributes sets the "attributes" field.
func (sruo *ServiceRevisionUpdateOne) SetAttributes(pr property.Values) *ServiceRevisionUpdateOne {
	sruo.mutation.SetAttributes(pr)
	return sruo
}

// ClearAttributes clears the value of the "attributes" field.
func (sruo *ServiceRevisionUpdateOne) ClearAttributes() *ServiceRevisionUpdateOne {
	sruo.mutation.ClearAttributes()
	return sruo
}

// SetVariables sets the "variables" field.
func (sruo *ServiceRevisionUpdateOne) SetVariables(c crypto.Map[string, string]) *ServiceRevisionUpdateOne {
	sruo.mutation.SetVariables(c)
	return sruo
}

// SetInputPlan sets the "input_plan" field.
func (sruo *ServiceRevisionUpdateOne) SetInputPlan(s string) *ServiceRevisionUpdateOne {
	sruo.mutation.SetInputPlan(s)
	return sruo
}

// SetOutput sets the "output" field.
func (sruo *ServiceRevisionUpdateOne) SetOutput(s string) *ServiceRevisionUpdateOne {
	sruo.mutation.SetOutput(s)
	return sruo
}

// SetDeployerType sets the "deployer_type" field.
func (sruo *ServiceRevisionUpdateOne) SetDeployerType(s string) *ServiceRevisionUpdateOne {
	sruo.mutation.SetDeployerType(s)
	return sruo
}

// SetNillableDeployerType sets the "deployer_type" field if the given value is not nil.
func (sruo *ServiceRevisionUpdateOne) SetNillableDeployerType(s *string) *ServiceRevisionUpdateOne {
	if s != nil {
		sruo.SetDeployerType(*s)
	}
	return sruo
}

// SetDuration sets the "duration" field.
func (sruo *ServiceRevisionUpdateOne) SetDuration(i int) *ServiceRevisionUpdateOne {
	sruo.mutation.ResetDuration()
	sruo.mutation.SetDuration(i)
	return sruo
}

// SetNillableDuration sets the "duration" field if the given value is not nil.
func (sruo *ServiceRevisionUpdateOne) SetNillableDuration(i *int) *ServiceRevisionUpdateOne {
	if i != nil {
		sruo.SetDuration(*i)
	}
	return sruo
}

// AddDuration adds i to the "duration" field.
func (sruo *ServiceRevisionUpdateOne) AddDuration(i int) *ServiceRevisionUpdateOne {
	sruo.mutation.AddDuration(i)
	return sruo
}

// SetPreviousRequiredProviders sets the "previous_required_providers" field.
func (sruo *ServiceRevisionUpdateOne) SetPreviousRequiredProviders(tr []types.ProviderRequirement) *ServiceRevisionUpdateOne {
	sruo.mutation.SetPreviousRequiredProviders(tr)
	return sruo
}

// AppendPreviousRequiredProviders appends tr to the "previous_required_providers" field.
func (sruo *ServiceRevisionUpdateOne) AppendPreviousRequiredProviders(tr []types.ProviderRequirement) *ServiceRevisionUpdateOne {
	sruo.mutation.AppendPreviousRequiredProviders(tr)
	return sruo
}

// SetTags sets the "tags" field.
func (sruo *ServiceRevisionUpdateOne) SetTags(s []string) *ServiceRevisionUpdateOne {
	sruo.mutation.SetTags(s)
	return sruo
}

// AppendTags appends s to the "tags" field.
func (sruo *ServiceRevisionUpdateOne) AppendTags(s []string) *ServiceRevisionUpdateOne {
	sruo.mutation.AppendTags(s)
	return sruo
}

// Mutation returns the ServiceRevisionMutation object of the builder.
func (sruo *ServiceRevisionUpdateOne) Mutation() *ServiceRevisionMutation {
	return sruo.mutation
}

// Where appends a list predicates to the ServiceRevisionUpdate builder.
func (sruo *ServiceRevisionUpdateOne) Where(ps ...predicate.ServiceRevision) *ServiceRevisionUpdateOne {
	sruo.mutation.Where(ps...)
	return sruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (sruo *ServiceRevisionUpdateOne) Select(field string, fields ...string) *ServiceRevisionUpdateOne {
	sruo.fields = append([]string{field}, fields...)
	return sruo
}

// Save executes the query and returns the updated ServiceRevision entity.
func (sruo *ServiceRevisionUpdateOne) Save(ctx context.Context) (*ServiceRevision, error) {
	return withHooks(ctx, sruo.sqlSave, sruo.mutation, sruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (sruo *ServiceRevisionUpdateOne) SaveX(ctx context.Context) *ServiceRevision {
	node, err := sruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (sruo *ServiceRevisionUpdateOne) Exec(ctx context.Context) error {
	_, err := sruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sruo *ServiceRevisionUpdateOne) ExecX(ctx context.Context) {
	if err := sruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sruo *ServiceRevisionUpdateOne) check() error {
	if v, ok := sruo.mutation.TemplateVersion(); ok {
		if err := servicerevision.TemplateVersionValidator(v); err != nil {
			return &ValidationError{Name: "template_version", err: fmt.Errorf(`model: validator failed for field "ServiceRevision.template_version": %w`, err)}
		}
	}
	if _, ok := sruo.mutation.ProjectID(); sruo.mutation.ProjectCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "ServiceRevision.project"`)
	}
	if _, ok := sruo.mutation.EnvironmentID(); sruo.mutation.EnvironmentCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "ServiceRevision.environment"`)
	}
	if _, ok := sruo.mutation.ServiceID(); sruo.mutation.ServiceCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "ServiceRevision.service"`)
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
func (sruo *ServiceRevisionUpdateOne) Set(obj *ServiceRevision) *ServiceRevisionUpdateOne {
	h := func(n ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			mt := m.(*ServiceRevisionMutation)
			db, err := mt.Client().ServiceRevision.Get(ctx, *mt.id)
			if err != nil {
				return nil, fmt.Errorf("failed getting ServiceRevision with id: %v", *mt.id)
			}

			// Without Default.
			if obj.Status != "" {
				if db.Status != obj.Status {
					sruo.SetStatus(obj.Status)
				}
			}
			if obj.StatusMessage != "" {
				if db.StatusMessage != obj.StatusMessage {
					sruo.SetStatusMessage(obj.StatusMessage)
				}
			}
			if db.Type != obj.Type {
				sruo.SetType(obj.Type)
			}
			if db.TemplateVersion != obj.TemplateVersion {
				sruo.SetTemplateVersion(obj.TemplateVersion)
			}
			if !reflect.ValueOf(obj.Attributes).IsZero() {
				if !reflect.DeepEqual(db.Attributes, obj.Attributes) {
					sruo.SetAttributes(obj.Attributes)
				}
			} else {
				sruo.ClearAttributes()
			}
			if !reflect.DeepEqual(db.Variables, obj.Variables) {
				sruo.SetVariables(obj.Variables)
			}
			if db.InputPlan != obj.InputPlan {
				sruo.SetInputPlan(obj.InputPlan)
			}
			if db.Output != obj.Output {
				sruo.SetOutput(obj.Output)
			}
			if db.DeployerType != obj.DeployerType {
				sruo.SetDeployerType(obj.DeployerType)
			}
			if db.Duration != obj.Duration {
				sruo.SetDuration(obj.Duration)
			}
			if !reflect.DeepEqual(db.PreviousRequiredProviders, obj.PreviousRequiredProviders) {
				sruo.SetPreviousRequiredProviders(obj.PreviousRequiredProviders)
			}
			if !reflect.DeepEqual(db.Tags, obj.Tags) {
				sruo.SetTags(obj.Tags)
			}

			// With Default.

			// Record the given object.
			sruo.object = obj

			return n.Mutate(ctx, m)
		})
	}

	sruo.hooks = append(sruo.hooks, h)

	return sruo
}

// getClientSet returns the ClientSet for the given builder.
func (sruo *ServiceRevisionUpdateOne) getClientSet() (mc ClientSet) {
	if _, ok := sruo.config.driver.(*txDriver); ok {
		tx := &Tx{config: sruo.config}
		tx.init()
		mc = tx
	} else {
		cli := &Client{config: sruo.config}
		cli.init()
		mc = cli
	}
	return mc
}

// SaveE calls the given function after updated the ServiceRevision entity,
// which is always good for cascading update operations.
func (sruo *ServiceRevisionUpdateOne) SaveE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *ServiceRevision) error) (*ServiceRevision, error) {
	obj, err := sruo.Save(ctx)
	if err != nil &&
		(sruo.object == nil || !errors.Is(err, stdsql.ErrNoRows)) {
		return nil, err
	}

	if len(cbs) == 0 {
		return obj, err
	}

	mc := sruo.getClientSet()

	if obj == nil {
		obj = sruo.object
	} else if x := sruo.object; x != nil {
		if _, set := sruo.mutation.Field(servicerevision.FieldStatus); set {
			obj.Status = x.Status
		}
		if _, set := sruo.mutation.Field(servicerevision.FieldStatusMessage); set {
			obj.StatusMessage = x.StatusMessage
		}
		if _, set := sruo.mutation.Field(servicerevision.FieldType); set {
			obj.Type = x.Type
		}
		if _, set := sruo.mutation.Field(servicerevision.FieldTemplateVersion); set {
			obj.TemplateVersion = x.TemplateVersion
		}
		if _, set := sruo.mutation.Field(servicerevision.FieldAttributes); set {
			obj.Attributes = x.Attributes
		}
		if _, set := sruo.mutation.Field(servicerevision.FieldVariables); set {
			obj.Variables = x.Variables
		}
		if _, set := sruo.mutation.Field(servicerevision.FieldInputPlan); set {
			obj.InputPlan = x.InputPlan
		}
		if _, set := sruo.mutation.Field(servicerevision.FieldOutput); set {
			obj.Output = x.Output
		}
		if _, set := sruo.mutation.Field(servicerevision.FieldDeployerType); set {
			obj.DeployerType = x.DeployerType
		}
		if _, set := sruo.mutation.Field(servicerevision.FieldDuration); set {
			obj.Duration = x.Duration
		}
		if _, set := sruo.mutation.Field(servicerevision.FieldPreviousRequiredProviders); set {
			obj.PreviousRequiredProviders = x.PreviousRequiredProviders
		}
		if _, set := sruo.mutation.Field(servicerevision.FieldTags); set {
			obj.Tags = x.Tags
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
func (sruo *ServiceRevisionUpdateOne) SaveEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *ServiceRevision) error) *ServiceRevision {
	obj, err := sruo.SaveE(ctx, cbs...)
	if err != nil {
		panic(err)
	}
	return obj
}

// ExecE calls the given function after executed the query,
// which is always good for cascading update operations.
func (sruo *ServiceRevisionUpdateOne) ExecE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *ServiceRevision) error) error {
	_, err := sruo.SaveE(ctx, cbs...)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sruo *ServiceRevisionUpdateOne) ExecEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *ServiceRevision) error) {
	if err := sruo.ExecE(ctx, cbs...); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (sruo *ServiceRevisionUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ServiceRevisionUpdateOne {
	sruo.modifiers = append(sruo.modifiers, modifiers...)
	return sruo
}

func (sruo *ServiceRevisionUpdateOne) sqlSave(ctx context.Context) (_node *ServiceRevision, err error) {
	if err := sruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(servicerevision.Table, servicerevision.Columns, sqlgraph.NewFieldSpec(servicerevision.FieldID, field.TypeString))
	id, ok := sruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`model: missing "ServiceRevision.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := sruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, servicerevision.FieldID)
		for _, f := range fields {
			if !servicerevision.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
			}
			if f != servicerevision.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := sruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sruo.mutation.Status(); ok {
		_spec.SetField(servicerevision.FieldStatus, field.TypeString, value)
	}
	if sruo.mutation.StatusCleared() {
		_spec.ClearField(servicerevision.FieldStatus, field.TypeString)
	}
	if value, ok := sruo.mutation.StatusMessage(); ok {
		_spec.SetField(servicerevision.FieldStatusMessage, field.TypeString, value)
	}
	if sruo.mutation.StatusMessageCleared() {
		_spec.ClearField(servicerevision.FieldStatusMessage, field.TypeString)
	}
	if value, ok := sruo.mutation.GetType(); ok {
		_spec.SetField(servicerevision.FieldType, field.TypeString, value)
	}
	if value, ok := sruo.mutation.TemplateVersion(); ok {
		_spec.SetField(servicerevision.FieldTemplateVersion, field.TypeString, value)
	}
	if value, ok := sruo.mutation.Attributes(); ok {
		_spec.SetField(servicerevision.FieldAttributes, field.TypeOther, value)
	}
	if sruo.mutation.AttributesCleared() {
		_spec.ClearField(servicerevision.FieldAttributes, field.TypeOther)
	}
	if value, ok := sruo.mutation.Variables(); ok {
		_spec.SetField(servicerevision.FieldVariables, field.TypeOther, value)
	}
	if value, ok := sruo.mutation.InputPlan(); ok {
		_spec.SetField(servicerevision.FieldInputPlan, field.TypeString, value)
	}
	if value, ok := sruo.mutation.Output(); ok {
		_spec.SetField(servicerevision.FieldOutput, field.TypeString, value)
	}
	if value, ok := sruo.mutation.DeployerType(); ok {
		_spec.SetField(servicerevision.FieldDeployerType, field.TypeString, value)
	}
	if value, ok := sruo.mutation.Duration(); ok {
		_spec.SetField(servicerevision.FieldDuration, field.TypeInt, value)
	}
	if value, ok := sruo.mutation.AddedDuration(); ok {
		_spec.AddField(servicerevision.FieldDuration, field.TypeInt, value)
	}
	if value, ok := sruo.mutation.PreviousRequiredProviders(); ok {
		_spec.SetField(servicerevision.FieldPreviousRequiredProviders, field.TypeJSON, value)
	}
	if value, ok := sruo.mutation.AppendedPreviousRequiredProviders(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, servicerevision.FieldPreviousRequiredProviders, value)
		})
	}
	if value, ok := sruo.mutation.Tags(); ok {
		_spec.SetField(servicerevision.FieldTags, field.TypeJSON, value)
	}
	if value, ok := sruo.mutation.AppendedTags(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, servicerevision.FieldTags, value)
		})
	}
	_spec.Node.Schema = sruo.schemaConfig.ServiceRevision
	ctx = internal.NewSchemaConfigContext(ctx, sruo.schemaConfig)
	_spec.AddModifiers(sruo.modifiers...)
	_node = &ServiceRevision{config: sruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, sruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{servicerevision.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	sruo.mutation.done = true
	return _node, nil
}
