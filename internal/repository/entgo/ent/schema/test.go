package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
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
		field.String("code").NotEmpty().MaxLen(100).Unique().Immutable(),
		field.Bool("published").Default(true),
		field.Strings("available_locales").Optional(),
	}
}

// Edges of the Test.
func (Test) Edges() []ent.Edge {
	return []ent.Edge{
		// has
		edge.To("takes", Take.Type),
		edge.To("questions", Question.Type),
		edge.To("translations", TestTranslation.Type),
		edge.To("scales", Scale.Type),
		edge.To("display", TestDisplay.Type).Unique(),
		edge.To("tags", Tag.Type),
		// belongs to
		// edge.From("author", User.Type).Ref("tests").Unique(), // not needed for now
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
		field.String("details").Optional(),
		field.String("instruction").Optional(),
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
		LocaleMixin{}, // holds locale names
	}
}
