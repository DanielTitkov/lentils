package app

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/google/uuid"

	"github.com/DanielTitkov/lentils/internal/domain"
)

func (a *App) GetChallengeByContent(ctx context.Context, content string) (*domain.Challenge, error) {
	return a.repo.GetChallengeByContent(ctx, content)
}

func (a *App) GetChallengeByID(ctx context.Context, id uuid.UUID, userID uuid.UUID) (*domain.Challenge, error) {
	return a.repo.GetChallengeByID(ctx, id, userID)
}

func (a *App) GetRandomFinishedChallenges(ctx context.Context) ([]*domain.Challenge, error) {
	return a.repo.GetRandomFinishedChallenges(ctx, a.Cfg.App.HomeChallengePageLimit)
}

func (a *App) GetRandomPendingChallenges(ctx context.Context) ([]*domain.Challenge, error) {
	return a.repo.GetRandomPendingChallenges(ctx, a.Cfg.App.HomeChallengePageLimit)
}

func (a *App) GetClosingChallenges(ctx context.Context) ([]*domain.Challenge, error) {
	return a.repo.GetClosingChallenges(ctx, a.Cfg.App.HomeChallengePageLimit)
}

func (a *App) GetRandomOngoingChallenges(ctx context.Context, userID uuid.UUID, limit int) ([]*domain.Challenge, error) {
	if limit == 0 || limit > a.Cfg.App.HomeChallengePageLimit {
		limit = a.Cfg.App.HomeChallengePageLimit
	}

	return a.repo.GetRandomOngoingChallenges(ctx, limit, userID)
}

func (a *App) GetRandomTrueChallenges(ctx context.Context, limit int) ([]*domain.Challenge, error) {
	if limit == 0 || limit > a.Cfg.App.HomeChallengePageLimit {
		limit = a.Cfg.App.HomeChallengePageLimit
	}

	return a.repo.GetRandomTrueChallenges(ctx, limit)
}

func (a *App) GetRandomFalseChallenges(ctx context.Context, limit int) ([]*domain.Challenge, error) {
	if limit == 0 || limit > a.Cfg.App.HomeChallengePageLimit {
		limit = a.Cfg.App.HomeChallengePageLimit
	}

	return a.repo.GetRandomFalseChallenges(ctx, limit)
}

func (a *App) FilterUserChallenges(ctx context.Context, args *domain.FilterChallengesArgs) ([]*domain.Challenge, int, error) {
	if err := args.Validate(true); err != nil {
		return nil, 0, err
	}

	return a.repo.FilterUserChallenges(ctx, args)
}

func (a *App) FilterChallenges(ctx context.Context, args *domain.FilterChallengesArgs) ([]*domain.Challenge, int, error) {
	if err := args.Validate(false); err != nil {
		return nil, 0, err
	}

	return a.repo.FilterChallenges(ctx, args)
}

func (a *App) GetChallengesByAuthor(ctx context.Context, authorID uuid.UUID, limit, offset int) ([]*domain.Challenge, int, error) {
	return a.repo.GetChallengesByAuthor(ctx, authorID, limit, offset)
}

func (a *App) SetChallengeOutcome(ctx context.Context, id uuid.UUID, outcome bool, proofs []*domain.Proof) error {
	if len(proofs) < a.Cfg.App.MinProofCount {
		return fmt.Errorf("minimal proof count is %d, but got %d", a.Cfg.App.MinProofCount, len(proofs))
	}

	return a.repo.SetChallengeOutcome(ctx, id, outcome, proofs)
}

func (a *App) UpdateChallengeByID(ctx context.Context, id uuid.UUID, args *domain.CreateChallengeArgs) (*domain.Challenge, error) {
	start, end, err := args.GetStartEndTime()
	if err != nil {
		return nil, err
	}

	ch := &domain.Challenge{
		ID:          id,
		Content:     args.Content,
		Description: args.Description,
		StartTime:   start,
		EndTime:     end,
		Published:   args.Published,
	}

	return a.repo.UpdateChallengeByID(ctx, ch)
}

func (a *App) CreateChallengeFromArgs(ctx context.Context, args domain.CreateChallengeArgs, strict bool) (*domain.Challenge, error) {
	startTime, err := time.Parse(args.TimeLayout, args.StartTime)
	if err != nil {
		return nil, err
	}

	endTime, err := time.Parse(args.TimeLayout, args.EndTime)
	if err != nil {
		return nil, err
	}

	challenge := &domain.Challenge{
		Type:        args.Type,
		Content:     args.Content,
		Description: args.Description,
		StartTime:   startTime,
		EndTime:     endTime,
		Outcome:     args.Outcome,
		AuthorID:    args.AuthorID,
		Published:   args.Published,
	}

	if strict {
		challenge, err = a.repo.CreateChallenge(ctx, challenge)
	} else {
		challenge, err = a.repo.CreateOrUpdateChallengeByContent(ctx, challenge)
	}
	if err != nil {
		return nil, err
	}

	return challenge, nil

}

func (a *App) loadChallengePresets() error {
	a.log.Info("loading challenge presets", fmt.Sprint(a.Cfg.Data.Presets.ChallengePresetsPaths))
	for _, path := range a.Cfg.Data.Presets.ChallengePresetsPaths {
		a.log.Debug("reading from file", path)
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		var presets []domain.CreateChallengeArgs
		err = json.Unmarshal(data, &presets)
		if err != nil {
			return err
		}

		for _, args := range presets {
			ctx := context.Background()

			args.Published = true // all preset challenges go to publish
			challenge, err := a.CreateChallengeFromArgs(ctx, args, false)
			if err != nil {
				return err
			}

			a.log.Debug("loaded challenge", fmt.Sprintf("%+v", challenge))
		}
	}

	return nil
}
