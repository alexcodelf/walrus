// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus". DO NOT EDIT.

package workflowstage

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"

	"github.com/seal-io/walrus/pkg/dao/model/internal"
	"github.com/seal-io/walrus/pkg/dao/model/predicate"
	"github.com/seal-io/walrus/pkg/dao/types/object"
)

// ID filters vertices based on their ID field.
func ID(id object.ID) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id object.ID) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id object.ID) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...object.ID) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...object.ID) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id object.ID) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id object.ID) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id object.ID) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id object.ID) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldEQ(FieldName, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldEQ(FieldDescription, v))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldEQ(FieldUpdateTime, v))
}

// ProjectID applies equality check predicate on the "project_id" field. It's identical to ProjectIDEQ.
func ProjectID(v object.ID) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldEQ(FieldProjectID, v))
}

// WorkflowID applies equality check predicate on the "workflow_id" field. It's identical to WorkflowIDEQ.
func WorkflowID(v object.ID) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldEQ(FieldWorkflowID, v))
}

// Duration applies equality check predicate on the "duration" field. It's identical to DurationEQ.
func Duration(v int) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldEQ(FieldDuration, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldContainsFold(FieldName, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldContainsFold(FieldDescription, v))
}

// LabelsIsNil applies the IsNil predicate on the "labels" field.
func LabelsIsNil() predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldIsNull(FieldLabels))
}

// LabelsNotNil applies the NotNil predicate on the "labels" field.
func LabelsNotNil() predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldNotNull(FieldLabels))
}

// AnnotationsIsNil applies the IsNil predicate on the "annotations" field.
func AnnotationsIsNil() predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldIsNull(FieldAnnotations))
}

// AnnotationsNotNil applies the NotNil predicate on the "annotations" field.
func AnnotationsNotNil() predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldNotNull(FieldAnnotations))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldLTE(FieldUpdateTime, v))
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldIsNull(FieldStatus))
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldNotNull(FieldStatus))
}

// ProjectIDEQ applies the EQ predicate on the "project_id" field.
func ProjectIDEQ(v object.ID) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldEQ(FieldProjectID, v))
}

// ProjectIDNEQ applies the NEQ predicate on the "project_id" field.
func ProjectIDNEQ(v object.ID) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldNEQ(FieldProjectID, v))
}

// ProjectIDIn applies the In predicate on the "project_id" field.
func ProjectIDIn(vs ...object.ID) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldIn(FieldProjectID, vs...))
}

// ProjectIDNotIn applies the NotIn predicate on the "project_id" field.
func ProjectIDNotIn(vs ...object.ID) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldNotIn(FieldProjectID, vs...))
}

// ProjectIDGT applies the GT predicate on the "project_id" field.
func ProjectIDGT(v object.ID) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldGT(FieldProjectID, v))
}

// ProjectIDGTE applies the GTE predicate on the "project_id" field.
func ProjectIDGTE(v object.ID) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldGTE(FieldProjectID, v))
}

// ProjectIDLT applies the LT predicate on the "project_id" field.
func ProjectIDLT(v object.ID) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldLT(FieldProjectID, v))
}

// ProjectIDLTE applies the LTE predicate on the "project_id" field.
func ProjectIDLTE(v object.ID) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldLTE(FieldProjectID, v))
}

// ProjectIDContains applies the Contains predicate on the "project_id" field.
func ProjectIDContains(v object.ID) predicate.WorkflowStage {
	vc := string(v)
	return predicate.WorkflowStage(sql.FieldContains(FieldProjectID, vc))
}

// ProjectIDHasPrefix applies the HasPrefix predicate on the "project_id" field.
func ProjectIDHasPrefix(v object.ID) predicate.WorkflowStage {
	vc := string(v)
	return predicate.WorkflowStage(sql.FieldHasPrefix(FieldProjectID, vc))
}

// ProjectIDHasSuffix applies the HasSuffix predicate on the "project_id" field.
func ProjectIDHasSuffix(v object.ID) predicate.WorkflowStage {
	vc := string(v)
	return predicate.WorkflowStage(sql.FieldHasSuffix(FieldProjectID, vc))
}

// ProjectIDEqualFold applies the EqualFold predicate on the "project_id" field.
func ProjectIDEqualFold(v object.ID) predicate.WorkflowStage {
	vc := string(v)
	return predicate.WorkflowStage(sql.FieldEqualFold(FieldProjectID, vc))
}

// ProjectIDContainsFold applies the ContainsFold predicate on the "project_id" field.
func ProjectIDContainsFold(v object.ID) predicate.WorkflowStage {
	vc := string(v)
	return predicate.WorkflowStage(sql.FieldContainsFold(FieldProjectID, vc))
}

// WorkflowIDEQ applies the EQ predicate on the "workflow_id" field.
func WorkflowIDEQ(v object.ID) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldEQ(FieldWorkflowID, v))
}

// WorkflowIDNEQ applies the NEQ predicate on the "workflow_id" field.
func WorkflowIDNEQ(v object.ID) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldNEQ(FieldWorkflowID, v))
}

// WorkflowIDIn applies the In predicate on the "workflow_id" field.
func WorkflowIDIn(vs ...object.ID) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldIn(FieldWorkflowID, vs...))
}

// WorkflowIDNotIn applies the NotIn predicate on the "workflow_id" field.
func WorkflowIDNotIn(vs ...object.ID) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldNotIn(FieldWorkflowID, vs...))
}

// WorkflowIDGT applies the GT predicate on the "workflow_id" field.
func WorkflowIDGT(v object.ID) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldGT(FieldWorkflowID, v))
}

