package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/DanielTitkov/lentils/internal/domain"
	"github.com/google/uuid"
)

// Sample holds the schema definition for the Sample entity.
type Sample struct {
	ent.Schema
}

// Fields of the Sample.
func (Sample) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("code").NotEmpty().Unique().Immutable(),
		field.JSON("criteria", domain.SampleCriteria{}),
	}
}

// Edges of the Sample.
func (Sample) Edges() []ent.Edge {
	return []ent.Edge{
		// has
		edge.To("norms", Norm.Type),
		// belongs to
	}
}

func (Sample) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
