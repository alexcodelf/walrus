package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	"github.com/seal-io/walrus/pkg/dao/schema/mixin"
	"github.com/seal-io/walrus/pkg/dao/types/object"
)

type WorkflowStep struct {
	ent.Schema
}

func (WorkflowStep) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Metadata(),
		mixin.Status(),
	}
}

func (WorkflowStep) Fields() []ent.Field {
	return []ent.Field{
		field.String("type").
			Comment("Type of the workflow step.").
			NotEmpty().
			Immutable(),
		object.IDField("workflow_id").
			Comment("ID of the workflow that this workflow step belongs to.").
			Immutable(),
		object.IDField("stage_id").
			Comment("ID of the stage that this workflow step belongs to."),
		field.JSON("spec", map[string]any{}).
			Comment("Spec of the workflow step.").
			Optional(),
		field.JSON("input", map[string]any{}).
			Comment("Input of the workflow step.").
			Optional(),
		field.JSON("output", map[string]any{}).
			Comment("Output of the workflow step.").
			Optional(),
		field.JSON("dependencies", []object.ID{}).
			Comment("ID list of the workflow steps that this workflow step depends on.").
			Default([]object.ID{}),
	}
}
