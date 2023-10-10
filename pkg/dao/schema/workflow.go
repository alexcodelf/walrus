package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/seal-io/walrus/pkg/dao/entx"
	"github.com/seal-io/walrus/pkg/dao/schema/intercept"
	"github.com/seal-io/walrus/pkg/dao/schema/mixin"
	"github.com/seal-io/walrus/pkg/dao/types/object"
)

type Workflow struct {
	ent.Schema
}

func (Workflow) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Metadata(),
		mixin.Status(),
	}
}

func (Workflow) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").
			Unique(),
	}
}

func (Workflow) Fields() []ent.Field {
	return []ent.Field{
		object.IDField("project_id").
			Comment("ID of the project that this workflow belongs to.").
			Immutable(),
		object.IDField("environment_id").
			Comment("ID of the environment that this workflow belongs to.").
			Optional().
			Immutable(),
		field.String("display_name").
			Comment("Display name is the human readable name that is shown to the user.").
			NotEmpty(),
		field.String("type").
			Comment("Type of the workflow.").
			NotEmpty().
			Immutable(),
		field.JSON("workflow_stage_ids", []object.ID{}).
			Comment("ID list of the stages that belong to this workflow.").
			Default([]object.ID{}),
		field.Int("parallelism").
			Comment("Number of task pods that can be executed in parallel of workflow.").
			Positive().
			Default(10),
	}
}

func (Workflow) Edges() []ent.Edge {
	return []ent.Edge{
		// Workflow 1-* WorkflowStages.
		edge.To("stages", WorkflowStage.Type).
			Comment("Stages that belong to this workflow.").
			Annotations(
				entsql.OnDelete(entsql.Cascade),
				entx.SkipIO()),

		// Workflow 1-* WorkflowExecutions.
		edge.To("executions", WorkflowExecution.Type).
			Comment("Workflow executions that belong to this workflow.").
			Annotations(
				entsql.OnDelete(entsql.Cascade),
				entx.SkipIO()),
	}
}

func (Workflow) Interceptors() []ent.Interceptor {
	return []ent.Interceptor{
		intercept.ByProject("project_id"),
	}
}
