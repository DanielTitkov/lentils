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
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return entToDomainTake(t), nil
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

	t, err = t.Update().
		SetProgress(tk.Progress).
		SetPage(tk.Page).
		SetStatus(take.Status(tk.Status)).
		SetMeta(tk.Meta).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return entToDomainTake(t), nil
}

func entToDomainTake(t *ent.Take) *domain.Take {
	var uid, tid uuid.UUID
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
		Meta:       t.Meta,
		UserID:     uid,
		TestID:     tid,
		CreateTime: t.CreateTime,
		UpdateTime: t.UpdateTime,
	}
}
