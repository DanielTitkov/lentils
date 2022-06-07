package app

import (
	"context"

	"github.com/DanielTitkov/lentils/internal/configs"
	"github.com/DanielTitkov/lentils/internal/domain"
	"github.com/DanielTitkov/lentils/logger"
	"github.com/google/uuid"
	"github.com/gorilla/sessions"
)

type (
	App struct {
		Cfg           configs.Config
		log           *logger.Logger
		repo          Repository
		systemSummary *domain.SystemSymmary
		Store         sessions.Store
	}
	Repository interface {
		// challenge
		CreateChallenge(context.Context, *domain.Challenge) (*domain.Challenge, error)
		UpdateChallengeByID(context.Context, *domain.Challenge) (*domain.Challenge, error)
		SetChallengeOutcome(ctx context.Context, id uuid.UUID, outcome bool, proofs []*domain.Proof) error
		CreateOrUpdateChallengeByContent(context.Context, *domain.Challenge) (*domain.Challenge, error)
		GetChallengeByContent(context.Context, string) (*domain.Challenge, error)
		GetChallengeByID(ctx context.Context, id uuid.UUID, userID uuid.UUID) (*domain.Challenge, error)
		GetRandomFinishedChallenges(ctx context.Context, limit int) ([]*domain.Challenge, error)
		GetRandomOngoingChallenges(ctx context.Context, limit int, userID uuid.UUID) ([]*domain.Challenge, error)
		GetClosingChallenges(ctx context.Context, limit int) ([]*domain.Challenge, error)
		GetRandomPendingChallenges(ctx context.Context, limit int) ([]*domain.Challenge, error)
		FilterUserChallenges(ctx context.Context, args *domain.FilterChallengesArgs) ([]*domain.Challenge, int, error)
		FilterChallenges(context.Context, *domain.FilterChallengesArgs) ([]*domain.Challenge, int, error)
		GetRandomTrueChallenges(ctx context.Context, limit int) ([]*domain.Challenge, error)
		GetRandomFalseChallenges(ctx context.Context, limit int) ([]*domain.Challenge, error)
		GetChallengesByAuthor(ctx context.Context, authorID uuid.UUID, limit, offset int) ([]*domain.Challenge, int, error)

		// user
		IfEmailRegistered(context.Context, string) (bool, error)
		GetUserByEmail(context.Context, string) (*domain.User, error)
		GetUserByID(context.Context, uuid.UUID) (*domain.User, error)
		CreateUser(context.Context, *domain.User) (*domain.User, error)
		GetUserSummary(ctx context.Context, userID uuid.UUID) (*domain.UserSummary, error)

		// user session
		IfSessionRegistered(context.Context, *domain.UserSession) (bool, error)
		CreateUserSession(context.Context, *domain.UserSession) (*domain.UserSession, error)
		CreateOrUpdateUserSession(context.Context, *domain.UserSession) (*domain.UserSession, error)
		UpdateUserSessionLastActivityBySID(context.Context, string) error
		GetUserBySession(context.Context, *domain.UserSession) (*domain.User, error)

		// badge
		CreateOrUpdateBadgeByType(context.Context, *domain.Badge) (*domain.Badge, error)

		// prediction
		CreatePrediction(context.Context, *domain.Prediction) (*domain.Prediction, error)
		GetPredictionByUserAndChallenge(ctx context.Context, userID, challengeID uuid.UUID) (*domain.Prediction, error)

		// for system summary
		GetChallengeCount(ctx context.Context) (int, error)
		GetOngoingChallengeCount(ctx context.Context) (int, error)
		GetFinishedChallengeCount(ctx context.Context) (int, error)
		GetUserCount(ctx context.Context) (int, error)
		GetPredictionCount(ctx context.Context) (int, error)
	}
)

func New(
	cfg configs.Config,
	logger *logger.Logger,
	repo Repository,
	store sessions.Store,
) (*App, error) {
	app := App{
		Cfg:   cfg,
		log:   logger,
		repo:  repo,
		Store: store,
	}

	err := app.loadChallengePresets()
	if err != nil {
		return nil, err
	}

	err = app.loadUserPresets()
	if err != nil {
		return nil, err
	}

	err = app.loadPredictionPresets()
	if err != nil {
		return nil, err
	}

	app.log.Info("finished loading presets", "")

	// init app jobs, caches and preload data (if any)
	go app.UpdateSystemSummaryJob() // TODO: move to jobs?

	return &app, nil
}
