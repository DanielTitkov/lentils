package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// Prediction holds the schema definition for the Prediction entity.
type Prediction struct {
	ent.Schema
}

// Fields of the Prediction.
func (Prediction) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Bool("prognosis").Immutable(),
		field.JSON("meta", make(map[string]interface{})).Optional(),
	}
}

// Edges of the Prediction.
func (Prediction) Edges() []ent.Edge {
	return []ent.Edge{
		// belongs to
		edge.From("challenge", Challenge.Type).Ref("predictions").Unique().Required(),
		edge.From("user", User.Type).Ref("predictions").Unique().Required(),
	}
}

func (Prediction) Indexes() []ent.Index {
	return []ent.Index{
		// one prediction per challenge for user
		index.Edges("challenge", "user").Unique(),
	}
}

func (Prediction) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
