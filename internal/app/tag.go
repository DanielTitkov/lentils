package app

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/DanielTitkov/lentils/internal/domain"
	"gopkg.in/yaml.v2"
)

func (a *App) CreateOrUpdateTagFromArgs(ctx context.Context, args domain.CreateTagArgs) error {
	return a.repo.CreateOrUpdateTagFromArgs(ctx, &args)
}

func (a *App) GetTags(ctx context.Context, locale string) ([]*domain.Tag, error) {
	return a.repo.GetTagsByCodes(ctx, locale)
}

func (a *App) loadTagPresets() error {
	a.log.Info("loading tag presets", fmt.Sprint(a.Cfg.Data.Presets.TestPresetsPaths))
	for _, path := range a.Cfg.Data.Presets.TagPresetsPaths {
		a.log.Debug("reading from file", path)
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		var tags []domain.CreateTagArgs
		err = yaml.Unmarshal(data, &tags)
		if err != nil {
			return err
		}

		for _, tag := range tags {
			if err := tag.ValidateTranslations(); err != nil {
				return err
			}

			err = a.CreateOrUpdateTagFromArgs(context.Background(), tag)
			if err != nil {
				a.log.Error("failed to load tag", err)
				continue
			}

			a.log.Debug("loaded tag", fmt.Sprintf("%+v", tag.Code))
		}

	}

	return nil
}
