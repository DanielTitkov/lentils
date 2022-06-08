// Code generated by entc, DO NOT EDIT.

package testtranslation

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.TestTranslation {
	return predicate.TestTranslation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.TestTranslation {
	return predicate.TestTranslation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.TestTranslation {
	return predicate.TestTranslation(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.TestTranslation {
	return predicate.TestTranslation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.TestTranslation {
	return predicate.TestTranslation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.TestTranslation {
	return predicate.TestTranslation(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.TestTranslation {
	return predicate.TestTranslation(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.TestTranslation {
	return predicate.TestTranslation(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.TestTranslation {
	return predicate.TestTranslation(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.TestTranslation {
	return predicate.TestTranslation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.TestTranslation {
	return predicate.TestTranslation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// Instruction applies equality check predicate on the "instruction" field. It's identical to InstructionEQ.
func Instruction(v string) predicate.TestTranslation {
	return predicate.TestTranslation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInstruction), v))
	})
}

// LocaleEQ applies the EQ predicate on the "locale" field.
func LocaleEQ(v Locale) predicate.TestTranslation {
	return predicate.TestTranslation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLocale), v))
	})
}

// LocaleNEQ applies the NEQ predicate on the "locale" field.
func LocaleNEQ(v Locale) predicate.TestTranslation {
	return predicate.TestTranslation(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLocale), v))
	})
}

// LocaleIn applies the In predicate on the "locale" field.
func LocaleIn(vs ...Locale) predicate.TestTranslation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TestTranslation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLocale), v...))
	})
}

// LocaleNotIn applies the NotIn predicate on the "locale" field.
func LocaleNotIn(vs ...Locale) predicate.TestTranslation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TestTranslation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLocale), v...))
	})
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.TestTranslation {
	return predicate.TestTranslation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.TestTranslation {
	return predicate.TestTranslation(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTitle), v))
	})
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.TestTranslation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TestTranslation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTitle), v...))
	})
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.TestTranslation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TestTranslation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTitle), v...))
	})
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.TestTranslation {
	return predicate.TestTranslation(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTitle), v))
	})
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.TestTranslation {
	return predicate.TestTranslation(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTitle), v))
	})
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.TestTranslation {
	return predicate.TestTranslation(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTitle), v))
	})
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.TestTranslation {
	return predicate.TestTranslation(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTitle), v))
	})
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.TestTranslation {
	return predicate.TestTranslation(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTitle), v))
	})
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.TestTranslation {
	return predicate.TestTranslation(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTitle), v))
	})
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.TestTranslation {
	return predicate.TestTranslation(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTitle), v))
	})
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.TestTranslation {
	return predicate.TestTranslation(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTitle), v))
	})
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.TestTranslation {
	return predicate.TestTranslation(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTitle), v))
	})
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.TestTranslation {
	return predicate.TestTranslation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.TestTranslation {
	return predicate.TestTranslation(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDescription), v))
	})
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.TestTranslation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TestTranslation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDescription), v...))
	})
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.TestTranslation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TestTranslation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDescription), v...))
	})
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.TestTranslation {
	return predicate.TestTranslation(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDescription), v))
	})
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.TestTranslation {
	return predicate.TestTranslation(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDescription), v))
	})
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.TestTranslation {
	return predicate.TestTranslation(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDescription), v))
	})
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.TestTranslation {
	return predicate.TestTranslation(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDescription), v))
	})
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.TestTranslation {
	return predicate.TestTranslation(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDescription), v))
	})
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.TestTranslation {
	return predicate.TestTranslation(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDescription), v))
	})
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.TestTranslation {
	return predicate.TestTranslation(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDescription), v))
	})
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.TestTranslation {
	return predicate.TestTranslation(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDescription)))
	})
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.TestTranslation {
	return predicate.TestTranslation(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDescription)))
	})
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.TestTranslation {
	return predicate.TestTranslation(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDescription), v))
	})
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.TestTranslation {
	return predicate.TestTranslation(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDescription), v))
	})
}

