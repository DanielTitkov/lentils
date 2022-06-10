package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// TestDisplay holds the schema definition for the TestDisplay entity.
type TestDisplay struct {
	ent.Schema
}

// Fields of the TestDisplay.
func (TestDisplay) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Bool("randomize_order").Default(false),
		field.Int("questions_per_page").Default(1).Positive(),
	}
}

// Edges of the TestDisplay.
func (TestDisplay) Edges() []ent.Edge {
	return []ent.Edge{
		// has
		// belongs to
		edge.From("test", Test.Type).Ref("display").Unique().Required(),
	}
}
