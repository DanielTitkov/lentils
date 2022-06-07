package entgo

import (
	"context"

	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/challenge"

	"github.com/google/uuid"

	"github.com/DanielTitkov/lentils/internal/domain"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/prediction"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/user"
)

func (r *EntgoRepository) GetPredictionCount(ctx context.Context) (int, error) {
	return r.client.Prediction.Query().Count(ctx)
}

func (r *EntgoRepository) CreatePrediction(ctx context.Context, pred *domain.Prediction) (*domain.Prediction, error) {
	p, err := r.client.Prediction.
		Create().
		SetUserID(pred.UserID).
		SetChallengeID(pred.ChallengeID).
		SetPrognosis(pred.Prognosis).
		SetMeta(pred.Meta).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return entToDomainPrediction(p), nil
}

func (r *EntgoRepository) GetPredictionByUserAndChallenge(ctx context.Context, userID, challengeID uuid.UUID) (*domain.Prediction, error) {
	p, err := r.client.Prediction.
		Query().
		Where(prediction.And(
			prediction.HasUserWith(user.IDEQ(userID)),
			prediction.HasChallengeWith(challenge.IDEQ(challengeID)),
		)).
		WithChallenge().
		WithUser().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return entToDomainPrediction(p), nil
}

func entToDomainPrediction(p *ent.Prediction) *domain.Prediction {
	var userID, challengeID uuid.UUID
	if p.Edges.User != nil {
		userID = p.Edges.User.ID
	}

	if p.Edges.Challenge != nil {
		challengeID = p.Edges.Challenge.ID
	}

	return &domain.Prediction{
		ID:          p.ID,
		UserID:      userID,
		ChallengeID: challengeID,
		Prognosis:   p.Prognosis,
		Meta:        p.Meta,
	}
}
