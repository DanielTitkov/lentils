package schema

import (
	"github.com/DanielTitkov/lentils/internal/domain"

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
		field.Int64("seed").Immutable(),
		field.Int("progress").Default(0),
		field.Int("page").Default(0),
		field.Time("start_time").Optional().Nillable(),
		field.Time("end_time").Optional().Nillable(),
		field.Bool("suspicious").Default(false),
		field.Enum("status").Values(
			domain.TestStepIntro,
			domain.TestStepQuestions,
			domain.TestStepFinish,
			domain.TestStepResult,
		).Default(domain.TestStepIntro),
		field.JSON("meta", make(map[string]interface{})).Optional(),
	}
}

// Edges of the Take.
func (Take) Edges() []ent.Edge {
	return []ent.Edge{
		// has
		edge.To("responses", Response.Type),
		edge.To("results", Result.Type),
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
