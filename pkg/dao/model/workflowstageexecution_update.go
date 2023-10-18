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
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"

	"github.com/seal-io/walrus/pkg/dao/model/internal"
	"github.com/seal-io/walrus/pkg/dao/model/predicate"
	"github.com/seal-io/walrus/pkg/dao/model/workflowstageexecution"
	"github.com/seal-io/walrus/pkg/dao/model/workflowstepexecution"
	"github.com/seal-io/walrus/pkg/dao/types/object"
	"github.com/seal-io/walrus/pkg/dao/types/status"
)

// WorkflowStageExecutionUpdate is the builder for updating WorkflowStageExecution entities.
type WorkflowStageExecutionUpdate struct {
	config
	hooks     []Hook
	mutation  *WorkflowStageExecutionMutation
	modifiers []func(*sql.UpdateBuilder)
	object    *WorkflowStageExecution
}

// Where appends a list predicates to the WorkflowStageExecutionUpdate builder.
func (wseu *WorkflowStageExecutionUpdate) Where(ps ...predicate.WorkflowStageExecution) *WorkflowStageExecutionUpdate {
	wseu.mutation.Where(ps...)
	return wseu
}

// SetDescription sets the "description" field.
func (wseu *WorkflowStageExecutionUpdate) SetDescription(s string) *WorkflowStageExecutionUpdate {
	wseu.mutation.SetDescription(s)
	return wseu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (wseu *WorkflowStageExecutionUpdate) SetNillableDescription(s *string) *WorkflowStageExecutionUpdate {
	if s != nil {
		wseu.SetDescription(*s)
	}
	return wseu
}

// ClearDescription clears the value of the "description" field.
func (wseu *WorkflowStageExecutionUpdate) ClearDescription() *WorkflowStageExecutionUpdate {
	wseu.mutation.ClearDescription()
	return wseu
}

// SetLabels sets the "labels" field.
func (wseu *WorkflowStageExecutionUpdate) SetLabels(m map[string]string) *WorkflowStageExecutionUpdate {
	wseu.mutation.SetLabels(m)
	return wseu
}

// ClearLabels clears the value of the "labels" field.
func (wseu *WorkflowStageExecutionUpdate) ClearLabels() *WorkflowStageExecutionUpdate {
	wseu.mutation.ClearLabels()
	return wseu
}

// SetAnnotations sets the "annotations" field.
func (wseu *WorkflowStageExecutionUpdate) SetAnnotations(m map[string]string) *WorkflowStageExecutionUpdate {
	wseu.mutation.SetAnnotations(m)
	return wseu
}

// ClearAnnotations clears the value of the "annotations" field.
func (wseu *WorkflowStageExecutionUpdate) ClearAnnotations() *WorkflowStageExecutionUpdate {
	wseu.mutation.ClearAnnotations()
	return wseu
}

// SetUpdateTime sets the "update_time" field.
func (wseu *WorkflowStageExecutionUpdate) SetUpdateTime(t time.Time) *WorkflowStageExecutionUpdate {
	wseu.mutation.SetUpdateTime(t)
	return wseu
}

// SetStatus sets the "status" field.
func (wseu *WorkflowStageExecutionUpdate) SetStatus(s status.Status) *WorkflowStageExecutionUpdate {
	wseu.mutation.SetStatus(s)
	return wseu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (wseu *WorkflowStageExecutionUpdate) SetNillableStatus(s *status.Status) *WorkflowStageExecutionUpdate {
	if s != nil {
		wseu.SetStatus(*s)
	}
	return wseu
}

// ClearStatus clears the value of the "status" field.
func (wseu *WorkflowStageExecutionUpdate) ClearStatus() *WorkflowStageExecutionUpdate {
	wseu.mutation.ClearStatus()
	return wseu
}

// SetDuration sets the "duration" field.
func (wseu *WorkflowStageExecutionUpdate) SetDuration(i int) *WorkflowStageExecutionUpdate {
	wseu.mutation.ResetDuration()
	wseu.mutation.SetDuration(i)
	return wseu
}

// SetNillableDuration sets the "duration" field if the given value is not nil.
func (wseu *WorkflowStageExecutionUpdate) SetNillableDuration(i *int) *WorkflowStageExecutionUpdate {
	if i != nil {
		wseu.SetDuration(*i)
	}
	return wseu
}

// AddDuration adds i to the "duration" field.
func (wseu *WorkflowStageExecutionUpdate) AddDuration(i int) *WorkflowStageExecutionUpdate {
	wseu.mutation.AddDuration(i)
	return wseu
}

// SetStepExecutionIds sets the "step_execution_ids" field.
func (wseu *WorkflowStageExecutionUpdate) SetStepExecutionIds(o []object.ID) *WorkflowStageExecutionUpdate {
	wseu.mutation.SetStepExecutionIds(o)
	return wseu
}

// AppendStepExecutionIds appends o to the "step_execution_ids" field.
func (wseu *WorkflowStageExecutionUpdate) AppendStepExecutionIds(o []object.ID) *WorkflowStageExecutionUpdate {
	wseu.mutation.AppendStepExecutionIds(o)
	return wseu
}

// SetRecord sets the "record" field.
func (wseu *WorkflowStageExecutionUpdate) SetRecord(s string) *WorkflowStageExecutionUpdate {
	wseu.mutation.SetRecord(s)
	return wseu
}

// SetNillableRecord sets the "record" field if the given value is not nil.
func (wseu *WorkflowStageExecutionUpdate) SetNillableRecord(s *string) *WorkflowStageExecutionUpdate {
	if s != nil {
		wseu.SetRecord(*s)
	}
	return wseu
}

// SetInput sets the "input" field.
func (wseu *WorkflowStageExecutionUpdate) SetInput(s string) *WorkflowStageExecutionUpdate {
	wseu.mutation.SetInput(s)
	return wseu
}

// SetNillableInput sets the "input" field if the given value is not nil.
func (wseu *WorkflowStageExecutionUpdate) SetNillableInput(s *string) *WorkflowStageExecutionUpdate {
	if s != nil {
		wseu.SetInput(*s)
	}
	return wseu
}

// AddStepIDs adds the "steps" edge to the WorkflowStepExecution entity by IDs.
func (wseu *WorkflowStageExecutionUpdate) AddStepIDs(ids ...object.ID) *WorkflowStageExecutionUpdate {
	wseu.mutation.AddStepIDs(ids...)
	return wseu
}

// AddSteps adds the "steps" edges to the WorkflowStepExecution entity.
func (wseu *WorkflowStageExecutionUpdate) AddSteps(w ...*WorkflowStepExecution) *WorkflowStageExecutionUpdate {
	ids := make([]object.ID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wseu.AddStepIDs(ids...)
}

// Mutation returns the WorkflowStageExecutionMutation object of the builder.
func (wseu *WorkflowStageExecutionUpdate) Mutation() *WorkflowStageExecutionMutation {
	return wseu.mutation
}

// ClearSteps clears all "steps" edges to the WorkflowStepExecution entity.
func (wseu *WorkflowStageExecutionUpdate) ClearSteps() *WorkflowStageExecutionUpdate {
	wseu.mutation.ClearSteps()
	return wseu
}

// RemoveStepIDs removes the "steps" edge to WorkflowStepExecution entities by IDs.
func (wseu *WorkflowStageExecutionUpdate) RemoveStepIDs(ids ...object.ID) *WorkflowStageExecutionUpdate {
	wseu.mutation.RemoveStepIDs(ids...)
	return wseu
}

// RemoveSteps removes "steps" edges to WorkflowStepExecution entities.
func (wseu *WorkflowStageExecutionUpdate) RemoveSteps(w ...*WorkflowStepExecution) *WorkflowStageExecutionUpdate {
	ids := make([]object.ID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wseu.RemoveStepIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wseu *WorkflowStageExecutionUpdate) Save(ctx context.Context) (int, error) {
	if err := wseu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, wseu.sqlSave, wseu.mutation, wseu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wseu *WorkflowStageExecutionUpdate) SaveX(ctx context.Context) int {
	affected, err := wseu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wseu *WorkflowStageExecutionUpdate) Exec(ctx context.Context) error {
	_, err := wseu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wseu *WorkflowStageExecutionUpdate) ExecX(ctx context.Context) {
	if err := wseu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wseu *WorkflowStageExecutionUpdate) defaults() error {
	if _, ok := wseu.mutation.UpdateTime(); !ok {
		if workflowstageexecution.UpdateDefaultUpdateTime == nil {
			return fmt.Errorf("model: uninitialized workflowstageexecution.UpdateDefaultUpdateTime (forgotten import model/runtime?)")
		}
		v := workflowstageexecution.UpdateDefaultUpdateTime()
		wseu.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (wseu *WorkflowStageExecutionUpdate) check() error {
	if v, ok := wseu.mutation.Duration(); ok {
		if err := workflowstageexecution.DurationValidator(v); err != nil {
			return &ValidationError{Name: "duration", err: fmt.Errorf(`model: validator failed for field "WorkflowStageExecution.duration": %w`, err)}
		}
	}
	if _, ok := wseu.mutation.ProjectID(); wseu.mutation.ProjectCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "WorkflowStageExecution.project"`)
	}
	if _, ok := wseu.mutation.StageID(); wseu.mutation.StageCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "WorkflowStageExecution.stage"`)
	}
	if _, ok := wseu.mutation.WorkflowExecutionID(); wseu.mutation.WorkflowExecutionCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "WorkflowStageExecution.workflow_execution"`)
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
func (wseu *WorkflowStageExecutionUpdate) Set(obj *WorkflowStageExecution) *WorkflowStageExecutionUpdate {
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
	wseu.SetDuration(obj.Duration)
	wseu.SetStepExecutionIds(obj.StepExecutionIds)
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
func (wseu *WorkflowStageExecutionUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *WorkflowStageExecutionUpdate {
	wseu.modifiers = append(wseu.modifiers, modifiers...)
	return wseu
}

func (wseu *WorkflowStageExecutionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := wseu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(workflowstageexecution.Table, workflowstageexecution.Columns, sqlgraph.NewFieldSpec(workflowstageexecution.FieldID, field.TypeString))
	if ps := wseu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wseu.mutation.Description(); ok {
		_spec.SetField(workflowstageexecution.FieldDescription, field.TypeString, value)
	}
	if wseu.mutation.DescriptionCleared() {
		_spec.ClearField(workflowstageexecution.FieldDescription, field.TypeString)
	}
	if value, ok := wseu.mutation.Labels(); ok {
		_spec.SetField(workflowstageexecution.FieldLabels, field.TypeJSON, value)
	}
	if wseu.mutation.LabelsCleared() {
		_spec.ClearField(workflowstageexecution.FieldLabels, field.TypeJSON)
	}
	if value, ok := wseu.mutation.Annotations(); ok {
		_spec.SetField(workflowstageexecution.FieldAnnotations, field.TypeJSON, value)
	}
	if wseu.mutation.AnnotationsCleared() {
		_spec.ClearField(workflowstageexecution.FieldAnnotations, field.TypeJSON)
	}
	if value, ok := wseu.mutation.UpdateTime(); ok {
		_spec.SetField(workflowstageexecution.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := wseu.mutation.Status(); ok {
		_spec.SetField(workflowstageexecution.FieldStatus, field.TypeJSON, value)
	}
	if wseu.mutation.StatusCleared() {
		_spec.ClearField(workflowstageexecution.FieldStatus, field.TypeJSON)
	}
	if value, ok := wseu.mutation.Duration(); ok {
		_spec.SetField(workflowstageexecution.FieldDuration, field.TypeInt, value)
	}
	if value, ok := wseu.mutation.AddedDuration(); ok {
		_spec.AddField(workflowstageexecution.FieldDuration, field.TypeInt, value)
	}
	if value, ok := wseu.mutation.StepExecutionIds(); ok {
		_spec.SetField(workflowstageexecution.FieldStepExecutionIds, field.TypeJSON, value)
	}
	if value, ok := wseu.mutation.AppendedStepExecutionIds(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, workflowstageexecution.FieldStepExecutionIds, value)
		})
	}
	if value, ok := wseu.mutation.Record(); ok {
		_spec.SetField(workflowstageexecution.FieldRecord, field.TypeString, value)
	}
	if value, ok := wseu.mutation.Input(); ok {
		_spec.SetField(workflowstageexecution.FieldInput, field.TypeString, value)
	}
	if wseu.mutation.StepsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflowstageexecution.StepsTable,
			Columns: []string{workflowstageexecution.StepsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflowstepexecution.FieldID, field.TypeString),
			},
		}
		edge.Schema = wseu.schemaConfig.WorkflowStepExecution
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wseu.mutation.RemovedStepsIDs(); len(nodes) > 0 && !wseu.mutation.StepsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflowstageexecution.StepsTable,
			Columns: []string{workflowstageexecution.StepsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflowstepexecution.FieldID, field.TypeString),
			},
		}
		edge.Schema = wseu.schemaConfig.WorkflowStepExecution
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wseu.mutation.StepsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflowstageexecution.StepsTable,
			Columns: []string{workflowstageexecution.StepsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflowstepexecution.FieldID, field.TypeString),
			},
		}
		edge.Schema = wseu.schemaConfig.WorkflowStepExecution
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = wseu.schemaConfig.WorkflowStageExecution
	ctx = internal.NewSchemaConfigContext(ctx, wseu.schemaConfig)
	_spec.AddModifiers(wseu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, wseu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{workflowstageexecution.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	wseu.mutation.done = true
	return n, nil
}

// WorkflowStageExecutionUpdateOne is the builder for updating a single WorkflowStageExecution entity.
type WorkflowStageExecutionUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *WorkflowStageExecutionMutation
	modifiers []func(*sql.UpdateBuilder)
	object    *WorkflowStageExecution
}

// SetDescription sets the "description" field.
func (wseuo *WorkflowStageExecutionUpdateOne) SetDescription(s string) *WorkflowStageExecutionUpdateOne {
	wseuo.mutation.SetDescription(s)
	return wseuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (wseuo *WorkflowStageExecutionUpdateOne) SetNillableDescription(s *string) *WorkflowStageExecutionUpdateOne {
	if s != nil {
		wseuo.SetDescription(*s)
	}
	return wseuo
}

// ClearDescription clears the value of the "description" field.
func (wseuo *WorkflowStageExecutionUpdateOne) ClearDescription() *WorkflowStageExecutionUpdateOne {
	wseuo.mutation.ClearDescription()
	return wseuo
}

// SetLabels sets the "labels" field.
func (wseuo *WorkflowStageExecutionUpdateOne) SetLabels(m map[string]string) *WorkflowStageExecutionUpdateOne {
	wseuo.mutation.SetLabels(m)
	return wseuo
}

// ClearLabels clears the value of the "labels" field.
func (wseuo *WorkflowStageExecutionUpdateOne) ClearLabels() *WorkflowStageExecutionUpdateOne {
	wseuo.mutation.ClearLabels()
	return wseuo
}

// SetAnnotations sets the "annotations" field.
func (wseuo *WorkflowStageExecutionUpdateOne) SetAnnotations(m map[string]string) *WorkflowStageExecutionUpdateOne {
	wseuo.mutation.SetAnnotations(m)
	return wseuo
}

// ClearAnnotations clears the value of the "annotations" field.
func (wseuo *WorkflowStageExecutionUpdateOne) ClearAnnotations() *WorkflowStageExecutionUpdateOne {
	wseuo.mutation.ClearAnnotations()
	return wseuo
}

// SetUpdateTime sets the "update_time" field.
func (wseuo *WorkflowStageExecutionUpdateOne) SetUpdateTime(t time.Time) *WorkflowStageExecutionUpdateOne {
	wseuo.mutation.SetUpdateTime(t)
	return wseuo
}

// SetStatus sets the "status" field.
func (wseuo *WorkflowStageExecutionUpdateOne) SetStatus(s status.Status) *WorkflowStageExecutionUpdateOne {
	wseuo.mutation.SetStatus(s)
	return wseuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (wseuo *WorkflowStageExecutionUpdateOne) SetNillableStatus(s *status.Status) *WorkflowStageExecutionUpdateOne {
	if s != nil {
		wseuo.SetStatus(*s)
	}
	return wseuo
}

// ClearStatus clears the value of the "status" field.
func (wseuo *WorkflowStageExecutionUpdateOne) ClearStatus() *WorkflowStageExecutionUpdateOne {
	wseuo.mutation.ClearStatus()
	return wseuo
}

// SetDuration sets the "duration" field.
func (wseuo *WorkflowStageExecutionUpdateOne) SetDuration(i int) *WorkflowStageExecutionUpdateOne {
	wseuo.mutation.ResetDuration()
	wseuo.mutation.SetDuration(i)
	return wseuo
}

// SetNillableDuration sets the "duration" field if the given value is not nil.
func (wseuo *WorkflowStageExecutionUpdateOne) SetNillableDuration(i *int) *WorkflowStageExecutionUpdateOne {
	if i != nil {
		wseuo.SetDuration(*i)
	}
	return wseuo
}

// AddDuration adds i to the "duration" field.
func (wseuo *WorkflowStageExecutionUpdateOne) AddDuration(i int) *WorkflowStageExecutionUpdateOne {
	wseuo.mutation.AddDuration(i)
	return wseuo
}

// SetStepExecutionIds sets the "step_execution_ids" field.
func (wseuo *WorkflowStageExecutionUpdateOne) SetStepExecutionIds(o []object.ID) *WorkflowStageExecutionUpdateOne {
	wseuo.mutation.SetStepExecutionIds(o)
	return wseuo
}

// AppendStepExecutionIds appends o to the "step_execution_ids" field.
func (wseuo *WorkflowStageExecutionUpdateOne) AppendStepExecutionIds(o []object.ID) *WorkflowStageExecutionUpdateOne {
	wseuo.mutation.AppendStepExecutionIds(o)
	return wseuo
}

// SetRecord sets the "record" field.
func (wseuo *WorkflowStageExecutionUpdateOne) SetRecord(s string) *WorkflowStageExecutionUpdateOne {
	wseuo.mutation.SetRecord(s)
	return wseuo
}

// SetNillableRecord sets the "record" field if the given value is not nil.
func (wseuo *WorkflowStageExecutionUpdateOne) SetNillableRecord(s *string) *WorkflowStageExecutionUpdateOne {
	if s != nil {
		wseuo.SetRecord(*s)
	}
	return wseuo
}

// SetInput sets the "input" field.
func (wseuo *WorkflowStageExecutionUpdateOne) SetInput(s string) *WorkflowStageExecutionUpdateOne {
	wseuo.mutation.SetInput(s)
	return wseuo
}

// SetNillableInput sets the "input" field if the given value is not nil.
func (wseuo *WorkflowStageExecutionUpdateOne) SetNillableInput(s *string) *WorkflowStageExecutionUpdateOne {
	if s != nil {
		wseuo.SetInput(*s)
	}
	return wseuo
}

// AddStepIDs adds the "steps" edge to the WorkflowStepExecution entity by IDs.
func (wseuo *WorkflowStageExecutionUpdateOne) AddStepIDs(ids ...object.ID) *WorkflowStageExecutionUpdateOne {
	wseuo.mutation.AddStepIDs(ids...)
	return wseuo
}

// AddSteps adds the "steps" edges to the WorkflowStepExecution entity.
func (wseuo *WorkflowStageExecutionUpdateOne) AddSteps(w ...*WorkflowStepExecution) *WorkflowStageExecutionUpdateOne {
	ids := make([]object.ID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wseuo.AddStepIDs(ids...)
}

// Mutation returns the WorkflowStageExecutionMutation object of the builder.
func (wseuo *WorkflowStageExecutionUpdateOne) Mutation() *WorkflowStageExecutionMutation {
	return wseuo.mutation
}

// ClearSteps clears all "steps" edges to the WorkflowStepExecution entity.
func (wseuo *WorkflowStageExecutionUpdateOne) ClearSteps() *WorkflowStageExecutionUpdateOne {
	wseuo.mutation.ClearSteps()
	return wseuo
}

// RemoveStepIDs removes the "steps" edge to WorkflowStepExecution entities by IDs.
func (wseuo *WorkflowStageExecutionUpdateOne) RemoveStepIDs(ids ...object.ID) *WorkflowStageExecutionUpdateOne {
	wseuo.mutation.RemoveStepIDs(ids...)
	return wseuo
}

// RemoveSteps removes "steps" edges to WorkflowStepExecution entities.
func (wseuo *WorkflowStageExecutionUpdateOne) RemoveSteps(w ...*WorkflowStepExecution) *WorkflowStageExecutionUpdateOne {
	ids := make([]object.ID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wseuo.RemoveStepIDs(ids...)
}

// Where appends a list predicates to the WorkflowStageExecutionUpdate builder.
func (wseuo *WorkflowStageExecutionUpdateOne) Where(ps ...predicate.WorkflowStageExecution) *WorkflowStageExecutionUpdateOne {
	wseuo.mutation.Where(ps...)
	return wseuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wseuo *WorkflowStageExecutionUpdateOne) Select(field string, fields ...string) *WorkflowStageExecutionUpdateOne {
	wseuo.fields = append([]string{field}, fields...)
	return wseuo
}

// Save executes the query and returns the updated WorkflowStageExecution entity.
func (wseuo *WorkflowStageExecutionUpdateOne) Save(ctx context.Context) (*WorkflowStageExecution, error) {
	if err := wseuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, wseuo.sqlSave, wseuo.mutation, wseuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wseuo *WorkflowStageExecutionUpdateOne) SaveX(ctx context.Context) *WorkflowStageExecution {
	node, err := wseuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wseuo *WorkflowStageExecutionUpdateOne) Exec(ctx context.Context) error {
	_, err := wseuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wseuo *WorkflowStageExecutionUpdateOne) ExecX(ctx context.Context) {
	if err := wseuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wseuo *WorkflowStageExecutionUpdateOne) defaults() error {
	if _, ok := wseuo.mutation.UpdateTime(); !ok {
		if workflowstageexecution.UpdateDefaultUpdateTime == nil {
			return fmt.Errorf("model: uninitialized workflowstageexecution.UpdateDefaultUpdateTime (forgotten import model/runtime?)")
		}
		v := workflowstageexecution.UpdateDefaultUpdateTime()
		wseuo.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (wseuo *WorkflowStageExecutionUpdateOne) check() error {
	if v, ok := wseuo.mutation.Duration(); ok {
		if err := workflowstageexecution.DurationValidator(v); err != nil {
			return &ValidationError{Name: "duration", err: fmt.Errorf(`model: validator failed for field "WorkflowStageExecution.duration": %w`, err)}
		}
	}
	if _, ok := wseuo.mutation.ProjectID(); wseuo.mutation.ProjectCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "WorkflowStageExecution.project"`)
	}
	if _, ok := wseuo.mutation.StageID(); wseuo.mutation.StageCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "WorkflowStageExecution.stage"`)
	}
	if _, ok := wseuo.mutation.WorkflowExecutionID(); wseuo.mutation.WorkflowExecutionCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "WorkflowStageExecution.workflow_execution"`)
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
func (wseuo *WorkflowStageExecutionUpdateOne) Set(obj *WorkflowStageExecution) *WorkflowStageExecutionUpdateOne {
	h := func(n ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			mt := m.(*WorkflowStageExecutionMutation)
			db, err := mt.Client().WorkflowStageExecution.Get(ctx, *mt.id)
			if err != nil {
				return nil, fmt.Errorf("failed getting WorkflowStageExecution with id: %v", *mt.id)
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
			if db.Duration != obj.Duration {
				wseuo.SetDuration(obj.Duration)
			}
			if !reflect.DeepEqual(db.StepExecutionIds, obj.StepExecutionIds) {
				wseuo.SetStepExecutionIds(obj.StepExecutionIds)
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
func (wseuo *WorkflowStageExecutionUpdateOne) getClientSet() (mc ClientSet) {
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

// SaveE calls the given function after updated the WorkflowStageExecution entity,
// which is always good for cascading update operations.
func (wseuo *WorkflowStageExecutionUpdateOne) SaveE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *WorkflowStageExecution) error) (*WorkflowStageExecution, error) {
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
		if _, set := wseuo.mutation.Field(workflowstageexecution.FieldDescription); set {
			obj.Description = x.Description
		}
		if _, set := wseuo.mutation.Field(workflowstageexecution.FieldLabels); set {
			obj.Labels = x.Labels
		}
		if _, set := wseuo.mutation.Field(workflowstageexecution.FieldAnnotations); set {
			obj.Annotations = x.Annotations
		}
		if _, set := wseuo.mutation.Field(workflowstageexecution.FieldStatus); set {
			obj.Status = x.Status
		}
		if _, set := wseuo.mutation.Field(workflowstageexecution.FieldDuration); set {
			obj.Duration = x.Duration
		}
		if _, set := wseuo.mutation.Field(workflowstageexecution.FieldStepExecutionIds); set {
			obj.StepExecutionIds = x.StepExecutionIds
		}
		if _, set := wseuo.mutation.Field(workflowstageexecution.FieldRecord); set {
			obj.Record = x.Record
		}
		if _, set := wseuo.mutation.Field(workflowstageexecution.FieldInput); set {
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
func (wseuo *WorkflowStageExecutionUpdateOne) SaveEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *WorkflowStageExecution) error) *WorkflowStageExecution {
	obj, err := wseuo.SaveE(ctx, cbs...)
	if err != nil {
		panic(err)
	}
	return obj
}

// ExecE calls the given function after executed the query,
// which is always good for cascading update operations.
func (wseuo *WorkflowStageExecutionUpdateOne) ExecE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *WorkflowStageExecution) error) error {
	_, err := wseuo.SaveE(ctx, cbs...)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wseuo *WorkflowStageExecutionUpdateOne) ExecEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *WorkflowStageExecution) error) {
	if err := wseuo.ExecE(ctx, cbs...); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (wseuo *WorkflowStageExecutionUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *WorkflowStageExecutionUpdateOne {
	wseuo.modifiers = append(wseuo.modifiers, modifiers...)
	return wseuo
}

func (wseuo *WorkflowStageExecutionUpdateOne) sqlSave(ctx context.Context) (_node *WorkflowStageExecution, err error) {
	if err := wseuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(workflowstageexecution.Table, workflowstageexecution.Columns, sqlgraph.NewFieldSpec(workflowstageexecution.FieldID, field.TypeString))
	id, ok := wseuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`model: missing "WorkflowStageExecution.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := wseuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, workflowstageexecution.FieldID)
		for _, f := range fields {
			if !workflowstageexecution.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
			}
			if f != workflowstageexecution.FieldID {
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
		_spec.SetField(workflowstageexecution.FieldDescription, field.TypeString, value)
	}
	if wseuo.mutation.DescriptionCleared() {
		_spec.ClearField(workflowstageexecution.FieldDescription, field.TypeString)
	}
	if value, ok := wseuo.mutation.Labels(); ok {
		_spec.SetField(workflowstageexecution.FieldLabels, field.TypeJSON, value)
	}
	if wseuo.mutation.LabelsCleared() {
		_spec.ClearField(workflowstageexecution.FieldLabels, field.TypeJSON)
	}
	if value, ok := wseuo.mutation.Annotations(); ok {
		_spec.SetField(workflowstageexecution.FieldAnnotations, field.TypeJSON, value)
	}
	if wseuo.mutation.AnnotationsCleared() {
		_spec.ClearField(workflowstageexecution.FieldAnnotations, field.TypeJSON)
	}
	if value, ok := wseuo.mutation.UpdateTime(); ok {
		_spec.SetField(workflowstageexecution.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := wseuo.mutation.Status(); ok {
		_spec.SetField(workflowstageexecution.FieldStatus, field.TypeJSON, value)
	}
	if wseuo.mutation.StatusCleared() {
		_spec.ClearField(workflowstageexecution.FieldStatus, field.TypeJSON)
	}
	if value, ok := wseuo.mutation.Duration(); ok {
		_spec.SetField(workflowstageexecution.FieldDuration, field.TypeInt, value)
	}
	if value, ok := wseuo.mutation.AddedDuration(); ok {
		_spec.AddField(workflowstageexecution.FieldDuration, field.TypeInt, value)
	}
	if value, ok := wseuo.mutation.StepExecutionIds(); ok {
		_spec.SetField(workflowstageexecution.FieldStepExecutionIds, field.TypeJSON, value)
	}
	if value, ok := wseuo.mutation.AppendedStepExecutionIds(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, workflowstageexecution.FieldStepExecutionIds, value)
		})
	}
	if value, ok := wseuo.mutation.Record(); ok {
		_spec.SetField(workflowstageexecution.FieldRecord, field.TypeString, value)
	}
	if value, ok := wseuo.mutation.Input(); ok {
		_spec.SetField(workflowstageexecution.FieldInput, field.TypeString, value)
	}
	if wseuo.mutation.StepsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflowstageexecution.StepsTable,
			Columns: []string{workflowstageexecution.StepsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflowstepexecution.FieldID, field.TypeString),
			},
		}
		edge.Schema = wseuo.schemaConfig.WorkflowStepExecution
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wseuo.mutation.RemovedStepsIDs(); len(nodes) > 0 && !wseuo.mutation.StepsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflowstageexecution.StepsTable,
			Columns: []string{workflowstageexecution.StepsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflowstepexecution.FieldID, field.TypeString),
			},
		}
		edge.Schema = wseuo.schemaConfig.WorkflowStepExecution
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wseuo.mutation.StepsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflowstageexecution.StepsTable,
			Columns: []string{workflowstageexecution.StepsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflowstepexecution.FieldID, field.TypeString),
			},
		}
		edge.Schema = wseuo.schemaConfig.WorkflowStepExecution
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = wseuo.schemaConfig.WorkflowStageExecution
	ctx = internal.NewSchemaConfigContext(ctx, wseuo.schemaConfig)
	_spec.AddModifiers(wseuo.modifiers...)
	_node = &WorkflowStageExecution{config: wseuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wseuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{workflowstageexecution.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	wseuo.mutation.done = true
	return _node, nil
}
