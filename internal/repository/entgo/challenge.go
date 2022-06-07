package entgo

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/prediction"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/user"

	"github.com/google/uuid"

	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/challenge"

	"github.com/DanielTitkov/lentils/internal/domain"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent"
)

func (r *EntgoRepository) GetChallengeCount(ctx context.Context) (int, error) {
	return r.client.Challenge.Query().Count(ctx)
}

func (r *EntgoRepository) GetOngoingChallengeCount(ctx context.Context) (int, error) {
	return r.client.Challenge.
		Query().
		Where(
			challenge.And(
				challenge.StartTimeLT(time.Now()),
				challenge.EndTimeGT(time.Now()),
				challenge.PublishedEQ(true),
			),
		).
		Count(ctx)
}

func (r *EntgoRepository) GetFinishedChallengeCount(ctx context.Context) (int, error) {
	return r.client.Challenge.
		Query().
		Where(
			challenge.And(
				challenge.StartTimeLT(time.Now()),
				challenge.EndTimeLT(time.Now()),
				challenge.PublishedEQ(true),
			),
		).
		Count(ctx)
}

func (r *EntgoRepository) GetChallengeByContent(ctx context.Context, content string) (*domain.Challenge, error) {
	c, err := r.client.Challenge.
		Query().
		Where(challenge.ContentEQ(content)).
		WithPredictions().
		WithProofs().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return entToDomainChallenge(c, nil), nil
}

func (r *EntgoRepository) GetChallengeByID(ctx context.Context, id uuid.UUID, userID uuid.UUID) (*domain.Challenge, error) {
	c, err := r.client.Challenge.
		Query().
		Where(challenge.IDEQ(id)).
		WithPredictions().
		WithAuthor().
		WithProofs().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	var userPrediction *domain.Prediction
	if userID != uuid.Nil {
		pred, err := r.client.Prediction.
			Query().
			Where(
				prediction.And(
					prediction.HasUserWith(
						user.IDEQ(userID),
					),
					prediction.HasChallengeWith(
						challenge.IDEQ(id),
					),
				),
			).
			Only(ctx)
		if err != nil && !ent.IsNotFound(err) {
			return nil, err
		}
		if pred != nil {
			userPrediction = entToDomainPrediction(pred)
		}
	}

	return entToDomainChallenge(c, userPrediction), nil
}

func (r *EntgoRepository) GetRandomFinishedChallenges(ctx context.Context, limit int) ([]*domain.Challenge, error) {
	chs, err := r.client.Challenge.
		Query().
		Where(
			challenge.And(
				challenge.StartTimeLT(time.Now()),
				challenge.EndTimeLT(time.Now()),
				challenge.PublishedEQ(true),
			),
		).
		WithPredictions().
		WithProofs().
		All(ctx)
	if err != nil {
		return nil, err
	}

	var sampledChs []*ent.Challenge
	if len(chs) > limit {
		for i := 0; i < limit; i++ {
			rnd := rand.New(rand.NewSource(time.Now().Unix()))
			index := rnd.Intn(len(chs))
			sampledChs = append(sampledChs, chs[index])
			chs = append(chs[:index], chs[index+1:]...)
		}
	} else {
		sampledChs = chs
	}

	var res []*domain.Challenge
	for _, ch := range sampledChs {
		res = append(res, entToDomainChallenge(ch, nil))
	}

	return res, nil
}

func (r *EntgoRepository) GetRandomFalseChallenges(ctx context.Context, limit int) ([]*domain.Challenge, error) {
	chs, err := r.client.Challenge.
		Query().
		Where(
			challenge.And(
				challenge.StartTimeLT(time.Now()),
				challenge.EndTimeLT(time.Now()),
				challenge.Outcome(false),
				challenge.PublishedEQ(true),
			),
		).
		WithPredictions().
		WithProofs().
		All(ctx)
	if err != nil {
		return nil, err
	}

	var sampledChs []*ent.Challenge
	if len(chs) > limit {
		for i := 0; i < limit; i++ {
			rnd := rand.New(rand.NewSource(time.Now().Unix()))
			index := rnd.Intn(len(chs))
			sampledChs = append(sampledChs, chs[index])
			chs = append(chs[:index], chs[index+1:]...)
		}
	} else {
		sampledChs = chs
	}

	var res []*domain.Challenge
	for _, ch := range sampledChs {
		res = append(res, entToDomainChallenge(ch, nil))
	}

	return res, nil
}

