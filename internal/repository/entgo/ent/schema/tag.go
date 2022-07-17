package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/tinygodsdev/orrery/internal/domain"
	"github.com/google/uuid"
)

// Tag holds the schema definition for the Tag entity.
type Tag struct {
	ent.Schema
}

// Fields of the Tag.
func (Tag) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("code").NotEmpty().Unique().Immutable(),
		field.Enum("type").Values(
			domain.TagTypeTheme,
			domain.TagTypeLen,
			domain.TagTypeFeature,
		).Default(domain.TagTypeFeature),
	}
}

// Edges of the Tag.
func (Tag) Edges() []ent.Edge {
	return []ent.Edge{
		// has
		edge.To("translations", TagTranslation.Type),
		// belongs to
		edge.From("test", Test.Type).Ref("tags"),
	}
}

func (Tag) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// TRANSLATIONS //

// TagTranslation holds the schema definition for the TagTranslation entity.
type TagTranslation struct {
	ent.Schema
}

// Fields of the TagTranslation.
func (TagTranslation) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("content").NotEmpty(),
	}
}

// Edges of the TagTranslation.
func (TagTranslation) Edges() []ent.Edge {
	return []ent.Edge{
		// belongs to
		edge.From("tag", Tag.Type).Ref("translations").Unique(),
	}
}

func (TagTranslation) Indexes() []ent.Index {
	return []ent.Index{
		// one translation per locale per tag
		index.Edges("tag").Fields("locale").Unique(),
	}
}

func (TagTranslation) Mixin() []ent.Mixin {
	return []ent.Mixin{
		LocaleMixin{}, // holds locale names
	}
}
