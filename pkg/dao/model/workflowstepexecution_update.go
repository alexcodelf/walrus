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
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"github.com/seal-io/walrus/pkg/dao/model/internal"
	"github.com/seal-io/walrus/pkg/dao/model/predicate"
	"github.com/seal-io/walrus/pkg/dao/model/workflowstepexecution"
	"github.com/seal-io/walrus/pkg/dao/types/status"
)

// WorkflowStepExecutionUpdate is the builder for updating WorkflowStepExecution entities.
type WorkflowStepExecutionUpdate struct {
	config
	hooks     []Hook
	mutation  *WorkflowStepExecutionMutation
	modifiers []func(*sql.UpdateBuilder)
	object    *WorkflowStepExecution
}

// Where appends a list predicates to the WorkflowStepExecutionUpdate builder.
func (wseu *WorkflowStepExecutionUpdate) Where(ps ...predicate.WorkflowStepExecution) *WorkflowStepExecutionUpdate {
	wseu.mutation.Where(ps...)
	return wseu
}

// SetDescription sets the "description" field.
func (wseu *WorkflowStepExecutionUpdate) SetDescription(s string) *WorkflowStepExecutionUpdate {
	wseu.mutation.SetDescription(s)
	return wseu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (wseu *WorkflowStepExecutionUpdate) SetNillableDescription(s *string) *WorkflowStepExecutionUpdate {
	if s != nil {
		wseu.SetDescription(*s)
	}
	return wseu
}

// ClearDescription clears the value of the "description" field.
func (wseu *WorkflowStepExecutionUpdate) ClearDescription() *WorkflowStepExecutionUpdate {
	wseu.mutation.ClearDescription()
	return wseu
}

// SetLabels sets the "labels" field.
func (wseu *WorkflowStepExecutionUpdate) SetLabels(m map[string]string) *WorkflowStepExecutionUpdate {
	wseu.mutation.SetLabels(m)
	return wseu
}

// ClearLabels clears the value of the "labels" field.
func (wseu *WorkflowStepExecutionUpdate) ClearLabels() *WorkflowStepExecutionUpdate {
	wseu.mutation.ClearLabels()
	return wseu
}

// SetAnnotations sets the "annotations" field.
func (wseu *WorkflowStepExecutionUpdate) SetAnnotations(m map[string]string) *WorkflowStepExecutionUpdate {
	wseu.mutation.SetAnnotations(m)
	return wseu
}

// ClearAnnotations clears the value of the "annotations" field.
func (wseu *WorkflowStepExecutionUpdate) ClearAnnotations() *WorkflowStepExecutionUpdate {
	wseu.mutation.ClearAnnotations()
	return wseu
}

// SetUpdateTime sets the "update_time" field.
func (wseu *WorkflowStepExecutionUpdate) SetUpdateTime(t time.Time) *WorkflowStepExecutionUpdate {
	wseu.mutation.SetUpdateTime(t)
	return wseu
}

// SetStatus sets the "status" field.
func (wseu *WorkflowStepExecutionUpdate) SetStatus(s status.Status) *WorkflowStepExecutionUpdate {
	wseu.mutation.SetStatus(s)
	return wseu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (wseu *WorkflowStepExecutionUpdate) SetNillableStatus(s *status.Status) *WorkflowStepExecutionUpdate {
	if s != nil {
		wseu.SetStatus(*s)
	}
	return wseu
}

// ClearStatus clears the value of the "status" field.
func (wseu *WorkflowStepExecutionUpdate) ClearStatus() *WorkflowStepExecutionUpdate {
	wseu.mutation.ClearStatus()
	return wseu
}

// SetSpec sets the "spec" field.
func (wseu *WorkflowStepExecutionUpdate) SetSpec(m map[string]any) *WorkflowStepExecutionUpdate {
	wseu.mutation.SetSpec(m)
	return wseu
}

// ClearSpec clears the value of the "spec" field.
func (wseu *WorkflowStepExecutionUpdate) ClearSpec() *WorkflowStepExecutionUpdate {
	wseu.mutation.ClearSpec()
	return wseu
}

// SetTimes sets the "times" field.
func (wseu *WorkflowStepExecutionUpdate) SetTimes(i int) *WorkflowStepExecutionUpdate {
	wseu.mutation.ResetTimes()
	wseu.mutation.SetTimes(i)
	return wseu
}

// SetNillableTimes sets the "times" field if the given value is not nil.
func (wseu *WorkflowStepExecutionUpdate) SetNillableTimes(i *int) *WorkflowStepExecutionUpdate {
	if i != nil {
		wseu.SetTimes(*i)
	}
	return wseu
}

// AddTimes adds i to the "times" field.
func (wseu *WorkflowStepExecutionUpdate) AddTimes(i int) *WorkflowStepExecutionUpdate {
	wseu.mutation.AddTimes(i)
	return wseu
}

// SetDuration sets the "duration" field.
func (wseu *WorkflowStepExecutionUpdate) SetDuration(i int) *WorkflowStepExecutionUpdate {
	wseu.mutation.ResetDuration()
	wseu.mutation.SetDuration(i)
	return wseu
}

// SetNillableDuration sets the "duration" field if the given value is not nil.
func (wseu *WorkflowStepExecutionUpdate) SetNillableDuration(i *int) *WorkflowStepExecutionUpdate {
	if i != nil {
		wseu.SetDuration(*i)
	}
	return wseu
}

// AddDuration adds i to the "duration" field.
func (wseu *WorkflowStepExecutionUpdate) AddDuration(i int) *WorkflowStepExecutionUpdate {
	wseu.mutation.AddDuration(i)
	return wseu
}

// SetRecord sets the "record" field.
func (wseu *WorkflowStepExecutionUpdate) SetRecord(s string) *WorkflowStepExecutionUpdate {
	wseu.mutation.SetRecord(s)
	return wseu
}

// SetNillableRecord sets the "record" field if the given value is not nil.
func (wseu *WorkflowStepExecutionUpdate) SetNillableRecord(s *string) *WorkflowStepExecutionUpdate {
	if s != nil {
		wseu.SetRecord(*s)
	}
	return wseu
}

// SetInput sets the "input" field.
func (wseu *WorkflowStepExecutionUpdate) SetInput(s string) *WorkflowStepExecutionUpdate {
	wseu.mutation.SetInput(s)
	return wseu
}

// SetNillableInput sets the "input" field if the given value is not nil.
func (wseu *WorkflowStepExecutionUpdate) SetNillableInput(s *string) *WorkflowStepExecutionUpdate {
	if s != nil {
		wseu.SetInput(*s)
	}
	return wseu
}

// Mutation returns the WorkflowStepExecutionMutation object of the builder.
func (wseu *WorkflowStepExecutionUpdate) Mutation() *WorkflowStepExecutionMutation {
	return wseu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wseu *WorkflowStepExecutionUpdate) Save(ctx context.Context) (int, error) {
	if err := wseu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, wseu.sqlSave, wseu.mutation, wseu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wseu *WorkflowStepExecutionUpdate) SaveX(ctx context.Context) int {
	affected, err := wseu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wseu *WorkflowStepExecutionUpdate) Exec(ctx context.Context) error {
	_, err := wseu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wseu *WorkflowStepExecutionUpdate) ExecX(ctx context.Context) {
	if err := wseu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wseu *WorkflowStepExecutionUpdate) defaults() error {
	if _, ok := wseu.mutation.UpdateTime(); !ok {
		if workflowstepexecution.UpdateDefaultUpdateTime == nil {
			return fmt.Errorf("model: uninitialized workflowstepexecution.UpdateDefaultUpdateTime (forgotten import model/runtime?)")
		}
		v := workflowstepexecution.UpdateDefaultUpdateTime()
		wseu.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (wseu *WorkflowStepExecutionUpdate) check() error {
	if v, ok := wseu.mutation.Times(); ok {
		if err := workflowstepexecution.TimesValidator(v); err != nil {
			return &ValidationError{Name: "times", err: fmt.Errorf(`model: validator failed for field "WorkflowStepExecution.times": %w`, err)}
		}
	}
	if v, ok := wseu.mutation.Duration(); ok {
		if err := workflowstepexecution.DurationValidator(v); err != nil {
			return &ValidationError{Name: "duration", err: fmt.Errorf(`model: validator failed for field "WorkflowStepExecution.duration": %w`, err)}
		}
	}
	if _, ok := wseu.mutation.ProjectID(); wseu.mutation.ProjectCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "WorkflowStepExecution.project"`)
	}
	if _, ok := wseu.mutation.StepID(); wseu.mutation.StepCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "WorkflowStepExecution.step"`)
	}
	if _, ok := wseu.mutation.StageExecutionID(); wseu.mutation.StageExecutionCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "WorkflowStepExecution.stage_execution"`)
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
func (wseu *WorkflowStepExecutionUpdate) Set(obj *WorkflowStepExecution) *WorkflowStepExecutionUpdate {
	// Without Default.
	if obj.Description != "" {
		wseu.SetDescription(obj.Description)
	} else {
		wseu.ClearDescription()
	}
	if !reflect.ValueOf(obj.Labels).IsZero() {
		wseu.SetLabels(obj.Labels)
	} else {
		wseu.ClearLabels()
	}
	if !reflect.ValueOf(obj.Annotations).IsZero() {
		wseu.SetAnnotations(obj.Annotations)
	}
	if !reflect.ValueOf(obj.Status).IsZero() {
		wseu.SetStatus(obj.Status)
	}
	if !reflect.ValueOf(obj.Spec).IsZero() {
		wseu.SetSpec(obj.Spec)
	} else {
		wseu.ClearSpec()
	}
	wseu.SetTimes(obj.Times)
	wseu.SetDuration(obj.Duration)
	wseu.SetRecord(obj.Record)
	wseu.SetInput(obj.Input)

	// With Default.
	if obj.UpdateTime != nil {
		wseu.SetUpdateTime(*obj.UpdateTime)
	}

	// Record the given object.
	wseu.object = obj

	return wseu
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (wseu *WorkflowStepExecutionUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *WorkflowStepExecutionUpdate {
	wseu.modifiers = append(wseu.modifiers, modifiers...)
	return wseu
}

func (wseu *WorkflowStepExecutionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := wseu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(workflowstepexecution.Table, workflowstepexecution.Columns, sqlgraph.NewFieldSpec(workflowstepexecution.FieldID, field.TypeString))
	if ps := wseu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wseu.mutation.Description(); ok {
		_spec.SetField(workflowstepexecution.FieldDescription, field.TypeString, value)
	}
	if wseu.mutation.DescriptionCleared() {
		_spec.ClearField(workflowstepexecution.FieldDescription, field.TypeString)
	}
	if value, ok := wseu.mutation.Labels(); ok {
		_spec.SetField(workflowstepexecution.FieldLabels, field.TypeJSON, value)
	}
	if wseu.mutation.LabelsCleared() {
		_spec.ClearField(workflowstepexecution.FieldLabels, field.TypeJSON)
	}
	if value, ok := wseu.mutation.Annotations(); ok {
		_spec.SetField(workflowstepexecution.FieldAnnotations, field.TypeJSON, value)
	}
	if wseu.mutation.AnnotationsCleared() {
		_spec.ClearField(workflowstepexecution.FieldAnnotations, field.TypeJSON)
	}
	if value, ok := wseu.mutation.UpdateTime(); ok {
		_spec.SetField(workflowstepexecution.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := wseu.mutation.Status(); ok {
		_spec.SetField(workflowstepexecution.FieldStatus, field.TypeJSON, value)
	}
	if wseu.mutation.StatusCleared() {
		_spec.ClearField(workflowstepexecution.FieldStatus, field.TypeJSON)
	}
	if value, ok := wseu.mutation.Spec(); ok {
		_spec.SetField(workflowstepexecution.FieldSpec, field.TypeJSON, value)
	}
	if wseu.mutation.SpecCleared() {
		_spec.ClearField(workflowstepexecution.FieldSpec, field.TypeJSON)
	}
	if value, ok := wseu.mutation.Times(); ok {
		_spec.SetField(workflowstepexecution.FieldTimes, field.TypeInt, value)
	}
	if value, ok := wseu.mutation.AddedTimes(); ok {
		_spec.AddField(workflowstepexecution.FieldTimes, field.TypeInt, value)
	}
	if value, ok := wseu.mutation.Duration(); ok {
		_spec.SetField(workflowstepexecution.FieldDuration, field.TypeInt, value)
	}
	if value, ok := wseu.mutation.AddedDuration(); ok {
		_spec.AddField(workflowstepexecution.FieldDuration, field.TypeInt, value)
	}
	if value, ok := wseu.mutation.Record(); ok {
		_spec.SetField(workflowstepexecution.FieldRecord, field.TypeString, value)
	}
	if value, ok := wseu.mutation.Input(); ok {
		_spec.SetField(workflowstepexecution.FieldInput, field.TypeString, value)
	}
	_spec.Node.Schema = wseu.schemaConfig.WorkflowStepExecution
	ctx = internal.NewSchemaConfigContext(ctx, wseu.schemaConfig)
	_spec.AddModifiers(wseu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, wseu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{workflowstepexecution.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	wseu.mutation.done = true
	return n, nil
}

// WorkflowStepExecutionUpdateOne is the builder for updating a single WorkflowStepExecution entity.
type WorkflowStepExecutionUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *WorkflowStepExecutionMutation
	modifiers []func(*sql.UpdateBuilder)
	object    *WorkflowStepExecution
}

// SetDescription sets the "description" field.
func (wseuo *WorkflowStepExecutionUpdateOne) SetDescription(s string) *WorkflowStepExecutionUpdateOne {
	wseuo.mutation.SetDescription(s)
	return wseuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (wseuo *WorkflowStepExecutionUpdateOne) SetNillableDescription(s *string) *WorkflowStepExecutionUpdateOne {
	if s != nil {
		wseuo.SetDescription(*s)
	}
	return wseuo
}

// ClearDescription clears the value of the "description" field.
func (wseuo *WorkflowStepExecutionUpdateOne) ClearDescription() *WorkflowStepExecutionUpdateOne {
	wseuo.mutation.ClearDescription()
	return wseuo
}

// SetLabels sets the "labels" field.
func (wseuo *WorkflowStepExecutionUpdateOne) SetLabels(m map[string]string) *WorkflowStepExecutionUpdateOne {
	wseuo.mutation.SetLabels(m)
	return wseuo
}

// ClearLabels clears the value of the "labels" field.
func (wseuo *WorkflowStepExecutionUpdateOne) ClearLabels() *WorkflowStepExecutionUpdateOne {
	wseuo.mutation.ClearLabels()
	return wseuo
}

// SetAnnotations sets the "annotations" field.
func (wseuo *WorkflowStepExecutionUpdateOne) SetAnnotations(m map[string]string) *WorkflowStepExecutionUpdateOne {
	wseuo.mutation.SetAnnotations(m)
	return wseuo
}

// ClearAnnotations clears the value of the "annotations" field.
func (wseuo *WorkflowStepExecutionUpdateOne) ClearAnnotations() *WorkflowStepExecutionUpdateOne {
	wseuo.mutation.ClearAnnotations()
	return wseuo
}

// SetUpdateTime sets the "update_time" field.
func (wseuo *WorkflowStepExecutionUpdateOne) SetUpdateTime(t time.Time) *WorkflowStepExecutionUpdateOne {
	wseuo.mutation.SetUpdateTime(t)
	return wseuo
}

// SetStatus sets the "status" field.
func (wseuo *WorkflowStepExecutionUpdateOne) SetStatus(s status.Status) *WorkflowStepExecutionUpdateOne {
	wseuo.mutation.SetStatus(s)
	return wseuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (wseuo *WorkflowStepExecutionUpdateOne) SetNillableStatus(s *status.Status) *WorkflowStepExecutionUpdateOne {
	if s != nil {
		wseuo.SetStatus(*s)
	}
	return wseuo
}

// ClearStatus clears the value of the "status" field.
func (wseuo *WorkflowStepExecutionUpdateOne) ClearStatus() *WorkflowStepExecutionUpdateOne {
	wseuo.mutation.ClearStatus()
	return wseuo
}

// SetSpec sets the "spec" field.
func (wseuo *WorkflowStepExecutionUpdateOne) SetSpec(m map[string]any) *WorkflowStepExecutionUpdateOne {
	wseuo.mutation.SetSpec(m)
	return wseuo
}

// ClearSpec clears the value of the "spec" field.
func (wseuo *WorkflowStepExecutionUpdateOne) ClearSpec() *WorkflowStepExecutionUpdateOne {
	wseuo.mutation.ClearSpec()
	return wseuo
}

// SetTimes sets the "times" field.
func (wseuo *WorkflowStepExecutionUpdateOne) SetTimes(i int) *WorkflowStepExecutionUpdateOne {
	wseuo.mutation.ResetTimes()
	wseuo.mutation.SetTimes(i)
	return wseuo
}

// SetNillableTimes sets the "times" field if the given value is not nil.
func (wseuo *WorkflowStepExecutionUpdateOne) SetNillableTimes(i *int) *WorkflowStepExecutionUpdateOne {
	if i != nil {
		wseuo.SetTimes(*i)
	}
	return wseuo
}

// AddTimes adds i to the "times" field.
func (wseuo *WorkflowStepExecutionUpdateOne) AddTimes(i int) *WorkflowStepExecutionUpdateOne {
	wseuo.mutation.AddTimes(i)
	return wseuo
}

// SetDuration sets the "duration" field.
func (wseuo *WorkflowStepExecutionUpdateOne) SetDuration(i int) *WorkflowStepExecutionUpdateOne {
	wseuo.mutation.ResetDuration()
	wseuo.mutation.SetDuration(i)
	return wseuo
}

// SetNillableDuration sets the "duration" field if the given value is not nil.
func (wseuo *WorkflowStepExecutionUpdateOne) SetNillableDuration(i *int) *WorkflowStepExecutionUpdateOne {
	if i != nil {
		wseuo.SetDuration(*i)
	}
	return wseuo
}

// AddDuration adds i to the "duration" field.
func (wseuo *WorkflowStepExecutionUpdateOne) AddDuration(i int) *WorkflowStepExecutionUpdateOne {
	wseuo.mutation.AddDuration(i)
	return wseuo
}

// SetRecord sets the "record" field.
func (wseuo *WorkflowStepExecutionUpdateOne) SetRecord(s string) *WorkflowStepExecutionUpdateOne {
	wseuo.mutation.SetRecord(s)
	return wseuo
}

// SetNillableRecord sets the "record" field if the given value is not nil.
func (wseuo *WorkflowStepExecutionUpdateOne) SetNillableRecord(s *string) *WorkflowStepExecutionUpdateOne {
	if s != nil {
		wseuo.SetRecord(*s)
	}
	return wseuo
}

// SetInput sets the "input" field.
func (wseuo *WorkflowStepExecutionUpdateOne) SetInput(s string) *WorkflowStepExecutionUpdateOne {
	wseuo.mutation.SetInput(s)
	return wseuo
}

// SetNillableInput sets the "input" field if the given value is not nil.
func (wseuo *WorkflowStepExecutionUpdateOne) SetNillableInput(s *string) *WorkflowStepExecutionUpdateOne {
	if s != nil {
		wseuo.SetInput(*s)
	}
	return wseuo
}

// Mutation returns the WorkflowStepExecutionMutation object of the builder.
func (wseuo *WorkflowStepExecutionUpdateOne) Mutation() *WorkflowStepExecutionMutation {
	return wseuo.mutation
}

// Where appends a list predicates to the WorkflowStepExecutionUpdate builder.
func (wseuo *WorkflowStepExecutionUpdateOne) Where(ps ...predicate.WorkflowStepExecution) *WorkflowStepExecutionUpdateOne {
	wseuo.mutation.Where(ps...)
	return wseuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wseuo *WorkflowStepExecutionUpdateOne) Select(field string, fields ...string) *WorkflowStepExecutionUpdateOne {
	wseuo.fields = append([]string{field}, fields...)
	return wseuo
}

// Save executes the query and returns the updated WorkflowStepExecution entity.
func (wseuo *WorkflowStepExecutionUpdateOne) Save(ctx context.Context) (*WorkflowStepExecution, error) {
	if err := wseuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, wseuo.sqlSave, wseuo.mutation, wseuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wseuo *WorkflowStepExecutionUpdateOne) SaveX(ctx context.Context) *WorkflowStepExecution {
	node, err := wseuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wseuo *WorkflowStepExecutionUpdateOne) Exec(ctx context.Context) error {
	_, err := wseuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wseuo *WorkflowStepExecutionUpdateOne) ExecX(ctx context.Context) {
	if err := wseuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wseuo *WorkflowStepExecutionUpdateOne) defaults() error {
	if _, ok := wseuo.mutation.UpdateTime(); !ok {
		if workflowstepexecution.UpdateDefaultUpdateTime == nil {
			return fmt.Errorf("model: uninitialized workflowstepexecution.UpdateDefaultUpdateTime (forgotten import model/runtime?)")
		}
		v := workflowstepexecution.UpdateDefaultUpdateTime()
		wseuo.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (wseuo *WorkflowStepExecutionUpdateOne) check() error {
	if v, ok := wseuo.mutation.Times(); ok {
		if err := workflowstepexecution.TimesValidator(v); err != nil {
			return &ValidationError{Name: "times", err: fmt.Errorf(`model: validator failed for field "WorkflowStepExecution.times": %w`, err)}
		}
	}
	if v, ok := wseuo.mutation.Duration(); ok {
		if err := workflowstepexecution.DurationValidator(v); err != nil {
			return &ValidationError{Name: "duration", err: fmt.Errorf(`model: validator failed for field "WorkflowStepExecution.duration": %w`, err)}
		}
	}
	if _, ok := wseuo.mutation.ProjectID(); wseuo.mutation.ProjectCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "WorkflowStepExecution.project"`)
	}
	if _, ok := wseuo.mutation.StepID(); wseuo.mutation.StepCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "WorkflowStepExecution.step"`)
	}
	if _, ok := wseuo.mutation.StageExecutionID(); wseuo.mutation.StageExecutionCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "WorkflowStepExecution.stage_execution"`)
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
func (wseuo *WorkflowStepExecutionUpdateOne) Set(obj *WorkflowStepExecution) *WorkflowStepExecutionUpdateOne {
	h := func(n ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			mt := m.(*WorkflowStepExecutionMutation)
			db, err := mt.Client().WorkflowStepExecution.Get(ctx, *mt.id)
			if err != nil {
				return nil, fmt.Errorf("failed getting WorkflowStepExecution with id: %v", *mt.id)
			}

			// Without Default.
			if obj.Description != "" {
				if db.Description != obj.Description {
					wseuo.SetDescription(obj.Description)
				}
			} else {
				wseuo.ClearDescription()
			}
			if !reflect.ValueOf(obj.Labels).IsZero() {
				if !reflect.DeepEqual(db.Labels, obj.Labels) {
					wseuo.SetLabels(obj.Labels)
				}
			} else {
				wseuo.ClearLabels()
			}
			if !reflect.ValueOf(obj.Annotations).IsZero() {
				if !reflect.DeepEqual(db.Annotations, obj.Annotations) {
					wseuo.SetAnnotations(obj.Annotations)
				}
			}
			if !reflect.ValueOf(obj.Status).IsZero() {
				if !db.Status.Equal(obj.Status) {
					wseuo.SetStatus(obj.Status)
				}
			}
			if !reflect.ValueOf(obj.Spec).IsZero() {
				if !reflect.DeepEqual(db.Spec, obj.Spec) {
					wseuo.SetSpec(obj.Spec)
				}
			} else {
				wseuo.ClearSpec()
			}
			if db.Times != obj.Times {
				wseuo.SetTimes(obj.Times)
			}
			if db.Duration != obj.Duration {
				wseuo.SetDuration(obj.Duration)
			}
			if db.Record != obj.Record {
				wseuo.SetRecord(obj.Record)
			}
			if db.Input != obj.Input {
				wseuo.SetInput(obj.Input)
			}

			// With Default.
			if (obj.UpdateTime != nil) && (!reflect.DeepEqual(db.UpdateTime, obj.UpdateTime)) {
				wseuo.SetUpdateTime(*obj.UpdateTime)
			}

			// Record the given object.
			wseuo.object = obj

			return n.Mutate(ctx, m)
		})
	}

	wseuo.hooks = append(wseuo.hooks, h)

	return wseuo
}

