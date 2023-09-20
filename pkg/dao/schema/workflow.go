package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

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
		field.Int("duration").
			Comment("Duration of the workflow. N/M format," +
				"N is number of stages completed, M is total number of stages.").
			Positive().
			Immutable(),
		field.String("progress").
			Comment("Progress of the workflow.").
			Optional(),
		field.JSON("stages", []object.ID{}).
			Comment("ID list of the stages that belong to this workflow.").
			Default([]object.ID{}),
	}
}

func (Workflow) Edges() []ent.Edge {
	return []ent.Edge{}
}