func (r *EntgoRepository) GetRandomTrueChallenges(ctx context.Context, limit int) ([]*domain.Challenge, error) {
	chs, err := r.client.Challenge.
		Query().
		Where(
			challenge.And(
				challenge.StartTimeLT(time.Now()),
				challenge.EndTimeLT(time.Now()),
				challenge.Outcome(true),
				challenge.PublishedEQ(true),
			),
		).
		WithPredictions().
		WithProofs().
		All(ctx)
	if err != nil {
		return nil, err
	}

	var sampledChs []*ent.Challenge
	if len(chs) > limit {
		for i := 0; i < limit; i++ {
			rnd := rand.New(rand.NewSource(time.Now().Unix()))
			index := rnd.Intn(len(chs))
			sampledChs = append(sampledChs, chs[index])
			chs = append(chs[:index], chs[index+1:]...)
		}
	} else {
		sampledChs = chs
	}

	var res []*domain.Challenge
	for _, ch := range sampledChs {
		res = append(res, entToDomainChallenge(ch, nil))
	}

	return res, nil
}

// GetClosingChallenges returns challenges that are to be closed soon.
func (r *EntgoRepository) GetClosingChallenges(ctx context.Context, limit int) ([]*domain.Challenge, error) {
	chs, err := r.client.Challenge.
		Query().
		Where(
			challenge.And(
				challenge.StartTimeLT(time.Now()),
				challenge.EndTimeGT(time.Now()),
				challenge.OutcomeIsNil(),
				challenge.PublishedEQ(true),
			),
		).
		WithPredictions().
		Order(ent.Asc(challenge.FieldEndTime)).
		Limit(limit).
		All(ctx)
	if err != nil {
		return nil, err
	}

	var res []*domain.Challenge
	for _, ch := range chs {
		res = append(res, entToDomainChallenge(ch, nil))
	}

	return res, nil
}

func (r *EntgoRepository) GetRandomPendingChallenges(ctx context.Context, limit int) ([]*domain.Challenge, error) {
	chs, err := r.client.Challenge.
		Query().
		Where(
			challenge.And(
				challenge.StartTimeLT(time.Now()),
				challenge.EndTimeLT(time.Now()),
				challenge.OutcomeIsNil(),
				challenge.PublishedEQ(true),
			),
		).
		WithPredictions().
		All(ctx)
	if err != nil {
		return nil, err
	}

	var sampledChs []*ent.Challenge
	if len(chs) > limit {
		for i := 0; i < limit; i++ {
			rnd := rand.New(rand.NewSource(time.Now().Unix()))
			index := rnd.Intn(len(chs))
			sampledChs = append(sampledChs, chs[index])
			chs = append(chs[:index], chs[index+1:]...)
		}
	} else {
		sampledChs = chs
	}

	var res []*domain.Challenge
	for _, ch := range sampledChs {
		res = append(res, entToDomainChallenge(ch, nil))
	}

	return res, nil
}

func (r *EntgoRepository) GetRandomOngoingChallenges(ctx context.Context, limit int, userID uuid.UUID) ([]*domain.Challenge, error) {
	chs, err := r.client.Challenge.
		Query().
		Where(
			challenge.And(
				challenge.StartTimeLT(time.Now()),
				challenge.EndTimeGT(time.Now()),
				challenge.PublishedEQ(true),
				challenge.Not(
					challenge.HasPredictionsWith(
						prediction.HasUserWith(
							user.IDEQ(userID),
						),
					),
				),
			),
		).
		WithPredictions().
		All(ctx)
	if err != nil {
		return nil, err
	}

	var sampledChs []*ent.Challenge
	if len(chs) > limit {
		for i := 0; i < limit; i++ {
			rnd := rand.New(rand.NewSource(time.Now().Unix()))
			index := rnd.Intn(len(chs))
			sampledChs = append(sampledChs, chs[index])
			chs = append(chs[:index], chs[index+1:]...)
		}
	} else {
		sampledChs = chs
	}

	var res []*domain.Challenge
	for _, ch := range sampledChs {
		res = append(res, entToDomainChallenge(ch, nil))
	}

	return res, nil
}

// GetPopularOngoingChallenges returns ongoing challenges that have more predictions
func (r *EntgoRepository) GetPopularOngoingChallenges(ctx context.Context, limit int) ([]*domain.Challenge, error) {
	return nil, nil
}

