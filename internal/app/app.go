package app

import (
	"context"
	"time"

	"github.com/DanielTitkov/lentils/internal/util"

	"github.com/DanielTitkov/lentils/internal/configs"
	"github.com/DanielTitkov/lentils/internal/domain"
	"github.com/DanielTitkov/lentils/internal/logger"
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
		locales       []string // locale count is not very big so no need to have map
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
		GetTestByCode(ctx context.Context, code string, locale string) (*domain.Test, error)
		GetTakeData(ctx context.Context, tk *domain.Take, locale string) (*domain.Test, error)

		// take
		CreateTake(ctx context.Context, tk *domain.Take) (*domain.Take, error)
		UpdateTake(ctx context.Context, tk *domain.Take) (*domain.Take, error)

		// response
		AddOrUpdateResponse(ctx context.Context, takeID uuid.UUID, itm *domain.Item) (*domain.Response, error)

		// sample
		CreateOrUpdateSample(ctx context.Context, smp *domain.Sample) (*domain.Sample, error)

		// norm
		CreateOrUpdateNorm(ctx context.Context, nrm *domain.Norm) (*domain.Norm, error)

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
	defer util.InfoExecutionTime(time.Now(), "app.New", logger)
	app := App{
		Cfg:   cfg,
		log:   logger,
		repo:  repo,
		Store: store,
		locales: []string{
			domain.LocaleEn,
			domain.LocaleRu,
		},
	}

	err := app.loadUserPresets()
	if err != nil {
		return nil, err
	}

	err = app.loadTestPresets()
	if err != nil {
		return nil, err
	}

	err = app.initSamples()
	if err != nil {
		return nil, err
	}

	app.log.Info("finished loading presets", "")

	// init app jobs, caches and preload data (if any)
	go app.UpdateSystemSummaryJob() // TODO: move to jobs?

	return &app, nil
}

func (a *App) IsValidLocale(locale string) bool {
	for _, l := range a.locales {
		if l == locale {
			return true
		}
	}

	return false
}

func (a *App) AreValidLocales(locales []string) bool {
	for _, l := range locales {
		if ok := a.IsValidLocale(l); !ok {
			return false
		}
	}

	return true
}
