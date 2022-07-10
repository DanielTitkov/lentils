package app

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/multierr"

	"github.com/DanielTitkov/orrery/internal/domain"

	"github.com/DanielTitkov/orrery/internal/util"
	"github.com/montanaflynn/stats"
)

func (a *App) UpdateNorms(ctx context.Context) (errs error) {
	defer util.InfoExecutionTime(time.Now(), "app.UpdateNorms", a.log)
	// get all samples
	samples, err := a.repo.GetSamples(ctx)
	if err != nil {
		a.log.Error("failed to get samples", err)
		if a.IsDev() {
			return err
		} else {
			errs = multierr.Append(errs, err)
		}
	}

	a.log.Info("loaded samples", fmt.Sprintf("%d samples in total", len(samples)))

	// for each sample generate norm for each scale
	for _, sample := range samples {
		a.log.Info("processing sample", fmt.Sprintf("%+v", sample))
		data, err := a.repo.GetDataForNormCalculation(ctx, sample.Criteria)
		if err != nil {
			a.log.Error("failed to get sample data", err)
			if a.IsDev() {
				return err
			} else {
				errs = multierr.Append(errs, err)
				continue
			}
		}
		a.log.Debug("loaded data for sample", fmt.Sprintf("%d results", len(data)))
		for _, scale := range data {
			if len(scale.Results) == 0 {
				continue
			}

			base := len(scale.Results)
			mean, err := stats.Mean(scale.Results)
			if err != nil {
				a.log.Error("failed to calculate mean for sample", err)
				if a.IsDev() {
					return err
				} else {
					errs = multierr.Append(errs, err)
					continue
				}
			}
			sd, err := stats.StandardDeviationSample(scale.Results)
			if err != nil {
				a.log.Error("failed to calculate sd for sample", err)
				if a.IsDev() {
					return err
				} else {
					errs = multierr.Append(errs, err)
					continue
				}
			}

			norm := &domain.Norm{
				SampleID: sample.ID,
				ScaleID:  scale.ScaleID,
				Name:     fmt.Sprintf("%s__%s", sample.Code, scale.ScaleCode),
				Base:     base,
				Mean:     mean,
				Sigma:    sd,
			}

			rank := sample.NormRank(norm)
			norm.Rank = rank

			_, err = a.repo.CreateOrUpdateNorm(ctx, norm)
			if err != nil {
				a.log.Error("failed to save norm", err)
				if a.IsDev() {
					return err
				} else {
					errs = multierr.Append(errs, err)
					continue
				}
			}
			// a.log.Debug("updated norm", fmt.Sprintf("%+v", norm)) // FIXME
		}
	}

	return errs
}
