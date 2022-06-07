package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/DanielTitkov/lentils/internal/domain"
	"github.com/google/uuid"
)

// Challenge holds the schema definition for the Challenge entity.
type Challenge struct {
	ent.Schema
}

// Fields of the Challenge.
func (Challenge) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("content").NotEmpty().MaxLen(140).Unique(),
		field.String("description").Optional().MaxLen(280),
		field.Bool("outcome").Optional().Nillable(),
		field.Bool("published").Default(true),
		field.Time("start_time"),
		field.Time("end_time"),
		field.Enum("type").Values(
			domain.ChallengeTypeBool,
		).Immutable().Default(domain.ChallengeTypeBool),
	}
}

// Edges of the Challenge.
func (Challenge) Edges() []ent.Edge {
	return []ent.Edge{
		// has
		edge.To("predictions", Prediction.Type),
		edge.To("proofs", Proof.Type),
		// belongs to
		edge.From("author", User.Type).Ref("challenges").Unique(),
	}
}

func (Challenge) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
