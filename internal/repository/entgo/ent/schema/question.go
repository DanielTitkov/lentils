package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/DanielTitkov/orrery/internal/domain"
	"github.com/google/uuid"
)

// Question holds the schema definition for the Question entity.
type Question struct {
	ent.Schema
}

// Fields of the Question.
func (Question) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Int("order").Default(10),
		field.String("code").NotEmpty().Unique().Immutable(),
		field.Enum("type").Values(
			domain.QuestionTypeSimple,
		).Default(domain.QuestionTypeSimple),
	}
}

// Edges of the Question.
func (Question) Edges() []ent.Edge {
	return []ent.Edge{
		// has
		edge.To("items", Item.Type),
		edge.To("translations", QuestionTranslation.Type),
		// belongs to
		edge.From("test", Test.Type).Ref("questions"),
	}
}

func (Question) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// TRANSLATIONS //

// QuestionTranslation holds the schema definition for the QuestionTranslation entity.
type QuestionTranslation struct {
	ent.Schema
}

// Fields of the QuestionTranslation.
func (QuestionTranslation) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("content").Optional(),
		field.String("header_content").Optional(),
		field.String("footer_content").Optional(),
	}
}

// Edges of the QuestionTranslation.
func (QuestionTranslation) Edges() []ent.Edge {
	return []ent.Edge{
		// belongs to
		edge.From("question", Question.Type).Ref("translations").Unique(),
	}
}

func (QuestionTranslation) Indexes() []ent.Index {
	return []ent.Index{
		// one translation per locale per question
		index.Edges("question").Fields("locale").Unique(),
	}
}

func (QuestionTranslation) Mixin() []ent.Mixin {
	return []ent.Mixin{
		LocaleMixin{}, // holds locale names
	}
}
