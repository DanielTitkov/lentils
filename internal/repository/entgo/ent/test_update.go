// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/predicate"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/question"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/scale"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/take"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/test"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/testtranslation"
	"github.com/google/uuid"
)

// TestUpdate is the builder for updating Test entities.
type TestUpdate struct {
	config
	hooks    []Hook
	mutation *TestMutation
}

// Where appends a list predicates to the TestUpdate builder.
func (tu *TestUpdate) Where(ps ...predicate.Test) *TestUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetUpdateTime sets the "update_time" field.
func (tu *TestUpdate) SetUpdateTime(t time.Time) *TestUpdate {
	tu.mutation.SetUpdateTime(t)
	return tu
}

// SetCode sets the "code" field.
func (tu *TestUpdate) SetCode(s string) *TestUpdate {
	tu.mutation.SetCode(s)
	return tu
}

// SetPublished sets the "published" field.
func (tu *TestUpdate) SetPublished(b bool) *TestUpdate {
	tu.mutation.SetPublished(b)
	return tu
}

// SetNillablePublished sets the "published" field if the given value is not nil.
func (tu *TestUpdate) SetNillablePublished(b *bool) *TestUpdate {
	if b != nil {
		tu.SetPublished(*b)
	}
	return tu
}

// AddTakeIDs adds the "takes" edge to the Take entity by IDs.
func (tu *TestUpdate) AddTakeIDs(ids ...uuid.UUID) *TestUpdate {
	tu.mutation.AddTakeIDs(ids...)
	return tu
}

// AddTakes adds the "takes" edges to the Take entity.
func (tu *TestUpdate) AddTakes(t ...*Take) *TestUpdate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.AddTakeIDs(ids...)
}

// AddQuestionIDs adds the "questions" edge to the Question entity by IDs.
func (tu *TestUpdate) AddQuestionIDs(ids ...uuid.UUID) *TestUpdate {
	tu.mutation.AddQuestionIDs(ids...)
	return tu
}

// AddQuestions adds the "questions" edges to the Question entity.
func (tu *TestUpdate) AddQuestions(q ...*Question) *TestUpdate {
	ids := make([]uuid.UUID, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return tu.AddQuestionIDs(ids...)
}

// AddTranslationIDs adds the "translations" edge to the TestTranslation entity by IDs.
func (tu *TestUpdate) AddTranslationIDs(ids ...uuid.UUID) *TestUpdate {
	tu.mutation.AddTranslationIDs(ids...)
	return tu
}

// AddTranslations adds the "translations" edges to the TestTranslation entity.
func (tu *TestUpdate) AddTranslations(t ...*TestTranslation) *TestUpdate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.AddTranslationIDs(ids...)
}

// AddScaleIDs adds the "scales" edge to the Scale entity by IDs.
func (tu *TestUpdate) AddScaleIDs(ids ...uuid.UUID) *TestUpdate {
	tu.mutation.AddScaleIDs(ids...)
	return tu
}

// AddScales adds the "scales" edges to the Scale entity.
func (tu *TestUpdate) AddScales(s ...*Scale) *TestUpdate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return tu.AddScaleIDs(ids...)
}

// Mutation returns the TestMutation object of the builder.
func (tu *TestUpdate) Mutation() *TestMutation {
	return tu.mutation
}

// ClearTakes clears all "takes" edges to the Take entity.
func (tu *TestUpdate) ClearTakes() *TestUpdate {
	tu.mutation.ClearTakes()
	return tu
}

// RemoveTakeIDs removes the "takes" edge to Take entities by IDs.
func (tu *TestUpdate) RemoveTakeIDs(ids ...uuid.UUID) *TestUpdate {
	tu.mutation.RemoveTakeIDs(ids...)
	return tu
}

// RemoveTakes removes "takes" edges to Take entities.
func (tu *TestUpdate) RemoveTakes(t ...*Take) *TestUpdate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.RemoveTakeIDs(ids...)
}

// ClearQuestions clears all "questions" edges to the Question entity.
func (tu *TestUpdate) ClearQuestions() *TestUpdate {
	tu.mutation.ClearQuestions()
	return tu
}

// RemoveQuestionIDs removes the "questions" edge to Question entities by IDs.
func (tu *TestUpdate) RemoveQuestionIDs(ids ...uuid.UUID) *TestUpdate {
	tu.mutation.RemoveQuestionIDs(ids...)
	return tu
}