// InstructionEQ applies the EQ predicate on the "instruction" field.
func InstructionEQ(v string) predicate.TestTranslation {
	return predicate.TestTranslation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInstruction), v))
	})
}

// InstructionNEQ applies the NEQ predicate on the "instruction" field.
func InstructionNEQ(v string) predicate.TestTranslation {
	return predicate.TestTranslation(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldInstruction), v))
	})
}

// InstructionIn applies the In predicate on the "instruction" field.
func InstructionIn(vs ...string) predicate.TestTranslation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TestTranslation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldInstruction), v...))
	})
}

// InstructionNotIn applies the NotIn predicate on the "instruction" field.
func InstructionNotIn(vs ...string) predicate.TestTranslation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TestTranslation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldInstruction), v...))
	})
}

// InstructionGT applies the GT predicate on the "instruction" field.
func InstructionGT(v string) predicate.TestTranslation {
	return predicate.TestTranslation(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldInstruction), v))
	})
}

// InstructionGTE applies the GTE predicate on the "instruction" field.
func InstructionGTE(v string) predicate.TestTranslation {
	return predicate.TestTranslation(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldInstruction), v))
	})
}

// InstructionLT applies the LT predicate on the "instruction" field.
func InstructionLT(v string) predicate.TestTranslation {
	return predicate.TestTranslation(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldInstruction), v))
	})
}

// InstructionLTE applies the LTE predicate on the "instruction" field.
func InstructionLTE(v string) predicate.TestTranslation {
	return predicate.TestTranslation(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldInstruction), v))
	})
}

// InstructionContains applies the Contains predicate on the "instruction" field.
func InstructionContains(v string) predicate.TestTranslation {
	return predicate.TestTranslation(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldInstruction), v))
	})
}

// InstructionHasPrefix applies the HasPrefix predicate on the "instruction" field.
func InstructionHasPrefix(v string) predicate.TestTranslation {
	return predicate.TestTranslation(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldInstruction), v))
	})
}

// InstructionHasSuffix applies the HasSuffix predicate on the "instruction" field.
func InstructionHasSuffix(v string) predicate.TestTranslation {
	return predicate.TestTranslation(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldInstruction), v))
	})
}

// InstructionIsNil applies the IsNil predicate on the "instruction" field.
func InstructionIsNil() predicate.TestTranslation {
	return predicate.TestTranslation(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldInstruction)))
	})
}

// InstructionNotNil applies the NotNil predicate on the "instruction" field.
func InstructionNotNil() predicate.TestTranslation {
	return predicate.TestTranslation(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldInstruction)))
	})
}

// InstructionEqualFold applies the EqualFold predicate on the "instruction" field.
func InstructionEqualFold(v string) predicate.TestTranslation {
	return predicate.TestTranslation(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldInstruction), v))
	})
}

// InstructionContainsFold applies the ContainsFold predicate on the "instruction" field.
func InstructionContainsFold(v string) predicate.TestTranslation {
	return predicate.TestTranslation(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldInstruction), v))
	})
}

// HasTest applies the HasEdge predicate on the "test" edge.
func HasTest() predicate.TestTranslation {
	return predicate.TestTranslation(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TestTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TestTable, TestColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTestWith applies the HasEdge predicate on the "test" edge with a given conditions (other predicates).
func HasTestWith(preds ...predicate.Test) predicate.TestTranslation {
	return predicate.TestTranslation(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TestInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TestTable, TestColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.TestTranslation) predicate.TestTranslation {
	return predicate.TestTranslation(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.TestTranslation) predicate.TestTranslation {
	return predicate.TestTranslation(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.TestTranslation) predicate.TestTranslation {
	return predicate.TestTranslation(func(s *sql.Selector) {
		p(s.Not())
	})
}
