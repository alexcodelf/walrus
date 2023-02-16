// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// GENERATED, DO NOT EDIT.

package applicationmodulerelationship

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"

	"github.com/seal-io/seal/pkg/dao/model/internal"
	"github.com/seal-io/seal/pkg/dao/model/predicate"
	"github.com/seal-io/seal/pkg/dao/types"
)

// CreateTime applies equality check predicate on the "createTime" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.ApplicationModuleRelationship {
	return predicate.ApplicationModuleRelationship(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "updateTime" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.ApplicationModuleRelationship {
	return predicate.ApplicationModuleRelationship(sql.FieldEQ(FieldUpdateTime, v))
}

// ApplicationID applies equality check predicate on the "application_id" field. It's identical to ApplicationIDEQ.
func ApplicationID(v types.ID) predicate.ApplicationModuleRelationship {
	return predicate.ApplicationModuleRelationship(sql.FieldEQ(FieldApplicationID, v))
}

// ModuleID applies equality check predicate on the "module_id" field. It's identical to ModuleIDEQ.
func ModuleID(v string) predicate.ApplicationModuleRelationship {
	return predicate.ApplicationModuleRelationship(sql.FieldEQ(FieldModuleID, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.ApplicationModuleRelationship {
	return predicate.ApplicationModuleRelationship(sql.FieldEQ(FieldName, v))
}

// CreateTimeEQ applies the EQ predicate on the "createTime" field.
func CreateTimeEQ(v time.Time) predicate.ApplicationModuleRelationship {
	return predicate.ApplicationModuleRelationship(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "createTime" field.
func CreateTimeNEQ(v time.Time) predicate.ApplicationModuleRelationship {
	return predicate.ApplicationModuleRelationship(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "createTime" field.
func CreateTimeIn(vs ...time.Time) predicate.ApplicationModuleRelationship {
	return predicate.ApplicationModuleRelationship(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "createTime" field.
func CreateTimeNotIn(vs ...time.Time) predicate.ApplicationModuleRelationship {
	return predicate.ApplicationModuleRelationship(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "createTime" field.
func CreateTimeGT(v time.Time) predicate.ApplicationModuleRelationship {
	return predicate.ApplicationModuleRelationship(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "createTime" field.
func CreateTimeGTE(v time.Time) predicate.ApplicationModuleRelationship {
	return predicate.ApplicationModuleRelationship(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "createTime" field.
func CreateTimeLT(v time.Time) predicate.ApplicationModuleRelationship {
	return predicate.ApplicationModuleRelationship(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "createTime" field.
func CreateTimeLTE(v time.Time) predicate.ApplicationModuleRelationship {
	return predicate.ApplicationModuleRelationship(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "updateTime" field.
func UpdateTimeEQ(v time.Time) predicate.ApplicationModuleRelationship {
	return predicate.ApplicationModuleRelationship(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "updateTime" field.
func UpdateTimeNEQ(v time.Time) predicate.ApplicationModuleRelationship {
	return predicate.ApplicationModuleRelationship(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "updateTime" field.
func UpdateTimeIn(vs ...time.Time) predicate.ApplicationModuleRelationship {
	return predicate.ApplicationModuleRelationship(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "updateTime" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.ApplicationModuleRelationship {
	return predicate.ApplicationModuleRelationship(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "updateTime" field.
func UpdateTimeGT(v time.Time) predicate.ApplicationModuleRelationship {
	return predicate.ApplicationModuleRelationship(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "updateTime" field.
func UpdateTimeGTE(v time.Time) predicate.ApplicationModuleRelationship {
	return predicate.ApplicationModuleRelationship(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "updateTime" field.
func UpdateTimeLT(v time.Time) predicate.ApplicationModuleRelationship {
	return predicate.ApplicationModuleRelationship(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "updateTime" field.
func UpdateTimeLTE(v time.Time) predicate.ApplicationModuleRelationship {
	return predicate.ApplicationModuleRelationship(sql.FieldLTE(FieldUpdateTime, v))
}

// ApplicationIDEQ applies the EQ predicate on the "application_id" field.
func ApplicationIDEQ(v types.ID) predicate.ApplicationModuleRelationship {
	return predicate.ApplicationModuleRelationship(sql.FieldEQ(FieldApplicationID, v))
}

// ApplicationIDNEQ applies the NEQ predicate on the "application_id" field.
func ApplicationIDNEQ(v types.ID) predicate.ApplicationModuleRelationship {
	return predicate.ApplicationModuleRelationship(sql.FieldNEQ(FieldApplicationID, v))
}

// ApplicationIDIn applies the In predicate on the "application_id" field.
func ApplicationIDIn(vs ...types.ID) predicate.ApplicationModuleRelationship {
	return predicate.ApplicationModuleRelationship(sql.FieldIn(FieldApplicationID, vs...))
}

// ApplicationIDNotIn applies the NotIn predicate on the "application_id" field.
func ApplicationIDNotIn(vs ...types.ID) predicate.ApplicationModuleRelationship {
	return predicate.ApplicationModuleRelationship(sql.FieldNotIn(FieldApplicationID, vs...))
}

// ApplicationIDGT applies the GT predicate on the "application_id" field.
func ApplicationIDGT(v types.ID) predicate.ApplicationModuleRelationship {
	return predicate.ApplicationModuleRelationship(sql.FieldGT(FieldApplicationID, v))
}

// ApplicationIDGTE applies the GTE predicate on the "application_id" field.
func ApplicationIDGTE(v types.ID) predicate.ApplicationModuleRelationship {
	return predicate.ApplicationModuleRelationship(sql.FieldGTE(FieldApplicationID, v))
}

// ApplicationIDLT applies the LT predicate on the "application_id" field.
func ApplicationIDLT(v types.ID) predicate.ApplicationModuleRelationship {
	return predicate.ApplicationModuleRelationship(sql.FieldLT(FieldApplicationID, v))
}

// ApplicationIDLTE applies the LTE predicate on the "application_id" field.
func ApplicationIDLTE(v types.ID) predicate.ApplicationModuleRelationship {
	return predicate.ApplicationModuleRelationship(sql.FieldLTE(FieldApplicationID, v))
}

// ApplicationIDContains applies the Contains predicate on the "application_id" field.
func ApplicationIDContains(v types.ID) predicate.ApplicationModuleRelationship {
	vc := string(v)
	return predicate.ApplicationModuleRelationship(sql.FieldContains(FieldApplicationID, vc))
}

// ApplicationIDHasPrefix applies the HasPrefix predicate on the "application_id" field.
func ApplicationIDHasPrefix(v types.ID) predicate.ApplicationModuleRelationship {
	vc := string(v)
	return predicate.ApplicationModuleRelationship(sql.FieldHasPrefix(FieldApplicationID, vc))
}

// ApplicationIDHasSuffix applies the HasSuffix predicate on the "application_id" field.
func ApplicationIDHasSuffix(v types.ID) predicate.ApplicationModuleRelationship {
	vc := string(v)
	return predicate.ApplicationModuleRelationship(sql.FieldHasSuffix(FieldApplicationID, vc))
}

// ApplicationIDEqualFold applies the EqualFold predicate on the "application_id" field.
func ApplicationIDEqualFold(v types.ID) predicate.ApplicationModuleRelationship {
	vc := string(v)
	return predicate.ApplicationModuleRelationship(sql.FieldEqualFold(FieldApplicationID, vc))
}

// ApplicationIDContainsFold applies the ContainsFold predicate on the "application_id" field.
func ApplicationIDContainsFold(v types.ID) predicate.ApplicationModuleRelationship {
	vc := string(v)
	return predicate.ApplicationModuleRelationship(sql.FieldContainsFold(FieldApplicationID, vc))
}

// ModuleIDEQ applies the EQ predicate on the "module_id" field.
func ModuleIDEQ(v string) predicate.ApplicationModuleRelationship {
	return predicate.ApplicationModuleRelationship(sql.FieldEQ(FieldModuleID, v))
}

// ModuleIDNEQ applies the NEQ predicate on the "module_id" field.
func ModuleIDNEQ(v string) predicate.ApplicationModuleRelationship {
	return predicate.ApplicationModuleRelationship(sql.FieldNEQ(FieldModuleID, v))
}

// ModuleIDIn applies the In predicate on the "module_id" field.
func ModuleIDIn(vs ...string) predicate.ApplicationModuleRelationship {
	return predicate.ApplicationModuleRelationship(sql.FieldIn(FieldModuleID, vs...))
}

// ModuleIDNotIn applies the NotIn predicate on the "module_id" field.
func ModuleIDNotIn(vs ...string) predicate.ApplicationModuleRelationship {
	return predicate.ApplicationModuleRelationship(sql.FieldNotIn(FieldModuleID, vs...))
}

// ModuleIDGT applies the GT predicate on the "module_id" field.
func ModuleIDGT(v string) predicate.ApplicationModuleRelationship {
	return predicate.ApplicationModuleRelationship(sql.FieldGT(FieldModuleID, v))
}

// ModuleIDGTE applies the GTE predicate on the "module_id" field.
func ModuleIDGTE(v string) predicate.ApplicationModuleRelationship {
	return predicate.ApplicationModuleRelationship(sql.FieldGTE(FieldModuleID, v))
}

// ModuleIDLT applies the LT predicate on the "module_id" field.
func ModuleIDLT(v string) predicate.ApplicationModuleRelationship {
	return predicate.ApplicationModuleRelationship(sql.FieldLT(FieldModuleID, v))
}

// ModuleIDLTE applies the LTE predicate on the "module_id" field.
func ModuleIDLTE(v string) predicate.ApplicationModuleRelationship {
	return predicate.ApplicationModuleRelationship(sql.FieldLTE(FieldModuleID, v))
}

// ModuleIDContains applies the Contains predicate on the "module_id" field.
func ModuleIDContains(v string) predicate.ApplicationModuleRelationship {
	return predicate.ApplicationModuleRelationship(sql.FieldContains(FieldModuleID, v))
}

// ModuleIDHasPrefix applies the HasPrefix predicate on the "module_id" field.
func ModuleIDHasPrefix(v string) predicate.ApplicationModuleRelationship {
	return predicate.ApplicationModuleRelationship(sql.FieldHasPrefix(FieldModuleID, v))
}

// ModuleIDHasSuffix applies the HasSuffix predicate on the "module_id" field.
func ModuleIDHasSuffix(v string) predicate.ApplicationModuleRelationship {
	return predicate.ApplicationModuleRelationship(sql.FieldHasSuffix(FieldModuleID, v))
}

// ModuleIDEqualFold applies the EqualFold predicate on the "module_id" field.
func ModuleIDEqualFold(v string) predicate.ApplicationModuleRelationship {
	return predicate.ApplicationModuleRelationship(sql.FieldEqualFold(FieldModuleID, v))
}

// ModuleIDContainsFold applies the ContainsFold predicate on the "module_id" field.
func ModuleIDContainsFold(v string) predicate.ApplicationModuleRelationship {
	return predicate.ApplicationModuleRelationship(sql.FieldContainsFold(FieldModuleID, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.ApplicationModuleRelationship {
	return predicate.ApplicationModuleRelationship(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.ApplicationModuleRelationship {
	return predicate.ApplicationModuleRelationship(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.ApplicationModuleRelationship {
	return predicate.ApplicationModuleRelationship(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.ApplicationModuleRelationship {
	return predicate.ApplicationModuleRelationship(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.ApplicationModuleRelationship {
	return predicate.ApplicationModuleRelationship(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.ApplicationModuleRelationship {
	return predicate.ApplicationModuleRelationship(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.ApplicationModuleRelationship {
	return predicate.ApplicationModuleRelationship(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.ApplicationModuleRelationship {
	return predicate.ApplicationModuleRelationship(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.ApplicationModuleRelationship {
	return predicate.ApplicationModuleRelationship(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.ApplicationModuleRelationship {
	return predicate.ApplicationModuleRelationship(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.ApplicationModuleRelationship {
	return predicate.ApplicationModuleRelationship(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.ApplicationModuleRelationship {
	return predicate.ApplicationModuleRelationship(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.ApplicationModuleRelationship {
	return predicate.ApplicationModuleRelationship(sql.FieldContainsFold(FieldName, v))
}

// VariablesIsNil applies the IsNil predicate on the "variables" field.
func VariablesIsNil() predicate.ApplicationModuleRelationship {
	return predicate.ApplicationModuleRelationship(sql.FieldIsNull(FieldVariables))
}

// VariablesNotNil applies the NotNil predicate on the "variables" field.
func VariablesNotNil() predicate.ApplicationModuleRelationship {
	return predicate.ApplicationModuleRelationship(sql.FieldNotNull(FieldVariables))
}

// HasApplication applies the HasEdge predicate on the "application" edge.
func HasApplication() predicate.ApplicationModuleRelationship {
	return predicate.ApplicationModuleRelationship(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, ApplicationColumn),
			sqlgraph.Edge(sqlgraph.M2O, false, ApplicationTable, ApplicationColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Application
		step.Edge.Schema = schemaConfig.ApplicationModuleRelationship
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasApplicationWith applies the HasEdge predicate on the "application" edge with a given conditions (other predicates).
func HasApplicationWith(preds ...predicate.Application) predicate.ApplicationModuleRelationship {
	return predicate.ApplicationModuleRelationship(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, ApplicationColumn),
			sqlgraph.To(ApplicationInverseTable, ApplicationFieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ApplicationTable, ApplicationColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Application
		step.Edge.Schema = schemaConfig.ApplicationModuleRelationship
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasModule applies the HasEdge predicate on the "module" edge.
func HasModule() predicate.ApplicationModuleRelationship {
	return predicate.ApplicationModuleRelationship(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, ModuleColumn),
			sqlgraph.Edge(sqlgraph.M2O, false, ModuleTable, ModuleColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Module
		step.Edge.Schema = schemaConfig.ApplicationModuleRelationship
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasModuleWith applies the HasEdge predicate on the "module" edge with a given conditions (other predicates).
func HasModuleWith(preds ...predicate.Module) predicate.ApplicationModuleRelationship {
	return predicate.ApplicationModuleRelationship(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, ModuleColumn),
			sqlgraph.To(ModuleInverseTable, ModuleFieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ModuleTable, ModuleColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Module
		step.Edge.Schema = schemaConfig.ApplicationModuleRelationship
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ApplicationModuleRelationship) predicate.ApplicationModuleRelationship {
	return predicate.ApplicationModuleRelationship(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ApplicationModuleRelationship) predicate.ApplicationModuleRelationship {
	return predicate.ApplicationModuleRelationship(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ApplicationModuleRelationship) predicate.ApplicationModuleRelationship {
	return predicate.ApplicationModuleRelationship(func(s *sql.Selector) {
		p(s.Not())
	})
}