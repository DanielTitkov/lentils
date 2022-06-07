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

// Test holds the schema definition for the Test entity.
type Test struct {
	ent.Schema
}

// Fields of the Test.
func (Test) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("code").NotEmpty().MaxLen(100).Unique(),
		field.String("content").NotEmpty().MaxLen(140).Unique(),
		field.String("description").Optional().MaxLen(280),
		field.Bool("published").Default(true),
	}
}

// Edges of the Test.
func (Test) Edges() []ent.Edge {
	return []ent.Edge{
		// has
		edge.To("translations", TestTranslation.Type),
		// belongs to
		edge.From("author", User.Type).Ref("tests").Unique(),
	}
}

func (Test) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// TRANSLATIONS //

// TestTranslations holds the schema definition for the TestTranslations entity.
type TestTranslation struct {
	ent.Schema
}

// Fields of the TestTranslations.
func (TestTranslation) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("title").NotEmpty().MaxLen(140),
		field.String("description").Optional(),
		field.String("instruction").Optional(),
		field.Enum("locale").Values(
			domain.LocaleEn,
			domain.LocaleRu,
		),
	}
}

// Edges of the TestTranslations.
func (TestTranslation) Edges() []ent.Edge {
	return []ent.Edge{
		// belongs to
		edge.From("test", Test.Type).Ref("translations").Unique(),
	}
}

func (TestTranslation) Indexes() []ent.Index {
	return []ent.Index{
		// one translation per locale per test
		index.Edges("test").Fields("locale").Unique(),
	}
}

func (TestTranslation) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
