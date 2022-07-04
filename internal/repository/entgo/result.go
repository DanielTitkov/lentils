package entgo

import (
	"context"

	"github.com/DanielTitkov/orrery/internal/domain"
	"github.com/DanielTitkov/orrery/internal/repository/entgo/ent"
	"github.com/DanielTitkov/orrery/internal/repository/entgo/ent/result"
	"github.com/DanielTitkov/orrery/internal/repository/entgo/ent/scale"
	"github.com/DanielTitkov/orrery/internal/repository/entgo/ent/take"
	"github.com/google/uuid"
)

func (r *EntgoRepository) CreateOrUpdateResult(ctx context.Context, res *domain.Result) (*domain.Result, error) {
	rslt, err := r.client.Result.Query().
		Where(
			result.HasScaleWith(scale.IDEQ(res.ScaleID)),
			result.HasTakeWith(take.IDEQ(res.TakeID)),
		).
		Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return nil, err
		}

		// result not found, create result
		rslt, err := r.client.Result.Create().
			SetScaleID(res.ScaleID).
			SetTakeID(res.TakeID).
			SetRawScore(res.RawScore).
			SetFinalScore(res.FinalScore).
			SetMeta(res.Meta).
			Save(ctx)
		if err != nil {
			return nil, err
		}

		return entToDomainResult(rslt, res.ScaleID, res.TakeID), nil
	}

	// update result
	rslt, err = rslt.Update().
		SetRawScore(res.RawScore).
		SetFinalScore(res.FinalScore).
		SetMeta(res.Meta).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return entToDomainResult(rslt, res.ScaleID, res.TakeID), nil
}

func entToDomainResult(r *ent.Result, scaleID, takeID uuid.UUID) *domain.Result {
	if r.Edges.Scale != nil {
		scaleID = r.Edges.Scale.ID
	}

	if r.Edges.Take != nil {
		takeID = r.Edges.Take.ID
	}

	return &domain.Result{
		ID:         r.ID,
		ScaleID:    scaleID,
		TakeID:     takeID,
		RawScore:   r.RawScore,
		FinalScore: r.FinalScore,
		CreateTime: r.CreateTime,
		UpdateTime: r.UpdateTime,
		Meta:       r.Meta,
	}
}
