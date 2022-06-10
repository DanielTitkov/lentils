package entgo

import (
	"github.com/DanielTitkov/lentils/internal/domain"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent"
)

func entToDomainQuestion(q *ent.Question, locale string) *domain.Question {
	content := "no content for this locale: " + locale
	headerContent := "no header content for this locale: " + locale
	footerContent := "no footer content for this locale: " + locale

	if q.Edges.Translations != nil {
		if len(q.Edges.Translations) == 1 {
			trans := q.Edges.Translations[0]
			content = trans.Content
			headerContent = trans.HeaderContent
			footerContent = trans.FooterContent
		}
	}

	var items []*domain.Item
	if q.Edges.Items != nil {
		for _, itm := range q.Edges.Items {
			items = append(items, entToDomainItem(itm, locale))
		}
	}

	return &domain.Question{
		ID:            q.ID,
		Code:          q.Code,
		Order:         q.Order,
		Content:       content,
		HeaderContent: headerContent,
		FooterContent: footerContent,
		Items:         items,
	}
}
