package entgo

import (
	"context"

	"github.com/DanielTitkov/orrery/internal/repository/entgo/ent/result"
	"github.com/DanielTitkov/orrery/internal/repository/entgo/ent/take"

	"github.com/DanielTitkov/orrery/internal/domain"
	"github.com/DanielTitkov/orrery/internal/repository/entgo/ent"
)

// GetDataForNormCalculation return data for 1 norm for all scales
func (r *EntgoRepository) GetDataForNormCalculation(ctx context.Context, crit domain.SampleCriteria) ([]*domain.NormCalculationData, error) {
	scales, err := r.client.Scale.Query().
		WithResults(func(q *ent.ResultQuery) {
			if crit.NotSuspicious {
				q.Where(result.HasTakeWith(take.SuspiciousEQ(false)))
			}

			if crit.Locale != "" {
				q.Where(result.HasTakeWith(take.InLocaleEQ(take.InLocale(crit.Locale))))
			}
		}).
		All(ctx)
	if err != nil {
		return nil, err
	}

	var res []*domain.NormCalculationData
	for _, scale := range scales {
		res = append(res, entToDomainNormCalculationData(scale))
	}

	return res, nil
}

func entToDomainNormCalculationData(s *ent.Scale) *domain.NormCalculationData {
	var results []float64
	if s.Edges.Results != nil {
		for _, r := range s.Edges.Results {
			results = append(results, r.RawScore)
		}
	}

	return &domain.NormCalculationData{
		ScaleID:   s.ID,
		ScaleCode: s.Code,
		Results:   results,
	}
}

func entToDomainScale(s *ent.Scale, locale string) *domain.Scale {
	title := "no title for this locale: " + locale
	description := "no description for this locale: " + locale
	abbreviation := "no abbreviation for this locale: " + locale

	if s.Edges.Translations != nil {
		if len(s.Edges.Translations) == 1 {
			trans := s.Edges.Translations[0]
			title = trans.Title
			description = trans.Description
			abbreviation = trans.Abbreviation
		}
	}

	var interpretations []*domain.Interpretation
	if s.Edges.Interpretations != nil {
		for _, in := range s.Edges.Interpretations {
			interpretations = append(interpretations, entToDomainInterpretation(in, locale))
		}
	}

	var items []*domain.Item
	if s.Edges.Items != nil {
		for _, itm := range s.Edges.Items {
			items = append(items, entToDomainItem(itm, locale))
		}
	}

	return &domain.Scale{
		ID:              s.ID,
		Code:            s.Code,
		Type:            s.Type.String(),
		Global:          s.Global,
		Title:           title,
		Description:     description,
		Abbreviation:    abbreviation,
		Interpretations: interpretations,
		Items:           items,
	}
}
