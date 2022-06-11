package app

import (
	"context"

	"github.com/DanielTitkov/lentils/internal/domain"
)

func (a *App) BeginTest(ctx context.Context, take *domain.Take) (*domain.Take, error) {
	// change take status to questions
	err := take.Begin()
	if err != nil {
		return nil, err
	}

	take, err = a.repo.UpdateTake(ctx, take)
	if err != nil {
		return nil, err
	}

	return take, nil
}
