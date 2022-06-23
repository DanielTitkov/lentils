package app

import (
	"context"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/DanielTitkov/lentils/internal/domain"
	"gopkg.in/yaml.v2"
)

func (a *App) BeginTest(ctx context.Context, test *domain.Test) (*domain.Test, error) {
	// change take status to questions
	err := test.Take.Begin()
	if err != nil {
		a.log.Error("begin test failed (begin test)", err)
		return nil, err
	}

	test.Take, err = a.repo.UpdateTake(ctx, test.Take)
	if err != nil {
		a.log.Error("begin test failed (update take)", err)
		return nil, err
	}

	return test, nil
}

func (a *App) EndTest(ctx context.Context, test *domain.Test) (*domain.Test, error) {
	err := test.Take.End()
	if err != nil {
		a.log.Error("end test failed (end test)", err)
		return nil, err
	}

	test.Take, err = a.repo.UpdateTake(ctx, test.Take)
	if err != nil {
		a.log.Error("end test failed (update take)", err)
		return nil, err
	}

	return test, nil
}

func (a *App) CreateOrUpdateTestFromArgs(ctx context.Context, args domain.CreateTestArgs) error {
	return a.repo.CreateOrUpdateTestFromArgs(ctx, &args)
}

func (a *App) GetTestsForLocale(ctx context.Context, locale string) ([]*domain.Test, error) {
	tests, err := a.repo.GetTests(ctx, locale)
	if err != nil {
		a.log.Error("get tests for locale failed", err)
		return nil, err
	}

	return tests, nil
}

func (a *App) GetTestByCode(ctx context.Context, code string, locale string) (*domain.Test, error) {
	// TODO: check locale
	return a.repo.GetTestByCode(ctx, code, locale)
}

func (a *App) PrepareTestResult(ctx context.Context, test *domain.Test, locale string) (*domain.Test, error) {
	test, err := a.repo.GetTakeData(ctx, test.Take, locale)
	if err != nil {
		a.log.Error("prepare test results failed (get data)", err)
		return nil, err
	}

	if err := test.CalculateResult(); err != nil {
		a.log.Error("prepare test results failed (calculate)", err)
		return nil, err
	}

	// save results to db for further use in norm calculation

	return test, nil
}

func (a *App) PrepareTest(ctx context.Context, code string, locale string, args *domain.PrepareTestArgs) (*domain.Test, error) {
	if ok := a.IsValidLocale(locale); !ok {
		return nil, fmt.Errorf("got unknown locale: %s", locale)
	}

	test, err := a.repo.GetTestByCode(ctx, code, locale)
	if err != nil {
		return nil, err
	}

	// TODO: in take is loaded from db, use old seed
	seed := time.Now().Unix()

	takeMeta := make(map[string]interface{})
	takeMeta["session"] = args.Session
	take, err := a.repo.CreateTake(ctx, &domain.Take{
		Seed:   seed,
		UserID: args.UserID,
		TestID: test.ID,
		Meta:   takeMeta,
	})
	if err != nil {
		return nil, err
	}

	// assign take
	test.Take = take

	// establish questions order (random or fixed)
	test.OrderQuestions(seed)

	return test, nil
}

func (a *App) loadTestPresets() error {
	a.log.Info("loading test presets", fmt.Sprint(a.Cfg.Data.Presets.TestPresetsPaths))
	// FIXME: for production MUST NOT FAIL on single wrong test
	for _, path := range a.Cfg.Data.Presets.TestPresetsPaths {
		a.log.Debug("reading from file", path)
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		var test domain.CreateTestArgs
		err = yaml.Unmarshal(data, &test)
		if err != nil {
			return err
		}

		if ok := a.AreValidLocales(test.AvailableLocales); !ok {
			return fmt.Errorf("locales are not valid: %v", test.AvailableLocales)
		}

		if err := test.ValidateTranslations(); err != nil {
			return err
		}

		// switch test.Generate.Method {
		// case domain.GenerateQuestionsNone:
		// case domain.GenerateQuestionsEachItem:
		// default:
		// }

		err = a.CreateOrUpdateTestFromArgs(context.Background(), test)
		if err != nil {
			a.log.Error("failed to load test", err)
			continue
		}

		a.log.Debug("loaded test", fmt.Sprintf("%+v", test.Code))
	}

	return nil
}

func (a *App) SaveTestResults(ctx context.Context, test *domain.Test) error {
	for _, s := range test.Scales {
		if s.Result == nil {
			a.log.Warn("got scale with nil result with is unexpected", fmt.Sprintf("%+v", s))
			continue
		}
	}

	return nil
}

// func generateQuestionsEachItem(test *domain.CreateTestArgs) error {
// 	var questions []domain.CreateQuestionArgs
// 	for _, s := range

// 	return nil
// }
