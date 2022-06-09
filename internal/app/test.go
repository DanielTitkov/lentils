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

func (a *App) loadTestPresets() error {
	a.log.Info("loading test presets", fmt.Sprint(a.Cfg.Data.Presets.TestPresetsPaths))
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

		ctx := context.Background()

		err = a.CreateOrUpdateTestFromArgs(ctx, test)
		if err != nil {
			a.log.Error("failed to load test", err)
			continue
		}

		a.log.Debug("loaded test", fmt.Sprintf("%+v", test))
	}

	return nil
}
