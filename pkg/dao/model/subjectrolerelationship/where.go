// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "ent". DO NOT EDIT.

package subjectrolerelationship

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"

	"github.com/seal-io/seal/pkg/dao/model/internal"
	"github.com/seal-io/seal/pkg/dao/model/predicate"
	"github.com/seal-io/seal/pkg/dao/types/oid"
)

// ID filters vertices based on their ID field.
func ID(id oid.ID) predicate.SubjectRoleRelationship {
	return predicate.SubjectRoleRelationship(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id oid.ID) predicate.SubjectRoleRelationship {
	return predicate.SubjectRoleRelationship(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id oid.ID) predicate.SubjectRoleRelationship {
	return predicate.SubjectRoleRelationship(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...oid.ID) predicate.SubjectRoleRelationship {
	return predicate.SubjectRoleRelationship(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...oid.ID) predicate.SubjectRoleRelationship {
	return predicate.SubjectRoleRelationship(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id oid.ID) predicate.SubjectRoleRelationship {
	return predicate.SubjectRoleRelationship(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id oid.ID) predicate.SubjectRoleRelationship {
	return predicate.SubjectRoleRelationship(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id oid.ID) predicate.SubjectRoleRelationship {
	return predicate.SubjectRoleRelationship(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id oid.ID) predicate.SubjectRoleRelationship {
	return predicate.SubjectRoleRelationship(sql.FieldLTE(FieldID, id))
}

// ProjectID applies equality check predicate on the "projectID" field. It's identical to ProjectIDEQ.
func ProjectID(v oid.ID) predicate.SubjectRoleRelationship {
	return predicate.SubjectRoleRelationship(sql.FieldEQ(FieldProjectID, v))
}

// CreateTime applies equality check predicate on the "createTime" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.SubjectRoleRelationship {
	return predicate.SubjectRoleRelationship(sql.FieldEQ(FieldCreateTime, v))
}

// SubjectID applies equality check predicate on the "subject_id" field. It's identical to SubjectIDEQ.
func SubjectID(v oid.ID) predicate.SubjectRoleRelationship {
	return predicate.SubjectRoleRelationship(sql.FieldEQ(FieldSubjectID, v))
}

// RoleID applies equality check predicate on the "role_id" field. It's identical to RoleIDEQ.
func RoleID(v string) predicate.SubjectRoleRelationship {
	return predicate.SubjectRoleRelationship(sql.FieldEQ(FieldRoleID, v))
}

// ProjectIDEQ applies the EQ predicate on the "projectID" field.
func ProjectIDEQ(v oid.ID) predicate.SubjectRoleRelationship {
	return predicate.SubjectRoleRelationship(sql.FieldEQ(FieldProjectID, v))
}

// ProjectIDNEQ applies the NEQ predicate on the "projectID" field.
func ProjectIDNEQ(v oid.ID) predicate.SubjectRoleRelationship {
	return predicate.SubjectRoleRelationship(sql.FieldNEQ(FieldProjectID, v))
}

// ProjectIDIn applies the In predicate on the "projectID" field.
func ProjectIDIn(vs ...oid.ID) predicate.SubjectRoleRelationship {
	return predicate.SubjectRoleRelationship(sql.FieldIn(FieldProjectID, vs...))
}

// ProjectIDNotIn applies the NotIn predicate on the "projectID" field.
func ProjectIDNotIn(vs ...oid.ID) predicate.SubjectRoleRelationship {
	return predicate.SubjectRoleRelationship(sql.FieldNotIn(FieldProjectID, vs...))
}

// ProjectIDGT applies the GT predicate on the "projectID" field.
func ProjectIDGT(v oid.ID) predicate.SubjectRoleRelationship {
	return predicate.SubjectRoleRelationship(sql.FieldGT(FieldProjectID, v))
}

// ProjectIDGTE applies the GTE predicate on the "projectID" field.
func ProjectIDGTE(v oid.ID) predicate.SubjectRoleRelationship {
	return predicate.SubjectRoleRelationship(sql.FieldGTE(FieldProjectID, v))
}

// ProjectIDLT applies the LT predicate on the "projectID" field.
func ProjectIDLT(v oid.ID) predicate.SubjectRoleRelationship {
	return predicate.SubjectRoleRelationship(sql.FieldLT(FieldProjectID, v))
}

// ProjectIDLTE applies the LTE predicate on the "projectID" field.
func ProjectIDLTE(v oid.ID) predicate.SubjectRoleRelationship {
	return predicate.SubjectRoleRelationship(sql.FieldLTE(FieldProjectID, v))
}

// ProjectIDContains applies the Contains predicate on the "projectID" field.
func ProjectIDContains(v oid.ID) predicate.SubjectRoleRelationship {
	vc := string(v)
	return predicate.SubjectRoleRelationship(sql.FieldContains(FieldProjectID, vc))
}

// ProjectIDHasPrefix applies the HasPrefix predicate on the "projectID" field.
func ProjectIDHasPrefix(v oid.ID) predicate.SubjectRoleRelationship {
	vc := string(v)
	return predicate.SubjectRoleRelationship(sql.FieldHasPrefix(FieldProjectID, vc))
}

// ProjectIDHasSuffix applies the HasSuffix predicate on the "projectID" field.
func ProjectIDHasSuffix(v oid.ID) predicate.SubjectRoleRelationship {
	vc := string(v)
	return predicate.SubjectRoleRelationship(sql.FieldHasSuffix(FieldProjectID, vc))
}

// ProjectIDIsNil applies the IsNil predicate on the "projectID" field.
func ProjectIDIsNil() predicate.SubjectRoleRelationship {
	return predicate.SubjectRoleRelationship(sql.FieldIsNull(FieldProjectID))
}

// ProjectIDNotNil applies the NotNil predicate on the "projectID" field.
func ProjectIDNotNil() predicate.SubjectRoleRelationship {
	return predicate.SubjectRoleRelationship(sql.FieldNotNull(FieldProjectID))
}

// ProjectIDEqualFold applies the EqualFold predicate on the "projectID" field.
func ProjectIDEqualFold(v oid.ID) predicate.SubjectRoleRelationship {
	vc := string(v)
	return predicate.SubjectRoleRelationship(sql.FieldEqualFold(FieldProjectID, vc))
}

// ProjectIDContainsFold applies the ContainsFold predicate on the "projectID" field.
func ProjectIDContainsFold(v oid.ID) predicate.SubjectRoleRelationship {
	vc := string(v)
	return predicate.SubjectRoleRelationship(sql.FieldContainsFold(FieldProjectID, vc))
}

// CreateTimeEQ applies the EQ predicate on the "createTime" field.
func CreateTimeEQ(v time.Time) predicate.SubjectRoleRelationship {
	return predicate.SubjectRoleRelationship(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "createTime" field.
func CreateTimeNEQ(v time.Time) predicate.SubjectRoleRelationship {
	return predicate.SubjectRoleRelationship(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "createTime" field.
func CreateTimeIn(vs ...time.Time) predicate.SubjectRoleRelationship {
	return predicate.SubjectRoleRelationship(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "createTime" field.
func CreateTimeNotIn(vs ...time.Time) predicate.SubjectRoleRelationship {
	return predicate.SubjectRoleRelationship(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "createTime" field.
func CreateTimeGT(v time.Time) predicate.SubjectRoleRelationship {
	return predicate.SubjectRoleRelationship(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "createTime" field.
func CreateTimeGTE(v time.Time) predicate.SubjectRoleRelationship {
	return predicate.SubjectRoleRelationship(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "createTime" field.
func CreateTimeLT(v time.Time) predicate.SubjectRoleRelationship {
	return predicate.SubjectRoleRelationship(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "createTime" field.
func CreateTimeLTE(v time.Time) predicate.SubjectRoleRelationship {
	return predicate.SubjectRoleRelationship(sql.FieldLTE(FieldCreateTime, v))
}

// SubjectIDEQ applies the EQ predicate on the "subject_id" field.
func SubjectIDEQ(v oid.ID) predicate.SubjectRoleRelationship {
	return predicate.SubjectRoleRelationship(sql.FieldEQ(FieldSubjectID, v))
}

// SubjectIDNEQ applies the NEQ predicate on the "subject_id" field.
func SubjectIDNEQ(v oid.ID) predicate.SubjectRoleRelationship {
	return predicate.SubjectRoleRelationship(sql.FieldNEQ(FieldSubjectID, v))
}

// SubjectIDIn applies the In predicate on the "subject_id" field.
func SubjectIDIn(vs ...oid.ID) predicate.SubjectRoleRelationship {
	return predicate.SubjectRoleRelationship(sql.FieldIn(FieldSubjectID, vs...))
}

// SubjectIDNotIn applies the NotIn predicate on the "subject_id" field.
func SubjectIDNotIn(vs ...oid.ID) predicate.SubjectRoleRelationship {
	return predicate.SubjectRoleRelationship(sql.FieldNotIn(FieldSubjectID, vs...))
}

// SubjectIDGT applies the GT predicate on the "subject_id" field.
func SubjectIDGT(v oid.ID) predicate.SubjectRoleRelationship {
	return predicate.SubjectRoleRelationship(sql.FieldGT(FieldSubjectID, v))
}

// SubjectIDGTE applies the GTE predicate on the "subject_id" field.
func SubjectIDGTE(v oid.ID) predicate.SubjectRoleRelationship {
	return predicate.SubjectRoleRelationship(sql.FieldGTE(FieldSubjectID, v))
}

// SubjectIDLT applies the LT predicate on the "subject_id" field.
func SubjectIDLT(v oid.ID) predicate.SubjectRoleRelationship {
	return predicate.SubjectRoleRelationship(sql.FieldLT(FieldSubjectID, v))
}

// SubjectIDLTE applies the LTE predicate on the "subject_id" field.
func SubjectIDLTE(v oid.ID) predicate.SubjectRoleRelationship {
	return predicate.SubjectRoleRelationship(sql.FieldLTE(FieldSubjectID, v))
}

// SubjectIDContains applies the Contains predicate on the "subject_id" field.
func SubjectIDContains(v oid.ID) predicate.SubjectRoleRelationship {
	vc := string(v)
	return predicate.SubjectRoleRelationship(sql.FieldContains(FieldSubjectID, vc))
}

// SubjectIDHasPrefix applies the HasPrefix predicate on the "subject_id" field.
func SubjectIDHasPrefix(v oid.ID) predicate.SubjectRoleRelationship {
	vc := string(v)
	return predicate.SubjectRoleRelationship(sql.FieldHasPrefix(FieldSubjectID, vc))
}

// SubjectIDHasSuffix applies the HasSuffix predicate on the "subject_id" field.
func SubjectIDHasSuffix(v oid.ID) predicate.SubjectRoleRelationship {
	vc := string(v)
	return predicate.SubjectRoleRelationship(sql.FieldHasSuffix(FieldSubjectID, vc))
}

// SubjectIDEqualFold applies the EqualFold predicate on the "subject_id" field.
func SubjectIDEqualFold(v oid.ID) predicate.SubjectRoleRelationship {
	vc := string(v)
	return predicate.SubjectRoleRelationship(sql.FieldEqualFold(FieldSubjectID, vc))
}

// SubjectIDContainsFold applies the ContainsFold predicate on the "subject_id" field.
func SubjectIDContainsFold(v oid.ID) predicate.SubjectRoleRelationship {
	vc := string(v)
	return predicate.SubjectRoleRelationship(sql.FieldContainsFold(FieldSubjectID, vc))
}

// RoleIDEQ applies the EQ predicate on the "role_id" field.
func RoleIDEQ(v string) predicate.SubjectRoleRelationship {
	return predicate.SubjectRoleRelationship(sql.FieldEQ(FieldRoleID, v))
}

// RoleIDNEQ applies the NEQ predicate on the "role_id" field.
func RoleIDNEQ(v string) predicate.SubjectRoleRelationship {
	return predicate.SubjectRoleRelationship(sql.FieldNEQ(FieldRoleID, v))
}

// RoleIDIn applies the In predicate on the "role_id" field.
func RoleIDIn(vs ...string) predicate.SubjectRoleRelationship {
	return predicate.SubjectRoleRelationship(sql.FieldIn(FieldRoleID, vs...))
}

// RoleIDNotIn applies the NotIn predicate on the "role_id" field.
func RoleIDNotIn(vs ...string) predicate.SubjectRoleRelationship {
	return predicate.SubjectRoleRelationship(sql.FieldNotIn(FieldRoleID, vs...))
}

// RoleIDGT applies the GT predicate on the "role_id" field.
func RoleIDGT(v string) predicate.SubjectRoleRelationship {
	return predicate.SubjectRoleRelationship(sql.FieldGT(FieldRoleID, v))
}

// RoleIDGTE applies the GTE predicate on the "role_id" field.
func RoleIDGTE(v string) predicate.SubjectRoleRelationship {
	return predicate.SubjectRoleRelationship(sql.FieldGTE(FieldRoleID, v))
}

// RoleIDLT applies the LT predicate on the "role_id" field.
func RoleIDLT(v string) predicate.SubjectRoleRelationship {
	return predicate.SubjectRoleRelationship(sql.FieldLT(FieldRoleID, v))
}

// RoleIDLTE applies the LTE predicate on the "role_id" field.
func RoleIDLTE(v string) predicate.SubjectRoleRelationship {
	return predicate.SubjectRoleRelationship(sql.FieldLTE(FieldRoleID, v))
}

// RoleIDContains applies the Contains predicate on the "role_id" field.
func RoleIDContains(v string) predicate.SubjectRoleRelationship {
	return predicate.SubjectRoleRelationship(sql.FieldContains(FieldRoleID, v))
}

// RoleIDHasPrefix applies the HasPrefix predicate on the "role_id" field.
func RoleIDHasPrefix(v string) predicate.SubjectRoleRelationship {
	return predicate.SubjectRoleRelationship(sql.FieldHasPrefix(FieldRoleID, v))
}

// RoleIDHasSuffix applies the HasSuffix predicate on the "role_id" field.
func RoleIDHasSuffix(v string) predicate.SubjectRoleRelationship {
	return predicate.SubjectRoleRelationship(sql.FieldHasSuffix(FieldRoleID, v))
}

// RoleIDEqualFold applies the EqualFold predicate on the "role_id" field.
func RoleIDEqualFold(v string) predicate.SubjectRoleRelationship {
	return predicate.SubjectRoleRelationship(sql.FieldEqualFold(FieldRoleID, v))
}

// RoleIDContainsFold applies the ContainsFold predicate on the "role_id" field.
func RoleIDContainsFold(v string) predicate.SubjectRoleRelationship {
	return predicate.SubjectRoleRelationship(sql.FieldContainsFold(FieldRoleID, v))
}

// HasProject applies the HasEdge predicate on the "project" edge.
func HasProject() predicate.SubjectRoleRelationship {
	return predicate.SubjectRoleRelationship(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProjectTable, ProjectColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Project
		step.Edge.Schema = schemaConfig.SubjectRoleRelationship
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProjectWith applies the HasEdge predicate on the "project" edge with a given conditions (other predicates).
func HasProjectWith(preds ...predicate.Project) predicate.SubjectRoleRelationship {
	return predicate.SubjectRoleRelationship(func(s *sql.Selector) {
		step := newProjectStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Project
		step.Edge.Schema = schemaConfig.SubjectRoleRelationship
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSubject applies the HasEdge predicate on the "subject" edge.
func HasSubject() predicate.SubjectRoleRelationship {
	return predicate.SubjectRoleRelationship(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, SubjectTable, SubjectColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Subject
		step.Edge.Schema = schemaConfig.SubjectRoleRelationship
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSubjectWith applies the HasEdge predicate on the "subject" edge with a given conditions (other predicates).
func HasSubjectWith(preds ...predicate.Subject) predicate.SubjectRoleRelationship {
	return predicate.SubjectRoleRelationship(func(s *sql.Selector) {
		step := newSubjectStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Subject
		step.Edge.Schema = schemaConfig.SubjectRoleRelationship
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRole applies the HasEdge predicate on the "role" edge.
func HasRole() predicate.SubjectRoleRelationship {
	return predicate.SubjectRoleRelationship(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, RoleTable, RoleColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Role
		step.Edge.Schema = schemaConfig.SubjectRoleRelationship
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRoleWith applies the HasEdge predicate on the "role" edge with a given conditions (other predicates).
func HasRoleWith(preds ...predicate.Role) predicate.SubjectRoleRelationship {
	return predicate.SubjectRoleRelationship(func(s *sql.Selector) {
		step := newRoleStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Role
		step.Edge.Schema = schemaConfig.SubjectRoleRelationship
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SubjectRoleRelationship) predicate.SubjectRoleRelationship {
	return predicate.SubjectRoleRelationship(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SubjectRoleRelationship) predicate.SubjectRoleRelationship {
	return predicate.SubjectRoleRelationship(func(s *sql.Selector) {
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
func Not(p predicate.SubjectRoleRelationship) predicate.SubjectRoleRelationship {
	return predicate.SubjectRoleRelationship(func(s *sql.Selector) {
		p(s.Not())
	})
}
