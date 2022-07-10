package app

import (
	"context"
	"fmt"
	"time"

	"github.com/DanielTitkov/orrery/internal/domain"
)

func (a *App) GetSystemSummary(ctx context.Context) (*domain.SystemSymmary, error) {
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

	a.systemSummary = &domain.SystemSymmary{
		Users:      userCount,
		CreateTime: time.Now(),
	}

	a.log.Debug("system summary updated", fmt.Sprintf("%+v", a.systemSummary))
	return nil
}
