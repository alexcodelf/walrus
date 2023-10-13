// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus". DO NOT EDIT.

package project

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"golang.org/x/exp/slices"
)

const (
	// Label holds the string label denoting the project type in the database.
	Label = "project"
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
	// EdgeEnvironments holds the string denoting the environments edge name in mutations.
	EdgeEnvironments = "environments"
	// EdgeConnectors holds the string denoting the connectors edge name in mutations.
	EdgeConnectors = "connectors"
	// EdgeSubjectRoles holds the string denoting the subject_roles edge name in mutations.
	EdgeSubjectRoles = "subject_roles"
	// EdgeServices holds the string denoting the services edge name in mutations.
	EdgeServices = "services"
	// EdgeServiceResources holds the string denoting the service_resources edge name in mutations.
	EdgeServiceResources = "service_resources"
	// EdgeServiceRevisions holds the string denoting the service_revisions edge name in mutations.
	EdgeServiceRevisions = "service_revisions"
	// EdgeVariables holds the string denoting the variables edge name in mutations.
	EdgeVariables = "variables"
	// EdgeWorkflows holds the string denoting the workflows edge name in mutations.
	EdgeWorkflows = "workflows"
	// EdgeWorkflowStages holds the string denoting the workflow_stages edge name in mutations.
	EdgeWorkflowStages = "workflow_stages"
	// EdgeWorkflowSteps holds the string denoting the workflow_steps edge name in mutations.
	EdgeWorkflowSteps = "workflow_steps"
	// EdgeWorkflowExecutions holds the string denoting the workflow_executions edge name in mutations.
	EdgeWorkflowExecutions = "workflow_executions"
	// EdgeWorkflowStageExecutions holds the string denoting the workflow_stage_executions edge name in mutations.
	EdgeWorkflowStageExecutions = "workflow_stage_executions"
	// EdgeWorkflowStepExecutions holds the string denoting the workflow_step_executions edge name in mutations.
	EdgeWorkflowStepExecutions = "workflow_step_executions"
	// Table holds the table name of the project in the database.
	Table = "projects"
	// EnvironmentsTable is the table that holds the environments relation/edge.
	EnvironmentsTable = "environments"
	// EnvironmentsInverseTable is the table name for the Environment entity.
	// It exists in this package in order to avoid circular dependency with the "environment" package.
	EnvironmentsInverseTable = "environments"
	// EnvironmentsColumn is the table column denoting the environments relation/edge.
	EnvironmentsColumn = "project_id"
	// ConnectorsTable is the table that holds the connectors relation/edge.
	ConnectorsTable = "connectors"
	// ConnectorsInverseTable is the table name for the Connector entity.
	// It exists in this package in order to avoid circular dependency with the "connector" package.
	ConnectorsInverseTable = "connectors"
	// ConnectorsColumn is the table column denoting the connectors relation/edge.
	ConnectorsColumn = "project_id"
	// SubjectRolesTable is the table that holds the subject_roles relation/edge.
	SubjectRolesTable = "subject_role_relationships"
	// SubjectRolesInverseTable is the table name for the SubjectRoleRelationship entity.
	// It exists in this package in order to avoid circular dependency with the "subjectrolerelationship" package.
	SubjectRolesInverseTable = "subject_role_relationships"
	// SubjectRolesColumn is the table column denoting the subject_roles relation/edge.
	SubjectRolesColumn = "project_id"
	// ServicesTable is the table that holds the services relation/edge.
	ServicesTable = "services"
	// ServicesInverseTable is the table name for the Service entity.
	// It exists in this package in order to avoid circular dependency with the "service" package.
	ServicesInverseTable = "services"
	// ServicesColumn is the table column denoting the services relation/edge.
	ServicesColumn = "project_id"
	// ServiceResourcesTable is the table that holds the service_resources relation/edge.
	ServiceResourcesTable = "service_resources"
	// ServiceResourcesInverseTable is the table name for the ServiceResource entity.
	// It exists in this package in order to avoid circular dependency with the "serviceresource" package.
	ServiceResourcesInverseTable = "service_resources"
	// ServiceResourcesColumn is the table column denoting the service_resources relation/edge.
	ServiceResourcesColumn = "project_id"
	// ServiceRevisionsTable is the table that holds the service_revisions relation/edge.
	ServiceRevisionsTable = "service_revisions"
	// ServiceRevisionsInverseTable is the table name for the ServiceRevision entity.
	// It exists in this package in order to avoid circular dependency with the "servicerevision" package.
	ServiceRevisionsInverseTable = "service_revisions"
	// ServiceRevisionsColumn is the table column denoting the service_revisions relation/edge.
	ServiceRevisionsColumn = "project_id"
	// VariablesTable is the table that holds the variables relation/edge.
	VariablesTable = "variables"
	// VariablesInverseTable is the table name for the Variable entity.
	// It exists in this package in order to avoid circular dependency with the "variable" package.
	VariablesInverseTable = "variables"
	// VariablesColumn is the table column denoting the variables relation/edge.
	VariablesColumn = "project_id"
	// WorkflowsTable is the table that holds the workflows relation/edge.
	WorkflowsTable = "workflows"
	// WorkflowsInverseTable is the table name for the Workflow entity.
	// It exists in this package in order to avoid circular dependency with the "workflow" package.
	WorkflowsInverseTable = "workflows"
	// WorkflowsColumn is the table column denoting the workflows relation/edge.
	WorkflowsColumn = "project_id"
	// WorkflowStagesTable is the table that holds the workflow_stages relation/edge.
	WorkflowStagesTable = "workflow_stages"
	// WorkflowStagesInverseTable is the table name for the WorkflowStage entity.
	// It exists in this package in order to avoid circular dependency with the "workflowstage" package.
	WorkflowStagesInverseTable = "workflow_stages"
	// WorkflowStagesColumn is the table column denoting the workflow_stages relation/edge.
	WorkflowStagesColumn = "project_id"
	// WorkflowStepsTable is the table that holds the workflow_steps relation/edge.
	WorkflowStepsTable = "workflow_steps"
	// WorkflowStepsInverseTable is the table name for the WorkflowStep entity.
	// It exists in this package in order to avoid circular dependency with the "workflowstep" package.
	WorkflowStepsInverseTable = "workflow_steps"
	// WorkflowStepsColumn is the table column denoting the workflow_steps relation/edge.
	WorkflowStepsColumn = "project_id"
	// WorkflowExecutionsTable is the table that holds the workflow_executions relation/edge.
	WorkflowExecutionsTable = "workflow_executions"
	// WorkflowExecutionsInverseTable is the table name for the WorkflowExecution entity.
	// It exists in this package in order to avoid circular dependency with the "workflowexecution" package.
	WorkflowExecutionsInverseTable = "workflow_executions"
	// WorkflowExecutionsColumn is the table column denoting the workflow_executions relation/edge.
	WorkflowExecutionsColumn = "project_id"
	// WorkflowStageExecutionsTable is the table that holds the workflow_stage_executions relation/edge.
	WorkflowStageExecutionsTable = "workflow_stage_executions"
	// WorkflowStageExecutionsInverseTable is the table name for the WorkflowStageExecution entity.
	// It exists in this package in order to avoid circular dependency with the "workflowstageexecution" package.
	WorkflowStageExecutionsInverseTable = "workflow_stage_executions"
	// WorkflowStageExecutionsColumn is the table column denoting the workflow_stage_executions relation/edge.
	WorkflowStageExecutionsColumn = "project_id"
	// WorkflowStepExecutionsTable is the table that holds the workflow_step_executions relation/edge.
	WorkflowStepExecutionsTable = "workflow_step_executions"
	// WorkflowStepExecutionsInverseTable is the table name for the WorkflowStepExecution entity.
	// It exists in this package in order to avoid circular dependency with the "workflowstepexecution" package.
	WorkflowStepExecutionsInverseTable = "workflow_step_executions"
	// WorkflowStepExecutionsColumn is the table column denoting the workflow_step_executions relation/edge.
	WorkflowStepExecutionsColumn = "project_id"
)

