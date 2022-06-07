package app

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/DanielTitkov/lentils/internal/domain"
)

func (a *App) GetSystemSummary(ctx context.Context) (*domain.SystemSymmary, error) {
	if a.systemSummary == nil {
		a.log.Debug("system summary requested but not found, gathering...", "")
		err := a.updateSystemSummary(ctx)
		if err != nil {
			return nil, err
		}
	}

	return a.systemSummary, nil
}

func (a *App) UpdateSystemSummaryJob() {
	for {
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(a.Cfg.App.SystemSummaryTimeout)*time.Millisecond)
		processDone := make(chan bool)
		go func() {
			err := a.updateSystemSummary(ctx)
			if err != nil {
				a.log.Error("failed to update system summary", err)
			}
			processDone <- true
		}()

		select {
		case <-ctx.Done():
			a.log.Error("failed to update system summary", errors.New("timeout reached"))
		case <-processDone:
		}

		cancel()
		time.Sleep(time.Minute * time.Duration(a.Cfg.App.SystemSummaryInterval))
	}
}

func (a *App) updateSystemSummary(ctx context.Context) error {
	a.log.Debug("updating system summary", "")

	userCount, err := a.repo.GetUserCount(ctx)
	if err != nil {
		return err
	}

	challengeCount, err := a.repo.GetChallengeCount(ctx)
	if err != nil {
		return err
	}

	ongoingChallengeCount, err := a.repo.GetOngoingChallengeCount(ctx)
	if err != nil {
		return err
	}

	finishedChallengeCount, err := a.repo.GetFinishedChallengeCount(ctx)
	if err != nil {
		return err
	}

	predictionCount, err := a.repo.GetPredictionCount(ctx)
	if err != nil {
		return err
	}

	a.systemSummary = &domain.SystemSymmary{
		Users:              userCount,
		Predictions:        predictionCount,
		Challenges:         challengeCount,
		OngoingChallenges:  ongoingChallengeCount,
		FinishedChallenges: finishedChallengeCount,
		CreateTime:         time.Now(),
	}

	a.log.Debug("system summary updated", fmt.Sprintf("%+v", a.systemSummary))
	return nil
}
