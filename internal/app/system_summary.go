package app

import (
	"context"
	"fmt"
	"time"

	"github.com/tinygodsdev/orrery/internal/domain"
)

func (a *App) GetSystemSummary(ctx context.Context) (*domain.SystemSummary, error) {
	if a.systemSummary == nil {
		a.log.Debug("system summary requested but not found, gathering...", "")
		err := a.UpdateSystemSummary(ctx)
		if err != nil {
			return nil, err
		}
	}

	return a.systemSummary, nil
}

func (a *App) UpdateSystemSummary(ctx context.Context) error {
	a.log.Debug("updating system summary", "")

	userCount, err := a.repo.GetUserCount(ctx)
	if err != nil {
		return err
	}

	testCount, err := a.repo.TestCount(ctx)
	if err != nil {
		return err
	}

	finishedTakeCount, err := a.repo.FinishedTakeCount(ctx)
	if err != nil {
		return err
	}

	responseCount, err := a.repo.ResponseCount(ctx)
	if err != nil {
		return err
	}

	a.systemSummary = &domain.SystemSummary{
		Users:         userCount,
		CreateTime:    time.Now(),
		Tests:         testCount,
		FinishedTakes: finishedTakeCount,
		Responses:     responseCount,
	}

	a.log.Debug("system summary updated", fmt.Sprintf("%+v", a.systemSummary))
	return nil
}
