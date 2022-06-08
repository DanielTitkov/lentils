// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/predicate"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/questiontranslation"
)

// QuestionTranslationDelete is the builder for deleting a QuestionTranslation entity.
type QuestionTranslationDelete struct {
	config
	hooks    []Hook
	mutation *QuestionTranslationMutation
}

// Where appends a list predicates to the QuestionTranslationDelete builder.
func (qtd *QuestionTranslationDelete) Where(ps ...predicate.QuestionTranslation) *QuestionTranslationDelete {
	qtd.mutation.Where(ps...)
	return qtd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (qtd *QuestionTranslationDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(qtd.hooks) == 0 {
		affected, err = qtd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*QuestionTranslationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			qtd.mutation = mutation
			affected, err = qtd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(qtd.hooks) - 1; i >= 0; i-- {
			if qtd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = qtd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, qtd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (qtd *QuestionTranslationDelete) ExecX(ctx context.Context) int {
	n, err := qtd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (qtd *QuestionTranslationDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: questiontranslation.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: questiontranslation.FieldID,
			},
		},
	}
	if ps := qtd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, qtd.driver, _spec)
}

// QuestionTranslationDeleteOne is the builder for deleting a single QuestionTranslation entity.
type QuestionTranslationDeleteOne struct {
	qtd *QuestionTranslationDelete
}

// Exec executes the deletion query.
func (qtdo *QuestionTranslationDeleteOne) Exec(ctx context.Context) error {
	n, err := qtdo.qtd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{questiontranslation.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (qtdo *QuestionTranslationDeleteOne) ExecX(ctx context.Context) {
	qtdo.qtd.ExecX(ctx)
}
