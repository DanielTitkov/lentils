package entgo

import (
	"context"

	"github.com/DanielTitkov/orrery/internal/repository/entgo/ent/take"

	"github.com/google/uuid"

	"github.com/DanielTitkov/orrery/internal/domain"
	"github.com/DanielTitkov/orrery/internal/repository/entgo/ent"
)

func (r *EntgoRepository) FinishedTakeCount(ctx context.Context) (int, error) {
	return r.client.Take.Query().Where(take.StatusEQ(take.StatusFinish)).Count(ctx)
}

func (r *EntgoRepository) GetTake(ctx context.Context, takeID uuid.UUID) (*domain.Take, error) {
	tk, err := r.client.Take.
		Query().
		Where(take.IDEQ(takeID)).
		WithTest(). // TODO maybe migrate to fk field for simpler queries
		WithUser().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return entToDomainTake(tk, uuid.Nil, uuid.Nil), nil
}

func (r *EntgoRepository) CreateTake(ctx context.Context, tk *domain.Take) (*domain.Take, error) {
	t, err := r.client.Take.Create().
		SetTestID(tk.TestID).
		SetUserID(tk.UserID).
		SetMeta(tk.Meta).
		SetNillableStartTime(tk.StartTime).
		SetInLocale(take.InLocale(tk.InLocale)).
		SetSeed(tk.Seed).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return entToDomainTake(t, tk.UserID, tk.TestID), nil
}

func (r *EntgoRepository) UpdateTakeMark(ctx context.Context, takeID uuid.UUID, mark int) error {
	_, err := r.client.Take.UpdateOneID(takeID).SetMark(mark).Save(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r *EntgoRepository) UpdateTake(ctx context.Context, tk *domain.Take) (*domain.Take, error) {
	t, err := r.client.Take.Query().
		Where(take.IDEQ(tk.ID)).
		WithUser().
		WithTest().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	updatedQuery := t.Update().
		SetProgress(tk.Progress).
		SetPage(tk.Page).
		SetInLocale(take.InLocale(tk.InLocale)).
		SetStatus(take.Status(tk.Status)).
		SetNillableStartTime(tk.StartTime).
		SetMeta(tk.Meta).
		SetSuspicious(tk.Suspicious)

	if tk.EndTime != nil && !tk.EndTime.IsZero() {
		updatedQuery.SetNillableEndTime(tk.EndTime)
	}

	t, err = updatedQuery.Save(ctx)
	if err != nil {
		return nil, err
	}

	return entToDomainTake(t, tk.UserID, tk.TestID), nil
}

func entToDomainTake(t *ent.Take, uid, tid uuid.UUID) *domain.Take {
	if t.Edges.User != nil {
		uid = t.Edges.User.ID
	}

	if t.Edges.Test != nil {
		tid = t.Edges.Test.ID
	}

	return &domain.Take{
		ID:         t.ID,
		Progress:   t.Progress,
		Status:     t.Status.String(),
		Page:       t.Page,
		Seed:       t.Seed,
		UserID:     uid,
		TestID:     tid,
		InLocale:   t.InLocale.String(),
		StartTime:  t.StartTime,
		EndTime:    t.EndTime,
		Suspicious: t.Suspicious,
		CreateTime: t.CreateTime,
		UpdateTime: t.UpdateTime,
		Mark:       t.Mark,
		Meta:       t.Meta,
	}
}
