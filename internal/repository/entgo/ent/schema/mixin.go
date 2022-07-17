package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/tinygodsdev/orrery/internal/domain"
)

// LocaleMixin implements the ent.Mixin for sharing
// locale fields with package schemas.
type LocaleMixin struct {
	mixin.Schema
}

func (LocaleMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("locale").Values(
			domain.Locales()...,
		).Immutable(),
	}
}
