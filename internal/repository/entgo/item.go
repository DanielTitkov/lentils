package entgo

import (
	"github.com/tinygodsdev/orrery/internal/domain"
	"github.com/tinygodsdev/orrery/internal/repository/entgo/ent"
	"github.com/google/uuid"
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

	var response *domain.Response
	if i.Edges.Responses != nil {
		if len(i.Edges.Responses) == 1 {
			response = entToDomainResponse(i.Edges.Responses[0], uuid.Nil, i.ID)
		}
	}

	return &domain.Item{
		ID:       i.ID,
		Code:     i.Code,
		Steps:    i.Steps,
		Reverse:  reverse,
		Content:  content,
		Response: response,
	}
}
