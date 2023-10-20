// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus". DO NOT EDIT.

package workflow

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"

	"github.com/seal-io/walrus/pkg/dao/model/internal"
	"github.com/seal-io/walrus/pkg/dao/model/predicate"
	"github.com/seal-io/walrus/pkg/dao/types/object"
)

// ID filters vertices based on their ID field.
func ID(id object.ID) predicate.Workflow {
	return predicate.Workflow(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id object.ID) predicate.Workflow {
	return predicate.Workflow(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id object.ID) predicate.Workflow {
	return predicate.Workflow(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...object.ID) predicate.Workflow {
	return predicate.Workflow(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...object.ID) predicate.Workflow {
	return predicate.Workflow(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id object.ID) predicate.Workflow {
	return predicate.Workflow(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id object.ID) predicate.Workflow {
	return predicate.Workflow(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id object.ID) predicate.Workflow {
	return predicate.Workflow(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id object.ID) predicate.Workflow {
	return predicate.Workflow(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Workflow {
	return predicate.Workflow(sql.FieldEQ(FieldName, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Workflow {
	return predicate.Workflow(sql.FieldEQ(FieldDescription, v))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Workflow {
	return predicate.Workflow(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Workflow {
	return predicate.Workflow(sql.FieldEQ(FieldUpdateTime, v))
}

// ProjectID applies equality check predicate on the "project_id" field. It's identical to ProjectIDEQ.
func ProjectID(v object.ID) predicate.Workflow {
	return predicate.Workflow(sql.FieldEQ(FieldProjectID, v))
}

// EnvironmentID applies equality check predicate on the "environment_id" field. It's identical to EnvironmentIDEQ.
func EnvironmentID(v object.ID) predicate.Workflow {
	return predicate.Workflow(sql.FieldEQ(FieldEnvironmentID, v))
}

// DisplayName applies equality check predicate on the "display_name" field. It's identical to DisplayNameEQ.
func DisplayName(v string) predicate.Workflow {
	return predicate.Workflow(sql.FieldEQ(FieldDisplayName, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.Workflow {
	return predicate.Workflow(sql.FieldEQ(FieldType, v))
}

// Parallelism applies equality check predicate on the "parallelism" field. It's identical to ParallelismEQ.
func Parallelism(v int) predicate.Workflow {
	return predicate.Workflow(sql.FieldEQ(FieldParallelism, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Workflow {
	return predicate.Workflow(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Workflow {
	return predicate.Workflow(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Workflow {
	return predicate.Workflow(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Workflow {
	return predicate.Workflow(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Workflow {
	return predicate.Workflow(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Workflow {
	return predicate.Workflow(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Workflow {
	return predicate.Workflow(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Workflow {
	return predicate.Workflow(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Workflow {
	return predicate.Workflow(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Workflow {
	return predicate.Workflow(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Workflow {
	return predicate.Workflow(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Workflow {
	return predicate.Workflow(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Workflow {
	return predicate.Workflow(sql.FieldContainsFold(FieldName, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Workflow {
	return predicate.Workflow(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Workflow {
	return predicate.Workflow(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Workflow {
	return predicate.Workflow(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Workflow {
	return predicate.Workflow(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Workflow {
	return predicate.Workflow(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Workflow {
	return predicate.Workflow(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Workflow {
	return predicate.Workflow(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Workflow {
	return predicate.Workflow(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Workflow {
	return predicate.Workflow(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Workflow {
	return predicate.Workflow(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Workflow {
	return predicate.Workflow(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Workflow {
	return predicate.Workflow(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Workflow {
	return predicate.Workflow(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Workflow {
	return predicate.Workflow(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Workflow {
	return predicate.Workflow(sql.FieldContainsFold(FieldDescription, v))
}

// LabelsIsNil applies the IsNil predicate on the "labels" field.
func LabelsIsNil() predicate.Workflow {
	return predicate.Workflow(sql.FieldIsNull(FieldLabels))
}

// LabelsNotNil applies the NotNil predicate on the "labels" field.
func LabelsNotNil() predicate.Workflow {
	return predicate.Workflow(sql.FieldNotNull(FieldLabels))
}

// AnnotationsIsNil applies the IsNil predicate on the "annotations" field.
func AnnotationsIsNil() predicate.Workflow {
	return predicate.Workflow(sql.FieldIsNull(FieldAnnotations))
}

// AnnotationsNotNil applies the NotNil predicate on the "annotations" field.
func AnnotationsNotNil() predicate.Workflow {
	return predicate.Workflow(sql.FieldNotNull(FieldAnnotations))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Workflow {
	return predicate.Workflow(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Workflow {
	return predicate.Workflow(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Workflow {
	return predicate.Workflow(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Workflow {
	return predicate.Workflow(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Workflow {
	return predicate.Workflow(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Workflow {
	return predicate.Workflow(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Workflow {
	return predicate.Workflow(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Workflow {
	return predicate.Workflow(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Workflow {
	return predicate.Workflow(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Workflow {
	return predicate.Workflow(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Workflow {
	return predicate.Workflow(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.Workflow {
	return predicate.Workflow(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.Workflow {
	return predicate.Workflow(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Workflow {
	return predicate.Workflow(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Workflow {
	return predicate.Workflow(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Workflow {
	return predicate.Workflow(sql.FieldLTE(FieldUpdateTime, v))
}

// ProjectIDEQ applies the EQ predicate on the "project_id" field.
func ProjectIDEQ(v object.ID) predicate.Workflow {
	return predicate.Workflow(sql.FieldEQ(FieldProjectID, v))
}

// ProjectIDNEQ applies the NEQ predicate on the "project_id" field.
func ProjectIDNEQ(v object.ID) predicate.Workflow {
	return predicate.Workflow(sql.FieldNEQ(FieldProjectID, v))
}

// ProjectIDIn applies the In predicate on the "project_id" field.
func ProjectIDIn(vs ...object.ID) predicate.Workflow {
	return predicate.Workflow(sql.FieldIn(FieldProjectID, vs...))
}

// ProjectIDNotIn applies the NotIn predicate on the "project_id" field.
func ProjectIDNotIn(vs ...object.ID) predicate.Workflow {
	return predicate.Workflow(sql.FieldNotIn(FieldProjectID, vs...))
}

// ProjectIDGT applies the GT predicate on the "project_id" field.
func ProjectIDGT(v object.ID) predicate.Workflow {
	return predicate.Workflow(sql.FieldGT(FieldProjectID, v))
}

// ProjectIDGTE applies the GTE predicate on the "project_id" field.
func ProjectIDGTE(v object.ID) predicate.Workflow {
	return predicate.Workflow(sql.FieldGTE(FieldProjectID, v))
}

// ProjectIDLT applies the LT predicate on the "project_id" field.
func ProjectIDLT(v object.ID) predicate.Workflow {
	return predicate.Workflow(sql.FieldLT(FieldProjectID, v))
}

// ProjectIDLTE applies the LTE predicate on the "project_id" field.
func ProjectIDLTE(v object.ID) predicate.Workflow {
	return predicate.Workflow(sql.FieldLTE(FieldProjectID, v))
}

// ProjectIDContains applies the Contains predicate on the "project_id" field.
func ProjectIDContains(v object.ID) predicate.Workflow {
	vc := string(v)
	return predicate.Workflow(sql.FieldContains(FieldProjectID, vc))
}

// ProjectIDHasPrefix applies the HasPrefix predicate on the "project_id" field.
func ProjectIDHasPrefix(v object.ID) predicate.Workflow {
	vc := string(v)
	return predicate.Workflow(sql.FieldHasPrefix(FieldProjectID, vc))
}

// ProjectIDHasSuffix applies the HasSuffix predicate on the "project_id" field.
func ProjectIDHasSuffix(v object.ID) predicate.Workflow {
	vc := string(v)
	return predicate.Workflow(sql.FieldHasSuffix(FieldProjectID, vc))
}

// ProjectIDEqualFold applies the EqualFold predicate on the "project_id" field.
func ProjectIDEqualFold(v object.ID) predicate.Workflow {
	vc := string(v)
	return predicate.Workflow(sql.FieldEqualFold(FieldProjectID, vc))
}

// ProjectIDContainsFold applies the ContainsFold predicate on the "project_id" field.
func ProjectIDContainsFold(v object.ID) predicate.Workflow {
	vc := string(v)
	return predicate.Workflow(sql.FieldContainsFold(FieldProjectID, vc))
}

// EnvironmentIDEQ applies the EQ predicate on the "environment_id" field.
func EnvironmentIDEQ(v object.ID) predicate.Workflow {
	return predicate.Workflow(sql.FieldEQ(FieldEnvironmentID, v))
}

// EnvironmentIDNEQ applies the NEQ predicate on the "environment_id" field.
func EnvironmentIDNEQ(v object.ID) predicate.Workflow {
	return predicate.Workflow(sql.FieldNEQ(FieldEnvironmentID, v))
}

// EnvironmentIDIn applies the In predicate on the "environment_id" field.
func EnvironmentIDIn(vs ...object.ID) predicate.Workflow {
	return predicate.Workflow(sql.FieldIn(FieldEnvironmentID, vs...))
}

// EnvironmentIDNotIn applies the NotIn predicate on the "environment_id" field.
func EnvironmentIDNotIn(vs ...object.ID) predicate.Workflow {
	return predicate.Workflow(sql.FieldNotIn(FieldEnvironmentID, vs...))
}

// EnvironmentIDGT applies the GT predicate on the "environment_id" field.
func EnvironmentIDGT(v object.ID) predicate.Workflow {
	return predicate.Workflow(sql.FieldGT(FieldEnvironmentID, v))
}

// EnvironmentIDGTE applies the GTE predicate on the "environment_id" field.
func EnvironmentIDGTE(v object.ID) predicate.Workflow {
	return predicate.Workflow(sql.FieldGTE(FieldEnvironmentID, v))
}

// EnvironmentIDLT applies the LT predicate on the "environment_id" field.
func EnvironmentIDLT(v object.ID) predicate.Workflow {
	return predicate.Workflow(sql.FieldLT(FieldEnvironmentID, v))
}

// EnvironmentIDLTE applies the LTE predicate on the "environment_id" field.
func EnvironmentIDLTE(v object.ID) predicate.Workflow {
	return predicate.Workflow(sql.FieldLTE(FieldEnvironmentID, v))
}

// EnvironmentIDContains applies the Contains predicate on the "environment_id" field.
func EnvironmentIDContains(v object.ID) predicate.Workflow {
	vc := string(v)
	return predicate.Workflow(sql.FieldContains(FieldEnvironmentID, vc))
}

// EnvironmentIDHasPrefix applies the HasPrefix predicate on the "environment_id" field.
func EnvironmentIDHasPrefix(v object.ID) predicate.Workflow {
	vc := string(v)
	return predicate.Workflow(sql.FieldHasPrefix(FieldEnvironmentID, vc))
}

// EnvironmentIDHasSuffix applies the HasSuffix predicate on the "environment_id" field.
func EnvironmentIDHasSuffix(v object.ID) predicate.Workflow {
	vc := string(v)
	return predicate.Workflow(sql.FieldHasSuffix(FieldEnvironmentID, vc))
}

// EnvironmentIDIsNil applies the IsNil predicate on the "environment_id" field.
func EnvironmentIDIsNil() predicate.Workflow {
	return predicate.Workflow(sql.FieldIsNull(FieldEnvironmentID))
}

// EnvironmentIDNotNil applies the NotNil predicate on the "environment_id" field.
func EnvironmentIDNotNil() predicate.Workflow {
	return predicate.Workflow(sql.FieldNotNull(FieldEnvironmentID))
}

// EnvironmentIDEqualFold applies the EqualFold predicate on the "environment_id" field.
func EnvironmentIDEqualFold(v object.ID) predicate.Workflow {
	vc := string(v)
	return predicate.Workflow(sql.FieldEqualFold(FieldEnvironmentID, vc))
}

// EnvironmentIDContainsFold applies the ContainsFold predicate on the "environment_id" field.
func EnvironmentIDContainsFold(v object.ID) predicate.Workflow {
	vc := string(v)
	return predicate.Workflow(sql.FieldContainsFold(FieldEnvironmentID, vc))
}

// DisplayNameEQ applies the EQ predicate on the "display_name" field.
func DisplayNameEQ(v string) predicate.Workflow {
	return predicate.Workflow(sql.FieldEQ(FieldDisplayName, v))
}

// DisplayNameNEQ applies the NEQ predicate on the "display_name" field.
func DisplayNameNEQ(v string) predicate.Workflow {
	return predicate.Workflow(sql.FieldNEQ(FieldDisplayName, v))
}

// DisplayNameIn applies the In predicate on the "display_name" field.
func DisplayNameIn(vs ...string) predicate.Workflow {
	return predicate.Workflow(sql.FieldIn(FieldDisplayName, vs...))
}

// DisplayNameNotIn applies the NotIn predicate on the "display_name" field.
func DisplayNameNotIn(vs ...string) predicate.Workflow {
	return predicate.Workflow(sql.FieldNotIn(FieldDisplayName, vs...))
}

// DisplayNameGT applies the GT predicate on the "display_name" field.
func DisplayNameGT(v string) predicate.Workflow {
	return predicate.Workflow(sql.FieldGT(FieldDisplayName, v))
}

// DisplayNameGTE applies the GTE predicate on the "display_name" field.
func DisplayNameGTE(v string) predicate.Workflow {
	return predicate.Workflow(sql.FieldGTE(FieldDisplayName, v))
}

// DisplayNameLT applies the LT predicate on the "display_name" field.
func DisplayNameLT(v string) predicate.Workflow {
	return predicate.Workflow(sql.FieldLT(FieldDisplayName, v))
}

// DisplayNameLTE applies the LTE predicate on the "display_name" field.
func DisplayNameLTE(v string) predicate.Workflow {
	return predicate.Workflow(sql.FieldLTE(FieldDisplayName, v))
}

// DisplayNameContains applies the Contains predicate on the "display_name" field.
func DisplayNameContains(v string) predicate.Workflow {
	return predicate.Workflow(sql.FieldContains(FieldDisplayName, v))
}

// DisplayNameHasPrefix applies the HasPrefix predicate on the "display_name" field.
func DisplayNameHasPrefix(v string) predicate.Workflow {
	return predicate.Workflow(sql.FieldHasPrefix(FieldDisplayName, v))
}

// DisplayNameHasSuffix applies the HasSuffix predicate on the "display_name" field.
func DisplayNameHasSuffix(v string) predicate.Workflow {
	return predicate.Workflow(sql.FieldHasSuffix(FieldDisplayName, v))
}

// DisplayNameEqualFold applies the EqualFold predicate on the "display_name" field.
func DisplayNameEqualFold(v string) predicate.Workflow {
	return predicate.Workflow(sql.FieldEqualFold(FieldDisplayName, v))
}

// DisplayNameContainsFold applies the ContainsFold predicate on the "display_name" field.
func DisplayNameContainsFold(v string) predicate.Workflow {
	return predicate.Workflow(sql.FieldContainsFold(FieldDisplayName, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.Workflow {
	return predicate.Workflow(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.Workflow {
	return predicate.Workflow(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.Workflow {
	return predicate.Workflow(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.Workflow {
	return predicate.Workflow(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.Workflow {
	return predicate.Workflow(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.Workflow {
	return predicate.Workflow(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.Workflow {
	return predicate.Workflow(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.Workflow {
	return predicate.Workflow(sql.FieldLTE(FieldType, v))
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.Workflow {
	return predicate.Workflow(sql.FieldContains(FieldType, v))
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.Workflow {
	return predicate.Workflow(sql.FieldHasPrefix(FieldType, v))
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.Workflow {
	return predicate.Workflow(sql.FieldHasSuffix(FieldType, v))
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.Workflow {
	return predicate.Workflow(sql.FieldEqualFold(FieldType, v))
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.Workflow {
	return predicate.Workflow(sql.FieldContainsFold(FieldType, v))
}

// ParallelismEQ applies the EQ predicate on the "parallelism" field.
func ParallelismEQ(v int) predicate.Workflow {
	return predicate.Workflow(sql.FieldEQ(FieldParallelism, v))
}

// ParallelismNEQ applies the NEQ predicate on the "parallelism" field.
func ParallelismNEQ(v int) predicate.Workflow {
	return predicate.Workflow(sql.FieldNEQ(FieldParallelism, v))
}

// ParallelismIn applies the In predicate on the "parallelism" field.
func ParallelismIn(vs ...int) predicate.Workflow {
	return predicate.Workflow(sql.FieldIn(FieldParallelism, vs...))
}

// ParallelismNotIn applies the NotIn predicate on the "parallelism" field.
func ParallelismNotIn(vs ...int) predicate.Workflow {
	return predicate.Workflow(sql.FieldNotIn(FieldParallelism, vs...))
}

// ParallelismGT applies the GT predicate on the "parallelism" field.
func ParallelismGT(v int) predicate.Workflow {
	return predicate.Workflow(sql.FieldGT(FieldParallelism, v))
}

// ParallelismGTE applies the GTE predicate on the "parallelism" field.
func ParallelismGTE(v int) predicate.Workflow {
	return predicate.Workflow(sql.FieldGTE(FieldParallelism, v))
}

// ParallelismLT applies the LT predicate on the "parallelism" field.
func ParallelismLT(v int) predicate.Workflow {
	return predicate.Workflow(sql.FieldLT(FieldParallelism, v))
}

// ParallelismLTE applies the LTE predicate on the "parallelism" field.
func ParallelismLTE(v int) predicate.Workflow {
	return predicate.Workflow(sql.FieldLTE(FieldParallelism, v))
}

// HasProject applies the HasEdge predicate on the "project" edge.
func HasProject() predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProjectTable, ProjectColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Project
		step.Edge.Schema = schemaConfig.Workflow
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProjectWith applies the HasEdge predicate on the "project" edge with a given conditions (other predicates).
func HasProjectWith(preds ...predicate.Project) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		step := newProjectStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Project
		step.Edge.Schema = schemaConfig.Workflow
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasStages applies the HasEdge predicate on the "stages" edge.
func HasStages() predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, StagesTable, StagesColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.WorkflowStage
		step.Edge.Schema = schemaConfig.WorkflowStage
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStagesWith applies the HasEdge predicate on the "stages" edge with a given conditions (other predicates).
func HasStagesWith(preds ...predicate.WorkflowStage) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		step := newStagesStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.WorkflowStage
		step.Edge.Schema = schemaConfig.WorkflowStage
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasExecutions applies the HasEdge predicate on the "executions" edge.
func HasExecutions() predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ExecutionsTable, ExecutionsColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.WorkflowExecution
		step.Edge.Schema = schemaConfig.WorkflowExecution
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasExecutionsWith applies the HasEdge predicate on the "executions" edge with a given conditions (other predicates).
func HasExecutionsWith(preds ...predicate.WorkflowExecution) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		step := newExecutionsStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.WorkflowExecution
		step.Edge.Schema = schemaConfig.WorkflowExecution
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Workflow) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Workflow) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
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
func Not(p predicate.Workflow) predicate.Workflow {
	return predicate.Workflow(func(s *sql.Selector) {
		p(s.Not())
	})
}
