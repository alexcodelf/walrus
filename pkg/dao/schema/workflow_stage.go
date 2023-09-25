package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"github.com/seal-io/walrus/pkg/dao/entx"
	"github.com/seal-io/walrus/pkg/dao/schema/mixin"
	"github.com/seal-io/walrus/pkg/dao/types/object"
)

type WorkflowStage struct {
	ent.Schema
}

func (WorkflowStage) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Metadata(),
		mixin.Status(),
	}
}

func (WorkflowStage) Fields() []ent.Field {
	return []ent.Field{
		object.IDField("workflow_id").
			Comment("ID of the workflow that this workflow stage belongs to.").
			Immutable(),
		field.JSON("workflow_step_ids", []object.ID{}).
			Comment("IDs of the workflow steps that belong to this workflow stage.").
			Immutable(),
		field.Int("duration").
			Comment("Duration of the workflow stage.").
			Positive().
			Default(0),
		field.JSON("dependencies", []object.ID{}).
			Comment("ID list of the workflow stages that this workflow stage depends on.").
			Default([]object.ID{}),
	}
}

func (WorkflowStage) Edges() []ent.Edge {
	return []ent.Edge{
		// WorkflowStage 1-* WorkflowSteps.
		edge.To("steps", WorkflowStep.Type).
			Comment("Workflow steps that belong to this workflow stage.").
			Annotations(
				entsql.OnDelete(entsql.Cascade),
				entx.SkipIO()),
		// WorkflowStage 1-* WorkflowStageExecutions.
		edge.To("workflow_stage_executions", WorkflowStageExecution.Type).
			Comment("Workflow stage executions that belong to this workflow stage.").
			Annotations(
				entsql.OnDelete(entsql.Cascade),
				entx.SkipIO()),

		// Workflow 1-* WorkflowStages.
		edge.From("workflow", Workflow.Type).
			Ref("stages").
			Field("workflow_id").
			Comment("Workflow that this workflow stage belongs to.").
			Required().
			Unique().
			Immutable().
			Annotations(
				entx.SkipIO()),
	}
}
