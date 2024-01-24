// SPDX-FileCopyrightText: 2024 Seal, Inc
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
	"github.com/seal-io/walrus/pkg/dao/model/resourcerevision"
	"github.com/seal-io/walrus/pkg/dao/types"
	"github.com/seal-io/walrus/pkg/dao/types/crypto"
	"github.com/seal-io/walrus/pkg/dao/types/property"
	"github.com/seal-io/walrus/pkg/dao/types/status"
)

// ResourceRevisionUpdate is the builder for updating ResourceRevision entities.
type ResourceRevisionUpdate struct {
	config
	hooks     []Hook
	mutation  *ResourceRevisionMutation
	modifiers []func(*sql.UpdateBuilder)
	object    *ResourceRevision
}

// Where appends a list predicates to the ResourceRevisionUpdate builder.
func (rru *ResourceRevisionUpdate) Where(ps ...predicate.ResourceRevision) *ResourceRevisionUpdate {
	rru.mutation.Where(ps...)
	return rru
}

// SetStatus sets the "status" field.
func (rru *ResourceRevisionUpdate) SetStatus(s status.Status) *ResourceRevisionUpdate {
	rru.mutation.SetStatus(s)
	return rru
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (rru *ResourceRevisionUpdate) SetNillableStatus(s *status.Status) *ResourceRevisionUpdate {
	if s != nil {
		rru.SetStatus(*s)
	}
	return rru
}

// ClearStatus clears the value of the "status" field.
func (rru *ResourceRevisionUpdate) ClearStatus() *ResourceRevisionUpdate {
	rru.mutation.ClearStatus()
	return rru
}

// SetTemplateVersion sets the "template_version" field.
func (rru *ResourceRevisionUpdate) SetTemplateVersion(s string) *ResourceRevisionUpdate {
	rru.mutation.SetTemplateVersion(s)
	return rru
}

// SetNillableTemplateVersion sets the "template_version" field if the given value is not nil.
func (rru *ResourceRevisionUpdate) SetNillableTemplateVersion(s *string) *ResourceRevisionUpdate {
	if s != nil {
		rru.SetTemplateVersion(*s)
	}
	return rru
}

// SetAttributes sets the "attributes" field.
func (rru *ResourceRevisionUpdate) SetAttributes(pr property.Values) *ResourceRevisionUpdate {
	rru.mutation.SetAttributes(pr)
	return rru
}

// ClearAttributes clears the value of the "attributes" field.
func (rru *ResourceRevisionUpdate) ClearAttributes() *ResourceRevisionUpdate {
	rru.mutation.ClearAttributes()
	return rru
}

// SetComputedAttributes sets the "computed_attributes" field.
func (rru *ResourceRevisionUpdate) SetComputedAttributes(pr property.Values) *ResourceRevisionUpdate {
	rru.mutation.SetComputedAttributes(pr)
	return rru
}

// ClearComputedAttributes clears the value of the "computed_attributes" field.
func (rru *ResourceRevisionUpdate) ClearComputedAttributes() *ResourceRevisionUpdate {
	rru.mutation.ClearComputedAttributes()
	return rru
}

// SetVariables sets the "variables" field.
func (rru *ResourceRevisionUpdate) SetVariables(c crypto.Map[string, string]) *ResourceRevisionUpdate {
	rru.mutation.SetVariables(c)
	return rru
}

// SetInputPlan sets the "input_plan" field.
func (rru *ResourceRevisionUpdate) SetInputPlan(s string) *ResourceRevisionUpdate {
	rru.mutation.SetInputPlan(s)
	return rru
}

// SetNillableInputPlan sets the "input_plan" field if the given value is not nil.
func (rru *ResourceRevisionUpdate) SetNillableInputPlan(s *string) *ResourceRevisionUpdate {
	if s != nil {
		rru.SetInputPlan(*s)
	}
	return rru
}

// SetOutput sets the "output" field.
func (rru *ResourceRevisionUpdate) SetOutput(s string) *ResourceRevisionUpdate {
	rru.mutation.SetOutput(s)
	return rru
}

// SetNillableOutput sets the "output" field if the given value is not nil.
func (rru *ResourceRevisionUpdate) SetNillableOutput(s *string) *ResourceRevisionUpdate {
	if s != nil {
		rru.SetOutput(*s)
	}
	return rru
}

// SetDeployerType sets the "deployer_type" field.
func (rru *ResourceRevisionUpdate) SetDeployerType(s string) *ResourceRevisionUpdate {
	rru.mutation.SetDeployerType(s)
	return rru
}

// SetNillableDeployerType sets the "deployer_type" field if the given value is not nil.
func (rru *ResourceRevisionUpdate) SetNillableDeployerType(s *string) *ResourceRevisionUpdate {
	if s != nil {
		rru.SetDeployerType(*s)
	}
	return rru
}

// SetDuration sets the "duration" field.
func (rru *ResourceRevisionUpdate) SetDuration(i int) *ResourceRevisionUpdate {
	rru.mutation.ResetDuration()
	rru.mutation.SetDuration(i)
	return rru
}

// SetNillableDuration sets the "duration" field if the given value is not nil.
func (rru *ResourceRevisionUpdate) SetNillableDuration(i *int) *ResourceRevisionUpdate {
	if i != nil {
		rru.SetDuration(*i)
	}
	return rru
}

// AddDuration adds i to the "duration" field.
func (rru *ResourceRevisionUpdate) AddDuration(i int) *ResourceRevisionUpdate {
	rru.mutation.AddDuration(i)
	return rru
}

// SetPreviousRequiredProviders sets the "previous_required_providers" field.
func (rru *ResourceRevisionUpdate) SetPreviousRequiredProviders(tr []types.ProviderRequirement) *ResourceRevisionUpdate {
	rru.mutation.SetPreviousRequiredProviders(tr)
	return rru
}

// AppendPreviousRequiredProviders appends tr to the "previous_required_providers" field.
func (rru *ResourceRevisionUpdate) AppendPreviousRequiredProviders(tr []types.ProviderRequirement) *ResourceRevisionUpdate {
	rru.mutation.AppendPreviousRequiredProviders(tr)
	return rru
}

// SetRecord sets the "record" field.
func (rru *ResourceRevisionUpdate) SetRecord(s string) *ResourceRevisionUpdate {
	rru.mutation.SetRecord(s)
	return rru
}

// SetNillableRecord sets the "record" field if the given value is not nil.
func (rru *ResourceRevisionUpdate) SetNillableRecord(s *string) *ResourceRevisionUpdate {
	if s != nil {
		rru.SetRecord(*s)
	}
	return rru
}

// ClearRecord clears the value of the "record" field.
func (rru *ResourceRevisionUpdate) ClearRecord() *ResourceRevisionUpdate {
	rru.mutation.ClearRecord()
	return rru
}

// SetChangeComment sets the "change_comment" field.
func (rru *ResourceRevisionUpdate) SetChangeComment(s string) *ResourceRevisionUpdate {
	rru.mutation.SetChangeComment(s)
	return rru
}

// SetNillableChangeComment sets the "change_comment" field if the given value is not nil.
func (rru *ResourceRevisionUpdate) SetNillableChangeComment(s *string) *ResourceRevisionUpdate {
	if s != nil {
		rru.SetChangeComment(*s)
	}
	return rru
}

// ClearChangeComment clears the value of the "change_comment" field.
func (rru *ResourceRevisionUpdate) ClearChangeComment() *ResourceRevisionUpdate {
	rru.mutation.ClearChangeComment()
	return rru
}

// SetCreatedBy sets the "created_by" field.
func (rru *ResourceRevisionUpdate) SetCreatedBy(s string) *ResourceRevisionUpdate {
	rru.mutation.SetCreatedBy(s)
	return rru
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (rru *ResourceRevisionUpdate) SetNillableCreatedBy(s *string) *ResourceRevisionUpdate {
	if s != nil {
		rru.SetCreatedBy(*s)
	}
	return rru
}

// Mutation returns the ResourceRevisionMutation object of the builder.
func (rru *ResourceRevisionUpdate) Mutation() *ResourceRevisionMutation {
	return rru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (rru *ResourceRevisionUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, rru.sqlSave, rru.mutation, rru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rru *ResourceRevisionUpdate) SaveX(ctx context.Context) int {
	affected, err := rru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (rru *ResourceRevisionUpdate) Exec(ctx context.Context) error {
	_, err := rru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rru *ResourceRevisionUpdate) ExecX(ctx context.Context) {
	if err := rru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rru *ResourceRevisionUpdate) check() error {
	if v, ok := rru.mutation.TemplateVersion(); ok {
		if err := resourcerevision.TemplateVersionValidator(v); err != nil {
			return &ValidationError{Name: "template_version", err: fmt.Errorf(`model: validator failed for field "ResourceRevision.template_version": %w`, err)}
		}
	}
	if _, ok := rru.mutation.ProjectID(); rru.mutation.ProjectCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "ResourceRevision.project"`)
	}
	if _, ok := rru.mutation.EnvironmentID(); rru.mutation.EnvironmentCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "ResourceRevision.environment"`)
	}
	if _, ok := rru.mutation.ResourceID(); rru.mutation.ResourceCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "ResourceRevision.resource"`)
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
func (rru *ResourceRevisionUpdate) Set(obj *ResourceRevision) *ResourceRevisionUpdate {
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
	rru.SetInputPlan(obj.InputPlan)
	rru.SetOutput(obj.Output)
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
func (rru *ResourceRevisionUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ResourceRevisionUpdate {
	rru.modifiers = append(rru.modifiers, modifiers...)
	return rru
}

func (rru *ResourceRevisionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := rru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(resourcerevision.Table, resourcerevision.Columns, sqlgraph.NewFieldSpec(resourcerevision.FieldID, field.TypeString))
	if ps := rru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rru.mutation.Status(); ok {
		_spec.SetField(resourcerevision.FieldStatus, field.TypeJSON, value)
	}
	if rru.mutation.StatusCleared() {
		_spec.ClearField(resourcerevision.FieldStatus, field.TypeJSON)
	}
	if value, ok := rru.mutation.TemplateVersion(); ok {
		_spec.SetField(resourcerevision.FieldTemplateVersion, field.TypeString, value)
	}
	if value, ok := rru.mutation.Attributes(); ok {
		_spec.SetField(resourcerevision.FieldAttributes, field.TypeOther, value)
	}
	if rru.mutation.AttributesCleared() {
		_spec.ClearField(resourcerevision.FieldAttributes, field.TypeOther)
	}
	if value, ok := rru.mutation.ComputedAttributes(); ok {
		_spec.SetField(resourcerevision.FieldComputedAttributes, field.TypeOther, value)
	}
	if rru.mutation.ComputedAttributesCleared() {
		_spec.ClearField(resourcerevision.FieldComputedAttributes, field.TypeOther)
	}
	if value, ok := rru.mutation.Variables(); ok {
		_spec.SetField(resourcerevision.FieldVariables, field.TypeOther, value)
	}
	if value, ok := rru.mutation.InputPlan(); ok {
		_spec.SetField(resourcerevision.FieldInputPlan, field.TypeString, value)
	}
	if value, ok := rru.mutation.Output(); ok {
		_spec.SetField(resourcerevision.FieldOutput, field.TypeString, value)
	}
	if value, ok := rru.mutation.DeployerType(); ok {
		_spec.SetField(resourcerevision.FieldDeployerType, field.TypeString, value)
	}
	if value, ok := rru.mutation.Duration(); ok {
		_spec.SetField(resourcerevision.FieldDuration, field.TypeInt, value)
	}
	if value, ok := rru.mutation.AddedDuration(); ok {
		_spec.AddField(resourcerevision.FieldDuration, field.TypeInt, value)
	}
	if value, ok := rru.mutation.PreviousRequiredProviders(); ok {
		_spec.SetField(resourcerevision.FieldPreviousRequiredProviders, field.TypeJSON, value)
	}
	if value, ok := rru.mutation.AppendedPreviousRequiredProviders(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, resourcerevision.FieldPreviousRequiredProviders, value)
		})
	}
	if value, ok := rru.mutation.Record(); ok {
		_spec.SetField(resourcerevision.FieldRecord, field.TypeString, value)
	}
	if rru.mutation.RecordCleared() {
		_spec.ClearField(resourcerevision.FieldRecord, field.TypeString)
	}
	if value, ok := rru.mutation.ChangeComment(); ok {
		_spec.SetField(resourcerevision.FieldChangeComment, field.TypeString, value)
	}
	if rru.mutation.ChangeCommentCleared() {
		_spec.ClearField(resourcerevision.FieldChangeComment, field.TypeString)
	}
	if value, ok := rru.mutation.CreatedBy(); ok {
		_spec.SetField(resourcerevision.FieldCreatedBy, field.TypeString, value)
	}
	_spec.Node.Schema = rru.schemaConfig.ResourceRevision
	ctx = internal.NewSchemaConfigContext(ctx, rru.schemaConfig)
	_spec.AddModifiers(rru.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, rru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{resourcerevision.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	rru.mutation.done = true
	return n, nil
}

// ResourceRevisionUpdateOne is the builder for updating a single ResourceRevision entity.
type ResourceRevisionUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *ResourceRevisionMutation
	modifiers []func(*sql.UpdateBuilder)
	object    *ResourceRevision
}

// SetStatus sets the "status" field.
func (rruo *ResourceRevisionUpdateOne) SetStatus(s status.Status) *ResourceRevisionUpdateOne {
	rruo.mutation.SetStatus(s)
	return rruo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (rruo *ResourceRevisionUpdateOne) SetNillableStatus(s *status.Status) *ResourceRevisionUpdateOne {
	if s != nil {
		rruo.SetStatus(*s)
	}
	return rruo
}

// ClearStatus clears the value of the "status" field.
func (rruo *ResourceRevisionUpdateOne) ClearStatus() *ResourceRevisionUpdateOne {
	rruo.mutation.ClearStatus()
	return rruo
}

// SetTemplateVersion sets the "template_version" field.
func (rruo *ResourceRevisionUpdateOne) SetTemplateVersion(s string) *ResourceRevisionUpdateOne {
	rruo.mutation.SetTemplateVersion(s)
	return rruo
}

// SetNillableTemplateVersion sets the "template_version" field if the given value is not nil.
func (rruo *ResourceRevisionUpdateOne) SetNillableTemplateVersion(s *string) *ResourceRevisionUpdateOne {
	if s != nil {
		rruo.SetTemplateVersion(*s)
	}
	return rruo
}

// SetAttributes sets the "attributes" field.
func (rruo *ResourceRevisionUpdateOne) SetAttributes(pr property.Values) *ResourceRevisionUpdateOne {
	rruo.mutation.SetAttributes(pr)
	return rruo
}

// ClearAttributes clears the value of the "attributes" field.
func (rruo *ResourceRevisionUpdateOne) ClearAttributes() *ResourceRevisionUpdateOne {
	rruo.mutation.ClearAttributes()
	return rruo
}

// SetComputedAttributes sets the "computed_attributes" field.
func (rruo *ResourceRevisionUpdateOne) SetComputedAttributes(pr property.Values) *ResourceRevisionUpdateOne {
	rruo.mutation.SetComputedAttributes(pr)
	return rruo
}

// ClearComputedAttributes clears the value of the "computed_attributes" field.
func (rruo *ResourceRevisionUpdateOne) ClearComputedAttributes() *ResourceRevisionUpdateOne {
	rruo.mutation.ClearComputedAttributes()
	return rruo
}

// SetVariables sets the "variables" field.
func (rruo *ResourceRevisionUpdateOne) SetVariables(c crypto.Map[string, string]) *ResourceRevisionUpdateOne {
	rruo.mutation.SetVariables(c)
	return rruo
}

// SetInputPlan sets the "input_plan" field.
func (rruo *ResourceRevisionUpdateOne) SetInputPlan(s string) *ResourceRevisionUpdateOne {
	rruo.mutation.SetInputPlan(s)
	return rruo
}

// SetNillableInputPlan sets the "input_plan" field if the given value is not nil.
func (rruo *ResourceRevisionUpdateOne) SetNillableInputPlan(s *string) *ResourceRevisionUpdateOne {
	if s != nil {
		rruo.SetInputPlan(*s)
	}
	return rruo
}

// SetOutput sets the "output" field.
func (rruo *ResourceRevisionUpdateOne) SetOutput(s string) *ResourceRevisionUpdateOne {
	rruo.mutation.SetOutput(s)
	return rruo
}

// SetNillableOutput sets the "output" field if the given value is not nil.
func (rruo *ResourceRevisionUpdateOne) SetNillableOutput(s *string) *ResourceRevisionUpdateOne {
	if s != nil {
		rruo.SetOutput(*s)
	}
	return rruo
}

// SetDeployerType sets the "deployer_type" field.
func (rruo *ResourceRevisionUpdateOne) SetDeployerType(s string) *ResourceRevisionUpdateOne {
	rruo.mutation.SetDeployerType(s)
	return rruo
}

// SetNillableDeployerType sets the "deployer_type" field if the given value is not nil.
func (rruo *ResourceRevisionUpdateOne) SetNillableDeployerType(s *string) *ResourceRevisionUpdateOne {
	if s != nil {
		rruo.SetDeployerType(*s)
	}
	return rruo
}

// SetDuration sets the "duration" field.
func (rruo *ResourceRevisionUpdateOne) SetDuration(i int) *ResourceRevisionUpdateOne {
	rruo.mutation.ResetDuration()
	rruo.mutation.SetDuration(i)
	return rruo
}

// SetNillableDuration sets the "duration" field if the given value is not nil.
func (rruo *ResourceRevisionUpdateOne) SetNillableDuration(i *int) *ResourceRevisionUpdateOne {
	if i != nil {
		rruo.SetDuration(*i)
	}
	return rruo
}

// AddDuration adds i to the "duration" field.
func (rruo *ResourceRevisionUpdateOne) AddDuration(i int) *ResourceRevisionUpdateOne {
	rruo.mutation.AddDuration(i)
	return rruo
}

// SetPreviousRequiredProviders sets the "previous_required_providers" field.
func (rruo *ResourceRevisionUpdateOne) SetPreviousRequiredProviders(tr []types.ProviderRequirement) *ResourceRevisionUpdateOne {
	rruo.mutation.SetPreviousRequiredProviders(tr)
	return rruo
}

// AppendPreviousRequiredProviders appends tr to the "previous_required_providers" field.
func (rruo *ResourceRevisionUpdateOne) AppendPreviousRequiredProviders(tr []types.ProviderRequirement) *ResourceRevisionUpdateOne {
	rruo.mutation.AppendPreviousRequiredProviders(tr)
	return rruo
}

// SetRecord sets the "record" field.
func (rruo *ResourceRevisionUpdateOne) SetRecord(s string) *ResourceRevisionUpdateOne {
	rruo.mutation.SetRecord(s)
	return rruo
}

// SetNillableRecord sets the "record" field if the given value is not nil.
func (rruo *ResourceRevisionUpdateOne) SetNillableRecord(s *string) *ResourceRevisionUpdateOne {
	if s != nil {
		rruo.SetRecord(*s)
	}
	return rruo
}

// ClearRecord clears the value of the "record" field.
func (rruo *ResourceRevisionUpdateOne) ClearRecord() *ResourceRevisionUpdateOne {
	rruo.mutation.ClearRecord()
	return rruo
}

// SetChangeComment sets the "change_comment" field.
func (rruo *ResourceRevisionUpdateOne) SetChangeComment(s string) *ResourceRevisionUpdateOne {
	rruo.mutation.SetChangeComment(s)
	return rruo
}

// SetNillableChangeComment sets the "change_comment" field if the given value is not nil.
func (rruo *ResourceRevisionUpdateOne) SetNillableChangeComment(s *string) *ResourceRevisionUpdateOne {
	if s != nil {
		rruo.SetChangeComment(*s)
	}
	return rruo
}

// ClearChangeComment clears the value of the "change_comment" field.
func (rruo *ResourceRevisionUpdateOne) ClearChangeComment() *ResourceRevisionUpdateOne {
	rruo.mutation.ClearChangeComment()
	return rruo
}

// SetCreatedBy sets the "created_by" field.
func (rruo *ResourceRevisionUpdateOne) SetCreatedBy(s string) *ResourceRevisionUpdateOne {
	rruo.mutation.SetCreatedBy(s)
	return rruo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (rruo *ResourceRevisionUpdateOne) SetNillableCreatedBy(s *string) *ResourceRevisionUpdateOne {
	if s != nil {
		rruo.SetCreatedBy(*s)
	}
	return rruo
}

// Mutation returns the ResourceRevisionMutation object of the builder.
func (rruo *ResourceRevisionUpdateOne) Mutation() *ResourceRevisionMutation {
	return rruo.mutation
}

// Where appends a list predicates to the ResourceRevisionUpdate builder.
func (rruo *ResourceRevisionUpdateOne) Where(ps ...predicate.ResourceRevision) *ResourceRevisionUpdateOne {
	rruo.mutation.Where(ps...)
	return rruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (rruo *ResourceRevisionUpdateOne) Select(field string, fields ...string) *ResourceRevisionUpdateOne {
	rruo.fields = append([]string{field}, fields...)
	return rruo
}

// Save executes the query and returns the updated ResourceRevision entity.
func (rruo *ResourceRevisionUpdateOne) Save(ctx context.Context) (*ResourceRevision, error) {
	return withHooks(ctx, rruo.sqlSave, rruo.mutation, rruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rruo *ResourceRevisionUpdateOne) SaveX(ctx context.Context) *ResourceRevision {
	node, err := rruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (rruo *ResourceRevisionUpdateOne) Exec(ctx context.Context) error {
	_, err := rruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rruo *ResourceRevisionUpdateOne) ExecX(ctx context.Context) {
	if err := rruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rruo *ResourceRevisionUpdateOne) check() error {
	if v, ok := rruo.mutation.TemplateVersion(); ok {
		if err := resourcerevision.TemplateVersionValidator(v); err != nil {
			return &ValidationError{Name: "template_version", err: fmt.Errorf(`model: validator failed for field "ResourceRevision.template_version": %w`, err)}
		}
	}
	if _, ok := rruo.mutation.ProjectID(); rruo.mutation.ProjectCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "ResourceRevision.project"`)
	}
	if _, ok := rruo.mutation.EnvironmentID(); rruo.mutation.EnvironmentCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "ResourceRevision.environment"`)
	}
	if _, ok := rruo.mutation.ResourceID(); rruo.mutation.ResourceCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "ResourceRevision.resource"`)
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
func (rruo *ResourceRevisionUpdateOne) Set(obj *ResourceRevision) *ResourceRevisionUpdateOne {
	h := func(n ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			mt := m.(*ResourceRevisionMutation)
			db, err := mt.Client().ResourceRevision.Get(ctx, *mt.id)
			if err != nil {
				return nil, fmt.Errorf("failed getting ResourceRevision with id: %v", *mt.id)
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
			if db.InputPlan != obj.InputPlan {
				rruo.SetInputPlan(obj.InputPlan)
			}
			if db.Output != obj.Output {
				rruo.SetOutput(obj.Output)
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
func (rruo *ResourceRevisionUpdateOne) getClientSet() (mc ClientSet) {
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

// SaveE calls the given function after updated the ResourceRevision entity,
// which is always good for cascading update operations.
func (rruo *ResourceRevisionUpdateOne) SaveE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *ResourceRevision) error) (*ResourceRevision, error) {
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
		if _, set := rruo.mutation.Field(resourcerevision.FieldStatus); set {
			obj.Status = x.Status
		}
		if _, set := rruo.mutation.Field(resourcerevision.FieldTemplateVersion); set {
			obj.TemplateVersion = x.TemplateVersion
		}
		if _, set := rruo.mutation.Field(resourcerevision.FieldAttributes); set {
			obj.Attributes = x.Attributes
		}
		if _, set := rruo.mutation.Field(resourcerevision.FieldComputedAttributes); set {
			obj.ComputedAttributes = x.ComputedAttributes
		}
		if _, set := rruo.mutation.Field(resourcerevision.FieldVariables); set {
			obj.Variables = x.Variables
		}
		if _, set := rruo.mutation.Field(resourcerevision.FieldInputPlan); set {
			obj.InputPlan = x.InputPlan
		}
		if _, set := rruo.mutation.Field(resourcerevision.FieldOutput); set {
			obj.Output = x.Output
		}
		if _, set := rruo.mutation.Field(resourcerevision.FieldDeployerType); set {
			obj.DeployerType = x.DeployerType
		}
		if _, set := rruo.mutation.Field(resourcerevision.FieldDuration); set {
			obj.Duration = x.Duration
		}
		if _, set := rruo.mutation.Field(resourcerevision.FieldPreviousRequiredProviders); set {
			obj.PreviousRequiredProviders = x.PreviousRequiredProviders
		}
		if _, set := rruo.mutation.Field(resourcerevision.FieldRecord); set {
			obj.Record = x.Record
		}
		if _, set := rruo.mutation.Field(resourcerevision.FieldChangeComment); set {
			obj.ChangeComment = x.ChangeComment
		}
		if _, set := rruo.mutation.Field(resourcerevision.FieldCreatedBy); set {
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
func (rruo *ResourceRevisionUpdateOne) SaveEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *ResourceRevision) error) *ResourceRevision {
	obj, err := rruo.SaveE(ctx, cbs...)
	if err != nil {
		panic(err)
	}
	return obj
}

// ExecE calls the given function after executed the query,
// which is always good for cascading update operations.
func (rruo *ResourceRevisionUpdateOne) ExecE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *ResourceRevision) error) error {
	_, err := rruo.SaveE(ctx, cbs...)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rruo *ResourceRevisionUpdateOne) ExecEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *ResourceRevision) error) {
	if err := rruo.ExecE(ctx, cbs...); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (rruo *ResourceRevisionUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ResourceRevisionUpdateOne {
	rruo.modifiers = append(rruo.modifiers, modifiers...)
	return rruo
}

func (rruo *ResourceRevisionUpdateOne) sqlSave(ctx context.Context) (_node *ResourceRevision, err error) {
	if err := rruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(resourcerevision.Table, resourcerevision.Columns, sqlgraph.NewFieldSpec(resourcerevision.FieldID, field.TypeString))
	id, ok := rruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`model: missing "ResourceRevision.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := rruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, resourcerevision.FieldID)
		for _, f := range fields {
			if !resourcerevision.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
			}
			if f != resourcerevision.FieldID {
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
		_spec.SetField(resourcerevision.FieldStatus, field.TypeJSON, value)
	}
	if rruo.mutation.StatusCleared() {
		_spec.ClearField(resourcerevision.FieldStatus, field.TypeJSON)
	}
	if value, ok := rruo.mutation.TemplateVersion(); ok {
		_spec.SetField(resourcerevision.FieldTemplateVersion, field.TypeString, value)
	}
	if value, ok := rruo.mutation.Attributes(); ok {
		_spec.SetField(resourcerevision.FieldAttributes, field.TypeOther, value)
	}
	if rruo.mutation.AttributesCleared() {
		_spec.ClearField(resourcerevision.FieldAttributes, field.TypeOther)
	}
	if value, ok := rruo.mutation.ComputedAttributes(); ok {
		_spec.SetField(resourcerevision.FieldComputedAttributes, field.TypeOther, value)
	}
	if rruo.mutation.ComputedAttributesCleared() {
		_spec.ClearField(resourcerevision.FieldComputedAttributes, field.TypeOther)
	}
	if value, ok := rruo.mutation.Variables(); ok {
		_spec.SetField(resourcerevision.FieldVariables, field.TypeOther, value)
	}
	if value, ok := rruo.mutation.InputPlan(); ok {
		_spec.SetField(resourcerevision.FieldInputPlan, field.TypeString, value)
	}
	if value, ok := rruo.mutation.Output(); ok {
		_spec.SetField(resourcerevision.FieldOutput, field.TypeString, value)
	}
	if value, ok := rruo.mutation.DeployerType(); ok {
		_spec.SetField(resourcerevision.FieldDeployerType, field.TypeString, value)
	}
	if value, ok := rruo.mutation.Duration(); ok {
		_spec.SetField(resourcerevision.FieldDuration, field.TypeInt, value)
	}
	if value, ok := rruo.mutation.AddedDuration(); ok {
		_spec.AddField(resourcerevision.FieldDuration, field.TypeInt, value)
	}
	if value, ok := rruo.mutation.PreviousRequiredProviders(); ok {
		_spec.SetField(resourcerevision.FieldPreviousRequiredProviders, field.TypeJSON, value)
	}
	if value, ok := rruo.mutation.AppendedPreviousRequiredProviders(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, resourcerevision.FieldPreviousRequiredProviders, value)
		})
	}
	if value, ok := rruo.mutation.Record(); ok {
		_spec.SetField(resourcerevision.FieldRecord, field.TypeString, value)
	}
	if rruo.mutation.RecordCleared() {
		_spec.ClearField(resourcerevision.FieldRecord, field.TypeString)
	}
	if value, ok := rruo.mutation.ChangeComment(); ok {
		_spec.SetField(resourcerevision.FieldChangeComment, field.TypeString, value)
	}
	if rruo.mutation.ChangeCommentCleared() {
		_spec.ClearField(resourcerevision.FieldChangeComment, field.TypeString)
	}
	if value, ok := rruo.mutation.CreatedBy(); ok {
		_spec.SetField(resourcerevision.FieldCreatedBy, field.TypeString, value)
	}
	_spec.Node.Schema = rruo.schemaConfig.ResourceRevision
	ctx = internal.NewSchemaConfigContext(ctx, rruo.schemaConfig)
	_spec.AddModifiers(rruo.modifiers...)
	_node = &ResourceRevision{config: rruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, rruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{resourcerevision.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	rruo.mutation.done = true
	return _node, nil
}