// RemoveQuestions removes "questions" edges to Question entities.
func (tu *TestUpdate) RemoveQuestions(q ...*Question) *TestUpdate {
	ids := make([]uuid.UUID, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return tu.RemoveQuestionIDs(ids...)
}

// ClearTranslations clears all "translations" edges to the TestTranslation entity.
func (tu *TestUpdate) ClearTranslations() *TestUpdate {
	tu.mutation.ClearTranslations()
	return tu
}

// RemoveTranslationIDs removes the "translations" edge to TestTranslation entities by IDs.
func (tu *TestUpdate) RemoveTranslationIDs(ids ...uuid.UUID) *TestUpdate {
	tu.mutation.RemoveTranslationIDs(ids...)
	return tu
}

// RemoveTranslations removes "translations" edges to TestTranslation entities.
func (tu *TestUpdate) RemoveTranslations(t ...*TestTranslation) *TestUpdate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.RemoveTranslationIDs(ids...)
}

// ClearScales clears all "scales" edges to the Scale entity.
func (tu *TestUpdate) ClearScales() *TestUpdate {
	tu.mutation.ClearScales()
	return tu
}

// RemoveScaleIDs removes the "scales" edge to Scale entities by IDs.
func (tu *TestUpdate) RemoveScaleIDs(ids ...uuid.UUID) *TestUpdate {
	tu.mutation.RemoveScaleIDs(ids...)
	return tu
}

