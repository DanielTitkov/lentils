package entgo

import (
	"context"

	"github.com/DanielTitkov/lentils/internal/domain"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent"
)

func (r *EntgoRepository) GetDataForNormCalculation(ctx context.Context, crit domain.SampleCriteria) error {

	// select s.id as scale_id, t.id as take_id, r.value, si.reverse, t.status from scales s
	// join scale_items si on si.scale_id = s.id
	// join responses r on si.item_id  = r.item_responses
	// join takes t on r.take_responses = t.id
	// where t.status = 'finish'
	// order by s.id, t.id

	return nil
}

func entToDomainScale(s *ent.Scale, locale string) *domain.Scale {
	title := "no title for this locale: " + locale
	description := "no description for this locale: " + locale

	if s.Edges.Translations != nil {
		if len(s.Edges.Translations) == 1 {
			trans := s.Edges.Translations[0]
			title = trans.Title
			description = trans.Description
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
		Interpretations: interpretations,
		Items:           items,
	}
}
