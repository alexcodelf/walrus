// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus". DO NOT EDIT.

package internal

import "context"

// SchemaConfig represents alternative schema names for all tables
// that can be passed at runtime.
type SchemaConfig struct {
	Catalog                          string // Catalog table.
	Connector                        string // Connector table.
	CostReport                       string // CostReport table.
	DistributeLock                   string // DistributeLock table.
	Environment                      string // Environment table.
	EnvironmentConnectorRelationship string // EnvironmentConnectorRelationship table.
	Perspective                      string // Perspective table.
	Project                          string // Project table.
	Role                             string // Role table.
	Service                          string // Service table.
	ServiceRelationship              string // ServiceRelationship table.
	ServiceResource                  string // ServiceResource table.
	ServiceResourceRelationship      string // ServiceResourceRelationship table.
	ServiceRevision                  string // ServiceRevision table.
	Setting                          string // Setting table.
	Subject                          string // Subject table.
	SubjectRoleRelationship          string // SubjectRoleRelationship table.
	Template                         string // Template table.
	TemplateVersion                  string // TemplateVersion table.
	Token                            string // Token table.
	Variable                         string // Variable table.
	Workflow                         string // Workflow table.
	WorkflowExecution                string // WorkflowExecution table.
	WorkflowStage                    string // WorkflowStage table.
	WorkflowStageExecution           string // WorkflowStageExecution table.
	WorkflowStep                     string // WorkflowStep table.
	WorkflowStepExecution            string // WorkflowStepExecution table.
	WorkflowStepTemplate             string // WorkflowStepTemplate table.
}

type schemaCtxKey struct{}

// SchemaConfigFromContext returns a SchemaConfig stored inside a context, or empty if there isn't one.
func SchemaConfigFromContext(ctx context.Context) SchemaConfig {
	config, _ := ctx.Value(schemaCtxKey{}).(SchemaConfig)
	return config
}

// NewSchemaConfigContext returns a new context with the given SchemaConfig attached.
func NewSchemaConfigContext(parent context.Context, config SchemaConfig) context.Context {
	return context.WithValue(parent, schemaCtxKey{}, config)
}