// RemoveScales removes "scales" edges to Scale entities.
func (tu *TestUpdate) RemoveScales(s ...*Scale) *TestUpdate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return tu.RemoveScaleIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TestUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	tu.defaults()
	if len(tu.hooks) == 0 {
		if err = tu.check(); err != nil {
			return 0, err
		}
		affected, err = tu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TestMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tu.check(); err != nil {
				return 0, err
			}
			tu.mutation = mutation
			affected, err = tu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tu.hooks) - 1; i >= 0; i-- {
			if tu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TestUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TestUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TestUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tu *TestUpdate) defaults() {
	if _, ok := tu.mutation.UpdateTime(); !ok {
		v := test.UpdateDefaultUpdateTime()
		tu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TestUpdate) check() error {
	if v, ok := tu.mutation.Code(); ok {
		if err := test.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "Test.code": %w`, err)}
		}
	}
	return nil
}

func (tu *TestUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   test.Table,
			Columns: test.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: test.FieldID,
			},
		},
	}
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: test.FieldUpdateTime,
		})
	}
	if value, ok := tu.mutation.Code(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: test.FieldCode,
		})
	}
	if value, ok := tu.mutation.Published(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: test.FieldPublished,
		})
	}
	if tu.mutation.TakesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   test.TakesTable,
			Columns: []string{test.TakesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: take.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedTakesIDs(); len(nodes) > 0 && !tu.mutation.TakesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   test.TakesTable,
			Columns: []string{test.TakesColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.TakesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   test.TakesTable,
			Columns: []string{test.TakesColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.QuestionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   test.QuestionsTable,
			Columns: test.QuestionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: question.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedQuestionsIDs(); len(nodes) > 0 && !tu.mutation.QuestionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   test.QuestionsTable,
			Columns: test.QuestionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: question.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.QuestionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   test.QuestionsTable,
			Columns: test.QuestionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: question.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.TranslationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   test.TranslationsTable,
			Columns: []string{test.TranslationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: testtranslation.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedTranslationsIDs(); len(nodes) > 0 && !tu.mutation.TranslationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   test.TranslationsTable,
			Columns: []string{test.TranslationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: testtranslation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.TranslationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   test.TranslationsTable,
			Columns: []string{test.TranslationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: testtranslation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.ScalesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   test.ScalesTable,
			Columns: test.ScalesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: scale.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedScalesIDs(); len(nodes) > 0 && !tu.mutation.ScalesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   test.ScalesTable,
			Columns: test.ScalesPrimaryKey,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.ScalesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   test.ScalesTable,
			Columns: test.ScalesPrimaryKey,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{test.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// TestUpdateOne is the builder for updating a single Test entity.
type TestUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TestMutation
}

// SetUpdateTime sets the "update_time" field.
func (tuo *TestUpdateOne) SetUpdateTime(t time.Time) *TestUpdateOne {
	tuo.mutation.SetUpdateTime(t)
	return tuo
}

// SetCode sets the "code" field.
func (tuo *TestUpdateOne) SetCode(s string) *TestUpdateOne {
	tuo.mutation.SetCode(s)
	return tuo
}

// SetPublished sets the "published" field.
func (tuo *TestUpdateOne) SetPublished(b bool) *TestUpdateOne {
	tuo.mutation.SetPublished(b)
	return tuo
}

// SetNillablePublished sets the "published" field if the given value is not nil.
func (tuo *TestUpdateOne) SetNillablePublished(b *bool) *TestUpdateOne {
	if b != nil {
		tuo.SetPublished(*b)
	}
	return tuo
}

// AddTakeIDs adds the "takes" edge to the Take entity by IDs.
func (tuo *TestUpdateOne) AddTakeIDs(ids ...uuid.UUID) *TestUpdateOne {
	tuo.mutation.AddTakeIDs(ids...)
	return tuo
}

// AddTakes adds the "takes" edges to the Take entity.
func (tuo *TestUpdateOne) AddTakes(t ...*Take) *TestUpdateOne {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.AddTakeIDs(ids...)
}

// AddQuestionIDs adds the "questions" edge to the Question entity by IDs.
func (tuo *TestUpdateOne) AddQuestionIDs(ids ...uuid.UUID) *TestUpdateOne {
	tuo.mutation.AddQuestionIDs(ids...)
	return tuo
}

// AddQuestions adds the "questions" edges to the Question entity.
func (tuo *TestUpdateOne) AddQuestions(q ...*Question) *TestUpdateOne {
	ids := make([]uuid.UUID, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return tuo.AddQuestionIDs(ids...)
}

// AddTranslationIDs adds the "translations" edge to the TestTranslation entity by IDs.
func (tuo *TestUpdateOne) AddTranslationIDs(ids ...uuid.UUID) *TestUpdateOne {
	tuo.mutation.AddTranslationIDs(ids...)
	return tuo
}

// AddTranslations adds the "translations" edges to the TestTranslation entity.
func (tuo *TestUpdateOne) AddTranslations(t ...*TestTranslation) *TestUpdateOne {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.AddTranslationIDs(ids...)
}

// AddScaleIDs adds the "scales" edge to the Scale entity by IDs.
func (tuo *TestUpdateOne) AddScaleIDs(ids ...uuid.UUID) *TestUpdateOne {
	tuo.mutation.AddScaleIDs(ids...)
	return tuo
}

// AddScales adds the "scales" edges to the Scale entity.
func (tuo *TestUpdateOne) AddScales(s ...*Scale) *TestUpdateOne {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return tuo.AddScaleIDs(ids...)
}

// Mutation returns the TestMutation object of the builder.
func (tuo *TestUpdateOne) Mutation() *TestMutation {
	return tuo.mutation
}

// ClearTakes clears all "takes" edges to the Take entity.
func (tuo *TestUpdateOne) ClearTakes() *TestUpdateOne {
	tuo.mutation.ClearTakes()
	return tuo
}

// RemoveTakeIDs removes the "takes" edge to Take entities by IDs.
func (tuo *TestUpdateOne) RemoveTakeIDs(ids ...uuid.UUID) *TestUpdateOne {
	tuo.mutation.RemoveTakeIDs(ids...)
	return tuo
}

// RemoveTakes removes "takes" edges to Take entities.
func (tuo *TestUpdateOne) RemoveTakes(t ...*Take) *TestUpdateOne {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.RemoveTakeIDs(ids...)
}

// ClearQuestions clears all "questions" edges to the Question entity.
func (tuo *TestUpdateOne) ClearQuestions() *TestUpdateOne {
	tuo.mutation.ClearQuestions()
	return tuo
}

// RemoveQuestionIDs removes the "questions" edge to Question entities by IDs.
func (tuo *TestUpdateOne) RemoveQuestionIDs(ids ...uuid.UUID) *TestUpdateOne {
	tuo.mutation.RemoveQuestionIDs(ids...)
	return tuo
}

// RemoveQuestions removes "questions" edges to Question entities.
func (tuo *TestUpdateOne) RemoveQuestions(q ...*Question) *TestUpdateOne {
	ids := make([]uuid.UUID, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return tuo.RemoveQuestionIDs(ids...)
}

// ClearTranslations clears all "translations" edges to the TestTranslation entity.
func (tuo *TestUpdateOne) ClearTranslations() *TestUpdateOne {
	tuo.mutation.ClearTranslations()
	return tuo
}

// RemoveTranslationIDs removes the "translations" edge to TestTranslation entities by IDs.
func (tuo *TestUpdateOne) RemoveTranslationIDs(ids ...uuid.UUID) *TestUpdateOne {
	tuo.mutation.RemoveTranslationIDs(ids...)
	return tuo
}

// RemoveTranslations removes "translations" edges to TestTranslation entities.
func (tuo *TestUpdateOne) RemoveTranslations(t ...*TestTranslation) *TestUpdateOne {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.RemoveTranslationIDs(ids...)
}

// ClearScales clears all "scales" edges to the Scale entity.
func (tuo *TestUpdateOne) ClearScales() *TestUpdateOne {
	tuo.mutation.ClearScales()
	return tuo
}

// RemoveScaleIDs removes the "scales" edge to Scale entities by IDs.
func (tuo *TestUpdateOne) RemoveScaleIDs(ids ...uuid.UUID) *TestUpdateOne {
	tuo.mutation.RemoveScaleIDs(ids...)
	return tuo
}

// RemoveScales removes "scales" edges to Scale entities.
func (tuo *TestUpdateOne) RemoveScales(s ...*Scale) *TestUpdateOne {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return tuo.RemoveScaleIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TestUpdateOne) Select(field string, fields ...string) *TestUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Test entity.
func (tuo *TestUpdateOne) Save(ctx context.Context) (*Test, error) {
	var (
		err  error
		node *Test
	)
	tuo.defaults()
	if len(tuo.hooks) == 0 {
		if err = tuo.check(); err != nil {
			return nil, err
		}
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TestMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tuo.check(); err != nil {
				return nil, err
			}
			tuo.mutation = mutation
			node, err = tuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuo.hooks) - 1; i >= 0; i-- {
			if tuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TestUpdateOne) SaveX(ctx context.Context) *Test {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TestUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TestUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tuo *TestUpdateOne) defaults() {
	if _, ok := tuo.mutation.UpdateTime(); !ok {
		v := test.UpdateDefaultUpdateTime()
		tuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TestUpdateOne) check() error {
	if v, ok := tuo.mutation.Code(); ok {
		if err := test.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "Test.code": %w`, err)}
		}
	}
	return nil
}

