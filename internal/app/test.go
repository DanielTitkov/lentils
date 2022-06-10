package app

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/DanielTitkov/lentils/internal/domain"
	"gopkg.in/yaml.v2"
)

func (a *App) CreateOrUpdateTestFromArgs(ctx context.Context, args domain.CreateTestArgs) error {
	return a.repo.CreateOrUpdateTestFromArgs(ctx, &args)
}

func (a *App) GetTestsForLocale(ctx context.Context, locale string) ([]*domain.Test, error) {
	tests, err := a.repo.GetTests(ctx, locale)
	if err != nil {
		return nil, err
	}

	return tests, nil
}

func (a *App) GetTestByCode(ctx context.Context, code string, locale string) (*domain.Test, error) {
	// TODO: check locale
	return a.repo.GetTestByCode(ctx, code, locale)
}

func (a *App) PrepareTest(ctx context.Context, code string, locale string, args *domain.PrepareTestArgs) (*domain.Test, error) {
	if ok := a.IsValidLocale(locale); !ok {
		return nil, fmt.Errorf("got unknown locale: %s", locale)
	}

	test, err := a.repo.GetTestByCode(ctx, code, locale)
	if err != nil {
		return nil, err
	}

	// TODO: probably move to test usecase
	if test.Display.RandomizeOrder {
	} else {
		test.OrderQuestions()
	}

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

// func generateQuestionsEachItem(test *domain.CreateTestArgs) error {
// 	var questions []domain.CreateQuestionArgs
// 	for _, s := range

// 	return nil
// }