// GetPopularFinishedChallenges returns finished challenges that have more predictions
func (r *EntgoRepository) GetPopularFinishedChallenges(ctx context.Context, limit int) ([]*domain.Challenge, error) {
	return nil, nil
}

func (r *EntgoRepository) CreateChallenge(ctx context.Context, ch *domain.Challenge) (*domain.Challenge, error) {
	query := r.client.Challenge.
		Create().
		SetContent(ch.Content).
		SetStartTime(ch.StartTime).
		SetEndTime(ch.EndTime).
		SetPublished(ch.Published).
		SetNillableOutcome(ch.Outcome).
		SetDescription(ch.Description)

	if ch.AuthorID != uuid.Nil {
		query.SetAuthorID(ch.AuthorID)
	}

	c, err := query.Save(ctx)
	if err != nil {
		return nil, err
	}
	return entToDomainChallenge(c, nil), nil
}

func (r *EntgoRepository) UpdateChallengeByID(ctx context.Context, ch *domain.Challenge) (*domain.Challenge, error) {
	c, err := r.client.Challenge.UpdateOneID(ch.ID).
		SetContent(ch.Content).
		SetDescription(ch.Description).
		SetStartTime(ch.StartTime).
		SetEndTime(ch.EndTime).
		SetPublished(ch.Published).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return entToDomainChallenge(c, nil), nil
}

func (r *EntgoRepository) CreateOrUpdateChallengeByContent(ctx context.Context, ch *domain.Challenge) (*domain.Challenge, error) {
	// query challenge by content
	c, err := r.client.Challenge.
		Query().
		Where(challenge.ContentEQ(ch.Content)).
		Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return nil, err
		}
		// create challenge
		chQuery := r.client.Challenge.
			Create().
			SetContent(ch.Content).
			SetStartTime(ch.StartTime).
			SetEndTime(ch.EndTime).
			SetPublished(ch.Published).
			SetNillableOutcome(ch.Outcome).
			SetDescription(ch.Description)

		if ch.AuthorID != uuid.Nil {
			chQuery.SetAuthorID(ch.AuthorID)
		}

		c, err = chQuery.Save(ctx)
		if err != nil {
			return nil, err
		}
		return entToDomainChallenge(c, nil), nil
	}

	if c.Outcome != nil {
		// it's not allowed to update challenges with set outcome
		return entToDomainChallenge(c, nil), nil
	}

	// update challenge
	chUpdateQuery := c.Update().
		SetStartTime(ch.StartTime).
		SetEndTime(ch.EndTime).
		SetPublished(ch.Published).
		SetNillableOutcome(ch.Outcome).
		SetDescription(ch.Description)

	c, err = chUpdateQuery.Save(ctx)
	if err != nil {
		return nil, err
	}

	return entToDomainChallenge(c, nil), nil
}

func (r *EntgoRepository) GetChallengesByAuthor(ctx context.Context, authorID uuid.UUID, limit, offset int) ([]*domain.Challenge, int, error) {
	query := r.client.Challenge.
		Query().
		Where(challenge.HasAuthorWith(user.IDEQ(authorID)))

	count, err := query.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	chs, err := query.
		Limit(limit).
		Offset(offset).
		All(ctx)
	if err != nil {
		return nil, 0, err
	}

	var res []*domain.Challenge
	for _, ch := range chs {
		res = append(res, entToDomainChallenge(ch, nil))
	}

	return res, count, nil
}

func (r *EntgoRepository) FilterChallenges(ctx context.Context, args *domain.FilterChallengesArgs) ([]*domain.Challenge, int, error) {
	query := r.client.Challenge.
		Query().
		WithPredictions().
		WithProofs()

	if args.Finished {
		query.Where(challenge.And(
			challenge.StartTimeLT(time.Now()),
			challenge.EndTimeLT(time.Now()),
		))
	}

	if args.Pending {
		query.Where(challenge.And(
			challenge.StartTimeLT(time.Now()),
			challenge.EndTimeLT(time.Now()),
			challenge.OutcomeIsNil(),
		))
	}

	if args.Unpublished {
		query.Where(challenge.PublishedEQ(false))
	} else {
		// in any other case we show only published challenges
		query.Where(challenge.PublishedEQ(true))
	}

	if args.Ongoing {
		query.Where(challenge.And(
			challenge.StartTimeLT(time.Now()),
			challenge.EndTimeGT(time.Now()),
		))
	}

	if args.Unvoted {
		query.Where(challenge.Not(
			challenge.HasPredictionsWith(
				prediction.HasUserWith(
					user.IDEQ(args.UserID),
				),
			),
		))
	}

	count, err := query.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	chs, err := query.
		Limit(args.Limit).
		Offset(args.Offset).
		All(ctx)
	if err != nil {
		return nil, 0, err
	}

	var res []*domain.Challenge
	for _, ch := range chs {
		res = append(res, entToDomainChallenge(ch, nil))
	}

	return res, count, nil
}

