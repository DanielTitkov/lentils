package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ScaleItem holds the edge schema definition for the ScaleItem edge.
type ScaleItem struct {
	ent.Schema
}

func (ScaleItem) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("scale_id", "item_id"),
	}
}

// Fields of the ScaleItem.
func (ScaleItem) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("reverse").Default(false),
		field.UUID("scale_id", uuid.UUID{}),
		field.UUID("item_id", uuid.UUID{}),
	}
}

// Edges of the ScaleItem.
func (ScaleItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("item", Item.Type).
			Unique().
			Required().
			Field("item_id"),
		edge.To("scale", Scale.Type).
			Unique().
			Required().
			Field("scale_id"),
	}
}
