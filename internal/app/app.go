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
		// user
		IfEmailRegistered(context.Context, string) (bool, error)
		GetUserByEmail(context.Context, string) (*domain.User, error)
		GetUserByID(context.Context, uuid.UUID) (*domain.User, error)
		CreateUser(context.Context, *domain.User) (*domain.User, error)

		// user session
		IfSessionRegistered(context.Context, *domain.UserSession) (bool, error)
		CreateUserSession(context.Context, *domain.UserSession) (*domain.UserSession, error)
		CreateOrUpdateUserSession(context.Context, *domain.UserSession) (*domain.UserSession, error)
		UpdateUserSessionLastActivityBySID(context.Context, string) error
		GetUserBySession(context.Context, *domain.UserSession) (*domain.User, error)

		// test
		CreateOrUpdateTestFromArgs(context.Context, *domain.CreateTestArgs) error
		GetTests(ctx context.Context, locale string) ([]*domain.Test, error)

		// for system summary
		GetUserCount(ctx context.Context) (int, error)
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

	err := app.loadUserPresets()
	if err != nil {
		return nil, err
	}

	err = app.loadTestPresets()
	if err != nil {
		return nil, err
	}

	app.log.Info("finished loading presets", "")

	// init app jobs, caches and preload data (if any)
	go app.UpdateSystemSummaryJob() // TODO: move to jobs?

	return &app, nil
}
