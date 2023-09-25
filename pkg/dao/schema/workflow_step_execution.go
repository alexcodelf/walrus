package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

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
		object.IDField("workflow_id").
			Comment("ID of the workflow that this workflow step execution belongs to.").
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
			Positive().
			Default(0),
		field.Int("duration").
			Comment("Duration of the workflow step execution.").
			Positive().
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
		// WorkflowStep 1-* WorkflowStepExecutions.
		edge.From("workflow_step", WorkflowStep.Type).
			Ref("executions").
			Field("workflow_step_id").
			Comment("Workflow step that this workflow step execution belongs to.").
			Required().
			Unique().
			Immutable(),
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
