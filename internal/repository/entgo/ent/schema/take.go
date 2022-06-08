package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// Take holds the schema definition for the Take entity.
type Take struct {
	ent.Schema
}

// Fields of the Take.
func (Take) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Int64("seed").Default(time.Now().Unix()).Immutable(),
		field.JSON("meta", make(map[string]interface{})).Optional(),
	}
}

// Edges of the Take.
func (Take) Edges() []ent.Edge {
	return []ent.Edge{
		// has
		edge.To("responses", Response.Type),
		// belongs to
		edge.From("test", Test.Type).Ref("takes").Unique().Required(),
		edge.From("user", User.Type).Ref("takes").Unique().Required(),
	}
}

func (Take) Indexes() []ent.Index {
	return []ent.Index{}
}

func (Take) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
