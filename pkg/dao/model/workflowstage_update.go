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
	"github.com/seal-io/walrus/pkg/dao/model/workflowstage"
	"github.com/seal-io/walrus/pkg/dao/model/workflowstageexecution"
	"github.com/seal-io/walrus/pkg/dao/model/workflowstep"
	"github.com/seal-io/walrus/pkg/dao/types/object"
	"github.com/seal-io/walrus/pkg/dao/types/status"
)

// WorkflowStageUpdate is the builder for updating WorkflowStage entities.
type WorkflowStageUpdate struct {
	config
	hooks     []Hook
	mutation  *WorkflowStageMutation
	modifiers []func(*sql.UpdateBuilder)
	object    *WorkflowStage
}

// Where appends a list predicates to the WorkflowStageUpdate builder.
func (wsu *WorkflowStageUpdate) Where(ps ...predicate.WorkflowStage) *WorkflowStageUpdate {
	wsu.mutation.Where(ps...)
	return wsu
}

// SetDescription sets the "description" field.
func (wsu *WorkflowStageUpdate) SetDescription(s string) *WorkflowStageUpdate {
	wsu.mutation.SetDescription(s)
	return wsu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (wsu *WorkflowStageUpdate) SetNillableDescription(s *string) *WorkflowStageUpdate {
	if s != nil {
		wsu.SetDescription(*s)
	}
	return wsu
}

// ClearDescription clears the value of the "description" field.
func (wsu *WorkflowStageUpdate) ClearDescription() *WorkflowStageUpdate {
	wsu.mutation.ClearDescription()
	return wsu
}

// SetLabels sets the "labels" field.
func (wsu *WorkflowStageUpdate) SetLabels(m map[string]string) *WorkflowStageUpdate {
	wsu.mutation.SetLabels(m)
	return wsu
}

// ClearLabels clears the value of the "labels" field.
func (wsu *WorkflowStageUpdate) ClearLabels() *WorkflowStageUpdate {
	wsu.mutation.ClearLabels()
	return wsu
}

// SetAnnotations sets the "annotations" field.
func (wsu *WorkflowStageUpdate) SetAnnotations(m map[string]string) *WorkflowStageUpdate {
	wsu.mutation.SetAnnotations(m)
	return wsu
}

// ClearAnnotations clears the value of the "annotations" field.
func (wsu *WorkflowStageUpdate) ClearAnnotations() *WorkflowStageUpdate {
	wsu.mutation.ClearAnnotations()
	return wsu
}

// SetUpdateTime sets the "update_time" field.
func (wsu *WorkflowStageUpdate) SetUpdateTime(t time.Time) *WorkflowStageUpdate {
	wsu.mutation.SetUpdateTime(t)
	return wsu
}

// SetStatus sets the "status" field.
func (wsu *WorkflowStageUpdate) SetStatus(s status.Status) *WorkflowStageUpdate {
	wsu.mutation.SetStatus(s)
	return wsu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (wsu *WorkflowStageUpdate) SetNillableStatus(s *status.Status) *WorkflowStageUpdate {
	if s != nil {
		wsu.SetStatus(*s)
	}
	return wsu
}

// ClearStatus clears the value of the "status" field.
func (wsu *WorkflowStageUpdate) ClearStatus() *WorkflowStageUpdate {
	wsu.mutation.ClearStatus()
	return wsu
}

// SetStepIds sets the "step_ids" field.
func (wsu *WorkflowStageUpdate) SetStepIds(o []object.ID) *WorkflowStageUpdate {
	wsu.mutation.SetStepIds(o)
	return wsu
}

// AppendStepIds appends o to the "step_ids" field.
func (wsu *WorkflowStageUpdate) AppendStepIds(o []object.ID) *WorkflowStageUpdate {
	wsu.mutation.AppendStepIds(o)
	return wsu
}

// SetDependencies sets the "dependencies" field.
func (wsu *WorkflowStageUpdate) SetDependencies(o []object.ID) *WorkflowStageUpdate {
	wsu.mutation.SetDependencies(o)
	return wsu
}

// AppendDependencies appends o to the "dependencies" field.
func (wsu *WorkflowStageUpdate) AppendDependencies(o []object.ID) *WorkflowStageUpdate {
	wsu.mutation.AppendDependencies(o)
	return wsu
}

// AddStepIDs adds the "steps" edge to the WorkflowStep entity by IDs.
func (wsu *WorkflowStageUpdate) AddStepIDs(ids ...object.ID) *WorkflowStageUpdate {
	wsu.mutation.AddStepIDs(ids...)
	return wsu
}

// AddSteps adds the "steps" edges to the WorkflowStep entity.
func (wsu *WorkflowStageUpdate) AddSteps(w ...*WorkflowStep) *WorkflowStageUpdate {
	ids := make([]object.ID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wsu.AddStepIDs(ids...)
}

// AddExecutionIDs adds the "executions" edge to the WorkflowStageExecution entity by IDs.
func (wsu *WorkflowStageUpdate) AddExecutionIDs(ids ...object.ID) *WorkflowStageUpdate {
	wsu.mutation.AddExecutionIDs(ids...)
	return wsu
}

// AddExecutions adds the "executions" edges to the WorkflowStageExecution entity.
func (wsu *WorkflowStageUpdate) AddExecutions(w ...*WorkflowStageExecution) *WorkflowStageUpdate {
	ids := make([]object.ID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wsu.AddExecutionIDs(ids...)
}

// Mutation returns the WorkflowStageMutation object of the builder.
func (wsu *WorkflowStageUpdate) Mutation() *WorkflowStageMutation {
	return wsu.mutation
}

// ClearSteps clears all "steps" edges to the WorkflowStep entity.
func (wsu *WorkflowStageUpdate) ClearSteps() *WorkflowStageUpdate {
	wsu.mutation.ClearSteps()
	return wsu
}

// RemoveStepIDs removes the "steps" edge to WorkflowStep entities by IDs.
func (wsu *WorkflowStageUpdate) RemoveStepIDs(ids ...object.ID) *WorkflowStageUpdate {
	wsu.mutation.RemoveStepIDs(ids...)
	return wsu
}

// RemoveSteps removes "steps" edges to WorkflowStep entities.
func (wsu *WorkflowStageUpdate) RemoveSteps(w ...*WorkflowStep) *WorkflowStageUpdate {
	ids := make([]object.ID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wsu.RemoveStepIDs(ids...)
}

// ClearExecutions clears all "executions" edges to the WorkflowStageExecution entity.
func (wsu *WorkflowStageUpdate) ClearExecutions() *WorkflowStageUpdate {
	wsu.mutation.ClearExecutions()
	return wsu
}

// RemoveExecutionIDs removes the "executions" edge to WorkflowStageExecution entities by IDs.
func (wsu *WorkflowStageUpdate) RemoveExecutionIDs(ids ...object.ID) *WorkflowStageUpdate {
	wsu.mutation.RemoveExecutionIDs(ids...)
	return wsu
}

// RemoveExecutions removes "executions" edges to WorkflowStageExecution entities.
func (wsu *WorkflowStageUpdate) RemoveExecutions(w ...*WorkflowStageExecution) *WorkflowStageUpdate {
	ids := make([]object.ID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wsu.RemoveExecutionIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wsu *WorkflowStageUpdate) Save(ctx context.Context) (int, error) {
	if err := wsu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, wsu.sqlSave, wsu.mutation, wsu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wsu *WorkflowStageUpdate) SaveX(ctx context.Context) int {
	affected, err := wsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wsu *WorkflowStageUpdate) Exec(ctx context.Context) error {
	_, err := wsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wsu *WorkflowStageUpdate) ExecX(ctx context.Context) {
	if err := wsu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wsu *WorkflowStageUpdate) defaults() error {
	if _, ok := wsu.mutation.UpdateTime(); !ok {
		if workflowstage.UpdateDefaultUpdateTime == nil {
			return fmt.Errorf("model: uninitialized workflowstage.UpdateDefaultUpdateTime (forgotten import model/runtime?)")
		}
		v := workflowstage.UpdateDefaultUpdateTime()
		wsu.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (wsu *WorkflowStageUpdate) check() error {
	if _, ok := wsu.mutation.ProjectID(); wsu.mutation.ProjectCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "WorkflowStage.project"`)
	}
	if _, ok := wsu.mutation.WorkflowID(); wsu.mutation.WorkflowCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "WorkflowStage.workflow"`)
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
func (wsu *WorkflowStageUpdate) Set(obj *WorkflowStage) *WorkflowStageUpdate {
	// Without Default.
	if obj.Description != "" {
		wsu.SetDescription(obj.Description)
	} else {
		wsu.ClearDescription()
	}
	if !reflect.ValueOf(obj.Labels).IsZero() {
		wsu.SetLabels(obj.Labels)
	} else {
		wsu.ClearLabels()
	}
	if !reflect.ValueOf(obj.Annotations).IsZero() {
		wsu.SetAnnotations(obj.Annotations)
	}
	if !reflect.ValueOf(obj.Status).IsZero() {
		wsu.SetStatus(obj.Status)
	}
	wsu.SetStepIds(obj.StepIds)
	wsu.SetDependencies(obj.Dependencies)

	// With Default.
	if obj.UpdateTime != nil {
		wsu.SetUpdateTime(*obj.UpdateTime)
	}

	// Record the given object.
	wsu.object = obj

	return wsu
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (wsu *WorkflowStageUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *WorkflowStageUpdate {
	wsu.modifiers = append(wsu.modifiers, modifiers...)
	return wsu
}

func (wsu *WorkflowStageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := wsu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(workflowstage.Table, workflowstage.Columns, sqlgraph.NewFieldSpec(workflowstage.FieldID, field.TypeString))
	if ps := wsu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wsu.mutation.Description(); ok {
		_spec.SetField(workflowstage.FieldDescription, field.TypeString, value)
	}
	if wsu.mutation.DescriptionCleared() {
		_spec.ClearField(workflowstage.FieldDescription, field.TypeString)
	}
	if value, ok := wsu.mutation.Labels(); ok {
		_spec.SetField(workflowstage.FieldLabels, field.TypeJSON, value)
	}
	if wsu.mutation.LabelsCleared() {
		_spec.ClearField(workflowstage.FieldLabels, field.TypeJSON)
	}
	if value, ok := wsu.mutation.Annotations(); ok {
		_spec.SetField(workflowstage.FieldAnnotations, field.TypeJSON, value)
	}
	if wsu.mutation.AnnotationsCleared() {
		_spec.ClearField(workflowstage.FieldAnnotations, field.TypeJSON)
	}
	if value, ok := wsu.mutation.UpdateTime(); ok {
		_spec.SetField(workflowstage.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := wsu.mutation.Status(); ok {
		_spec.SetField(workflowstage.FieldStatus, field.TypeJSON, value)
	}
	if wsu.mutation.StatusCleared() {
		_spec.ClearField(workflowstage.FieldStatus, field.TypeJSON)
	}
	if value, ok := wsu.mutation.StepIds(); ok {
		_spec.SetField(workflowstage.FieldStepIds, field.TypeJSON, value)
	}
	if value, ok := wsu.mutation.AppendedStepIds(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, workflowstage.FieldStepIds, value)
		})
	}
	if value, ok := wsu.mutation.Dependencies(); ok {
		_spec.SetField(workflowstage.FieldDependencies, field.TypeJSON, value)
	}
	if value, ok := wsu.mutation.AppendedDependencies(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, workflowstage.FieldDependencies, value)
		})
	}
	if wsu.mutation.StepsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflowstage.StepsTable,
			Columns: []string{workflowstage.StepsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflowstep.FieldID, field.TypeString),
			},
		}
		edge.Schema = wsu.schemaConfig.WorkflowStep
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wsu.mutation.RemovedStepsIDs(); len(nodes) > 0 && !wsu.mutation.StepsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflowstage.StepsTable,
			Columns: []string{workflowstage.StepsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflowstep.FieldID, field.TypeString),
			},
		}
		edge.Schema = wsu.schemaConfig.WorkflowStep
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wsu.mutation.StepsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflowstage.StepsTable,
			Columns: []string{workflowstage.StepsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflowstep.FieldID, field.TypeString),
			},
		}
		edge.Schema = wsu.schemaConfig.WorkflowStep
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wsu.mutation.ExecutionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflowstage.ExecutionsTable,
			Columns: []string{workflowstage.ExecutionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflowstageexecution.FieldID, field.TypeString),
			},
		}
		edge.Schema = wsu.schemaConfig.WorkflowStageExecution
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wsu.mutation.RemovedExecutionsIDs(); len(nodes) > 0 && !wsu.mutation.ExecutionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflowstage.ExecutionsTable,
			Columns: []string{workflowstage.ExecutionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflowstageexecution.FieldID, field.TypeString),
			},
		}
		edge.Schema = wsu.schemaConfig.WorkflowStageExecution
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wsu.mutation.ExecutionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflowstage.ExecutionsTable,
			Columns: []string{workflowstage.ExecutionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflowstageexecution.FieldID, field.TypeString),
			},
		}
		edge.Schema = wsu.schemaConfig.WorkflowStageExecution
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = wsu.schemaConfig.WorkflowStage
	ctx = internal.NewSchemaConfigContext(ctx, wsu.schemaConfig)
	_spec.AddModifiers(wsu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, wsu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{workflowstage.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	wsu.mutation.done = true
	return n, nil
}

// WorkflowStageUpdateOne is the builder for updating a single WorkflowStage entity.
type WorkflowStageUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *WorkflowStageMutation
	modifiers []func(*sql.UpdateBuilder)
	object    *WorkflowStage
}

