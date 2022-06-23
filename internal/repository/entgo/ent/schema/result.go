package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// Result holds the schema definition for the Result entity.
type Result struct {
	ent.Schema
}

// Fields of the Result.
func (Result) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Float("raw_score"),
		field.Float("final_score"),
		field.JSON("meta", make(map[string]interface{})).Optional(),
	}
}

// Edges of the Result.
func (Result) Edges() []ent.Edge {
	return []ent.Edge{
		// belongs to
		edge.From("scale", Scale.Type).Ref("results").Unique().Required(),
		edge.From("take", Take.Type).Ref("results").Unique().Required(),
	}
}

func (Result) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges("scale", "take").Unique(),
	}
}

func (Result) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
