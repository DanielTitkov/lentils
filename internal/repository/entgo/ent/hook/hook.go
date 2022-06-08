// Code generated by entc, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent"
)

// The InterpretationFunc type is an adapter to allow the use of ordinary
// function as Interpretation mutator.
type InterpretationFunc func(context.Context, *ent.InterpretationMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f InterpretationFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.InterpretationMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.InterpretationMutation", m)
	}
	return f(ctx, mv)
}

// The InterpretationTranslationFunc type is an adapter to allow the use of ordinary
// function as InterpretationTranslation mutator.
type InterpretationTranslationFunc func(context.Context, *ent.InterpretationTranslationMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f InterpretationTranslationFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.InterpretationTranslationMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.InterpretationTranslationMutation", m)
	}
	return f(ctx, mv)
}

// The ItemFunc type is an adapter to allow the use of ordinary
// function as Item mutator.
type ItemFunc func(context.Context, *ent.ItemMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ItemFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ItemMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ItemMutation", m)
	}
	return f(ctx, mv)
}

// The ItemTranslationFunc type is an adapter to allow the use of ordinary
// function as ItemTranslation mutator.
type ItemTranslationFunc func(context.Context, *ent.ItemTranslationMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ItemTranslationFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ItemTranslationMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ItemTranslationMutation", m)
	}
	return f(ctx, mv)
}

// The QuestionFunc type is an adapter to allow the use of ordinary
// function as Question mutator.
type QuestionFunc func(context.Context, *ent.QuestionMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f QuestionFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.QuestionMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.QuestionMutation", m)
	}
	return f(ctx, mv)
}

// The QuestionTranslationFunc type is an adapter to allow the use of ordinary
// function as QuestionTranslation mutator.
type QuestionTranslationFunc func(context.Context, *ent.QuestionTranslationMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f QuestionTranslationFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.QuestionTranslationMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.QuestionTranslationMutation", m)
	}
	return f(ctx, mv)
}

// The ResponseFunc type is an adapter to allow the use of ordinary
// function as Response mutator.
type ResponseFunc func(context.Context, *ent.ResponseMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ResponseFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ResponseMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ResponseMutation", m)
	}
	return f(ctx, mv)
}

// The ScaleFunc type is an adapter to allow the use of ordinary
// function as Scale mutator.
type ScaleFunc func(context.Context, *ent.ScaleMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ScaleFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ScaleMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ScaleMutation", m)
	}
	return f(ctx, mv)
}

// The ScaleTranslationFunc type is an adapter to allow the use of ordinary
// function as ScaleTranslation mutator.
type ScaleTranslationFunc func(context.Context, *ent.ScaleTranslationMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ScaleTranslationFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ScaleTranslationMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ScaleTranslationMutation", m)
	}
	return f(ctx, mv)
}

// The TakeFunc type is an adapter to allow the use of ordinary
// function as Take mutator.
type TakeFunc func(context.Context, *ent.TakeMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TakeFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.TakeMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TakeMutation", m)
	}
	return f(ctx, mv)
}

// The TestFunc type is an adapter to allow the use of ordinary
// function as Test mutator.
type TestFunc func(context.Context, *ent.TestMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TestFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.TestMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TestMutation", m)
	}
	return f(ctx, mv)
}

// The TestTranslationFunc type is an adapter to allow the use of ordinary
// function as TestTranslation mutator.
type TestTranslationFunc func(context.Context, *ent.TestTranslationMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TestTranslationFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.TestTranslationMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TestTranslationMutation", m)
	}
	return f(ctx, mv)
}

// The UserFunc type is an adapter to allow the use of ordinary
// function as User mutator.
type UserFunc func(context.Context, *ent.UserMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UserFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.UserMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UserMutation", m)
	}
	return f(ctx, mv)
}

// The UserSessionFunc type is an adapter to allow the use of ordinary
// function as UserSession mutator.
type UserSessionFunc func(context.Context, *ent.UserSessionMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UserSessionFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.UserSessionMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UserSessionMutation", m)
	}
	return f(ctx, mv)
}

// Condition is a hook condition function.
type Condition func(context.Context, ent.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op ent.Op) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
//
func If(hk ent.Hook, cond Condition) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, ent.Delete|ent.Create)
//
func On(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, ent.Update|ent.UpdateOne)
//
func Unless(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) ent.Hook {
	return func(ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(context.Context, ent.Mutation) (ent.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []ent.Hook {
//		return []ent.Hook{
//			Reject(ent.Delete|ent.Update),
//		}
//	}
//
func Reject(op ent.Op) ent.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []ent.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...ent.Hook) Chain {
	return Chain{append([]ent.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() ent.Hook {
	return func(mutator ent.Mutator) ent.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...ent.Hook) Chain {
	newHooks := make([]ent.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
