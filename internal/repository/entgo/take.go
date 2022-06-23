package entgo

import (
	"context"

	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/take"

	"github.com/google/uuid"

	"github.com/DanielTitkov/lentils/internal/domain"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent"
)

func (r *EntgoRepository) CreateTake(ctx context.Context, tk *domain.Take) (*domain.Take, error) {
	t, err := r.client.Take.Create().
		SetTestID(tk.TestID).
		SetUserID(tk.UserID).
		SetMeta(tk.Meta).
		SetNillableStartTime(tk.StartTime).
		SetSeed(tk.Seed).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return entToDomainTake(t, tk.UserID, tk.TestID), nil
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
		StartTime:  t.StartTime,
		EndTime:    t.EndTime,
		Suspicious: t.Suspicious,
		CreateTime: t.CreateTime,
		UpdateTime: t.UpdateTime,
		Meta:       t.Meta,
	}
}
