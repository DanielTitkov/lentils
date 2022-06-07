package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/DanielTitkov/lentils/internal/domain"
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
		field.String("email").NotEmpty().Unique(),
		field.String("picture").Optional().Default("https://www.gravatar.com/avatar/00000000000000000000000000000000?d=mp&f=y"),
		field.Bool("admin").Default(false),
		field.String("password_hash"),
		field.Enum("locale").Values(domain.LocaleEn, domain.LocaleRu).Default(domain.LocaleRu),
		field.JSON("meta", make(map[string]interface{})).Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		// has
		edge.To("predictions", Prediction.Type),
		edge.To("sessions", UserSession.Type),
		edge.To("badges", Badge.Type),
		edge.To("challenges", Challenge.Type),
		edge.To("tests", Test.Type),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
