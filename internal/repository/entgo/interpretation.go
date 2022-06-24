package entgo

import (
	"html/template"

	"github.com/DanielTitkov/lentils/internal/domain"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent"
)

func entToDomainInterpretation(in *ent.Interpretation, locale string) *domain.Interpretation {
	content := "no content for this locale: " + locale

	if in.Edges.Translations != nil {
		if len(in.Edges.Translations) == 1 {
			trans := in.Edges.Translations[0]
			content = trans.Content
		}
	}

	return &domain.Interpretation{
		ID:      in.ID,
		Range:   in.Range,
		Content: template.HTML(content),
	}
}
