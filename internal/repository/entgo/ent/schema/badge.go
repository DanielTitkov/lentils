package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Badge holds the schema definition for the Badge entity.
type Badge struct {
	ent.Schema
}

// Fields of the Badge.
func (Badge) Fields() []ent.Field {
	return []ent.Field{
		field.String("type").NotEmpty().Unique().Immutable(),
		field.Bool("active").Default(true),
		field.JSON("meta", make(map[string]interface{})).Optional(),
	}
}

// Edges of the Badge.
func (Badge) Edges() []ent.Edge {
	return []ent.Edge{
		// belongs to
		edge.From("users", User.Type).Ref("badges"),
	}
}

func (Badge) Indexes() []ent.Index {
	return []ent.Index{}
}

func (Badge) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