// SetDescription sets the "description" field.
func (wsuo *WorkflowStageUpdateOne) SetDescription(s string) *WorkflowStageUpdateOne {
	wsuo.mutation.SetDescription(s)
	return wsuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (wsuo *WorkflowStageUpdateOne) SetNillableDescription(s *string) *WorkflowStageUpdateOne {
	if s != nil {
		wsuo.SetDescription(*s)
	}
	return wsuo
}

// ClearDescription clears the value of the "description" field.
func (wsuo *WorkflowStageUpdateOne) ClearDescription() *WorkflowStageUpdateOne {
	wsuo.mutation.ClearDescription()
	return wsuo
}

// SetLabels sets the "labels" field.
func (wsuo *WorkflowStageUpdateOne) SetLabels(m map[string]string) *WorkflowStageUpdateOne {
	wsuo.mutation.SetLabels(m)
	return wsuo
}

// ClearLabels clears the value of the "labels" field.
func (wsuo *WorkflowStageUpdateOne) ClearLabels() *WorkflowStageUpdateOne {
	wsuo.mutation.ClearLabels()
	return wsuo
}

// SetAnnotations sets the "annotations" field.
func (wsuo *WorkflowStageUpdateOne) SetAnnotations(m map[string]string) *WorkflowStageUpdateOne {
	wsuo.mutation.SetAnnotations(m)
	return wsuo
}

// ClearAnnotations clears the value of the "annotations" field.
func (wsuo *WorkflowStageUpdateOne) ClearAnnotations() *WorkflowStageUpdateOne {
	wsuo.mutation.ClearAnnotations()
	return wsuo
}

// SetUpdateTime sets the "update_time" field.
func (wsuo *WorkflowStageUpdateOne) SetUpdateTime(t time.Time) *WorkflowStageUpdateOne {
	wsuo.mutation.SetUpdateTime(t)
	return wsuo
}

// SetStatus sets the "status" field.
func (wsuo *WorkflowStageUpdateOne) SetStatus(s status.Status) *WorkflowStageUpdateOne {
	wsuo.mutation.SetStatus(s)
	return wsuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (wsuo *WorkflowStageUpdateOne) SetNillableStatus(s *status.Status) *WorkflowStageUpdateOne {
	if s != nil {
		wsuo.SetStatus(*s)
	}
	return wsuo
}

// ClearStatus clears the value of the "status" field.
func (wsuo *WorkflowStageUpdateOne) ClearStatus() *WorkflowStageUpdateOne {
	wsuo.mutation.ClearStatus()
	return wsuo
}

// SetStepIds sets the "step_ids" field.
func (wsuo *WorkflowStageUpdateOne) SetStepIds(o []object.ID) *WorkflowStageUpdateOne {
	wsuo.mutation.SetStepIds(o)
	return wsuo
}

// AppendStepIds appends o to the "step_ids" field.
func (wsuo *WorkflowStageUpdateOne) AppendStepIds(o []object.ID) *WorkflowStageUpdateOne {
	wsuo.mutation.AppendStepIds(o)
	return wsuo
}

// SetDependencies sets the "dependencies" field.
func (wsuo *WorkflowStageUpdateOne) SetDependencies(o []object.ID) *WorkflowStageUpdateOne {
	wsuo.mutation.SetDependencies(o)
	return wsuo
}

// AppendDependencies appends o to the "dependencies" field.
func (wsuo *WorkflowStageUpdateOne) AppendDependencies(o []object.ID) *WorkflowStageUpdateOne {
	wsuo.mutation.AppendDependencies(o)
	return wsuo
}

// AddStepIDs adds the "steps" edge to the WorkflowStep entity by IDs.
func (wsuo *WorkflowStageUpdateOne) AddStepIDs(ids ...object.ID) *WorkflowStageUpdateOne {
	wsuo.mutation.AddStepIDs(ids...)
	return wsuo
}

// AddSteps adds the "steps" edges to the WorkflowStep entity.
func (wsuo *WorkflowStageUpdateOne) AddSteps(w ...*WorkflowStep) *WorkflowStageUpdateOne {
	ids := make([]object.ID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wsuo.AddStepIDs(ids...)
}

// AddExecutionIDs adds the "executions" edge to the WorkflowStageExecution entity by IDs.
func (wsuo *WorkflowStageUpdateOne) AddExecutionIDs(ids ...object.ID) *WorkflowStageUpdateOne {
	wsuo.mutation.AddExecutionIDs(ids...)
	return wsuo
}

// AddExecutions adds the "executions" edges to the WorkflowStageExecution entity.
func (wsuo *WorkflowStageUpdateOne) AddExecutions(w ...*WorkflowStageExecution) *WorkflowStageUpdateOne {
	ids := make([]object.ID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wsuo.AddExecutionIDs(ids...)
}

// Mutation returns the WorkflowStageMutation object of the builder.
func (wsuo *WorkflowStageUpdateOne) Mutation() *WorkflowStageMutation {
	return wsuo.mutation
}

// ClearSteps clears all "steps" edges to the WorkflowStep entity.
func (wsuo *WorkflowStageUpdateOne) ClearSteps() *WorkflowStageUpdateOne {
	wsuo.mutation.ClearSteps()
	return wsuo
}

// RemoveStepIDs removes the "steps" edge to WorkflowStep entities by IDs.
func (wsuo *WorkflowStageUpdateOne) RemoveStepIDs(ids ...object.ID) *WorkflowStageUpdateOne {
	wsuo.mutation.RemoveStepIDs(ids...)
	return wsuo
}

// RemoveSteps removes "steps" edges to WorkflowStep entities.
func (wsuo *WorkflowStageUpdateOne) RemoveSteps(w ...*WorkflowStep) *WorkflowStageUpdateOne {
	ids := make([]object.ID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wsuo.RemoveStepIDs(ids...)
}

// ClearExecutions clears all "executions" edges to the WorkflowStageExecution entity.
func (wsuo *WorkflowStageUpdateOne) ClearExecutions() *WorkflowStageUpdateOne {
	wsuo.mutation.ClearExecutions()
	return wsuo
}

// RemoveExecutionIDs removes the "executions" edge to WorkflowStageExecution entities by IDs.
func (wsuo *WorkflowStageUpdateOne) RemoveExecutionIDs(ids ...object.ID) *WorkflowStageUpdateOne {
	wsuo.mutation.RemoveExecutionIDs(ids...)
	return wsuo
}

// RemoveExecutions removes "executions" edges to WorkflowStageExecution entities.
func (wsuo *WorkflowStageUpdateOne) RemoveExecutions(w ...*WorkflowStageExecution) *WorkflowStageUpdateOne {
	ids := make([]object.ID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wsuo.RemoveExecutionIDs(ids...)
}

// Where appends a list predicates to the WorkflowStageUpdate builder.
func (wsuo *WorkflowStageUpdateOne) Where(ps ...predicate.WorkflowStage) *WorkflowStageUpdateOne {
	wsuo.mutation.Where(ps...)
	return wsuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wsuo *WorkflowStageUpdateOne) Select(field string, fields ...string) *WorkflowStageUpdateOne {
	wsuo.fields = append([]string{field}, fields...)
	return wsuo
}

// Save executes the query and returns the updated WorkflowStage entity.
func (wsuo *WorkflowStageUpdateOne) Save(ctx context.Context) (*WorkflowStage, error) {
	if err := wsuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, wsuo.sqlSave, wsuo.mutation, wsuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wsuo *WorkflowStageUpdateOne) SaveX(ctx context.Context) *WorkflowStage {
	node, err := wsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wsuo *WorkflowStageUpdateOne) Exec(ctx context.Context) error {
	_, err := wsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wsuo *WorkflowStageUpdateOne) ExecX(ctx context.Context) {
	if err := wsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wsuo *WorkflowStageUpdateOne) defaults() error {
	if _, ok := wsuo.mutation.UpdateTime(); !ok {
		if workflowstage.UpdateDefaultUpdateTime == nil {
			return fmt.Errorf("model: uninitialized workflowstage.UpdateDefaultUpdateTime (forgotten import model/runtime?)")
		}
		v := workflowstage.UpdateDefaultUpdateTime()
		wsuo.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (wsuo *WorkflowStageUpdateOne) check() error {
	if _, ok := wsuo.mutation.ProjectID(); wsuo.mutation.ProjectCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "WorkflowStage.project"`)
	}
	if _, ok := wsuo.mutation.WorkflowID(); wsuo.mutation.WorkflowCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "WorkflowStage.workflow"`)
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
func (wsuo *WorkflowStageUpdateOne) Set(obj *WorkflowStage) *WorkflowStageUpdateOne {
	h := func(n ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			mt := m.(*WorkflowStageMutation)
			db, err := mt.Client().WorkflowStage.Get(ctx, *mt.id)
			if err != nil {
				return nil, fmt.Errorf("failed getting WorkflowStage with id: %v", *mt.id)
			}

			// Without Default.
			if obj.Description != "" {
				if db.Description != obj.Description {
					wsuo.SetDescription(obj.Description)
				}
			} else {
				wsuo.ClearDescription()
			}
			if !reflect.ValueOf(obj.Labels).IsZero() {
				if !reflect.DeepEqual(db.Labels, obj.Labels) {
					wsuo.SetLabels(obj.Labels)
				}
			} else {
				wsuo.ClearLabels()
			}
			if !reflect.ValueOf(obj.Annotations).IsZero() {
				if !reflect.DeepEqual(db.Annotations, obj.Annotations) {
					wsuo.SetAnnotations(obj.Annotations)
				}
			}
			if !reflect.ValueOf(obj.Status).IsZero() {
				if !db.Status.Equal(obj.Status) {
					wsuo.SetStatus(obj.Status)
				}
			}
			if !reflect.DeepEqual(db.StepIds, obj.StepIds) {
				wsuo.SetStepIds(obj.StepIds)
			}
			if !reflect.DeepEqual(db.Dependencies, obj.Dependencies) {
				wsuo.SetDependencies(obj.Dependencies)
			}

			// With Default.
			if (obj.UpdateTime != nil) && (!reflect.DeepEqual(db.UpdateTime, obj.UpdateTime)) {
				wsuo.SetUpdateTime(*obj.UpdateTime)
			}

			// Record the given object.
			wsuo.object = obj

			return n.Mutate(ctx, m)
		})
	}

	wsuo.hooks = append(wsuo.hooks, h)

	return wsuo
}

// getClientSet returns the ClientSet for the given builder.
func (wsuo *WorkflowStageUpdateOne) getClientSet() (mc ClientSet) {
	if _, ok := wsuo.config.driver.(*txDriver); ok {
		tx := &Tx{config: wsuo.config}
		tx.init()
		mc = tx
	} else {
		cli := &Client{config: wsuo.config}
		cli.init()
		mc = cli
	}
	return mc
}

// SaveE calls the given function after updated the WorkflowStage entity,
// which is always good for cascading update operations.
func (wsuo *WorkflowStageUpdateOne) SaveE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *WorkflowStage) error) (*WorkflowStage, error) {
	obj, err := wsuo.Save(ctx)
	if err != nil &&
		(wsuo.object == nil || !errors.Is(err, stdsql.ErrNoRows)) {
		return nil, err
	}

	if len(cbs) == 0 {
		return obj, err
	}

	mc := wsuo.getClientSet()

	if obj == nil {
		obj = wsuo.object
	} else if x := wsuo.object; x != nil {
		if _, set := wsuo.mutation.Field(workflowstage.FieldDescription); set {
			obj.Description = x.Description
		}
		if _, set := wsuo.mutation.Field(workflowstage.FieldLabels); set {
			obj.Labels = x.Labels
		}
		if _, set := wsuo.mutation.Field(workflowstage.FieldAnnotations); set {
			obj.Annotations = x.Annotations
		}
		if _, set := wsuo.mutation.Field(workflowstage.FieldStatus); set {
			obj.Status = x.Status
		}
		if _, set := wsuo.mutation.Field(workflowstage.FieldStepIds); set {
			obj.StepIds = x.StepIds
		}
		if _, set := wsuo.mutation.Field(workflowstage.FieldDependencies); set {
			obj.Dependencies = x.Dependencies
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
func (wsuo *WorkflowStageUpdateOne) SaveEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *WorkflowStage) error) *WorkflowStage {
	obj, err := wsuo.SaveE(ctx, cbs...)
	if err != nil {
		panic(err)
	}
	return obj
}

// ExecE calls the given function after executed the query,
// which is always good for cascading update operations.
func (wsuo *WorkflowStageUpdateOne) ExecE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *WorkflowStage) error) error {
	_, err := wsuo.SaveE(ctx, cbs...)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wsuo *WorkflowStageUpdateOne) ExecEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *WorkflowStage) error) {
	if err := wsuo.ExecE(ctx, cbs...); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (wsuo *WorkflowStageUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *WorkflowStageUpdateOne {
	wsuo.modifiers = append(wsuo.modifiers, modifiers...)
	return wsuo
}

func (wsuo *WorkflowStageUpdateOne) sqlSave(ctx context.Context) (_node *WorkflowStage, err error) {
	if err := wsuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(workflowstage.Table, workflowstage.Columns, sqlgraph.NewFieldSpec(workflowstage.FieldID, field.TypeString))
	id, ok := wsuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`model: missing "WorkflowStage.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := wsuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, workflowstage.FieldID)
		for _, f := range fields {
			if !workflowstage.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
			}
			if f != workflowstage.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wsuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wsuo.mutation.Description(); ok {
		_spec.SetField(workflowstage.FieldDescription, field.TypeString, value)
	}
	if wsuo.mutation.DescriptionCleared() {
		_spec.ClearField(workflowstage.FieldDescription, field.TypeString)
	}
	if value, ok := wsuo.mutation.Labels(); ok {
		_spec.SetField(workflowstage.FieldLabels, field.TypeJSON, value)
	}
	if wsuo.mutation.LabelsCleared() {
		_spec.ClearField(workflowstage.FieldLabels, field.TypeJSON)
	}
	if value, ok := wsuo.mutation.Annotations(); ok {
		_spec.SetField(workflowstage.FieldAnnotations, field.TypeJSON, value)
	}
	if wsuo.mutation.AnnotationsCleared() {
		_spec.ClearField(workflowstage.FieldAnnotations, field.TypeJSON)
	}
	if value, ok := wsuo.mutation.UpdateTime(); ok {
		_spec.SetField(workflowstage.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := wsuo.mutation.Status(); ok {
		_spec.SetField(workflowstage.FieldStatus, field.TypeJSON, value)
	}
	if wsuo.mutation.StatusCleared() {
		_spec.ClearField(workflowstage.FieldStatus, field.TypeJSON)
	}
	if value, ok := wsuo.mutation.StepIds(); ok {
		_spec.SetField(workflowstage.FieldStepIds, field.TypeJSON, value)
	}
	if value, ok := wsuo.mutation.AppendedStepIds(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, workflowstage.FieldStepIds, value)
		})
	}
	if value, ok := wsuo.mutation.Dependencies(); ok {
		_spec.SetField(workflowstage.FieldDependencies, field.TypeJSON, value)
	}
	if value, ok := wsuo.mutation.AppendedDependencies(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, workflowstage.FieldDependencies, value)
		})
	}
	if wsuo.mutation.StepsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflowstage.StepsTable,
			Columns: []string{workflowstage.StepsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflowstep.FieldID, field.TypeString),
			},
		}
		edge.Schema = wsuo.schemaConfig.WorkflowStep
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wsuo.mutation.RemovedStepsIDs(); len(nodes) > 0 && !wsuo.mutation.StepsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflowstage.StepsTable,
			Columns: []string{workflowstage.StepsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflowstep.FieldID, field.TypeString),
			},
		}
		edge.Schema = wsuo.schemaConfig.WorkflowStep
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wsuo.mutation.StepsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflowstage.StepsTable,
			Columns: []string{workflowstage.StepsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflowstep.FieldID, field.TypeString),
			},
		}
		edge.Schema = wsuo.schemaConfig.WorkflowStep
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wsuo.mutation.ExecutionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflowstage.ExecutionsTable,
			Columns: []string{workflowstage.ExecutionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflowstageexecution.FieldID, field.TypeString),
			},
		}
		edge.Schema = wsuo.schemaConfig.WorkflowStageExecution
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wsuo.mutation.RemovedExecutionsIDs(); len(nodes) > 0 && !wsuo.mutation.ExecutionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflowstage.ExecutionsTable,
			Columns: []string{workflowstage.ExecutionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflowstageexecution.FieldID, field.TypeString),
			},
		}
		edge.Schema = wsuo.schemaConfig.WorkflowStageExecution
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wsuo.mutation.ExecutionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflowstage.ExecutionsTable,
			Columns: []string{workflowstage.ExecutionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflowstageexecution.FieldID, field.TypeString),
			},
		}
		edge.Schema = wsuo.schemaConfig.WorkflowStageExecution
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = wsuo.schemaConfig.WorkflowStage
	ctx = internal.NewSchemaConfigContext(ctx, wsuo.schemaConfig)
	_spec.AddModifiers(wsuo.modifiers...)
	_node = &WorkflowStage{config: wsuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wsuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{workflowstage.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	wsuo.mutation.done = true
	return _node, nil
}
