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
	"github.com/seal-io/walrus/pkg/dao/model/workflow"
	"github.com/seal-io/walrus/pkg/dao/model/workflowexecution"
	"github.com/seal-io/walrus/pkg/dao/model/workflowstage"
	"github.com/seal-io/walrus/pkg/dao/types/object"
	"github.com/seal-io/walrus/pkg/dao/types/status"
)

// WorkflowUpdate is the builder for updating Workflow entities.
type WorkflowUpdate struct {
	config
	hooks     []Hook
	mutation  *WorkflowMutation
	modifiers []func(*sql.UpdateBuilder)
	object    *Workflow
}

// Where appends a list predicates to the WorkflowUpdate builder.
func (wu *WorkflowUpdate) Where(ps ...predicate.Workflow) *WorkflowUpdate {
	wu.mutation.Where(ps...)
	return wu
}

// SetDescription sets the "description" field.
func (wu *WorkflowUpdate) SetDescription(s string) *WorkflowUpdate {
	wu.mutation.SetDescription(s)
	return wu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (wu *WorkflowUpdate) SetNillableDescription(s *string) *WorkflowUpdate {
	if s != nil {
		wu.SetDescription(*s)
	}
	return wu
}

// ClearDescription clears the value of the "description" field.
func (wu *WorkflowUpdate) ClearDescription() *WorkflowUpdate {
	wu.mutation.ClearDescription()
	return wu
}

// SetLabels sets the "labels" field.
func (wu *WorkflowUpdate) SetLabels(m map[string]string) *WorkflowUpdate {
	wu.mutation.SetLabels(m)
	return wu
}

// ClearLabels clears the value of the "labels" field.
func (wu *WorkflowUpdate) ClearLabels() *WorkflowUpdate {
	wu.mutation.ClearLabels()
	return wu
}

// SetAnnotations sets the "annotations" field.
func (wu *WorkflowUpdate) SetAnnotations(m map[string]string) *WorkflowUpdate {
	wu.mutation.SetAnnotations(m)
	return wu
}

// ClearAnnotations clears the value of the "annotations" field.
func (wu *WorkflowUpdate) ClearAnnotations() *WorkflowUpdate {
	wu.mutation.ClearAnnotations()
	return wu
}

// SetUpdateTime sets the "update_time" field.
func (wu *WorkflowUpdate) SetUpdateTime(t time.Time) *WorkflowUpdate {
	wu.mutation.SetUpdateTime(t)
	return wu
}

// SetStatus sets the "status" field.
func (wu *WorkflowUpdate) SetStatus(s status.Status) *WorkflowUpdate {
	wu.mutation.SetStatus(s)
	return wu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (wu *WorkflowUpdate) SetNillableStatus(s *status.Status) *WorkflowUpdate {
	if s != nil {
		wu.SetStatus(*s)
	}
	return wu
}

// ClearStatus clears the value of the "status" field.
func (wu *WorkflowUpdate) ClearStatus() *WorkflowUpdate {
	wu.mutation.ClearStatus()
	return wu
}

// SetDisplayName sets the "display_name" field.
func (wu *WorkflowUpdate) SetDisplayName(s string) *WorkflowUpdate {
	wu.mutation.SetDisplayName(s)
	return wu
}

// SetStageIds sets the "stage_ids" field.
func (wu *WorkflowUpdate) SetStageIds(o []object.ID) *WorkflowUpdate {
	wu.mutation.SetStageIds(o)
	return wu
}

// AppendStageIds appends o to the "stage_ids" field.
func (wu *WorkflowUpdate) AppendStageIds(o []object.ID) *WorkflowUpdate {
	wu.mutation.AppendStageIds(o)
	return wu
}

// SetParallelism sets the "parallelism" field.
func (wu *WorkflowUpdate) SetParallelism(i int) *WorkflowUpdate {
	wu.mutation.ResetParallelism()
	wu.mutation.SetParallelism(i)
	return wu
}

// SetNillableParallelism sets the "parallelism" field if the given value is not nil.
func (wu *WorkflowUpdate) SetNillableParallelism(i *int) *WorkflowUpdate {
	if i != nil {
		wu.SetParallelism(*i)
	}
	return wu
}

// AddParallelism adds i to the "parallelism" field.
func (wu *WorkflowUpdate) AddParallelism(i int) *WorkflowUpdate {
	wu.mutation.AddParallelism(i)
	return wu
}

// AddStageIDs adds the "stages" edge to the WorkflowStage entity by IDs.
func (wu *WorkflowUpdate) AddStageIDs(ids ...object.ID) *WorkflowUpdate {
	wu.mutation.AddStageIDs(ids...)
	return wu
}

// AddStages adds the "stages" edges to the WorkflowStage entity.
func (wu *WorkflowUpdate) AddStages(w ...*WorkflowStage) *WorkflowUpdate {
	ids := make([]object.ID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wu.AddStageIDs(ids...)
}

// AddExecutionIDs adds the "executions" edge to the WorkflowExecution entity by IDs.
func (wu *WorkflowUpdate) AddExecutionIDs(ids ...object.ID) *WorkflowUpdate {
	wu.mutation.AddExecutionIDs(ids...)
	return wu
}

// AddExecutions adds the "executions" edges to the WorkflowExecution entity.
func (wu *WorkflowUpdate) AddExecutions(w ...*WorkflowExecution) *WorkflowUpdate {
	ids := make([]object.ID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wu.AddExecutionIDs(ids...)
}

// Mutation returns the WorkflowMutation object of the builder.
func (wu *WorkflowUpdate) Mutation() *WorkflowMutation {
	return wu.mutation
}

// ClearStages clears all "stages" edges to the WorkflowStage entity.
func (wu *WorkflowUpdate) ClearStages() *WorkflowUpdate {
	wu.mutation.ClearStages()
	return wu
}

// RemoveStageIDs removes the "stages" edge to WorkflowStage entities by IDs.
func (wu *WorkflowUpdate) RemoveStageIDs(ids ...object.ID) *WorkflowUpdate {
	wu.mutation.RemoveStageIDs(ids...)
	return wu
}

// RemoveStages removes "stages" edges to WorkflowStage entities.
func (wu *WorkflowUpdate) RemoveStages(w ...*WorkflowStage) *WorkflowUpdate {
	ids := make([]object.ID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wu.RemoveStageIDs(ids...)
}

// ClearExecutions clears all "executions" edges to the WorkflowExecution entity.
func (wu *WorkflowUpdate) ClearExecutions() *WorkflowUpdate {
	wu.mutation.ClearExecutions()
	return wu
}

// RemoveExecutionIDs removes the "executions" edge to WorkflowExecution entities by IDs.
func (wu *WorkflowUpdate) RemoveExecutionIDs(ids ...object.ID) *WorkflowUpdate {
	wu.mutation.RemoveExecutionIDs(ids...)
	return wu
}

// RemoveExecutions removes "executions" edges to WorkflowExecution entities.
func (wu *WorkflowUpdate) RemoveExecutions(w ...*WorkflowExecution) *WorkflowUpdate {
	ids := make([]object.ID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wu.RemoveExecutionIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wu *WorkflowUpdate) Save(ctx context.Context) (int, error) {
	if err := wu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, wu.sqlSave, wu.mutation, wu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wu *WorkflowUpdate) SaveX(ctx context.Context) int {
	affected, err := wu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wu *WorkflowUpdate) Exec(ctx context.Context) error {
	_, err := wu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wu *WorkflowUpdate) ExecX(ctx context.Context) {
	if err := wu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wu *WorkflowUpdate) defaults() error {
	if _, ok := wu.mutation.UpdateTime(); !ok {
		if workflow.UpdateDefaultUpdateTime == nil {
			return fmt.Errorf("model: uninitialized workflow.UpdateDefaultUpdateTime (forgotten import model/runtime?)")
		}
		v := workflow.UpdateDefaultUpdateTime()
		wu.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (wu *WorkflowUpdate) check() error {
	if v, ok := wu.mutation.DisplayName(); ok {
		if err := workflow.DisplayNameValidator(v); err != nil {
			return &ValidationError{Name: "display_name", err: fmt.Errorf(`model: validator failed for field "Workflow.display_name": %w`, err)}
		}
	}
	if v, ok := wu.mutation.Parallelism(); ok {
		if err := workflow.ParallelismValidator(v); err != nil {
			return &ValidationError{Name: "parallelism", err: fmt.Errorf(`model: validator failed for field "Workflow.parallelism": %w`, err)}
		}
	}
	if _, ok := wu.mutation.ProjectID(); wu.mutation.ProjectCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "Workflow.project"`)
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
func (wu *WorkflowUpdate) Set(obj *Workflow) *WorkflowUpdate {
	// Without Default.
	if obj.Description != "" {
		wu.SetDescription(obj.Description)
	} else {
		wu.ClearDescription()
	}
	if !reflect.ValueOf(obj.Labels).IsZero() {
		wu.SetLabels(obj.Labels)
	} else {
		wu.ClearLabels()
	}
	if !reflect.ValueOf(obj.Annotations).IsZero() {
		wu.SetAnnotations(obj.Annotations)
	}
	if !reflect.ValueOf(obj.Status).IsZero() {
		wu.SetStatus(obj.Status)
	}
	wu.SetDisplayName(obj.DisplayName)
	wu.SetStageIds(obj.StageIds)
	wu.SetParallelism(obj.Parallelism)

	// With Default.
	if obj.UpdateTime != nil {
		wu.SetUpdateTime(*obj.UpdateTime)
	}

	// Record the given object.
	wu.object = obj

	return wu
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (wu *WorkflowUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *WorkflowUpdate {
	wu.modifiers = append(wu.modifiers, modifiers...)
	return wu
}

func (wu *WorkflowUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := wu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(workflow.Table, workflow.Columns, sqlgraph.NewFieldSpec(workflow.FieldID, field.TypeString))
	if ps := wu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wu.mutation.Description(); ok {
		_spec.SetField(workflow.FieldDescription, field.TypeString, value)
	}
	if wu.mutation.DescriptionCleared() {
		_spec.ClearField(workflow.FieldDescription, field.TypeString)
	}
	if value, ok := wu.mutation.Labels(); ok {
		_spec.SetField(workflow.FieldLabels, field.TypeJSON, value)
	}
	if wu.mutation.LabelsCleared() {
		_spec.ClearField(workflow.FieldLabels, field.TypeJSON)
	}
	if value, ok := wu.mutation.Annotations(); ok {
		_spec.SetField(workflow.FieldAnnotations, field.TypeJSON, value)
	}
	if wu.mutation.AnnotationsCleared() {
		_spec.ClearField(workflow.FieldAnnotations, field.TypeJSON)
	}
	if value, ok := wu.mutation.UpdateTime(); ok {
		_spec.SetField(workflow.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := wu.mutation.Status(); ok {
		_spec.SetField(workflow.FieldStatus, field.TypeJSON, value)
	}
	if wu.mutation.StatusCleared() {
		_spec.ClearField(workflow.FieldStatus, field.TypeJSON)
	}
	if wu.mutation.EnvironmentIDCleared() {
		_spec.ClearField(workflow.FieldEnvironmentID, field.TypeString)
	}
	if value, ok := wu.mutation.DisplayName(); ok {
		_spec.SetField(workflow.FieldDisplayName, field.TypeString, value)
	}
	if value, ok := wu.mutation.StageIds(); ok {
		_spec.SetField(workflow.FieldStageIds, field.TypeJSON, value)
	}
	if value, ok := wu.mutation.AppendedStageIds(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, workflow.FieldStageIds, value)
		})
	}
	if value, ok := wu.mutation.Parallelism(); ok {
		_spec.SetField(workflow.FieldParallelism, field.TypeInt, value)
	}
	if value, ok := wu.mutation.AddedParallelism(); ok {
		_spec.AddField(workflow.FieldParallelism, field.TypeInt, value)
	}
	if wu.mutation.StagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.StagesTable,
			Columns: []string{workflow.StagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflowstage.FieldID, field.TypeString),
			},
		}
		edge.Schema = wu.schemaConfig.WorkflowStage
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.RemovedStagesIDs(); len(nodes) > 0 && !wu.mutation.StagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.StagesTable,
			Columns: []string{workflow.StagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflowstage.FieldID, field.TypeString),
			},
		}
		edge.Schema = wu.schemaConfig.WorkflowStage
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.StagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.StagesTable,
			Columns: []string{workflow.StagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflowstage.FieldID, field.TypeString),
			},
		}
		edge.Schema = wu.schemaConfig.WorkflowStage
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wu.mutation.ExecutionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.ExecutionsTable,
			Columns: []string{workflow.ExecutionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflowexecution.FieldID, field.TypeString),
			},
		}
		edge.Schema = wu.schemaConfig.WorkflowExecution
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.RemovedExecutionsIDs(); len(nodes) > 0 && !wu.mutation.ExecutionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.ExecutionsTable,
			Columns: []string{workflow.ExecutionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflowexecution.FieldID, field.TypeString),
			},
		}
		edge.Schema = wu.schemaConfig.WorkflowExecution
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.ExecutionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.ExecutionsTable,
			Columns: []string{workflow.ExecutionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflowexecution.FieldID, field.TypeString),
			},
		}
		edge.Schema = wu.schemaConfig.WorkflowExecution
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = wu.schemaConfig.Workflow
	ctx = internal.NewSchemaConfigContext(ctx, wu.schemaConfig)
	_spec.AddModifiers(wu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, wu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{workflow.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	wu.mutation.done = true
	return n, nil
}

// WorkflowUpdateOne is the builder for updating a single Workflow entity.
type WorkflowUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *WorkflowMutation
	modifiers []func(*sql.UpdateBuilder)
	object    *Workflow
}

// SetDescription sets the "description" field.
func (wuo *WorkflowUpdateOne) SetDescription(s string) *WorkflowUpdateOne {
	wuo.mutation.SetDescription(s)
	return wuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (wuo *WorkflowUpdateOne) SetNillableDescription(s *string) *WorkflowUpdateOne {
	if s != nil {
		wuo.SetDescription(*s)
	}
	return wuo
}

// ClearDescription clears the value of the "description" field.
func (wuo *WorkflowUpdateOne) ClearDescription() *WorkflowUpdateOne {
	wuo.mutation.ClearDescription()
	return wuo
}

// SetLabels sets the "labels" field.
func (wuo *WorkflowUpdateOne) SetLabels(m map[string]string) *WorkflowUpdateOne {
	wuo.mutation.SetLabels(m)
	return wuo
}

// ClearLabels clears the value of the "labels" field.
func (wuo *WorkflowUpdateOne) ClearLabels() *WorkflowUpdateOne {
	wuo.mutation.ClearLabels()
	return wuo
}

// SetAnnotations sets the "annotations" field.
func (wuo *WorkflowUpdateOne) SetAnnotations(m map[string]string) *WorkflowUpdateOne {
	wuo.mutation.SetAnnotations(m)
	return wuo
}

// ClearAnnotations clears the value of the "annotations" field.
func (wuo *WorkflowUpdateOne) ClearAnnotations() *WorkflowUpdateOne {
	wuo.mutation.ClearAnnotations()
	return wuo
}

// SetUpdateTime sets the "update_time" field.
func (wuo *WorkflowUpdateOne) SetUpdateTime(t time.Time) *WorkflowUpdateOne {
	wuo.mutation.SetUpdateTime(t)
	return wuo
}

// SetStatus sets the "status" field.
func (wuo *WorkflowUpdateOne) SetStatus(s status.Status) *WorkflowUpdateOne {
	wuo.mutation.SetStatus(s)
	return wuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (wuo *WorkflowUpdateOne) SetNillableStatus(s *status.Status) *WorkflowUpdateOne {
	if s != nil {
		wuo.SetStatus(*s)
	}
	return wuo
}

// ClearStatus clears the value of the "status" field.
func (wuo *WorkflowUpdateOne) ClearStatus() *WorkflowUpdateOne {
	wuo.mutation.ClearStatus()
	return wuo
}

// SetDisplayName sets the "display_name" field.
func (wuo *WorkflowUpdateOne) SetDisplayName(s string) *WorkflowUpdateOne {
	wuo.mutation.SetDisplayName(s)
	return wuo
}

// SetStageIds sets the "stage_ids" field.
func (wuo *WorkflowUpdateOne) SetStageIds(o []object.ID) *WorkflowUpdateOne {
	wuo.mutation.SetStageIds(o)
	return wuo
}

// AppendStageIds appends o to the "stage_ids" field.
func (wuo *WorkflowUpdateOne) AppendStageIds(o []object.ID) *WorkflowUpdateOne {
	wuo.mutation.AppendStageIds(o)
	return wuo
}

// SetParallelism sets the "parallelism" field.
func (wuo *WorkflowUpdateOne) SetParallelism(i int) *WorkflowUpdateOne {
	wuo.mutation.ResetParallelism()
	wuo.mutation.SetParallelism(i)
	return wuo
}

// SetNillableParallelism sets the "parallelism" field if the given value is not nil.
func (wuo *WorkflowUpdateOne) SetNillableParallelism(i *int) *WorkflowUpdateOne {
	if i != nil {
		wuo.SetParallelism(*i)
	}
	return wuo
}

// AddParallelism adds i to the "parallelism" field.
func (wuo *WorkflowUpdateOne) AddParallelism(i int) *WorkflowUpdateOne {
	wuo.mutation.AddParallelism(i)
	return wuo
}

// AddStageIDs adds the "stages" edge to the WorkflowStage entity by IDs.
func (wuo *WorkflowUpdateOne) AddStageIDs(ids ...object.ID) *WorkflowUpdateOne {
	wuo.mutation.AddStageIDs(ids...)
	return wuo
}

// AddStages adds the "stages" edges to the WorkflowStage entity.
func (wuo *WorkflowUpdateOne) AddStages(w ...*WorkflowStage) *WorkflowUpdateOne {
	ids := make([]object.ID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wuo.AddStageIDs(ids...)
}

// AddExecutionIDs adds the "executions" edge to the WorkflowExecution entity by IDs.
func (wuo *WorkflowUpdateOne) AddExecutionIDs(ids ...object.ID) *WorkflowUpdateOne {
	wuo.mutation.AddExecutionIDs(ids...)
	return wuo
}

// AddExecutions adds the "executions" edges to the WorkflowExecution entity.
func (wuo *WorkflowUpdateOne) AddExecutions(w ...*WorkflowExecution) *WorkflowUpdateOne {
	ids := make([]object.ID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wuo.AddExecutionIDs(ids...)
}

// Mutation returns the WorkflowMutation object of the builder.
func (wuo *WorkflowUpdateOne) Mutation() *WorkflowMutation {
	return wuo.mutation
}

// ClearStages clears all "stages" edges to the WorkflowStage entity.
func (wuo *WorkflowUpdateOne) ClearStages() *WorkflowUpdateOne {
	wuo.mutation.ClearStages()
	return wuo
}

// RemoveStageIDs removes the "stages" edge to WorkflowStage entities by IDs.
func (wuo *WorkflowUpdateOne) RemoveStageIDs(ids ...object.ID) *WorkflowUpdateOne {
	wuo.mutation.RemoveStageIDs(ids...)
	return wuo
}

// RemoveStages removes "stages" edges to WorkflowStage entities.
func (wuo *WorkflowUpdateOne) RemoveStages(w ...*WorkflowStage) *WorkflowUpdateOne {
	ids := make([]object.ID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wuo.RemoveStageIDs(ids...)
}

// ClearExecutions clears all "executions" edges to the WorkflowExecution entity.
func (wuo *WorkflowUpdateOne) ClearExecutions() *WorkflowUpdateOne {
	wuo.mutation.ClearExecutions()
	return wuo
}

// RemoveExecutionIDs removes the "executions" edge to WorkflowExecution entities by IDs.
func (wuo *WorkflowUpdateOne) RemoveExecutionIDs(ids ...object.ID) *WorkflowUpdateOne {
	wuo.mutation.RemoveExecutionIDs(ids...)
	return wuo
}

// RemoveExecutions removes "executions" edges to WorkflowExecution entities.
func (wuo *WorkflowUpdateOne) RemoveExecutions(w ...*WorkflowExecution) *WorkflowUpdateOne {
	ids := make([]object.ID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wuo.RemoveExecutionIDs(ids...)
}

// Where appends a list predicates to the WorkflowUpdate builder.
func (wuo *WorkflowUpdateOne) Where(ps ...predicate.Workflow) *WorkflowUpdateOne {
	wuo.mutation.Where(ps...)
	return wuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wuo *WorkflowUpdateOne) Select(field string, fields ...string) *WorkflowUpdateOne {
	wuo.fields = append([]string{field}, fields...)
	return wuo
}

// Save executes the query and returns the updated Workflow entity.
func (wuo *WorkflowUpdateOne) Save(ctx context.Context) (*Workflow, error) {
	if err := wuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, wuo.sqlSave, wuo.mutation, wuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wuo *WorkflowUpdateOne) SaveX(ctx context.Context) *Workflow {
	node, err := wuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wuo *WorkflowUpdateOne) Exec(ctx context.Context) error {
	_, err := wuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wuo *WorkflowUpdateOne) ExecX(ctx context.Context) {
	if err := wuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wuo *WorkflowUpdateOne) defaults() error {
	if _, ok := wuo.mutation.UpdateTime(); !ok {
		if workflow.UpdateDefaultUpdateTime == nil {
			return fmt.Errorf("model: uninitialized workflow.UpdateDefaultUpdateTime (forgotten import model/runtime?)")
		}
		v := workflow.UpdateDefaultUpdateTime()
		wuo.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (wuo *WorkflowUpdateOne) check() error {
	if v, ok := wuo.mutation.DisplayName(); ok {
		if err := workflow.DisplayNameValidator(v); err != nil {
			return &ValidationError{Name: "display_name", err: fmt.Errorf(`model: validator failed for field "Workflow.display_name": %w`, err)}
		}
	}
	if v, ok := wuo.mutation.Parallelism(); ok {
		if err := workflow.ParallelismValidator(v); err != nil {
			return &ValidationError{Name: "parallelism", err: fmt.Errorf(`model: validator failed for field "Workflow.parallelism": %w`, err)}
		}
	}
	if _, ok := wuo.mutation.ProjectID(); wuo.mutation.ProjectCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "Workflow.project"`)
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
func (wuo *WorkflowUpdateOne) Set(obj *Workflow) *WorkflowUpdateOne {
	h := func(n ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			mt := m.(*WorkflowMutation)
			db, err := mt.Client().Workflow.Get(ctx, *mt.id)
			if err != nil {
				return nil, fmt.Errorf("failed getting Workflow with id: %v", *mt.id)
			}

			// Without Default.
			if obj.Description != "" {
				if db.Description != obj.Description {
					wuo.SetDescription(obj.Description)
				}
			} else {
				wuo.ClearDescription()
			}
			if !reflect.ValueOf(obj.Labels).IsZero() {
				if !reflect.DeepEqual(db.Labels, obj.Labels) {
					wuo.SetLabels(obj.Labels)
				}
			} else {
				wuo.ClearLabels()
			}
			if !reflect.ValueOf(obj.Annotations).IsZero() {
				if !reflect.DeepEqual(db.Annotations, obj.Annotations) {
					wuo.SetAnnotations(obj.Annotations)
				}
			}
			if !reflect.ValueOf(obj.Status).IsZero() {
				if !db.Status.Equal(obj.Status) {
					wuo.SetStatus(obj.Status)
				}
			}
			if db.DisplayName != obj.DisplayName {
				wuo.SetDisplayName(obj.DisplayName)
			}
			if !reflect.DeepEqual(db.StageIds, obj.StageIds) {
				wuo.SetStageIds(obj.StageIds)
			}
			if db.Parallelism != obj.Parallelism {
				wuo.SetParallelism(obj.Parallelism)
			}

			// With Default.
			if (obj.UpdateTime != nil) && (!reflect.DeepEqual(db.UpdateTime, obj.UpdateTime)) {
				wuo.SetUpdateTime(*obj.UpdateTime)
			}

			// Record the given object.
			wuo.object = obj

			return n.Mutate(ctx, m)
		})
	}

	wuo.hooks = append(wuo.hooks, h)

	return wuo
}

// getClientSet returns the ClientSet for the given builder.
func (wuo *WorkflowUpdateOne) getClientSet() (mc ClientSet) {
	if _, ok := wuo.config.driver.(*txDriver); ok {
		tx := &Tx{config: wuo.config}
		tx.init()
		mc = tx
	} else {
		cli := &Client{config: wuo.config}
		cli.init()
		mc = cli
	}
	return mc
}

// SaveE calls the given function after updated the Workflow entity,
// which is always good for cascading update operations.
func (wuo *WorkflowUpdateOne) SaveE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *Workflow) error) (*Workflow, error) {
	obj, err := wuo.Save(ctx)
	if err != nil &&
		(wuo.object == nil || !errors.Is(err, stdsql.ErrNoRows)) {
		return nil, err
	}

	if len(cbs) == 0 {
		return obj, err
	}

	mc := wuo.getClientSet()

	if obj == nil {
		obj = wuo.object
	} else if x := wuo.object; x != nil {
		if _, set := wuo.mutation.Field(workflow.FieldDescription); set {
			obj.Description = x.Description
		}
		if _, set := wuo.mutation.Field(workflow.FieldLabels); set {
			obj.Labels = x.Labels
		}
		if _, set := wuo.mutation.Field(workflow.FieldAnnotations); set {
			obj.Annotations = x.Annotations
		}
		if _, set := wuo.mutation.Field(workflow.FieldStatus); set {
			obj.Status = x.Status
		}
		if _, set := wuo.mutation.Field(workflow.FieldDisplayName); set {
			obj.DisplayName = x.DisplayName
		}
		if _, set := wuo.mutation.Field(workflow.FieldStageIds); set {
			obj.StageIds = x.StageIds
		}
		if _, set := wuo.mutation.Field(workflow.FieldParallelism); set {
			obj.Parallelism = x.Parallelism
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
func (wuo *WorkflowUpdateOne) SaveEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *Workflow) error) *Workflow {
	obj, err := wuo.SaveE(ctx, cbs...)
	if err != nil {
		panic(err)
	}
	return obj
}

// ExecE calls the given function after executed the query,
// which is always good for cascading update operations.
func (wuo *WorkflowUpdateOne) ExecE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *Workflow) error) error {
	_, err := wuo.SaveE(ctx, cbs...)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wuo *WorkflowUpdateOne) ExecEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *Workflow) error) {
	if err := wuo.ExecE(ctx, cbs...); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (wuo *WorkflowUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *WorkflowUpdateOne {
	wuo.modifiers = append(wuo.modifiers, modifiers...)
	return wuo
}

func (wuo *WorkflowUpdateOne) sqlSave(ctx context.Context) (_node *Workflow, err error) {
	if err := wuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(workflow.Table, workflow.Columns, sqlgraph.NewFieldSpec(workflow.FieldID, field.TypeString))
	id, ok := wuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`model: missing "Workflow.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := wuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, workflow.FieldID)
		for _, f := range fields {
			if !workflow.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
			}
			if f != workflow.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wuo.mutation.Description(); ok {
		_spec.SetField(workflow.FieldDescription, field.TypeString, value)
	}
	if wuo.mutation.DescriptionCleared() {
		_spec.ClearField(workflow.FieldDescription, field.TypeString)
	}
	if value, ok := wuo.mutation.Labels(); ok {
		_spec.SetField(workflow.FieldLabels, field.TypeJSON, value)
	}
	if wuo.mutation.LabelsCleared() {
		_spec.ClearField(workflow.FieldLabels, field.TypeJSON)
	}
	if value, ok := wuo.mutation.Annotations(); ok {
		_spec.SetField(workflow.FieldAnnotations, field.TypeJSON, value)
	}
	if wuo.mutation.AnnotationsCleared() {
		_spec.ClearField(workflow.FieldAnnotations, field.TypeJSON)
	}
	if value, ok := wuo.mutation.UpdateTime(); ok {
		_spec.SetField(workflow.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := wuo.mutation.Status(); ok {
		_spec.SetField(workflow.FieldStatus, field.TypeJSON, value)
	}
	if wuo.mutation.StatusCleared() {
		_spec.ClearField(workflow.FieldStatus, field.TypeJSON)
	}
	if wuo.mutation.EnvironmentIDCleared() {
		_spec.ClearField(workflow.FieldEnvironmentID, field.TypeString)
	}
	if value, ok := wuo.mutation.DisplayName(); ok {
		_spec.SetField(workflow.FieldDisplayName, field.TypeString, value)
	}
	if value, ok := wuo.mutation.StageIds(); ok {
		_spec.SetField(workflow.FieldStageIds, field.TypeJSON, value)
	}
	if value, ok := wuo.mutation.AppendedStageIds(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, workflow.FieldStageIds, value)
		})
	}
	if value, ok := wuo.mutation.Parallelism(); ok {
		_spec.SetField(workflow.FieldParallelism, field.TypeInt, value)
	}
	if value, ok := wuo.mutation.AddedParallelism(); ok {
		_spec.AddField(workflow.FieldParallelism, field.TypeInt, value)
	}
	if wuo.mutation.StagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.StagesTable,
			Columns: []string{workflow.StagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflowstage.FieldID, field.TypeString),
			},
		}
		edge.Schema = wuo.schemaConfig.WorkflowStage
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.RemovedStagesIDs(); len(nodes) > 0 && !wuo.mutation.StagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.StagesTable,
			Columns: []string{workflow.StagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflowstage.FieldID, field.TypeString),
			},
		}
		edge.Schema = wuo.schemaConfig.WorkflowStage
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.StagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.StagesTable,
			Columns: []string{workflow.StagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflowstage.FieldID, field.TypeString),
			},
		}
		edge.Schema = wuo.schemaConfig.WorkflowStage
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wuo.mutation.ExecutionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.ExecutionsTable,
			Columns: []string{workflow.ExecutionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflowexecution.FieldID, field.TypeString),
			},
		}
		edge.Schema = wuo.schemaConfig.WorkflowExecution
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.RemovedExecutionsIDs(); len(nodes) > 0 && !wuo.mutation.ExecutionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.ExecutionsTable,
			Columns: []string{workflow.ExecutionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflowexecution.FieldID, field.TypeString),
			},
		}
		edge.Schema = wuo.schemaConfig.WorkflowExecution
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.ExecutionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.ExecutionsTable,
			Columns: []string{workflow.ExecutionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workflowexecution.FieldID, field.TypeString),
			},
		}
		edge.Schema = wuo.schemaConfig.WorkflowExecution
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = wuo.schemaConfig.Workflow
	ctx = internal.NewSchemaConfigContext(ctx, wuo.schemaConfig)
	_spec.AddModifiers(wuo.modifiers...)
	_node = &Workflow{config: wuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{workflow.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	wuo.mutation.done = true
	return _node, nil
}