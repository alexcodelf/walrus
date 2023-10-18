package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"github.com/seal-io/walrus/pkg/dao/entx"
	"github.com/seal-io/walrus/pkg/dao/schema/intercept"
	"github.com/seal-io/walrus/pkg/dao/schema/mixin"
	"github.com/seal-io/walrus/pkg/dao/types/object"
)

type WorkflowExecution struct {
	ent.Schema
}

func (WorkflowExecution) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Metadata(),
		mixin.Status(),
	}
}

func (WorkflowExecution) Fields() []ent.Field {
	return []ent.Field{
		object.IDField("project_id").
			Comment("ID of the project to belong.").
			NotEmpty().
			Immutable(),
		object.IDField("workflow_id").
			Comment("ID of the workflow that this workflow execution belongs to.").
			NotEmpty().
			Immutable(),
		object.IDField("subject_id").
			Comment("ID of the subject that this workflow execution belongs to.").
			Immutable(),
		field.String("progress").
			Comment("Progress of the workflow. N/M format," +
				"N is number of stages completed, M is total number of stages."),
		field.Int("duration").
			Comment("Duration of the workflow execution.").
			NonNegative().
			Default(0),
		field.JSON("stage_execution_ids", []object.ID{}).
			Comment("ID list of the stage executions that belong to this workflow execution.").
			Default([]object.ID{}),
		field.Text("record").
			Comment("Log record of the workflow execution.").
			Default(""),
		// TODO encrypt this field.
		field.Text("input").
			Comment("Input of the workflow execution. It's the yaml file that defines the workflow execution.").
			Default(""),
	}
}

func (WorkflowExecution) Edges() []ent.Edge {
	return []ent.Edge{
		// Project 1-* WorkflowExecutions.
		edge.From("project", Project.Type).
			Ref("workflow_executions").
			Field("project_id").
			Comment("Project to which the workflow execution belongs.").
			Unique().
			Required().
			Immutable().
			Annotations(
				entx.ValidateContext(intercept.WithProjectInterceptor),
				entx.SkipInput()),
		// WorkflowExecution 1-* WorkflowStageExecutions.
		edge.To("stages", WorkflowStageExecution.Type).
			Comment("Workflow stage executions that belong to this workflow execution.").
			Annotations(
				entsql.OnDelete(entsql.Cascade),
				entx.SkipInput()),

		// Workflow 1-* WorkflowExecutions.
		edge.From("workflow", Workflow.Type).
			Ref("executions").
			Field("workflow_id").
			Comment("Workflow that this workflow execution belongs to.").
			Required().
			Unique().
			Immutable().
			Annotations(
				entx.SkipInput(entx.WithCreate(), entx.WithUpdate())),
	}
}

func (WorkflowExecution) Interceptors() []ent.Interceptor {
	return []ent.Interceptor{
		intercept.ByProject("project_id"),
	}
}