func (r *EntgoRepository) FilterUserChallenges(ctx context.Context, args *domain.FilterChallengesArgs) ([]*domain.Challenge, int, error) {
	query := r.client.Challenge.
		Query().
		Where(challenge.HasPredictionsWith(
			prediction.HasUserWith(
				user.IDEQ(args.UserID),
			),
		)).
		WithPredictions(func(q *ent.PredictionQuery) {
			q.Where(prediction.HasUserWith(
				user.IDEQ(args.UserID),
			))
		}).
		WithProofs()

	if args.Finished {
		query.Where(challenge.And(
			challenge.OutcomeNotNil(),
			challenge.StartTimeLT(time.Now()),
			challenge.EndTimeLT(time.Now()),
		))
	}

	if args.Pending {
		query.Where(challenge.And(
			challenge.StartTimeLT(time.Now()),
			challenge.EndTimeLT(time.Now()),
			challenge.OutcomeIsNil(),
		))
	}

	if args.Unpublished {
		query.Where(challenge.PublishedEQ(false))
	} else {
		// in any other case we show only published challenges
		query.Where(challenge.PublishedEQ(true))
	}

	if args.Ongoing {
		query.Where(challenge.And(
			challenge.OutcomeIsNil(),
		))
	}

	count, err := query.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	chs, err := query.
		Limit(args.Limit).
		Offset(args.Offset).
		All(ctx)
	if err != nil {
		return nil, 0, err
	}

	var res []*domain.Challenge
	for _, ch := range chs {
		if ch.Edges.Predictions != nil {
			if len(ch.Edges.Predictions) == 1 {
				res = append(res, entToDomainChallenge(
					ch,
					entToDomainPrediction(ch.Edges.Predictions[0])),
				)
			} else {
				return nil, 0, fmt.Errorf("expected exactly 1 prediction, but got %d", len(ch.Edges.Predictions))
			}
		}
	}

	return res, count, nil
}

func (r *EntgoRepository) SetChallengeOutcome(ctx context.Context, id uuid.UUID, outcome bool, proofs []*domain.Proof) error {
	tx, err := r.client.Tx(ctx)
	if err != nil {
		return fmt.Errorf("starting a transaction: %w", err)
	}

	// create proofs for challenge
	bulk := make([]*ent.ProofCreate, len(proofs))
	for i, p := range proofs {
		bulk[i] = tx.Proof.
			Create().
			SetChallengeID(id).
			SetContent(p.Content).
			SetLink(p.Link).
			SetMeta(p.Meta)
	}
	_, err = tx.Proof.CreateBulk(bulk...).Save(ctx)
	if err != nil {
		return rollback(tx, err)
	}

	// set challenge outcome
	_, err = tx.Challenge.
		UpdateOneID(id).
		SetOutcome(outcome).
		Save(ctx)
	if err != nil {
		return rollback(tx, err)
	}

	return tx.Commit()
}

func entToDomainChallenge(ch *ent.Challenge, userPrediction *domain.Prediction) *domain.Challenge {
	var predictions []*domain.Prediction
	if ch.Edges.Predictions != nil {
		for _, p := range ch.Edges.Predictions {
			predictions = append(predictions, entToDomainPrediction(p))
		}
	}

	var proofs []*domain.Proof
	if ch.Edges.Proofs != nil {
		for _, p := range ch.Edges.Proofs {
			proofs = append(proofs, entToDomainProof(p))
		}
	}

	var authorID uuid.UUID
	if ch.Edges.Author != nil {
		authorID = ch.Edges.Author.ID
	}

	return &domain.Challenge{
		ID:             ch.ID,
		Content:        ch.Content,
		Outcome:        ch.Outcome,
		Description:    ch.Description,
		StartTime:      ch.StartTime,
		EndTime:        ch.EndTime,
		Published:      ch.Published,
		Predictions:    predictions,
		AuthorID:       authorID,
		UserPrediction: userPrediction,
		Proofs:         proofs,
	}
}
