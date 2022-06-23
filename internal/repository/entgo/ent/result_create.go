// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/result"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/scale"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/take"
	"github.com/google/uuid"
)

// ResultCreate is the builder for creating a Result entity.
type ResultCreate struct {
	config
	mutation *ResultMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (rc *ResultCreate) SetCreateTime(t time.Time) *ResultCreate {
	rc.mutation.SetCreateTime(t)
	return rc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (rc *ResultCreate) SetNillableCreateTime(t *time.Time) *ResultCreate {
	if t != nil {
		rc.SetCreateTime(*t)
	}
	return rc
}

// SetUpdateTime sets the "update_time" field.
func (rc *ResultCreate) SetUpdateTime(t time.Time) *ResultCreate {
	rc.mutation.SetUpdateTime(t)
	return rc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (rc *ResultCreate) SetNillableUpdateTime(t *time.Time) *ResultCreate {
	if t != nil {
		rc.SetUpdateTime(*t)
	}
	return rc
}

// SetRawScore sets the "raw_score" field.
func (rc *ResultCreate) SetRawScore(f float64) *ResultCreate {
	rc.mutation.SetRawScore(f)
	return rc
}

// SetFinalScore sets the "final_score" field.
func (rc *ResultCreate) SetFinalScore(f float64) *ResultCreate {
	rc.mutation.SetFinalScore(f)
	return rc
}

// SetMeta sets the "meta" field.
func (rc *ResultCreate) SetMeta(m map[string]interface{}) *ResultCreate {
	rc.mutation.SetMeta(m)
	return rc
}

// SetID sets the "id" field.
func (rc *ResultCreate) SetID(u uuid.UUID) *ResultCreate {
	rc.mutation.SetID(u)
	return rc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (rc *ResultCreate) SetNillableID(u *uuid.UUID) *ResultCreate {
	if u != nil {
		rc.SetID(*u)
	}
	return rc
}

// SetScaleID sets the "scale" edge to the Scale entity by ID.
func (rc *ResultCreate) SetScaleID(id uuid.UUID) *ResultCreate {
	rc.mutation.SetScaleID(id)
	return rc
}

// SetScale sets the "scale" edge to the Scale entity.
func (rc *ResultCreate) SetScale(s *Scale) *ResultCreate {
	return rc.SetScaleID(s.ID)
}

// SetTakeID sets the "take" edge to the Take entity by ID.
func (rc *ResultCreate) SetTakeID(id uuid.UUID) *ResultCreate {
	rc.mutation.SetTakeID(id)
	return rc
}

// SetTake sets the "take" edge to the Take entity.
func (rc *ResultCreate) SetTake(t *Take) *ResultCreate {
	return rc.SetTakeID(t.ID)
}

// Mutation returns the ResultMutation object of the builder.
func (rc *ResultCreate) Mutation() *ResultMutation {
	return rc.mutation
}

// Save creates the Result in the database.
func (rc *ResultCreate) Save(ctx context.Context) (*Result, error) {
	var (
		err  error
		node *Result
	)
	rc.defaults()
	if len(rc.hooks) == 0 {
		if err = rc.check(); err != nil {
			return nil, err
		}
		node, err = rc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ResultMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rc.check(); err != nil {
				return nil, err
			}
			rc.mutation = mutation
			if node, err = rc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(rc.hooks) - 1; i >= 0; i-- {
			if rc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, rc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Result)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ResultMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rc *ResultCreate) SaveX(ctx context.Context) *Result {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *ResultCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *ResultCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rc *ResultCreate) defaults() {
	if _, ok := rc.mutation.CreateTime(); !ok {
		v := result.DefaultCreateTime()
		rc.mutation.SetCreateTime(v)
	}
	if _, ok := rc.mutation.UpdateTime(); !ok {
		v := result.DefaultUpdateTime()
		rc.mutation.SetUpdateTime(v)
	}
	if _, ok := rc.mutation.ID(); !ok {
		v := result.DefaultID()
		rc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *ResultCreate) check() error {
	if _, ok := rc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Result.create_time"`)}
	}
	if _, ok := rc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Result.update_time"`)}
	}
	if _, ok := rc.mutation.RawScore(); !ok {
		return &ValidationError{Name: "raw_score", err: errors.New(`ent: missing required field "Result.raw_score"`)}
	}
	if _, ok := rc.mutation.FinalScore(); !ok {
		return &ValidationError{Name: "final_score", err: errors.New(`ent: missing required field "Result.final_score"`)}
	}
	if _, ok := rc.mutation.ScaleID(); !ok {
		return &ValidationError{Name: "scale", err: errors.New(`ent: missing required edge "Result.scale"`)}
	}
	if _, ok := rc.mutation.TakeID(); !ok {
		return &ValidationError{Name: "take", err: errors.New(`ent: missing required edge "Result.take"`)}
	}
	return nil
}

func (rc *ResultCreate) sqlSave(ctx context.Context) (*Result, error) {
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (rc *ResultCreate) createSpec() (*Result, *sqlgraph.CreateSpec) {
	var (
		_node = &Result{config: rc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: result.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: result.FieldID,
			},
		}
	)
	if id, ok := rc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := rc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: result.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := rc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: result.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := rc.mutation.RawScore(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: result.FieldRawScore,
		})
		_node.RawScore = value
	}
	if value, ok := rc.mutation.FinalScore(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: result.FieldFinalScore,
		})
		_node.FinalScore = value
	}
	if value, ok := rc.mutation.Meta(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: result.FieldMeta,
		})
		_node.Meta = value
	}
	if nodes := rc.mutation.ScaleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   result.ScaleTable,
			Columns: []string{result.ScaleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: scale.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.scale_results = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.TakeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   result.TakeTable,
			Columns: []string{result.TakeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: take.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.take_results = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ResultCreateBulk is the builder for creating many Result entities in bulk.
type ResultCreateBulk struct {
	config
	builders []*ResultCreate
}

// Save creates the Result entities in the database.
func (rcb *ResultCreateBulk) Save(ctx context.Context) ([]*Result, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Result, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ResultMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *ResultCreateBulk) SaveX(ctx context.Context) []*Result {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *ResultCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *ResultCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}