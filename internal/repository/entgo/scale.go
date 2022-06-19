package entgo

import (
	"github.com/DanielTitkov/lentils/internal/domain"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent"
)

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
