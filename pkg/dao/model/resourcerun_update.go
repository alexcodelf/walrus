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
	"github.com/seal-io/walrus/pkg/dao/model/resourcerun"
	"github.com/seal-io/walrus/pkg/dao/types"
	"github.com/seal-io/walrus/pkg/dao/types/crypto"
	"github.com/seal-io/walrus/pkg/dao/types/property"
	"github.com/seal-io/walrus/pkg/dao/types/status"
)

// ResourceRunUpdate is the builder for updating ResourceRun entities.
type ResourceRunUpdate struct {
	config
	hooks     []Hook
	mutation  *ResourceRunMutation
	modifiers []func(*sql.UpdateBuilder)
	object    *ResourceRun
}

// Where appends a list predicates to the ResourceRunUpdate builder.
func (rru *ResourceRunUpdate) Where(ps ...predicate.ResourceRun) *ResourceRunUpdate {
	rru.mutation.Where(ps...)
	return rru
}

// SetStatus sets the "status" field.
func (rru *ResourceRunUpdate) SetStatus(s status.Status) *ResourceRunUpdate {
	rru.mutation.SetStatus(s)
	return rru
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (rru *ResourceRunUpdate) SetNillableStatus(s *status.Status) *ResourceRunUpdate {
	if s != nil {
		rru.SetStatus(*s)
	}
	return rru
}

// ClearStatus clears the value of the "status" field.
func (rru *ResourceRunUpdate) ClearStatus() *ResourceRunUpdate {
	rru.mutation.ClearStatus()
	return rru
}

// SetTemplateVersion sets the "template_version" field.
func (rru *ResourceRunUpdate) SetTemplateVersion(s string) *ResourceRunUpdate {
	rru.mutation.SetTemplateVersion(s)
	return rru
}

// SetNillableTemplateVersion sets the "template_version" field if the given value is not nil.
func (rru *ResourceRunUpdate) SetNillableTemplateVersion(s *string) *ResourceRunUpdate {
	if s != nil {
		rru.SetTemplateVersion(*s)
	}
	return rru
}

// SetAttributes sets the "attributes" field.
func (rru *ResourceRunUpdate) SetAttributes(pr property.Values) *ResourceRunUpdate {
	rru.mutation.SetAttributes(pr)
	return rru
}

// ClearAttributes clears the value of the "attributes" field.
func (rru *ResourceRunUpdate) ClearAttributes() *ResourceRunUpdate {
	rru.mutation.ClearAttributes()
	return rru
}

// SetComputedAttributes sets the "computed_attributes" field.
func (rru *ResourceRunUpdate) SetComputedAttributes(pr property.Values) *ResourceRunUpdate {
	rru.mutation.SetComputedAttributes(pr)
	return rru
}

// ClearComputedAttributes clears the value of the "computed_attributes" field.
func (rru *ResourceRunUpdate) ClearComputedAttributes() *ResourceRunUpdate {
	rru.mutation.ClearComputedAttributes()
	return rru
}

// SetVariables sets the "variables" field.
func (rru *ResourceRunUpdate) SetVariables(c crypto.Map[string, string]) *ResourceRunUpdate {
	rru.mutation.SetVariables(c)
	return rru
}

// SetInputConfigs sets the "input_configs" field.
func (rru *ResourceRunUpdate) SetInputConfigs(m map[string][]uint8) *ResourceRunUpdate {
	rru.mutation.SetInputConfigs(m)
	return rru
}

// SetDeployerType sets the "deployer_type" field.
func (rru *ResourceRunUpdate) SetDeployerType(s string) *ResourceRunUpdate {
	rru.mutation.SetDeployerType(s)
	return rru
}

// SetNillableDeployerType sets the "deployer_type" field if the given value is not nil.
func (rru *ResourceRunUpdate) SetNillableDeployerType(s *string) *ResourceRunUpdate {
	if s != nil {
		rru.SetDeployerType(*s)
	}
	return rru
}

// SetDuration sets the "duration" field.
func (rru *ResourceRunUpdate) SetDuration(i int) *ResourceRunUpdate {
	rru.mutation.ResetDuration()
	rru.mutation.SetDuration(i)
	return rru
}

// SetNillableDuration sets the "duration" field if the given value is not nil.
func (rru *ResourceRunUpdate) SetNillableDuration(i *int) *ResourceRunUpdate {
	if i != nil {
		rru.SetDuration(*i)
	}
	return rru
}

// AddDuration adds i to the "duration" field.
func (rru *ResourceRunUpdate) AddDuration(i int) *ResourceRunUpdate {
	rru.mutation.AddDuration(i)
	return rru
}

// SetPreviousRequiredProviders sets the "previous_required_providers" field.
func (rru *ResourceRunUpdate) SetPreviousRequiredProviders(tr []types.ProviderRequirement) *ResourceRunUpdate {
	rru.mutation.SetPreviousRequiredProviders(tr)
	return rru
}

// AppendPreviousRequiredProviders appends tr to the "previous_required_providers" field.
func (rru *ResourceRunUpdate) AppendPreviousRequiredProviders(tr []types.ProviderRequirement) *ResourceRunUpdate {
	rru.mutation.AppendPreviousRequiredProviders(tr)
	return rru
}

// SetRecord sets the "record" field.
func (rru *ResourceRunUpdate) SetRecord(s string) *ResourceRunUpdate {
	rru.mutation.SetRecord(s)
	return rru
}

// SetNillableRecord sets the "record" field if the given value is not nil.
func (rru *ResourceRunUpdate) SetNillableRecord(s *string) *ResourceRunUpdate {
	if s != nil {
		rru.SetRecord(*s)
	}
	return rru
}

// ClearRecord clears the value of the "record" field.
func (rru *ResourceRunUpdate) ClearRecord() *ResourceRunUpdate {
	rru.mutation.ClearRecord()
	return rru
}

// SetChangeComment sets the "change_comment" field.
func (rru *ResourceRunUpdate) SetChangeComment(s string) *ResourceRunUpdate {
	rru.mutation.SetChangeComment(s)
	return rru
}

// SetNillableChangeComment sets the "change_comment" field if the given value is not nil.
func (rru *ResourceRunUpdate) SetNillableChangeComment(s *string) *ResourceRunUpdate {
	if s != nil {
		rru.SetChangeComment(*s)
	}
	return rru
}

// ClearChangeComment clears the value of the "change_comment" field.
func (rru *ResourceRunUpdate) ClearChangeComment() *ResourceRunUpdate {
	rru.mutation.ClearChangeComment()
	return rru
}

// SetCreatedBy sets the "created_by" field.
func (rru *ResourceRunUpdate) SetCreatedBy(s string) *ResourceRunUpdate {
	rru.mutation.SetCreatedBy(s)
	return rru
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (rru *ResourceRunUpdate) SetNillableCreatedBy(s *string) *ResourceRunUpdate {
	if s != nil {
		rru.SetCreatedBy(*s)
	}
	return rru
}

// Mutation returns the ResourceRunMutation object of the builder.
func (rru *ResourceRunUpdate) Mutation() *ResourceRunMutation {
	return rru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (rru *ResourceRunUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, rru.sqlSave, rru.mutation, rru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rru *ResourceRunUpdate) SaveX(ctx context.Context) int {
	affected, err := rru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (rru *ResourceRunUpdate) Exec(ctx context.Context) error {
	_, err := rru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rru *ResourceRunUpdate) ExecX(ctx context.Context) {
	if err := rru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rru *ResourceRunUpdate) check() error {
	if v, ok := rru.mutation.TemplateVersion(); ok {
		if err := resourcerun.TemplateVersionValidator(v); err != nil {
			return &ValidationError{Name: "template_version", err: fmt.Errorf(`model: validator failed for field "ResourceRun.template_version": %w`, err)}
		}
	}
	if _, ok := rru.mutation.ProjectID(); rru.mutation.ProjectCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "ResourceRun.project"`)
	}
	if _, ok := rru.mutation.EnvironmentID(); rru.mutation.EnvironmentCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "ResourceRun.environment"`)
	}
	if _, ok := rru.mutation.ResourceID(); rru.mutation.ResourceCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "ResourceRun.resource"`)
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
func (rru *ResourceRunUpdate) Set(obj *ResourceRun) *ResourceRunUpdate {
	// Without Default.
	if !reflect.ValueOf(obj.Status).IsZero() {
		rru.SetStatus(obj.Status)
	}
	rru.SetTemplateVersion(obj.TemplateVersion)
	if !reflect.ValueOf(obj.Attributes).IsZero() {
		rru.SetAttributes(obj.Attributes)
	} else {
		rru.ClearAttributes()
	}
	if !reflect.ValueOf(obj.ComputedAttributes).IsZero() {
		rru.SetComputedAttributes(obj.ComputedAttributes)
	} else {
		rru.ClearComputedAttributes()
	}
	rru.SetVariables(obj.Variables)
	rru.SetInputConfigs(obj.InputConfigs)
	rru.SetDeployerType(obj.DeployerType)
	rru.SetDuration(obj.Duration)
	rru.SetPreviousRequiredProviders(obj.PreviousRequiredProviders)
	if obj.Record != "" {
		rru.SetRecord(obj.Record)
	} else {
		rru.ClearRecord()
	}
	if obj.ChangeComment != "" {
		rru.SetChangeComment(obj.ChangeComment)
	} else {
		rru.ClearChangeComment()
	}
	rru.SetCreatedBy(obj.CreatedBy)

	// With Default.

	// Record the given object.
	rru.object = obj

	return rru
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (rru *ResourceRunUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ResourceRunUpdate {
	rru.modifiers = append(rru.modifiers, modifiers...)
	return rru
}

func (rru *ResourceRunUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := rru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(resourcerun.Table, resourcerun.Columns, sqlgraph.NewFieldSpec(resourcerun.FieldID, field.TypeString))
	if ps := rru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rru.mutation.Status(); ok {
		_spec.SetField(resourcerun.FieldStatus, field.TypeJSON, value)
	}
	if rru.mutation.StatusCleared() {
		_spec.ClearField(resourcerun.FieldStatus, field.TypeJSON)
	}
	if value, ok := rru.mutation.TemplateVersion(); ok {
		_spec.SetField(resourcerun.FieldTemplateVersion, field.TypeString, value)
	}
	if value, ok := rru.mutation.Attributes(); ok {
		_spec.SetField(resourcerun.FieldAttributes, field.TypeOther, value)
	}
	if rru.mutation.AttributesCleared() {
		_spec.ClearField(resourcerun.FieldAttributes, field.TypeOther)
	}
	if value, ok := rru.mutation.ComputedAttributes(); ok {
		_spec.SetField(resourcerun.FieldComputedAttributes, field.TypeOther, value)
	}
	if rru.mutation.ComputedAttributesCleared() {
		_spec.ClearField(resourcerun.FieldComputedAttributes, field.TypeOther)
	}
	if value, ok := rru.mutation.Variables(); ok {
		_spec.SetField(resourcerun.FieldVariables, field.TypeOther, value)
	}
	if value, ok := rru.mutation.InputConfigs(); ok {
		_spec.SetField(resourcerun.FieldInputConfigs, field.TypeJSON, value)
	}
	if value, ok := rru.mutation.DeployerType(); ok {
		_spec.SetField(resourcerun.FieldDeployerType, field.TypeString, value)
	}
	if value, ok := rru.mutation.Duration(); ok {
		_spec.SetField(resourcerun.FieldDuration, field.TypeInt, value)
	}
	if value, ok := rru.mutation.AddedDuration(); ok {
		_spec.AddField(resourcerun.FieldDuration, field.TypeInt, value)
	}
	if value, ok := rru.mutation.PreviousRequiredProviders(); ok {
		_spec.SetField(resourcerun.FieldPreviousRequiredProviders, field.TypeJSON, value)
	}
	if value, ok := rru.mutation.AppendedPreviousRequiredProviders(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, resourcerun.FieldPreviousRequiredProviders, value)
		})
	}
	if value, ok := rru.mutation.Record(); ok {
		_spec.SetField(resourcerun.FieldRecord, field.TypeString, value)
	}
	if rru.mutation.RecordCleared() {
		_spec.ClearField(resourcerun.FieldRecord, field.TypeString)
	}
	if value, ok := rru.mutation.ChangeComment(); ok {
		_spec.SetField(resourcerun.FieldChangeComment, field.TypeString, value)
	}
	if rru.mutation.ChangeCommentCleared() {
		_spec.ClearField(resourcerun.FieldChangeComment, field.TypeString)
	}
	if value, ok := rru.mutation.CreatedBy(); ok {
		_spec.SetField(resourcerun.FieldCreatedBy, field.TypeString, value)
	}
	_spec.Node.Schema = rru.schemaConfig.ResourceRun
	ctx = internal.NewSchemaConfigContext(ctx, rru.schemaConfig)
	_spec.AddModifiers(rru.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, rru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{resourcerun.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	rru.mutation.done = true
	return n, nil
}

// ResourceRunUpdateOne is the builder for updating a single ResourceRun entity.
type ResourceRunUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *ResourceRunMutation
	modifiers []func(*sql.UpdateBuilder)
	object    *ResourceRun
}

// SetStatus sets the "status" field.
func (rruo *ResourceRunUpdateOne) SetStatus(s status.Status) *ResourceRunUpdateOne {
	rruo.mutation.SetStatus(s)
	return rruo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (rruo *ResourceRunUpdateOne) SetNillableStatus(s *status.Status) *ResourceRunUpdateOne {
	if s != nil {
		rruo.SetStatus(*s)
	}
	return rruo
}

// ClearStatus clears the value of the "status" field.
func (rruo *ResourceRunUpdateOne) ClearStatus() *ResourceRunUpdateOne {
	rruo.mutation.ClearStatus()
	return rruo
}

// SetTemplateVersion sets the "template_version" field.
func (rruo *ResourceRunUpdateOne) SetTemplateVersion(s string) *ResourceRunUpdateOne {
	rruo.mutation.SetTemplateVersion(s)
	return rruo
}

// SetNillableTemplateVersion sets the "template_version" field if the given value is not nil.
func (rruo *ResourceRunUpdateOne) SetNillableTemplateVersion(s *string) *ResourceRunUpdateOne {
	if s != nil {
		rruo.SetTemplateVersion(*s)
	}
	return rruo
}

// SetAttributes sets the "attributes" field.
func (rruo *ResourceRunUpdateOne) SetAttributes(pr property.Values) *ResourceRunUpdateOne {
	rruo.mutation.SetAttributes(pr)
	return rruo
}

// ClearAttributes clears the value of the "attributes" field.
func (rruo *ResourceRunUpdateOne) ClearAttributes() *ResourceRunUpdateOne {
	rruo.mutation.ClearAttributes()
	return rruo
}

// SetComputedAttributes sets the "computed_attributes" field.
func (rruo *ResourceRunUpdateOne) SetComputedAttributes(pr property.Values) *ResourceRunUpdateOne {
	rruo.mutation.SetComputedAttributes(pr)
	return rruo
}

// ClearComputedAttributes clears the value of the "computed_attributes" field.
func (rruo *ResourceRunUpdateOne) ClearComputedAttributes() *ResourceRunUpdateOne {
	rruo.mutation.ClearComputedAttributes()
	return rruo
}

// SetVariables sets the "variables" field.
func (rruo *ResourceRunUpdateOne) SetVariables(c crypto.Map[string, string]) *ResourceRunUpdateOne {
	rruo.mutation.SetVariables(c)
	return rruo
}

// SetInputConfigs sets the "input_configs" field.
func (rruo *ResourceRunUpdateOne) SetInputConfigs(m map[string][]uint8) *ResourceRunUpdateOne {
	rruo.mutation.SetInputConfigs(m)
	return rruo
}

// SetDeployerType sets the "deployer_type" field.
func (rruo *ResourceRunUpdateOne) SetDeployerType(s string) *ResourceRunUpdateOne {
	rruo.mutation.SetDeployerType(s)
	return rruo
}

// SetNillableDeployerType sets the "deployer_type" field if the given value is not nil.
func (rruo *ResourceRunUpdateOne) SetNillableDeployerType(s *string) *ResourceRunUpdateOne {
	if s != nil {
		rruo.SetDeployerType(*s)
	}
	return rruo
}

// SetDuration sets the "duration" field.
func (rruo *ResourceRunUpdateOne) SetDuration(i int) *ResourceRunUpdateOne {
	rruo.mutation.ResetDuration()
	rruo.mutation.SetDuration(i)
	return rruo
}

// SetNillableDuration sets the "duration" field if the given value is not nil.
func (rruo *ResourceRunUpdateOne) SetNillableDuration(i *int) *ResourceRunUpdateOne {
	if i != nil {
		rruo.SetDuration(*i)
	}
	return rruo
}

// AddDuration adds i to the "duration" field.
func (rruo *ResourceRunUpdateOne) AddDuration(i int) *ResourceRunUpdateOne {
	rruo.mutation.AddDuration(i)
	return rruo
}

// SetPreviousRequiredProviders sets the "previous_required_providers" field.
func (rruo *ResourceRunUpdateOne) SetPreviousRequiredProviders(tr []types.ProviderRequirement) *ResourceRunUpdateOne {
	rruo.mutation.SetPreviousRequiredProviders(tr)
	return rruo
}

// AppendPreviousRequiredProviders appends tr to the "previous_required_providers" field.
func (rruo *ResourceRunUpdateOne) AppendPreviousRequiredProviders(tr []types.ProviderRequirement) *ResourceRunUpdateOne {
	rruo.mutation.AppendPreviousRequiredProviders(tr)
	return rruo
}

// SetRecord sets the "record" field.
func (rruo *ResourceRunUpdateOne) SetRecord(s string) *ResourceRunUpdateOne {
	rruo.mutation.SetRecord(s)
	return rruo
}

// SetNillableRecord sets the "record" field if the given value is not nil.
func (rruo *ResourceRunUpdateOne) SetNillableRecord(s *string) *ResourceRunUpdateOne {
	if s != nil {
		rruo.SetRecord(*s)
	}
	return rruo
}

// ClearRecord clears the value of the "record" field.
func (rruo *ResourceRunUpdateOne) ClearRecord() *ResourceRunUpdateOne {
	rruo.mutation.ClearRecord()
	return rruo
}

// SetChangeComment sets the "change_comment" field.
func (rruo *ResourceRunUpdateOne) SetChangeComment(s string) *ResourceRunUpdateOne {
	rruo.mutation.SetChangeComment(s)
	return rruo
}

// SetNillableChangeComment sets the "change_comment" field if the given value is not nil.
func (rruo *ResourceRunUpdateOne) SetNillableChangeComment(s *string) *ResourceRunUpdateOne {
	if s != nil {
		rruo.SetChangeComment(*s)
	}
	return rruo
}

// ClearChangeComment clears the value of the "change_comment" field.
func (rruo *ResourceRunUpdateOne) ClearChangeComment() *ResourceRunUpdateOne {
	rruo.mutation.ClearChangeComment()
	return rruo
}

// SetCreatedBy sets the "created_by" field.
func (rruo *ResourceRunUpdateOne) SetCreatedBy(s string) *ResourceRunUpdateOne {
	rruo.mutation.SetCreatedBy(s)
	return rruo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (rruo *ResourceRunUpdateOne) SetNillableCreatedBy(s *string) *ResourceRunUpdateOne {
	if s != nil {
		rruo.SetCreatedBy(*s)
	}
	return rruo
}

// Mutation returns the ResourceRunMutation object of the builder.
func (rruo *ResourceRunUpdateOne) Mutation() *ResourceRunMutation {
	return rruo.mutation
}

// Where appends a list predicates to the ResourceRunUpdate builder.
func (rruo *ResourceRunUpdateOne) Where(ps ...predicate.ResourceRun) *ResourceRunUpdateOne {
	rruo.mutation.Where(ps...)
	return rruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (rruo *ResourceRunUpdateOne) Select(field string, fields ...string) *ResourceRunUpdateOne {
	rruo.fields = append([]string{field}, fields...)
	return rruo
}

// Save executes the query and returns the updated ResourceRun entity.
func (rruo *ResourceRunUpdateOne) Save(ctx context.Context) (*ResourceRun, error) {
	return withHooks(ctx, rruo.sqlSave, rruo.mutation, rruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rruo *ResourceRunUpdateOne) SaveX(ctx context.Context) *ResourceRun {
	node, err := rruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (rruo *ResourceRunUpdateOne) Exec(ctx context.Context) error {
	_, err := rruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rruo *ResourceRunUpdateOne) ExecX(ctx context.Context) {
	if err := rruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rruo *ResourceRunUpdateOne) check() error {
	if v, ok := rruo.mutation.TemplateVersion(); ok {
		if err := resourcerun.TemplateVersionValidator(v); err != nil {
			return &ValidationError{Name: "template_version", err: fmt.Errorf(`model: validator failed for field "ResourceRun.template_version": %w`, err)}
		}
	}
	if _, ok := rruo.mutation.ProjectID(); rruo.mutation.ProjectCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "ResourceRun.project"`)
	}
	if _, ok := rruo.mutation.EnvironmentID(); rruo.mutation.EnvironmentCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "ResourceRun.environment"`)
	}
	if _, ok := rruo.mutation.ResourceID(); rruo.mutation.ResourceCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "ResourceRun.resource"`)
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
func (rruo *ResourceRunUpdateOne) Set(obj *ResourceRun) *ResourceRunUpdateOne {
	h := func(n ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			mt := m.(*ResourceRunMutation)
			db, err := mt.Client().ResourceRun.Get(ctx, *mt.id)
			if err != nil {
				return nil, fmt.Errorf("failed getting ResourceRun with id: %v", *mt.id)
			}

			// Without Default.
			if !reflect.ValueOf(obj.Status).IsZero() {
				if !db.Status.Equal(obj.Status) {
					rruo.SetStatus(obj.Status)
				}
			}
			if db.TemplateVersion != obj.TemplateVersion {
				rruo.SetTemplateVersion(obj.TemplateVersion)
			}
			if !reflect.ValueOf(obj.Attributes).IsZero() {
				if !reflect.DeepEqual(db.Attributes, obj.Attributes) {
					rruo.SetAttributes(obj.Attributes)
				}
			} else {
				rruo.ClearAttributes()
			}
			if !reflect.ValueOf(obj.ComputedAttributes).IsZero() {
				if !reflect.DeepEqual(db.ComputedAttributes, obj.ComputedAttributes) {
					rruo.SetComputedAttributes(obj.ComputedAttributes)
				}
			} else {
				rruo.ClearComputedAttributes()
			}
			if !reflect.DeepEqual(db.Variables, obj.Variables) {
				rruo.SetVariables(obj.Variables)
			}
			if !reflect.DeepEqual(db.InputConfigs, obj.InputConfigs) {
				rruo.SetInputConfigs(obj.InputConfigs)
			}
			if db.DeployerType != obj.DeployerType {
				rruo.SetDeployerType(obj.DeployerType)
			}
			if db.Duration != obj.Duration {
				rruo.SetDuration(obj.Duration)
			}
			if !reflect.DeepEqual(db.PreviousRequiredProviders, obj.PreviousRequiredProviders) {
				rruo.SetPreviousRequiredProviders(obj.PreviousRequiredProviders)
			}
			if obj.Record != "" {
				if db.Record != obj.Record {
					rruo.SetRecord(obj.Record)
				}
			} else {
				rruo.ClearRecord()
			}
			if obj.ChangeComment != "" {
				if db.ChangeComment != obj.ChangeComment {
					rruo.SetChangeComment(obj.ChangeComment)
				}
			} else {
				rruo.ClearChangeComment()
			}
			if db.CreatedBy != obj.CreatedBy {
				rruo.SetCreatedBy(obj.CreatedBy)
			}

			// With Default.

			// Record the given object.
			rruo.object = obj

			return n.Mutate(ctx, m)
		})
	}

	rruo.hooks = append(rruo.hooks, h)

	return rruo
}

