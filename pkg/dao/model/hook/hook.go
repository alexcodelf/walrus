// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// GENERATED, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/seal-io/seal/pkg/dao/model"
)

// The AllocationCostFunc type is an adapter to allow the use of ordinary
// function as AllocationCost mutator.
type AllocationCostFunc func(context.Context, *model.AllocationCostMutation) (model.Value, error)

// Mutate calls f(ctx, m).
func (f AllocationCostFunc) Mutate(ctx context.Context, m model.Mutation) (model.Value, error) {
	if mv, ok := m.(*model.AllocationCostMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *model.AllocationCostMutation", m)
}

// The ApplicationFunc type is an adapter to allow the use of ordinary
// function as Application mutator.
type ApplicationFunc func(context.Context, *model.ApplicationMutation) (model.Value, error)

// Mutate calls f(ctx, m).
func (f ApplicationFunc) Mutate(ctx context.Context, m model.Mutation) (model.Value, error) {
	if mv, ok := m.(*model.ApplicationMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *model.ApplicationMutation", m)
}

// The ApplicationInstanceFunc type is an adapter to allow the use of ordinary
// function as ApplicationInstance mutator.
type ApplicationInstanceFunc func(context.Context, *model.ApplicationInstanceMutation) (model.Value, error)

// Mutate calls f(ctx, m).
func (f ApplicationInstanceFunc) Mutate(ctx context.Context, m model.Mutation) (model.Value, error) {
	if mv, ok := m.(*model.ApplicationInstanceMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *model.ApplicationInstanceMutation", m)
}

// The ApplicationModuleRelationshipFunc type is an adapter to allow the use of ordinary
// function as ApplicationModuleRelationship mutator.
type ApplicationModuleRelationshipFunc func(context.Context, *model.ApplicationModuleRelationshipMutation) (model.Value, error)

// Mutate calls f(ctx, m).
func (f ApplicationModuleRelationshipFunc) Mutate(ctx context.Context, m model.Mutation) (model.Value, error) {
	if mv, ok := m.(*model.ApplicationModuleRelationshipMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *model.ApplicationModuleRelationshipMutation", m)
}

// The ApplicationResourceFunc type is an adapter to allow the use of ordinary
// function as ApplicationResource mutator.
type ApplicationResourceFunc func(context.Context, *model.ApplicationResourceMutation) (model.Value, error)

// Mutate calls f(ctx, m).
func (f ApplicationResourceFunc) Mutate(ctx context.Context, m model.Mutation) (model.Value, error) {
	if mv, ok := m.(*model.ApplicationResourceMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *model.ApplicationResourceMutation", m)
}

// The ApplicationRevisionFunc type is an adapter to allow the use of ordinary
// function as ApplicationRevision mutator.
type ApplicationRevisionFunc func(context.Context, *model.ApplicationRevisionMutation) (model.Value, error)

// Mutate calls f(ctx, m).
func (f ApplicationRevisionFunc) Mutate(ctx context.Context, m model.Mutation) (model.Value, error) {
	if mv, ok := m.(*model.ApplicationRevisionMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *model.ApplicationRevisionMutation", m)
}

// The ClusterCostFunc type is an adapter to allow the use of ordinary
// function as ClusterCost mutator.
type ClusterCostFunc func(context.Context, *model.ClusterCostMutation) (model.Value, error)

// Mutate calls f(ctx, m).
func (f ClusterCostFunc) Mutate(ctx context.Context, m model.Mutation) (model.Value, error) {
	if mv, ok := m.(*model.ClusterCostMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *model.ClusterCostMutation", m)
}

// The ConnectorFunc type is an adapter to allow the use of ordinary
// function as Connector mutator.
type ConnectorFunc func(context.Context, *model.ConnectorMutation) (model.Value, error)

// Mutate calls f(ctx, m).
func (f ConnectorFunc) Mutate(ctx context.Context, m model.Mutation) (model.Value, error) {
	if mv, ok := m.(*model.ConnectorMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *model.ConnectorMutation", m)
}

// The EnvironmentFunc type is an adapter to allow the use of ordinary
// function as Environment mutator.
type EnvironmentFunc func(context.Context, *model.EnvironmentMutation) (model.Value, error)

// Mutate calls f(ctx, m).
func (f EnvironmentFunc) Mutate(ctx context.Context, m model.Mutation) (model.Value, error) {
	if mv, ok := m.(*model.EnvironmentMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *model.EnvironmentMutation", m)
}

// The EnvironmentConnectorRelationshipFunc type is an adapter to allow the use of ordinary
// function as EnvironmentConnectorRelationship mutator.
type EnvironmentConnectorRelationshipFunc func(context.Context, *model.EnvironmentConnectorRelationshipMutation) (model.Value, error)

// Mutate calls f(ctx, m).
func (f EnvironmentConnectorRelationshipFunc) Mutate(ctx context.Context, m model.Mutation) (model.Value, error) {
	if mv, ok := m.(*model.EnvironmentConnectorRelationshipMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *model.EnvironmentConnectorRelationshipMutation", m)
}

// The ModuleFunc type is an adapter to allow the use of ordinary
// function as Module mutator.
type ModuleFunc func(context.Context, *model.ModuleMutation) (model.Value, error)

// Mutate calls f(ctx, m).
func (f ModuleFunc) Mutate(ctx context.Context, m model.Mutation) (model.Value, error) {
	if mv, ok := m.(*model.ModuleMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *model.ModuleMutation", m)
}

// The ModuleVersionFunc type is an adapter to allow the use of ordinary
// function as ModuleVersion mutator.
type ModuleVersionFunc func(context.Context, *model.ModuleVersionMutation) (model.Value, error)

// Mutate calls f(ctx, m).
func (f ModuleVersionFunc) Mutate(ctx context.Context, m model.Mutation) (model.Value, error) {
	if mv, ok := m.(*model.ModuleVersionMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *model.ModuleVersionMutation", m)
}

// The PerspectiveFunc type is an adapter to allow the use of ordinary
// function as Perspective mutator.
type PerspectiveFunc func(context.Context, *model.PerspectiveMutation) (model.Value, error)

// Mutate calls f(ctx, m).
func (f PerspectiveFunc) Mutate(ctx context.Context, m model.Mutation) (model.Value, error) {
	if mv, ok := m.(*model.PerspectiveMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *model.PerspectiveMutation", m)
}

// The ProjectFunc type is an adapter to allow the use of ordinary
// function as Project mutator.
type ProjectFunc func(context.Context, *model.ProjectMutation) (model.Value, error)

// Mutate calls f(ctx, m).
func (f ProjectFunc) Mutate(ctx context.Context, m model.Mutation) (model.Value, error) {
	if mv, ok := m.(*model.ProjectMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *model.ProjectMutation", m)
}

// The RoleFunc type is an adapter to allow the use of ordinary
// function as Role mutator.
type RoleFunc func(context.Context, *model.RoleMutation) (model.Value, error)

// Mutate calls f(ctx, m).
func (f RoleFunc) Mutate(ctx context.Context, m model.Mutation) (model.Value, error) {
	if mv, ok := m.(*model.RoleMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *model.RoleMutation", m)
}

// The SecretFunc type is an adapter to allow the use of ordinary
// function as Secret mutator.
type SecretFunc func(context.Context, *model.SecretMutation) (model.Value, error)

// Mutate calls f(ctx, m).
func (f SecretFunc) Mutate(ctx context.Context, m model.Mutation) (model.Value, error) {
	if mv, ok := m.(*model.SecretMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *model.SecretMutation", m)
}

// The SettingFunc type is an adapter to allow the use of ordinary
// function as Setting mutator.
type SettingFunc func(context.Context, *model.SettingMutation) (model.Value, error)

// Mutate calls f(ctx, m).
func (f SettingFunc) Mutate(ctx context.Context, m model.Mutation) (model.Value, error) {
	if mv, ok := m.(*model.SettingMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *model.SettingMutation", m)
}

// The SubjectFunc type is an adapter to allow the use of ordinary
// function as Subject mutator.
type SubjectFunc func(context.Context, *model.SubjectMutation) (model.Value, error)

// Mutate calls f(ctx, m).
func (f SubjectFunc) Mutate(ctx context.Context, m model.Mutation) (model.Value, error) {
	if mv, ok := m.(*model.SubjectMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *model.SubjectMutation", m)
}

// The TokenFunc type is an adapter to allow the use of ordinary
// function as Token mutator.
type TokenFunc func(context.Context, *model.TokenMutation) (model.Value, error)

// Mutate calls f(ctx, m).
func (f TokenFunc) Mutate(ctx context.Context, m model.Mutation) (model.Value, error) {
	if mv, ok := m.(*model.TokenMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *model.TokenMutation", m)
}

// Condition is a hook condition function.
type Condition func(context.Context, model.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m model.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m model.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m model.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op model.Op) Condition {
	return func(_ context.Context, m model.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m model.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m model.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m model.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
func If(hk model.Hook, cond Condition) model.Hook {
	return func(next model.Mutator) model.Mutator {
		return model.MutateFunc(func(ctx context.Context, m model.Mutation) (model.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, model.Delete|model.Create)
func On(hk model.Hook, op model.Op) model.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, model.Update|model.UpdateOne)
func Unless(hk model.Hook, op model.Op) model.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) model.Hook {
	return func(model.Mutator) model.Mutator {
		return model.MutateFunc(func(context.Context, model.Mutation) (model.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []model.Hook {
//		return []model.Hook{
//			Reject(model.Delete|model.Update),
//		}
//	}
func Reject(op model.Op) model.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []model.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...model.Hook) Chain {
	return Chain{append([]model.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() model.Hook {
	return func(mutator model.Mutator) model.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...model.Hook) Chain {
	newHooks := make([]model.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
