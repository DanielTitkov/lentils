package entgo

import (
	"github.com/DanielTitkov/lentils/internal/domain"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent"
)

func entToDomainItem(i *ent.Item, locale string) *domain.Item {
	content := "no content for this locale: " + locale

	if i.Edges.Translations != nil {
		if len(i.Edges.Translations) == 1 {
			trans := i.Edges.Translations[0]
			content = trans.Content
		}
	}

	var reverse bool
	if i.Edges.ScaleItem != nil {
		if len(i.Edges.ScaleItem) == 1 {
			reverse = i.Edges.ScaleItem[0].Reverse
		}
	}

	return &domain.Item{
		ID:      i.ID,
		Code:    i.Code,
		Steps:   i.Steps,
		Reverse: reverse,
		Content: content,
	}
}
