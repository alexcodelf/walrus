// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus". DO NOT EDIT.

package workflowstepexecution

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"golang.org/x/exp/slices"
)

const (
	// Label holds the string label denoting the workflowstepexecution type in the database.
	Label = "workflow_step_execution"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldLabels holds the string denoting the labels field in the database.
	FieldLabels = "labels"
	// FieldAnnotations holds the string denoting the annotations field in the database.
	FieldAnnotations = "annotations"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldWorkflowStepID holds the string denoting the workflow_step_id field in the database.
	FieldWorkflowStepID = "workflow_step_id"
	// FieldWorkflowExecutionID holds the string denoting the workflow_execution_id field in the database.
	FieldWorkflowExecutionID = "workflow_execution_id"
	// FieldWorkflowStageExecutionID holds the string denoting the workflow_stage_execution_id field in the database.
	FieldWorkflowStageExecutionID = "workflow_stage_execution_id"
	// FieldProjectID holds the string denoting the project_id field in the database.
	FieldProjectID = "project_id"
	// FieldWorkflowID holds the string denoting the workflow_id field in the database.
	FieldWorkflowID = "workflow_id"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldSpec holds the string denoting the spec field in the database.
	FieldSpec = "spec"
	// FieldTimes holds the string denoting the times field in the database.
	FieldTimes = "times"
	// FieldExecuteTime holds the string denoting the execute_time field in the database.
	FieldExecuteTime = "execute_time"
	// FieldDuration holds the string denoting the duration field in the database.
	FieldDuration = "duration"
	// FieldRetryStrategy holds the string denoting the retrystrategy field in the database.
	FieldRetryStrategy = "retry_strategy"
	// FieldTimeout holds the string denoting the timeout field in the database.
	FieldTimeout = "timeout"
	// FieldOrder holds the string denoting the order field in the database.
	FieldOrder = "order"
	// FieldRecord holds the string denoting the record field in the database.
	FieldRecord = "record"
	// EdgeProject holds the string denoting the project edge name in mutations.
	EdgeProject = "project"
	// EdgeStageExecution holds the string denoting the stage_execution edge name in mutations.
	EdgeStageExecution = "stage_execution"
	// Table holds the table name of the workflowstepexecution in the database.
	Table = "workflow_step_executions"
	// ProjectTable is the table that holds the project relation/edge.
	ProjectTable = "workflow_step_executions"
	// ProjectInverseTable is the table name for the Project entity.
	// It exists in this package in order to avoid circular dependency with the "project" package.
	ProjectInverseTable = "projects"
	// ProjectColumn is the table column denoting the project relation/edge.
	ProjectColumn = "project_id"
	// StageExecutionTable is the table that holds the stage_execution relation/edge.
	StageExecutionTable = "workflow_step_executions"
	// StageExecutionInverseTable is the table name for the WorkflowStageExecution entity.
	// It exists in this package in order to avoid circular dependency with the "workflowstageexecution" package.
	StageExecutionInverseTable = "workflow_stage_executions"
	// StageExecutionColumn is the table column denoting the stage_execution relation/edge.
	StageExecutionColumn = "workflow_stage_execution_id"
)

// Columns holds all SQL columns for workflowstepexecution fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDescription,
	FieldLabels,
	FieldAnnotations,
	FieldCreateTime,
	FieldUpdateTime,
	FieldStatus,
	FieldWorkflowStepID,
	FieldWorkflowExecutionID,
	FieldWorkflowStageExecutionID,
	FieldProjectID,
	FieldWorkflowID,
	FieldType,
	FieldSpec,
	FieldTimes,
	FieldExecuteTime,
	FieldDuration,
	FieldRetryStrategy,
	FieldTimeout,
	FieldOrder,
	FieldRecord,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/seal-io/walrus/pkg/dao/model/runtime"
