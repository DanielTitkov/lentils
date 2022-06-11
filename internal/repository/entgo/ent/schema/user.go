package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("name").NotEmpty(),
		field.String("email").NotEmpty().Unique().Optional().Nillable(),
		field.String("picture").Optional().Default("https://www.gravatar.com/avatar/00000000000000000000000000000000?d=mp&f=y"),
		field.String("password_hash"),
		field.Bool("admin").Default(false),
		field.Bool("anonymous").Default(false),
		field.JSON("meta", make(map[string]interface{})).Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		// has
		// edge.To("tests", Test.Type), // not needed for now
		edge.To("sessions", UserSession.Type),
		edge.To("takes", Take.Type),
		// same type
		// if anonymous user registeres at some point his anonymous accounts
		// shoud be bound to his real account
		edge.To("aliases", User.Type).From("parent").Unique(),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		LocaleMixin{}, // holds locale names
	}
}
