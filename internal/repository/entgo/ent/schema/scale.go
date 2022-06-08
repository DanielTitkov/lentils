package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/DanielTitkov/lentils/internal/domain"
	"github.com/google/uuid"
)

// Scale holds the schema definition for the Scale entity.
type Scale struct {
	ent.Schema
}

// Fields of the Scale.
func (Scale) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("code").NotEmpty().Unique(),
		field.Bool("global").Default(false),
		field.Enum("type").Values(
			domain.ScaleTypeSten,
		).Default(domain.ScaleTypeSten),
	}
}

// Edges of the Scale.
func (Scale) Edges() []ent.Edge {
	return []ent.Edge{
		// has
		edge.To("items", Item.Type),
		edge.To("interpretations", Interpretation.Type),
		edge.To("translations", ScaleTranslation.Type),
		// belongs to
		edge.From("test", Test.Type).Ref("scales"),
	}
}

func (Scale) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// TRANSLATIONS //

// ScaleTranslation holds the schema definition for the ScaleTranslation entity.
type ScaleTranslation struct {
	ent.Schema
}

// Fields of the ScaleTranslation.
func (ScaleTranslation) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("title").NotEmpty(),
		field.String("description").Optional(),
	}
}

// Edges of the ScaleTranslation.
func (ScaleTranslation) Edges() []ent.Edge {
	return []ent.Edge{
		// belongs to
		edge.From("scale", Scale.Type).Ref("translations").Unique(),
	}
}

func (ScaleTranslation) Indexes() []ent.Index {
	return []ent.Index{
		// one translation per locale per scale
		index.Edges("scale").Fields("locale").Unique(),
	}
}

func (ScaleTranslation) Mixin() []ent.Mixin {
	return []ent.Mixin{
		LocaleMixin{}, // holds locale names
	}
}