func (tuo *TestUpdateOne) sqlSave(ctx context.Context) (_node *Test, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   test.Table,
			Columns: test.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: test.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Test.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, test.FieldID)
		for _, f := range fields {
			if !test.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != test.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: test.FieldUpdateTime,
		})
	}
	if value, ok := tuo.mutation.Code(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: test.FieldCode,
		})
	}
	if value, ok := tuo.mutation.Published(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: test.FieldPublished,
		})
	}
	if tuo.mutation.TakesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   test.TakesTable,
			Columns: []string{test.TakesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: take.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedTakesIDs(); len(nodes) > 0 && !tuo.mutation.TakesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   test.TakesTable,
			Columns: []string{test.TakesColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.TakesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   test.TakesTable,
			Columns: []string{test.TakesColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.QuestionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   test.QuestionsTable,
			Columns: test.QuestionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: question.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedQuestionsIDs(); len(nodes) > 0 && !tuo.mutation.QuestionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   test.QuestionsTable,
			Columns: test.QuestionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: question.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.QuestionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   test.QuestionsTable,
			Columns: test.QuestionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: question.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.TranslationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   test.TranslationsTable,
			Columns: []string{test.TranslationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: testtranslation.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedTranslationsIDs(); len(nodes) > 0 && !tuo.mutation.TranslationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   test.TranslationsTable,
			Columns: []string{test.TranslationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: testtranslation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.TranslationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   test.TranslationsTable,
			Columns: []string{test.TranslationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: testtranslation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.ScalesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   test.ScalesTable,
			Columns: test.ScalesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: scale.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedScalesIDs(); len(nodes) > 0 && !tuo.mutation.ScalesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   test.ScalesTable,
			Columns: test.ScalesPrimaryKey,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.ScalesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   test.ScalesTable,
			Columns: test.ScalesPrimaryKey,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Test{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{test.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}