// getClientSet returns the ClientSet for the given builder.
func (rruo *ResourceRunUpdateOne) getClientSet() (mc ClientSet) {
	if _, ok := rruo.config.driver.(*txDriver); ok {
		tx := &Tx{config: rruo.config}
		tx.init()
		mc = tx
	} else {
		cli := &Client{config: rruo.config}
		cli.init()
		mc = cli
	}
	return mc
}

// SaveE calls the given function after updated the ResourceRun entity,
// which is always good for cascading update operations.
func (rruo *ResourceRunUpdateOne) SaveE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *ResourceRun) error) (*ResourceRun, error) {
	obj, err := rruo.Save(ctx)
	if err != nil &&
		(rruo.object == nil || !errors.Is(err, stdsql.ErrNoRows)) {
		return nil, err
	}

	if len(cbs) == 0 {
		return obj, err
	}

	mc := rruo.getClientSet()

	if obj == nil {
		obj = rruo.object
	} else if x := rruo.object; x != nil {
		if _, set := rruo.mutation.Field(resourcerun.FieldStatus); set {
			obj.Status = x.Status
		}
		if _, set := rruo.mutation.Field(resourcerun.FieldTemplateVersion); set {
			obj.TemplateVersion = x.TemplateVersion
		}
		if _, set := rruo.mutation.Field(resourcerun.FieldAttributes); set {
			obj.Attributes = x.Attributes
		}
		if _, set := rruo.mutation.Field(resourcerun.FieldComputedAttributes); set {
			obj.ComputedAttributes = x.ComputedAttributes
		}
		if _, set := rruo.mutation.Field(resourcerun.FieldVariables); set {
			obj.Variables = x.Variables
		}
		if _, set := rruo.mutation.Field(resourcerun.FieldInputConfigs); set {
			obj.InputConfigs = x.InputConfigs
		}
		if _, set := rruo.mutation.Field(resourcerun.FieldDeployerType); set {
			obj.DeployerType = x.DeployerType
		}
		if _, set := rruo.mutation.Field(resourcerun.FieldDuration); set {
			obj.Duration = x.Duration
		}
		if _, set := rruo.mutation.Field(resourcerun.FieldPreviousRequiredProviders); set {
			obj.PreviousRequiredProviders = x.PreviousRequiredProviders
		}
		if _, set := rruo.mutation.Field(resourcerun.FieldRecord); set {
			obj.Record = x.Record
		}
		if _, set := rruo.mutation.Field(resourcerun.FieldChangeComment); set {
			obj.ChangeComment = x.ChangeComment
		}
		if _, set := rruo.mutation.Field(resourcerun.FieldCreatedBy); set {
			obj.CreatedBy = x.CreatedBy
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
func (rruo *ResourceRunUpdateOne) SaveEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *ResourceRun) error) *ResourceRun {
	obj, err := rruo.SaveE(ctx, cbs...)
	if err != nil {
		panic(err)
	}
	return obj
}

// ExecE calls the given function after executed the query,
// which is always good for cascading update operations.
func (rruo *ResourceRunUpdateOne) ExecE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *ResourceRun) error) error {
	_, err := rruo.SaveE(ctx, cbs...)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rruo *ResourceRunUpdateOne) ExecEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *ResourceRun) error) {
	if err := rruo.ExecE(ctx, cbs...); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (rruo *ResourceRunUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ResourceRunUpdateOne {
	rruo.modifiers = append(rruo.modifiers, modifiers...)
	return rruo
}

func (rruo *ResourceRunUpdateOne) sqlSave(ctx context.Context) (_node *ResourceRun, err error) {
	if err := rruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(resourcerun.Table, resourcerun.Columns, sqlgraph.NewFieldSpec(resourcerun.FieldID, field.TypeString))
	id, ok := rruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`model: missing "ResourceRun.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := rruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, resourcerun.FieldID)
		for _, f := range fields {
			if !resourcerun.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
			}
			if f != resourcerun.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := rruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rruo.mutation.Status(); ok {
		_spec.SetField(resourcerun.FieldStatus, field.TypeJSON, value)
	}
	if rruo.mutation.StatusCleared() {
		_spec.ClearField(resourcerun.FieldStatus, field.TypeJSON)
	}
	if value, ok := rruo.mutation.TemplateVersion(); ok {
		_spec.SetField(resourcerun.FieldTemplateVersion, field.TypeString, value)
	}
	if value, ok := rruo.mutation.Attributes(); ok {
		_spec.SetField(resourcerun.FieldAttributes, field.TypeOther, value)
	}
	if rruo.mutation.AttributesCleared() {
		_spec.ClearField(resourcerun.FieldAttributes, field.TypeOther)
	}
	if value, ok := rruo.mutation.ComputedAttributes(); ok {
		_spec.SetField(resourcerun.FieldComputedAttributes, field.TypeOther, value)
	}
	if rruo.mutation.ComputedAttributesCleared() {
		_spec.ClearField(resourcerun.FieldComputedAttributes, field.TypeOther)
	}
	if value, ok := rruo.mutation.Variables(); ok {
		_spec.SetField(resourcerun.FieldVariables, field.TypeOther, value)
	}
	if value, ok := rruo.mutation.InputConfigs(); ok {
		_spec.SetField(resourcerun.FieldInputConfigs, field.TypeJSON, value)
	}
	if value, ok := rruo.mutation.DeployerType(); ok {
		_spec.SetField(resourcerun.FieldDeployerType, field.TypeString, value)
	}
	if value, ok := rruo.mutation.Duration(); ok {
		_spec.SetField(resourcerun.FieldDuration, field.TypeInt, value)
	}
	if value, ok := rruo.mutation.AddedDuration(); ok {
		_spec.AddField(resourcerun.FieldDuration, field.TypeInt, value)
	}
	if value, ok := rruo.mutation.PreviousRequiredProviders(); ok {
		_spec.SetField(resourcerun.FieldPreviousRequiredProviders, field.TypeJSON, value)
	}
	if value, ok := rruo.mutation.AppendedPreviousRequiredProviders(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, resourcerun.FieldPreviousRequiredProviders, value)
		})
	}
	if value, ok := rruo.mutation.Record(); ok {
		_spec.SetField(resourcerun.FieldRecord, field.TypeString, value)
	}
	if rruo.mutation.RecordCleared() {
		_spec.ClearField(resourcerun.FieldRecord, field.TypeString)
	}
	if value, ok := rruo.mutation.ChangeComment(); ok {
		_spec.SetField(resourcerun.FieldChangeComment, field.TypeString, value)
	}
	if rruo.mutation.ChangeCommentCleared() {
		_spec.ClearField(resourcerun.FieldChangeComment, field.TypeString)
	}
	if value, ok := rruo.mutation.CreatedBy(); ok {
		_spec.SetField(resourcerun.FieldCreatedBy, field.TypeString, value)
	}
	_spec.Node.Schema = rruo.schemaConfig.ResourceRun
	ctx = internal.NewSchemaConfigContext(ctx, rruo.schemaConfig)
	_spec.AddModifiers(rruo.modifiers...)
	_node = &ResourceRun{config: rruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, rruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{resourcerun.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	rruo.mutation.done = true
	return _node, nil
}
