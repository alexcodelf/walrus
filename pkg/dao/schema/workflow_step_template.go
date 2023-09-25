package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	"github.com/seal-io/walrus/pkg/dao/schema/mixin"
	"github.com/seal-io/walrus/pkg/dao/types/property"
)

type WorkflowStepTemplate struct {
	ent.Schema
}

func (WorkflowStepTemplate) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Metadata(),
		mixin.Status(),
	}
}

func (WorkflowStepTemplate) Fields() []ent.Field {
	return []ent.Field{
		field.String("type").
			Comment("Type of the workflow step template.").
			NotEmpty().
			Immutable(),
		property.SchemasField("schema").
			Comment("Schema of the workflow step template.").
			Default(property.Schemas{}),
	}
}
