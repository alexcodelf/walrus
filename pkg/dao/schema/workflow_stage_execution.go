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

type WorkflowStageExecution struct {
	ent.Schema
}

func (WorkflowStageExecution) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Metadata(),
		mixin.Status(),
	}
}

func (WorkflowStageExecution) Fields() []ent.Field {
	return []ent.Field{
		field.Int("duration").
			Comment("Duration of the workflow stage execution.").
			Positive().
			Default(0),
		object.IDField("stage_id").
			Comment("ID of the workflow stage that this workflow stage execution belongs to.").
			Immutable(),
		object.IDField("workflow_execution_id").
			Comment("ID of the workflow execution that this workflow stage execution belongs to.").
			Immutable(),
		field.JSON("step_execution_ids", []object.ID{}).
			Comment("ID list of the workflow step executions that belong to this workflow stage execution.").
			Default([]object.ID{}),
		field.String("record").
			Comment("Log record of the workflow stage execution.").
			Default(""),
		field.Text("input").
			Comment("Input of the workflow stage execution." +
				" It's the yaml file that defines the workflow stage execution.").
			Default(""),
	}
}

func (WorkflowStageExecution) Edges() []ent.Edge {
	return []ent.Edge{
		// WorkflowStageExecution 1-* WorkflowStepExecutions.
		edge.To("step_executions", WorkflowStepExecution.Type).
			Comment("Workflow step executions that belong to this workflow stage execution.").
			Annotations(
				entsql.OnDelete(entsql.Cascade),
				entx.SkipIO()),
		// WorkflowStage 1-* WorkflowStageExecutions.
		edge.From("stage", WorkflowStage.Type).
			Ref("workflow_stage_executions").
			Field("stage_id").
			Comment("Workflow stage that this workflow stage execution belongs to.").
			Required().
			Unique().
			Immutable(),
		// WorkflowExecution 1-* WorkflowStageExecutions.
		edge.From("workflow_execution", WorkflowExecution.Type).
			Ref("workflow_stage_executions").
			Field("workflow_execution_id").
			Comment("Workflow execution that this workflow stage execution belongs to.").
			Required().
			Unique().
			Immutable(),
	}
}