// Columns holds all SQL columns for project fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDescription,
	FieldLabels,
	FieldAnnotations,
	FieldCreateTime,
	FieldUpdateTime,
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
)

// OrderOption defines the ordering options for the Project queries.
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

// ByEnvironmentsCount orders the results by environments count.
func ByEnvironmentsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newEnvironmentsStep(), opts...)
	}
}

// ByEnvironments orders the results by environments terms.
func ByEnvironments(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newEnvironmentsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByConnectorsCount orders the results by connectors count.
func ByConnectorsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newConnectorsStep(), opts...)
	}
}

// ByConnectors orders the results by connectors terms.
func ByConnectors(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newConnectorsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// BySubjectRolesCount orders the results by subject_roles count.
func BySubjectRolesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSubjectRolesStep(), opts...)
	}
}

// BySubjectRoles orders the results by subject_roles terms.
func BySubjectRoles(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSubjectRolesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByServicesCount orders the results by services count.
func ByServicesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newServicesStep(), opts...)
	}
}

// ByServices orders the results by services terms.
func ByServices(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newServicesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByServiceResourcesCount orders the results by service_resources count.
func ByServiceResourcesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newServiceResourcesStep(), opts...)
	}
}

// ByServiceResources orders the results by service_resources terms.
func ByServiceResources(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newServiceResourcesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByServiceRevisionsCount orders the results by service_revisions count.
func ByServiceRevisionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newServiceRevisionsStep(), opts...)
	}
}

