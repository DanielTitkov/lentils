package app

import (
	"context"
	"fmt"
	"time"

	"github.com/DanielTitkov/lentils/internal/domain"

	"github.com/DanielTitkov/lentils/internal/util"
	"github.com/montanaflynn/stats"
)

func (a *App) UpdateNorms(ctx context.Context) error {
	defer util.InfoExecutionTime(time.Now(), "app.UpdateNorms", a.log)
	// get all samples
	samples, err := a.repo.GetSamples(ctx)
	if err != nil {
		a.log.Error("failed to get samples", err)
		return err
	}

	a.log.Info("loaded samples", fmt.Sprintf("%d samples in total", len(samples)))

	// for each sample generate norm for each scale
	for _, sample := range samples {
		a.log.Info("processing sample", fmt.Sprintf("%+v", sample))
		data, err := a.repo.GetDataForNormCalculation(ctx, sample.Criteria)
		if err != nil {
			a.log.Error("failed to get sample data", err)
			return err
		}
		a.log.Debug("loaded data for sample", fmt.Sprintf("%+v", data))
		for _, scale := range data {
			if len(scale.Results) == 0 {
				continue
			}

			base := len(scale.Results)
			mean, err := stats.Mean(scale.Results)
			if err != nil {
				a.log.Error("failed to calculate mean for sample", err)
				return err
			}
			sd, err := stats.StandardDeviationSample(scale.Results)
			if err != nil {
				a.log.Error("failed to calculate sd for sample", err)
				return err
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

			norm, err = a.repo.CreateOrUpdateNorm(ctx, norm)
			if err != nil {
				a.log.Error("failed to save norm", err)
				return err
			}
			a.log.Debug("updated norm", fmt.Sprintf("%+v", norm))
		}
	}

	return nil
}
