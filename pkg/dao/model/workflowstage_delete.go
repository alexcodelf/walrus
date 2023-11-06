// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus". DO NOT EDIT.

package model

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"github.com/seal-io/walrus/pkg/dao/model/internal"
	"github.com/seal-io/walrus/pkg/dao/model/predicate"
	"github.com/seal-io/walrus/pkg/dao/model/workflowstage"
)

// WorkflowStageDelete is the builder for deleting a WorkflowStage entity.
type WorkflowStageDelete struct {
	config
	hooks    []Hook
	mutation *WorkflowStageMutation
}

// Where appends a list predicates to the WorkflowStageDelete builder.
func (wsd *WorkflowStageDelete) Where(ps ...predicate.WorkflowStage) *WorkflowStageDelete {
	wsd.mutation.Where(ps...)
	return wsd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (wsd *WorkflowStageDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, wsd.sqlExec, wsd.mutation, wsd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (wsd *WorkflowStageDelete) ExecX(ctx context.Context) int {
	n, err := wsd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (wsd *WorkflowStageDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(workflowstage.Table, sqlgraph.NewFieldSpec(workflowstage.FieldID, field.TypeString))
	_spec.Node.Schema = wsd.schemaConfig.WorkflowStage
	ctx = internal.NewSchemaConfigContext(ctx, wsd.schemaConfig)
	if ps := wsd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, wsd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	wsd.mutation.done = true
	return affected, err
}

// WorkflowStageDeleteOne is the builder for deleting a single WorkflowStage entity.
type WorkflowStageDeleteOne struct {
	wsd *WorkflowStageDelete
}

// Where appends a list predicates to the WorkflowStageDelete builder.
func (wsdo *WorkflowStageDeleteOne) Where(ps ...predicate.WorkflowStage) *WorkflowStageDeleteOne {
	wsdo.wsd.mutation.Where(ps...)
	return wsdo
}

// Exec executes the deletion query.
func (wsdo *WorkflowStageDeleteOne) Exec(ctx context.Context) error {
	n, err := wsdo.wsd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{workflowstage.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (wsdo *WorkflowStageDeleteOne) ExecX(ctx context.Context) {
	if err := wsdo.Exec(ctx); err != nil {
		panic(err)
	}
}