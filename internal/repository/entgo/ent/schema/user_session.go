package schema

import (
	"time"

	"entgo.io/ent/schema/index"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// UserSession holds the schema definition for the UserSession entity.
type UserSession struct {
	ent.Schema
}

// Fields of the UserSession.
func (UserSession) Fields() []ent.Field {
	return []ent.Field{
		field.String("sid").NotEmpty().Unique().Immutable(),
		field.String("ip"),
		field.String("user_agent"),
		field.Time("last_activity").Default(time.Now),
		field.Bool("active").Default(false),
		field.JSON("meta", make(map[string]interface{})).Optional(),
	}
}

// Edges of the UserSession.
func (UserSession) Edges() []ent.Edge {
	return []ent.Edge{
		// belongs to
		edge.From("user", User.Type).Ref("sessions").Unique().Required(),
	}
}

func (UserSession) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("active"),
	}
}

func (UserSession) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
