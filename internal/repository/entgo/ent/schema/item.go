package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// Item holds the schema definition for the Item entity.
type Item struct {
	ent.Schema
}

// Fields of the Item.
func (Item) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("code").NotEmpty().Unique().Immutable(),
		field.Int("steps").Default(2),
	}
}

// Edges of the Item.
func (Item) Edges() []ent.Edge {
	return []ent.Edge{
		// has
		edge.To("responses", Response.Type),
		edge.To("translations", ItemTranslation.Type),
		// belongs to
		edge.From("scale", Scale.Type).Ref("items").Through("scale_item", ScaleItem.Type),
		edge.From("question", Question.Type).Ref("items"),
	}
}

func (Item) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// TRANSLATIONS //

// ItemTranslations holds the schema definition for the ItemTranslations entity.
type ItemTranslation struct {
	ent.Schema
}

// Fields of the ItemTranslations.
func (ItemTranslation) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("content").NotEmpty(),
	}
}

// Edges of the ItemTranslations.
func (ItemTranslation) Edges() []ent.Edge {
	return []ent.Edge{
		// belongs to
		edge.From("item", Item.Type).Ref("translations").Unique(),
	}
}

func (ItemTranslation) Indexes() []ent.Index {
	return []ent.Index{
		// one translation per locale per item
		index.Edges("item").Fields("locale").Unique(),
	}
}

func (ItemTranslation) Mixin() []ent.Mixin {
	return []ent.Mixin{
		LocaleMixin{}, // holds locale names
	}
}
