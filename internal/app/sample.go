package app

import (
	"context"
	"fmt"

	"github.com/DanielTitkov/lentils/internal/domain"
)

func (a *App) initSamples() error {
	samples := []*domain.Sample{
		{
			Code:     domain.SampleAllCode,
			Criteria: domain.SampleCriteria{},
		},
		{
			Code: domain.SampleAllNonSuspiciousCode,
			Criteria: domain.SampleCriteria{
				NotSuspicious: true,
			},
		},
	}

	for _, s := range samples {
		smp, err := a.repo.CreateOrUpdateSample(context.TODO(), s)
		a.log.Info("created sample", fmt.Sprintf("%+v", smp))
		if err != nil {
			return err
		}
	}

	return nil
}
