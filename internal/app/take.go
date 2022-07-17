package app

import (
	"context"

	"github.com/tinygodsdev/orrery/internal/domain"
	"github.com/google/uuid"
)

func (a *App) GetTake(ctx context.Context, takeID uuid.UUID) (*domain.Take, error) {
	return a.repo.GetTake(ctx, takeID)
}

func (a *App) AddResponse(ctx context.Context, take *domain.Take, item *domain.Item) (*domain.Take, *domain.Response, error) {
	take, err := a.repo.UpdateTake(ctx, take)
	if err != nil {
		a.log.Error("add response failed (update take)", err)
		return nil, nil, err
	}

	resp, err := a.repo.AddOrUpdateResponse(ctx, take.ID, item)
	if err != nil {
		a.log.Error("add response failed (add response)", err)
		return nil, nil, err
	}

	return take, resp, nil
}

func (a *App) UpdateTakeMark(ctx context.Context, takeID uuid.UUID, mark int) error {
	return a.repo.UpdateTakeMark(ctx, takeID, mark)
}
