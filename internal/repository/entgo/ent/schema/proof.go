package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// Proof holds the schema definition for the Proof entity.
type Proof struct {
	ent.Schema
}

// Fields of the Proof.
func (Proof) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("content").Immutable().MaxLen(280),
		field.String("link").Immutable(),
		field.JSON("meta", make(map[string]interface{})).Optional(),
	}
}

// Edges of the Proof.
func (Proof) Edges() []ent.Edge {
	return []ent.Edge{
		// belongs to
		edge.From("challenge", Challenge.Type).Ref("proofs").Unique().Required(),
	}
}

func (Proof) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