// getClientSet returns the ClientSet for the given builder.
func (wseuo *WorkflowStepExecutionUpdateOne) getClientSet() (mc ClientSet) {
	if _, ok := wseuo.config.driver.(*txDriver); ok {
		tx := &Tx{config: wseuo.config}
		tx.init()
		mc = tx
	} else {
		cli := &Client{config: wseuo.config}
		cli.init()
		mc = cli
	}
	return mc
}

// SaveE calls the given function after updated the WorkflowStepExecution entity,
// which is always good for cascading update operations.
func (wseuo *WorkflowStepExecutionUpdateOne) SaveE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *WorkflowStepExecution) error) (*WorkflowStepExecution, error) {
	obj, err := wseuo.Save(ctx)
	if err != nil &&
		(wseuo.object == nil || !errors.Is(err, stdsql.ErrNoRows)) {
		return nil, err
	}

	if len(cbs) == 0 {
		return obj, err
	}

	mc := wseuo.getClientSet()

	if obj == nil {
		obj = wseuo.object
	} else if x := wseuo.object; x != nil {
		if _, set := wseuo.mutation.Field(workflowstepexecution.FieldDescription); set {
			obj.Description = x.Description
		}
		if _, set := wseuo.mutation.Field(workflowstepexecution.FieldLabels); set {
			obj.Labels = x.Labels
		}
		if _, set := wseuo.mutation.Field(workflowstepexecution.FieldAnnotations); set {
			obj.Annotations = x.Annotations
		}
		if _, set := wseuo.mutation.Field(workflowstepexecution.FieldStatus); set {
			obj.Status = x.Status
		}
		if _, set := wseuo.mutation.Field(workflowstepexecution.FieldSpec); set {
			obj.Spec = x.Spec
		}
		if _, set := wseuo.mutation.Field(workflowstepexecution.FieldTimes); set {
			obj.Times = x.Times
		}
		if _, set := wseuo.mutation.Field(workflowstepexecution.FieldDuration); set {
			obj.Duration = x.Duration
		}
		if _, set := wseuo.mutation.Field(workflowstepexecution.FieldRecord); set {
			obj.Record = x.Record
		}
		if _, set := wseuo.mutation.Field(workflowstepexecution.FieldInput); set {
			obj.Input = x.Input
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
func (wseuo *WorkflowStepExecutionUpdateOne) SaveEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *WorkflowStepExecution) error) *WorkflowStepExecution {
	obj, err := wseuo.SaveE(ctx, cbs...)
	if err != nil {
		panic(err)
	}
	return obj
}

// ExecE calls the given function after executed the query,
// which is always good for cascading update operations.
func (wseuo *WorkflowStepExecutionUpdateOne) ExecE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *WorkflowStepExecution) error) error {
	_, err := wseuo.SaveE(ctx, cbs...)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wseuo *WorkflowStepExecutionUpdateOne) ExecEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *WorkflowStepExecution) error) {
	if err := wseuo.ExecE(ctx, cbs...); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (wseuo *WorkflowStepExecutionUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *WorkflowStepExecutionUpdateOne {
	wseuo.modifiers = append(wseuo.modifiers, modifiers...)
	return wseuo
}

func (wseuo *WorkflowStepExecutionUpdateOne) sqlSave(ctx context.Context) (_node *WorkflowStepExecution, err error) {
	if err := wseuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(workflowstepexecution.Table, workflowstepexecution.Columns, sqlgraph.NewFieldSpec(workflowstepexecution.FieldID, field.TypeString))
	id, ok := wseuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`model: missing "WorkflowStepExecution.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := wseuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, workflowstepexecution.FieldID)
		for _, f := range fields {
			if !workflowstepexecution.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
			}
			if f != workflowstepexecution.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wseuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wseuo.mutation.Description(); ok {
		_spec.SetField(workflowstepexecution.FieldDescription, field.TypeString, value)
	}
	if wseuo.mutation.DescriptionCleared() {
		_spec.ClearField(workflowstepexecution.FieldDescription, field.TypeString)
	}
	if value, ok := wseuo.mutation.Labels(); ok {
		_spec.SetField(workflowstepexecution.FieldLabels, field.TypeJSON, value)
	}
	if wseuo.mutation.LabelsCleared() {
		_spec.ClearField(workflowstepexecution.FieldLabels, field.TypeJSON)
	}
	if value, ok := wseuo.mutation.Annotations(); ok {
		_spec.SetField(workflowstepexecution.FieldAnnotations, field.TypeJSON, value)
	}
	if wseuo.mutation.AnnotationsCleared() {
		_spec.ClearField(workflowstepexecution.FieldAnnotations, field.TypeJSON)
	}
	if value, ok := wseuo.mutation.UpdateTime(); ok {
		_spec.SetField(workflowstepexecution.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := wseuo.mutation.Status(); ok {
		_spec.SetField(workflowstepexecution.FieldStatus, field.TypeJSON, value)
	}
	if wseuo.mutation.StatusCleared() {
		_spec.ClearField(workflowstepexecution.FieldStatus, field.TypeJSON)
	}
	if value, ok := wseuo.mutation.Spec(); ok {
		_spec.SetField(workflowstepexecution.FieldSpec, field.TypeJSON, value)
	}
	if wseuo.mutation.SpecCleared() {
		_spec.ClearField(workflowstepexecution.FieldSpec, field.TypeJSON)
	}
	if value, ok := wseuo.mutation.Times(); ok {
		_spec.SetField(workflowstepexecution.FieldTimes, field.TypeInt, value)
	}
	if value, ok := wseuo.mutation.AddedTimes(); ok {
		_spec.AddField(workflowstepexecution.FieldTimes, field.TypeInt, value)
	}
	if value, ok := wseuo.mutation.Duration(); ok {
		_spec.SetField(workflowstepexecution.FieldDuration, field.TypeInt, value)
	}
	if value, ok := wseuo.mutation.AddedDuration(); ok {
		_spec.AddField(workflowstepexecution.FieldDuration, field.TypeInt, value)
	}
	if value, ok := wseuo.mutation.Record(); ok {
		_spec.SetField(workflowstepexecution.FieldRecord, field.TypeString, value)
	}
	if value, ok := wseuo.mutation.Input(); ok {
		_spec.SetField(workflowstepexecution.FieldInput, field.TypeString, value)
	}
	_spec.Node.Schema = wseuo.schemaConfig.WorkflowStepExecution
	ctx = internal.NewSchemaConfigContext(ctx, wseuo.schemaConfig)
	_spec.AddModifiers(wseuo.modifiers...)
	_node = &WorkflowStepExecution{config: wseuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wseuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{workflowstepexecution.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	wseuo.mutation.done = true
	return _node, nil
}
