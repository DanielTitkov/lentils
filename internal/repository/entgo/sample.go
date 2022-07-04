package entgo

import (
	"context"

	"github.com/DanielTitkov/orrery/internal/domain"
	"github.com/DanielTitkov/orrery/internal/repository/entgo/ent"
	"github.com/DanielTitkov/orrery/internal/repository/entgo/ent/sample"
)

func (r *EntgoRepository) GetSamples(ctx context.Context) ([]*domain.Sample, error) {
	samples, err := r.client.Sample.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	var res []*domain.Sample
	for _, s := range samples {
		res = append(res, entToDomainSample(s))
	}

	return res, nil
}

func (r *EntgoRepository) CreateOrUpdateSample(ctx context.Context, smp *domain.Sample) (*domain.Sample, error) {
	s, err := r.client.Sample.Query().
		Where(sample.CodeEQ(smp.Code)).
		Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return nil, err
		}

		// sample not found, create sample
		s, err := r.client.Sample.Create().
			SetCode(smp.Code).
			SetCriteria(smp.Criteria).
			Save(ctx)
		if err != nil {
			return nil, err
		}

		return entToDomainSample(s), nil
	}

	// update sample
	s, err = s.Update().
		SetCriteria(smp.Criteria).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return entToDomainSample(s), nil
}

func entToDomainSample(s *ent.Sample) *domain.Sample {
	return &domain.Sample{
		ID:       s.ID,
		Code:     s.Code,
		Criteria: s.Criteria,
	}
}
