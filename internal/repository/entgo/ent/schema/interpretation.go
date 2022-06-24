package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// Interpretation holds the schema definition for the Interpretation entity.
type Interpretation struct {
	ent.Schema
}

// Fields of the Interpretation.
func (Interpretation) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.JSON("range", [2]float64{}),
	}
}

// Edges of the Interpretation.
func (Interpretation) Edges() []ent.Edge {
	return []ent.Edge{
		// has
		edge.To("translations", InterpretationTranslation.Type).Annotations(
			entsql.Annotation{OnDelete: entsql.Cascade},
		),
		// belongs to
		edge.From("scale", Scale.Type).Ref("interpretations").Unique(),
	}
}

func (Interpretation) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// TRANSLATIONS //

// InterpretationTranslation holds the schema definition for the InterpretationTranslation entity.
type InterpretationTranslation struct {
	ent.Schema
}

// Fields of the InterpretationTranslation.
func (InterpretationTranslation) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("content").NotEmpty(),
	}
}

// Edges of the InterpretationTranslation.
func (InterpretationTranslation) Edges() []ent.Edge {
	return []ent.Edge{
		// belongs to
		edge.From("interpretation", Interpretation.Type).Ref("translations").Unique(),
	}
}

func (InterpretationTranslation) Indexes() []ent.Index {
	return []ent.Index{
		// one translation per locale per interpretation
		index.Edges("interpretation").Fields("locale").Unique(),
	}
}

func (InterpretationTranslation) Mixin() []ent.Mixin {
	return []ent.Mixin{
		LocaleMixin{}, // holds locale names
	}
}
