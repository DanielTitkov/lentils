package app

import (
	"context"

	"github.com/DanielTitkov/lentils/internal/domain"
)

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