// ByServiceRevisions orders the results by service_revisions terms.
func ByServiceRevisions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newServiceRevisionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByVariablesCount orders the results by variables count.
func ByVariablesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newVariablesStep(), opts...)
	}
}

// ByVariables orders the results by variables terms.
func ByVariables(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newVariablesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByWorkflowsCount orders the results by workflows count.
func ByWorkflowsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newWorkflowsStep(), opts...)
	}
}

// ByWorkflows orders the results by workflows terms.
func ByWorkflows(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newWorkflowsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByWorkflowStagesCount orders the results by workflow_stages count.
func ByWorkflowStagesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newWorkflowStagesStep(), opts...)
	}
}

// ByWorkflowStages orders the results by workflow_stages terms.
func ByWorkflowStages(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newWorkflowStagesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByWorkflowStepsCount orders the results by workflow_steps count.
func ByWorkflowStepsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newWorkflowStepsStep(), opts...)
	}
}

// ByWorkflowSteps orders the results by workflow_steps terms.
func ByWorkflowSteps(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newWorkflowStepsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByWorkflowExecutionsCount orders the results by workflow_executions count.
func ByWorkflowExecutionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newWorkflowExecutionsStep(), opts...)
	}
}

// ByWorkflowExecutions orders the results by workflow_executions terms.
func ByWorkflowExecutions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newWorkflowExecutionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByWorkflowStageExecutionsCount orders the results by workflow_stage_executions count.
func ByWorkflowStageExecutionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newWorkflowStageExecutionsStep(), opts...)
	}
}

// ByWorkflowStageExecutions orders the results by workflow_stage_executions terms.
func ByWorkflowStageExecutions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newWorkflowStageExecutionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByWorkflowStepExecutionsCount orders the results by workflow_step_executions count.
func ByWorkflowStepExecutionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newWorkflowStepExecutionsStep(), opts...)
	}
}

// ByWorkflowStepExecutions orders the results by workflow_step_executions terms.
func ByWorkflowStepExecutions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newWorkflowStepExecutionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newEnvironmentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(EnvironmentsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, EnvironmentsTable, EnvironmentsColumn),
	)
}
func newConnectorsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ConnectorsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ConnectorsTable, ConnectorsColumn),
	)
}
func newSubjectRolesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SubjectRolesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, SubjectRolesTable, SubjectRolesColumn),
	)
}
func newServicesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ServicesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ServicesTable, ServicesColumn),
	)
}
func newServiceResourcesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ServiceResourcesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ServiceResourcesTable, ServiceResourcesColumn),
	)
}
func newServiceRevisionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ServiceRevisionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ServiceRevisionsTable, ServiceRevisionsColumn),
	)
}
func newVariablesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(VariablesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, VariablesTable, VariablesColumn),
	)
}
func newWorkflowsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(WorkflowsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, WorkflowsTable, WorkflowsColumn),
	)
}
func newWorkflowStagesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(WorkflowStagesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, WorkflowStagesTable, WorkflowStagesColumn),
	)
}
func newWorkflowStepsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(WorkflowStepsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, WorkflowStepsTable, WorkflowStepsColumn),
	)
}
func newWorkflowExecutionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(WorkflowExecutionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, WorkflowExecutionsTable, WorkflowExecutionsColumn),
	)
}
func newWorkflowStageExecutionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(WorkflowStageExecutionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, WorkflowStageExecutionsTable, WorkflowStageExecutionsColumn),
	)
}
func newWorkflowStepExecutionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(WorkflowStepExecutionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, WorkflowStepExecutionsTable, WorkflowStepExecutionsColumn),
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