var (
	Hooks        [1]ent.Hook
	Interceptors [1]ent.Interceptor
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultLabels holds the default value on creation for the "labels" field.
	DefaultLabels map[string]string
	// DefaultAnnotations holds the default value on creation for the "annotations" field.
	DefaultAnnotations map[string]string
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// ProjectIDValidator is a validator for the "project_id" field. It is called by the builders before save.
	ProjectIDValidator func(string) error
	// WorkflowIDValidator is a validator for the "workflow_id" field. It is called by the builders before save.
	WorkflowIDValidator func(string) error
	// TypeValidator is a validator for the "type" field. It is called by the builders before save.
	TypeValidator func(string) error
	// DefaultTimes holds the default value on creation for the "times" field.
	DefaultTimes int
	// TimesValidator is a validator for the "times" field. It is called by the builders before save.
	TimesValidator func(int) error
	// DefaultDuration holds the default value on creation for the "duration" field.
	DefaultDuration int
	// DurationValidator is a validator for the "duration" field. It is called by the builders before save.
	DurationValidator func(int) error
	// DefaultTimeout holds the default value on creation for the "timeout" field.
	DefaultTimeout int
	// TimeoutValidator is a validator for the "timeout" field. It is called by the builders before save.
	TimeoutValidator func(int) error
	// DefaultOrder holds the default value on creation for the "order" field.
	DefaultOrder int
	// OrderValidator is a validator for the "order" field. It is called by the builders before save.
	OrderValidator func(int) error
	// DefaultRecord holds the default value on creation for the "record" field.
	DefaultRecord string
)

// OrderOption defines the ordering options for the WorkflowStepExecution queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByUpdateTime orders the results by the update_time field.
func ByUpdateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateTime, opts...).ToFunc()
}

// ByWorkflowStepID orders the results by the workflow_step_id field.
func ByWorkflowStepID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWorkflowStepID, opts...).ToFunc()
}

// ByWorkflowExecutionID orders the results by the workflow_execution_id field.
func ByWorkflowExecutionID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWorkflowExecutionID, opts...).ToFunc()
}

// ByWorkflowStageExecutionID orders the results by the workflow_stage_execution_id field.
func ByWorkflowStageExecutionID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWorkflowStageExecutionID, opts...).ToFunc()
}

// ByProjectID orders the results by the project_id field.
func ByProjectID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProjectID, opts...).ToFunc()
}

// ByWorkflowID orders the results by the workflow_id field.
func ByWorkflowID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWorkflowID, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByTimes orders the results by the times field.
func ByTimes(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTimes, opts...).ToFunc()
}

// ByExecuteTime orders the results by the execute_time field.
func ByExecuteTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExecuteTime, opts...).ToFunc()
}

// ByDuration orders the results by the duration field.
func ByDuration(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDuration, opts...).ToFunc()
}

// ByTimeout orders the results by the timeout field.
func ByTimeout(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTimeout, opts...).ToFunc()
}

// ByOrder orders the results by the order field.
func ByOrder(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrder, opts...).ToFunc()
}

// ByRecord orders the results by the record field.
func ByRecord(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRecord, opts...).ToFunc()
}

// ByProjectField orders the results by project field.
func ByProjectField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProjectStep(), sql.OrderByField(field, opts...))
	}
}

// ByStageExecutionField orders the results by stage_execution field.
func ByStageExecutionField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newStageExecutionStep(), sql.OrderByField(field, opts...))
	}
}
func newProjectStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProjectInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ProjectTable, ProjectColumn),
	)
}
func newStageExecutionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(StageExecutionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, StageExecutionTable, StageExecutionColumn),
	)
}

// WithoutFields returns the fields ignored the given list.
func WithoutFields(ignores ...string) []string {
	if len(ignores) == 0 {
		return slices.Clone(Columns)
	}

	var s = make(map[string]bool, len(ignores))
	for i := range ignores {
		s[ignores[i]] = true
	}

	var r = make([]string, 0, len(Columns)-len(s))
	for i := range Columns {
		if s[Columns[i]] {
			continue
		}
		r = append(r, Columns[i])
	}
	return r
}
