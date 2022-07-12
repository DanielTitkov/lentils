package app

import (
	"context"
	"time"

	"github.com/DanielTitkov/orrery/internal/util"

	"github.com/DanielTitkov/orrery/internal/configs"
	"github.com/DanielTitkov/orrery/internal/domain"
	"github.com/DanielTitkov/orrery/internal/logger"
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
		Errors        []error
		Events        []domain.Event
	}
	Repository interface {
		// user
		IfEmailRegistered(context.Context, string) (bool, error)
		GetUserByEmail(context.Context, string) (*domain.User, error)
		GetUserByID(context.Context, uuid.UUID) (*domain.User, error)
		CreateUser(context.Context, *domain.User) (*domain.User, error)
		UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error)

		// user session
		IfSessionRegistered(context.Context, *domain.UserSession) (bool, error)
		CreateOrUpdateUserSession(context.Context, *domain.UserSession) (*domain.UserSession, error)
		UpdateUserSessionLastActivityBySID(context.Context, string) error
		GetUserBySession(context.Context, *domain.UserSession) (*domain.User, error)

		// test
		CreateOrUpdateTestFromArgs(context.Context, *domain.CreateTestArgs) error
		GetTests(ctx context.Context, args *domain.QueryTestsArgs) ([]*domain.Test, error)
		GetTestByCode(ctx context.Context, code string, locale string) (*domain.Test, error)
		GetTakeData(ctx context.Context, take *domain.Take, locale string) (*domain.Test, error)
		GetDataForMarkCalculation(ctx context.Context) ([]*domain.Test, error)
		UpdateTestMark(ctx context.Context, testID uuid.UUID, mark float64) error

		// take
		GetTake(ctx context.Context, takeID uuid.UUID) (*domain.Take, error)
		CreateTake(ctx context.Context, take *domain.Take) (*domain.Take, error)
		UpdateTake(ctx context.Context, take *domain.Take) (*domain.Take, error)
		UpdateTakeMark(ctx context.Context, takeID uuid.UUID, mark int) error

		// response
		AddOrUpdateResponse(ctx context.Context, takeID uuid.UUID, itm *domain.Item) (*domain.Response, error)

		// sample
		CreateOrUpdateSample(ctx context.Context, sample *domain.Sample) (*domain.Sample, error)
		GetSamples(ctx context.Context) ([]*domain.Sample, error)

		// norm
		CreateOrUpdateNorm(ctx context.Context, norm *domain.Norm) (*domain.Norm, error)
		GetDataForNormCalculation(ctx context.Context, criteria domain.SampleCriteria) ([]*domain.NormCalculationData, error)
		GetScaleNorms(ctx context.Context, scaleID uuid.UUID, sampleIDs []uuid.UUID) ([]*domain.Norm, error)

		// result
		CreateOrUpdateResult(ctx context.Context, result *domain.Result) (*domain.Result, error)

		// for system summary
		GetUserCount(ctx context.Context) (int, error)

		// tag
		CreateOrUpdateTagFromArgs(ctx context.Context, args *domain.CreateTagArgs) error
		GetTagsByCodes(ctx context.Context, locale string, codes ...string) ([]*domain.Tag, error)
		GetTagIDsByCodes(ctx context.Context, codes ...string) ([]uuid.UUID, error)
	}
)

func New(
	cfg configs.Config,
	logger *logger.Logger,
	repo Repository,
	store sessions.Store,
) (*App, error) {
	start := time.Now()
	defer util.InfoExecutionTime(start, "app.New", logger)
	app := App{
		Cfg:     cfg,
		log:     logger,
		repo:    repo,
		Store:   store,
		locales: domain.Locales(),
	}

	err := app.loadTagPresets()
	if err != nil {
		app.log.Error("errors while loading tags", err)
		app.AddError(err)
		if app.IsDev() {
			return nil, err
		}
	}

	err = app.loadTestPresets()
	if err != nil {
		app.log.Error("errors while loading tests", err)
		app.AddError(err)
		if app.IsDev() {
			return nil, err
		}
	}

	app.log.Info("finished loading presets", "")

	err = app.initSamples()
	if err != nil {
		app.log.Error("errors while creating samples", err)
		app.AddError(err)
		if app.IsDev() {
			return nil, err
		}
	}

	app.AddEvent("app.New", start)

	return &app, nil
}

func (a *App) IsDev() bool {
	return a.Cfg.Env == domain.EnvDev
}

func (a *App) AddError(err error) {
	a.Errors = append(a.Errors, err)
}

func (a *App) AddEvent(name string, start time.Time) {
	if len(a.Events) >= domain.AppMaxEvents {
		a.Events = a.Events[:len(a.Events)-1]
	}
	// prepend event to show in reverse order
	a.Events = append([]domain.Event{domain.Event{
		Name:      name,
		StartTime: start,
		EndTime:   time.Now(),
		Elapsed:   time.Since(start),
	}}, a.Events...)
}
