// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/itemtranslation"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/predicate"
)

// ItemTranslationDelete is the builder for deleting a ItemTranslation entity.
type ItemTranslationDelete struct {
	config
	hooks    []Hook
	mutation *ItemTranslationMutation
}

// Where appends a list predicates to the ItemTranslationDelete builder.
func (itd *ItemTranslationDelete) Where(ps ...predicate.ItemTranslation) *ItemTranslationDelete {
	itd.mutation.Where(ps...)
	return itd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (itd *ItemTranslationDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(itd.hooks) == 0 {
		affected, err = itd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ItemTranslationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			itd.mutation = mutation
			affected, err = itd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(itd.hooks) - 1; i >= 0; i-- {
			if itd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = itd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, itd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (itd *ItemTranslationDelete) ExecX(ctx context.Context) int {
	n, err := itd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (itd *ItemTranslationDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: itemtranslation.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: itemtranslation.FieldID,
			},
		},
	}
	if ps := itd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, itd.driver, _spec)
}

// ItemTranslationDeleteOne is the builder for deleting a single ItemTranslation entity.
type ItemTranslationDeleteOne struct {
	itd *ItemTranslationDelete
}

// Exec executes the deletion query.
func (itdo *ItemTranslationDeleteOne) Exec(ctx context.Context) error {
	n, err := itdo.itd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{itemtranslation.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (itdo *ItemTranslationDeleteOne) ExecX(ctx context.Context) {
	itdo.itd.ExecX(ctx)
}
