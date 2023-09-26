// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus". DO NOT EDIT.

package servicerevision

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"golang.org/x/exp/slices"

	"github.com/seal-io/walrus/pkg/dao/types"
	"github.com/seal-io/walrus/pkg/dao/types/crypto"
)

const (
	// Label holds the string label denoting the servicerevision type in the database.
	Label = "service_revision"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldProjectID holds the string denoting the project_id field in the database.
	FieldProjectID = "project_id"
	// FieldEnvironmentID holds the string denoting the environment_id field in the database.
	FieldEnvironmentID = "environment_id"
	// FieldServiceID holds the string denoting the service_id field in the database.
	FieldServiceID = "service_id"
	// FieldTemplateName holds the string denoting the template_name field in the database.
	FieldTemplateName = "template_name"
	// FieldTemplateVersion holds the string denoting the template_version field in the database.
	FieldTemplateVersion = "template_version"
	// FieldAttributes holds the string denoting the attributes field in the database.
	FieldAttributes = "attributes"
	// FieldVariables holds the string denoting the variables field in the database.
	FieldVariables = "variables"
	// FieldInputPlanConfigs holds the string denoting the input_plan_configs field in the database.
	FieldInputPlanConfigs = "input_plan_configs"
	// FieldOutput holds the string denoting the output field in the database.
	FieldOutput = "output"
	// FieldDeployerType holds the string denoting the deployer_type field in the database.
	FieldDeployerType = "deployer_type"
	// FieldDuration holds the string denoting the duration field in the database.
	FieldDuration = "duration"
	// FieldPreviousRequiredProviders holds the string denoting the previous_required_providers field in the database.
	FieldPreviousRequiredProviders = "previous_required_providers"
	// FieldRecord holds the string denoting the record field in the database.
	FieldRecord = "record"
	// EdgeProject holds the string denoting the project edge name in mutations.
	EdgeProject = "project"
	// EdgeEnvironment holds the string denoting the environment edge name in mutations.
	EdgeEnvironment = "environment"
	// EdgeService holds the string denoting the service edge name in mutations.
	EdgeService = "service"
	// Table holds the table name of the servicerevision in the database.
	Table = "service_revisions"
	// ProjectTable is the table that holds the project relation/edge.
	ProjectTable = "service_revisions"
	// ProjectInverseTable is the table name for the Project entity.
	// It exists in this package in order to avoid circular dependency with the "project" package.
	ProjectInverseTable = "projects"
	// ProjectColumn is the table column denoting the project relation/edge.
	ProjectColumn = "project_id"
	// EnvironmentTable is the table that holds the environment relation/edge.
	EnvironmentTable = "service_revisions"
	// EnvironmentInverseTable is the table name for the Environment entity.
	// It exists in this package in order to avoid circular dependency with the "environment" package.
	EnvironmentInverseTable = "environments"
	// EnvironmentColumn is the table column denoting the environment relation/edge.
	EnvironmentColumn = "environment_id"
	// ServiceTable is the table that holds the service relation/edge.
	ServiceTable = "service_revisions"
	// ServiceInverseTable is the table name for the Service entity.
	// It exists in this package in order to avoid circular dependency with the "service" package.
	ServiceInverseTable = "services"
	// ServiceColumn is the table column denoting the service relation/edge.
	ServiceColumn = "service_id"
)

// Columns holds all SQL columns for servicerevision fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldStatus,
	FieldProjectID,
	FieldEnvironmentID,
	FieldServiceID,
	FieldTemplateName,
	FieldTemplateVersion,
	FieldAttributes,
	FieldVariables,
	FieldInputPlanConfigs,
	FieldOutput,
	FieldDeployerType,
	FieldDuration,
	FieldPreviousRequiredProviders,
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
	Hooks        [2]ent.Hook
	Interceptors [1]ent.Interceptor
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// ProjectIDValidator is a validator for the "project_id" field. It is called by the builders before save.
	ProjectIDValidator func(string) error
	// EnvironmentIDValidator is a validator for the "environment_id" field. It is called by the builders before save.
	EnvironmentIDValidator func(string) error
	// ServiceIDValidator is a validator for the "service_id" field. It is called by the builders before save.
	ServiceIDValidator func(string) error
	// TemplateNameValidator is a validator for the "template_name" field. It is called by the builders before save.
	TemplateNameValidator func(string) error
	// TemplateVersionValidator is a validator for the "template_version" field. It is called by the builders before save.
	TemplateVersionValidator func(string) error
	// DefaultVariables holds the default value on creation for the "variables" field.
	DefaultVariables crypto.Map[string, string]
	// DefaultDeployerType holds the default value on creation for the "deployer_type" field.
	DefaultDeployerType string
	// DefaultDuration holds the default value on creation for the "duration" field.
	DefaultDuration int
	// DefaultPreviousRequiredProviders holds the default value on creation for the "previous_required_providers" field.
	DefaultPreviousRequiredProviders []types.ProviderRequirement
)

// OrderOption defines the ordering options for the ServiceRevision queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByProjectID orders the results by the project_id field.
func ByProjectID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProjectID, opts...).ToFunc()
}

// ByEnvironmentID orders the results by the environment_id field.
func ByEnvironmentID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEnvironmentID, opts...).ToFunc()
}

// ByServiceID orders the results by the service_id field.
func ByServiceID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldServiceID, opts...).ToFunc()
}

// ByTemplateName orders the results by the template_name field.
func ByTemplateName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTemplateName, opts...).ToFunc()
}

// ByTemplateVersion orders the results by the template_version field.
func ByTemplateVersion(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTemplateVersion, opts...).ToFunc()
}

// ByAttributes orders the results by the attributes field.
func ByAttributes(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAttributes, opts...).ToFunc()
}

// ByVariables orders the results by the variables field.
func ByVariables(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVariables, opts...).ToFunc()
}

// ByOutput orders the results by the output field.
func ByOutput(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOutput, opts...).ToFunc()
}

// ByDeployerType orders the results by the deployer_type field.
func ByDeployerType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeployerType, opts...).ToFunc()
}

// ByDuration orders the results by the duration field.
func ByDuration(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDuration, opts...).ToFunc()
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

// ByEnvironmentField orders the results by environment field.
func ByEnvironmentField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newEnvironmentStep(), sql.OrderByField(field, opts...))
	}
}

// ByServiceField orders the results by service field.
func ByServiceField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newServiceStep(), sql.OrderByField(field, opts...))
	}
}
func newProjectStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProjectInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ProjectTable, ProjectColumn),
	)
}
func newEnvironmentStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(EnvironmentInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, EnvironmentTable, EnvironmentColumn),
	)
}
func newServiceStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ServiceInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ServiceTable, ServiceColumn),
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
