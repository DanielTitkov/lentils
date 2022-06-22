package entgo

import (
	"context"

	"github.com/google/uuid"

	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/sample"

	"github.com/DanielTitkov/lentils/internal/domain"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/norm"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/scale"
)

func (r *EntgoRepository) CreateOrUpdateNorm(ctx context.Context, nrm *domain.Norm) (*domain.Norm, error) {
	n, err := r.client.Norm.Query().
		Where(
			norm.HasScaleWith(scale.IDEQ(nrm.ScaleID)),
			norm.HasSampleWith(sample.IDEQ(nrm.SampleID)),
		).
		Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return nil, err
		}

		// norm not found, create norm
		n, err := r.client.Norm.Create().
			SetScaleID(nrm.ScaleID).
			SetSampleID(nrm.SampleID).
			SetName(nrm.Name).
			SetBase(nrm.Base).
			SetMean(nrm.Mean).
			SetSigma(nrm.Sigma).
			Save(ctx)
		if err != nil {
			return nil, err
		}

		return entToDomainNorm(n, nrm.ScaleID, nrm.SampleID), nil
	}

	// update sample
	n, err = n.Update().
		SetName(nrm.Name).
		SetBase(nrm.Base).
		SetMean(nrm.Mean).
		SetSigma(nrm.Sigma).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return entToDomainNorm(n, nrm.ScaleID, nrm.SampleID), nil
}

func entToDomainNorm(n *ent.Norm, scaleID, sampleID uuid.UUID) *domain.Norm {
	if n.Edges.Scale != nil {
		scaleID = n.Edges.Scale.ID
	}

	if n.Edges.Sample != nil {
		sampleID = n.Edges.Sample.ID
	}

	return &domain.Norm{
		ID:       n.ID,
		ScaleID:  scaleID,
		SampleID: sampleID,
		Name:     n.Name,
		Base:     n.Base,
		Mean:     n.Mean,
		Sigma:    n.Sigma,
		Meta:     n.Meta,
	}
}
