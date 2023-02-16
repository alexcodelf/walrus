// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// GENERATED, DO NOT EDIT.

package model

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"

	"github.com/seal-io/seal/pkg/dao/model/environmentconnectorrelationship"
	"github.com/seal-io/seal/pkg/dao/model/internal"
	"github.com/seal-io/seal/pkg/dao/model/predicate"
)

// EnvironmentConnectorRelationshipDelete is the builder for deleting a EnvironmentConnectorRelationship entity.
type EnvironmentConnectorRelationshipDelete struct {
	config
	hooks    []Hook
	mutation *EnvironmentConnectorRelationshipMutation
}

// Where appends a list predicates to the EnvironmentConnectorRelationshipDelete builder.
func (ecrd *EnvironmentConnectorRelationshipDelete) Where(ps ...predicate.EnvironmentConnectorRelationship) *EnvironmentConnectorRelationshipDelete {
	ecrd.mutation.Where(ps...)
	return ecrd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ecrd *EnvironmentConnectorRelationshipDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, EnvironmentConnectorRelationshipMutation](ctx, ecrd.sqlExec, ecrd.mutation, ecrd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ecrd *EnvironmentConnectorRelationshipDelete) ExecX(ctx context.Context) int {
	n, err := ecrd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ecrd *EnvironmentConnectorRelationshipDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: environmentconnectorrelationship.Table,
		},
	}
	_spec.Node.Schema = ecrd.schemaConfig.EnvironmentConnectorRelationship
	ctx = internal.NewSchemaConfigContext(ctx, ecrd.schemaConfig)
	if ps := ecrd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ecrd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ecrd.mutation.done = true
	return affected, err
}

// EnvironmentConnectorRelationshipDeleteOne is the builder for deleting a single EnvironmentConnectorRelationship entity.
type EnvironmentConnectorRelationshipDeleteOne struct {
	ecrd *EnvironmentConnectorRelationshipDelete
}

// Where appends a list predicates to the EnvironmentConnectorRelationshipDelete builder.
func (ecrdo *EnvironmentConnectorRelationshipDeleteOne) Where(ps ...predicate.EnvironmentConnectorRelationship) *EnvironmentConnectorRelationshipDeleteOne {
	ecrdo.ecrd.mutation.Where(ps...)
	return ecrdo
}

// Exec executes the deletion query.
func (ecrdo *EnvironmentConnectorRelationshipDeleteOne) Exec(ctx context.Context) error {
	n, err := ecrdo.ecrd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{environmentconnectorrelationship.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ecrdo *EnvironmentConnectorRelationshipDeleteOne) ExecX(ctx context.Context) {
	if err := ecrdo.Exec(ctx); err != nil {
		panic(err)
	}
}