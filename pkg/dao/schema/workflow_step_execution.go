package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"github.com/seal-io/walrus/pkg/dao/entx"
	"github.com/seal-io/walrus/pkg/dao/schema/intercept"
	"github.com/seal-io/walrus/pkg/dao/schema/mixin"
	"github.com/seal-io/walrus/pkg/dao/types/object"
)

type WorkflowStepExecution struct {
	ent.Schema
}

func (WorkflowStepExecution) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Metadata(),
		mixin.Status(),
	}
}

func (WorkflowStepExecution) Fields() []ent.Field {
	return []ent.Field{
		object.IDField("workflow_step_id").
			Comment("ID of the workflow step that this workflow step execution belongs to.").
			Immutable(),
		object.IDField("workflow_execution_id").
			Comment("ID of the workflow execution that this workflow step execution belongs to.").
			Immutable(),
		object.IDField("workflow_stage_execution_id").
			Comment("ID of the workflow stage execution that this workflow step execution belongs to.").
			Immutable(),
		object.IDField("project_id").
			Comment("ID of the project to belong.").
			NotEmpty().
			Immutable(),
		object.IDField("workflow_id").
			Comment("ID of the workflow that this workflow step execution belongs to.").
			NotEmpty().
			Immutable(),
		field.String("type").
			Comment("Type of the workflow step execution.").
			NotEmpty().
			Immutable(),
		field.JSON("spec", map[string]any{}).
			Comment("Spec of the workflow step execution.").
			Optional(),
		field.Int("times").
			Comment("Number of times that this workflow step execution has been executed.").
			NonNegative().
			Default(0),
		field.Int("duration").
			Comment("Duration of the workflow step execution.").
			NonNegative().
			Default(0),
		field.Text("record").
			Comment("Log record of the workflow step execution.").
			Default(""),
		field.Text("input").
			Comment("Input of the workflow step execution." +
				" It's the yaml file that defines the workflow step execution.").
			Default(""),
	}
}

func (WorkflowStepExecution) Edges() []ent.Edge {
	return []ent.Edge{
		// Project 1-* WorkflowStepExecutions.
		edge.From("project", Project.Type).
			Ref("workflow_step_executions").
			Field("project_id").
			Comment("Project to which the workflow step execution belongs.").
			Unique().
			Required().
			Immutable().
			Annotations(
				entx.ValidateContext(intercept.WithProjectInterceptor)),
		// WorkflowStep 1-* WorkflowStepExecutions.
		edge.From("workflow_step", WorkflowStep.Type).
			Ref("executions").
			Field("workflow_step_id").
			Comment("Workflow step that this workflow step execution belongs to.").
			Required().
			Unique().
			Immutable().
			Annotations(
				entx.SkipIO()),
		// WorkflowStageExecution 1-* WorkflowStepExecutions.
		edge.From("stage_execution", WorkflowStageExecution.Type).
			Ref("step_executions").
			Field("workflow_stage_execution_id").
			Comment("Workflow stage execution that this workflow step execution belongs to.").
			Required().
			Unique().
			Required().
			Immutable(),
	}
}

func (WorkflowStepExecution) Interceptors() []ent.Interceptor {
	return []ent.Interceptor{
		intercept.ByProject("project_id"),
	}
}
