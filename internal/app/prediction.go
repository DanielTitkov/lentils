package app

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/DanielTitkov/lentils/internal/domain"
	"github.com/google/uuid"
)

func (a *App) GetPredictionByUserAndChallenge(ctx context.Context, userID, challengeID uuid.UUID) (*domain.Prediction, error) {
	return a.repo.GetPredictionByUserAndChallenge(ctx, userID, challengeID)
}

func (a *App) CreatePrediction(ctx context.Context, pred *domain.Prediction) (*domain.Prediction, error) {
	if pred.UserID == uuid.Nil {
		return nil, errors.New("user ID must be not null")
	}

	if pred.ChallengeID == uuid.Nil {
		return nil, errors.New("challenge ID must be not null")
	}

	return a.repo.CreatePrediction(ctx, pred)
}

func (a *App) loadPredictionPresets() error {
	a.log.Info("loading prediction presets", fmt.Sprint(a.Cfg.Data.Presets.PredictionPresetsPaths))
	for _, path := range a.Cfg.Data.Presets.PredictionPresetsPaths {
		a.log.Debug("reading from file", path)
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		var presets []domain.CreatePredictionFromPresetArgs
		err = json.Unmarshal(data, &presets)
		if err != nil {
			return err
		}

		for _, preset := range presets {
			ctx := context.Background()

			user, err := a.GetUserByEmail(ctx, preset.UserEmail)
			if err != nil {
				return err
			}

			challenge, err := a.GetChallengeByContent(ctx, preset.ChallengeContent)
			if err != nil {
				return err
			}

			pred, _ := a.GetPredictionByUserAndChallenge(ctx, user.ID, challenge.ID)
			if pred != nil {
				a.log.Debug("prediction already exists", fmt.Sprintf("%+v", pred))
				continue
			}

			pred, err = a.CreatePrediction(ctx, &domain.Prediction{
				UserID:      user.ID,
				ChallengeID: challenge.ID,
				Prognosis:   preset.Prognosis,
				Meta: map[string]interface{}{
					"test": true,
				},
			})
			if err != nil {
				return err
			}

			a.log.Debug("loaded prediction present", fmt.Sprintf("prediction: %+v, user: %+v, challenge: %+v", pred, user, challenge))
		}
	}

	return nil
}
