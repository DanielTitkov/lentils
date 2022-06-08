package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// Response holds the schema definition for the Response entity.
type Response struct {
	ent.Schema
}

// Fields of the Response.
func (Response) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Int("value").Default(0),
		field.JSON("meta", make(map[string]interface{})).Optional(),
	}
}

// Edges of the Response.
func (Response) Edges() []ent.Edge {
	return []ent.Edge{
		// belongs to
		edge.From("item", Item.Type).Ref("responses").Unique().Required(),
		edge.From("take", Take.Type).Ref("responses").Unique().Required(),
	}
}

func (Response) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges("item", "take").Unique(),
	}
}

func (Response) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