// WorkflowIDGTE applies the GTE predicate on the "workflow_id" field.
func WorkflowIDGTE(v object.ID) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldGTE(FieldWorkflowID, v))
}

// WorkflowIDLT applies the LT predicate on the "workflow_id" field.
func WorkflowIDLT(v object.ID) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldLT(FieldWorkflowID, v))
}

// WorkflowIDLTE applies the LTE predicate on the "workflow_id" field.
func WorkflowIDLTE(v object.ID) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldLTE(FieldWorkflowID, v))
}

// WorkflowIDContains applies the Contains predicate on the "workflow_id" field.
func WorkflowIDContains(v object.ID) predicate.WorkflowStage {
	vc := string(v)
	return predicate.WorkflowStage(sql.FieldContains(FieldWorkflowID, vc))
}

// WorkflowIDHasPrefix applies the HasPrefix predicate on the "workflow_id" field.
func WorkflowIDHasPrefix(v object.ID) predicate.WorkflowStage {
	vc := string(v)
	return predicate.WorkflowStage(sql.FieldHasPrefix(FieldWorkflowID, vc))
}

// WorkflowIDHasSuffix applies the HasSuffix predicate on the "workflow_id" field.
func WorkflowIDHasSuffix(v object.ID) predicate.WorkflowStage {
	vc := string(v)
	return predicate.WorkflowStage(sql.FieldHasSuffix(FieldWorkflowID, vc))
}

// WorkflowIDEqualFold applies the EqualFold predicate on the "workflow_id" field.
func WorkflowIDEqualFold(v object.ID) predicate.WorkflowStage {
	vc := string(v)
	return predicate.WorkflowStage(sql.FieldEqualFold(FieldWorkflowID, vc))
}

// WorkflowIDContainsFold applies the ContainsFold predicate on the "workflow_id" field.
func WorkflowIDContainsFold(v object.ID) predicate.WorkflowStage {
	vc := string(v)
	return predicate.WorkflowStage(sql.FieldContainsFold(FieldWorkflowID, vc))
}

// DurationEQ applies the EQ predicate on the "duration" field.
func DurationEQ(v int) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldEQ(FieldDuration, v))
}

// DurationNEQ applies the NEQ predicate on the "duration" field.
func DurationNEQ(v int) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldNEQ(FieldDuration, v))
}

// DurationIn applies the In predicate on the "duration" field.
func DurationIn(vs ...int) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldIn(FieldDuration, vs...))
}

// DurationNotIn applies the NotIn predicate on the "duration" field.
func DurationNotIn(vs ...int) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldNotIn(FieldDuration, vs...))
}

// DurationGT applies the GT predicate on the "duration" field.
func DurationGT(v int) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldGT(FieldDuration, v))
}

// DurationGTE applies the GTE predicate on the "duration" field.
func DurationGTE(v int) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldGTE(FieldDuration, v))
}

// DurationLT applies the LT predicate on the "duration" field.
func DurationLT(v int) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldLT(FieldDuration, v))
}

// DurationLTE applies the LTE predicate on the "duration" field.
func DurationLTE(v int) predicate.WorkflowStage {
	return predicate.WorkflowStage(sql.FieldLTE(FieldDuration, v))
}

// HasSteps applies the HasEdge predicate on the "steps" edge.
func HasSteps() predicate.WorkflowStage {
	return predicate.WorkflowStage(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, StepsTable, StepsColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.WorkflowStep
		step.Edge.Schema = schemaConfig.WorkflowStep
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStepsWith applies the HasEdge predicate on the "steps" edge with a given conditions (other predicates).
func HasStepsWith(preds ...predicate.WorkflowStep) predicate.WorkflowStage {
	return predicate.WorkflowStage(func(s *sql.Selector) {
		step := newStepsStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.WorkflowStep
		step.Edge.Schema = schemaConfig.WorkflowStep
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasWorkflowStageExecutions applies the HasEdge predicate on the "workflow_stage_executions" edge.
func HasWorkflowStageExecutions() predicate.WorkflowStage {
	return predicate.WorkflowStage(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, WorkflowStageExecutionsTable, WorkflowStageExecutionsColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.WorkflowStageExecution
		step.Edge.Schema = schemaConfig.WorkflowStageExecution
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWorkflowStageExecutionsWith applies the HasEdge predicate on the "workflow_stage_executions" edge with a given conditions (other predicates).
func HasWorkflowStageExecutionsWith(preds ...predicate.WorkflowStageExecution) predicate.WorkflowStage {
	return predicate.WorkflowStage(func(s *sql.Selector) {
		step := newWorkflowStageExecutionsStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.WorkflowStageExecution
		step.Edge.Schema = schemaConfig.WorkflowStageExecution
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasWorkflow applies the HasEdge predicate on the "workflow" edge.
func HasWorkflow() predicate.WorkflowStage {
	return predicate.WorkflowStage(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, WorkflowTable, WorkflowColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Workflow
		step.Edge.Schema = schemaConfig.WorkflowStage
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWorkflowWith applies the HasEdge predicate on the "workflow" edge with a given conditions (other predicates).
func HasWorkflowWith(preds ...predicate.Workflow) predicate.WorkflowStage {
	return predicate.WorkflowStage(func(s *sql.Selector) {
		step := newWorkflowStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Workflow
		step.Edge.Schema = schemaConfig.WorkflowStage
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.WorkflowStage) predicate.WorkflowStage {
	return predicate.WorkflowStage(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.WorkflowStage) predicate.WorkflowStage {
	return predicate.WorkflowStage(func(s *sql.Selector) {
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
func Not(p predicate.WorkflowStage) predicate.WorkflowStage {
	return predicate.WorkflowStage(func(s *sql.Selector) {
		p(s.Not())
	})
}
