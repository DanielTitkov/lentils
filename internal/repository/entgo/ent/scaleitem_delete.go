// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/predicate"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/scaleitem"
)

// ScaleItemDelete is the builder for deleting a ScaleItem entity.
type ScaleItemDelete struct {
	config
	hooks    []Hook
	mutation *ScaleItemMutation
}

// Where appends a list predicates to the ScaleItemDelete builder.
func (sid *ScaleItemDelete) Where(ps ...predicate.ScaleItem) *ScaleItemDelete {
	sid.mutation.Where(ps...)
	return sid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sid *ScaleItemDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(sid.hooks) == 0 {
		affected, err = sid.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ScaleItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sid.mutation = mutation
			affected, err = sid.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(sid.hooks) - 1; i >= 0; i-- {
			if sid.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sid.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sid.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (sid *ScaleItemDelete) ExecX(ctx context.Context) int {
	n, err := sid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sid *ScaleItemDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: scaleitem.Table,
		},
	}
	if ps := sid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, sid.driver, _spec)
}

// ScaleItemDeleteOne is the builder for deleting a single ScaleItem entity.
type ScaleItemDeleteOne struct {
	sid *ScaleItemDelete
}

// Exec executes the deletion query.
func (sido *ScaleItemDeleteOne) Exec(ctx context.Context) error {
	n, err := sido.sid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{scaleitem.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sido *ScaleItemDeleteOne) ExecX(ctx context.Context) {
	sido.sid.ExecX(ctx)
}