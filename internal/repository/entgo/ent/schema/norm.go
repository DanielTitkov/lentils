package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// Norm holds the schema definition for the Norm entity.
type Norm struct {
	ent.Schema
}

// Fields of the Norm.
func (Norm) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("name").NotEmpty(),
		field.Int("base").Default(0).NonNegative(),
		field.Float("mean"),
		field.Float("sigma"),
		field.Int("rank").Default(0),
		field.JSON("meta", make(map[string]interface{})).Optional(),
	}
}

// Edges of the Norm.
func (Norm) Edges() []ent.Edge {
	return []ent.Edge{
		// belongs to
		edge.From("sample", Sample.Type).Ref("norms").Unique().Required(),
		edge.From("scale", Scale.Type).Ref("norms").Unique().Required(),
	}
}

func (Norm) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges("sample", "scale").Unique(),
	}
}

func (Norm) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
