package entgo

import (
	"context"
	"fmt"

	"github.com/DanielTitkov/orrery/internal/domain"
	"github.com/DanielTitkov/orrery/internal/repository/entgo/ent"
	"github.com/DanielTitkov/orrery/internal/repository/entgo/ent/tag"
	"github.com/DanielTitkov/orrery/internal/repository/entgo/ent/tagtranslation"
	"github.com/google/uuid"
)

func (r *EntgoRepository) GetTagsByCodes(ctx context.Context, locale string, codes ...string) ([]*domain.Tag, error) {
	query := r.client.Tag.Query().
		WithTranslations(func(q *ent.TagTranslationQuery) {
			q.Where(tagtranslation.LocaleEQ(tagtranslation.Locale(locale)))
		})
	if len(codes) != 0 {
		query.Where(tag.CodeIn(codes...))
	}
	tags, err := query.All(ctx)
	if err != nil {
		return nil, err
	}

	var res []*domain.Tag
	for _, t := range tags {
		res = append(res, entToDomainTag(t, locale))
	}

	return res, nil
}

func (r *EntgoRepository) GetTagIDsByCodes(ctx context.Context, codes ...string) ([]uuid.UUID, error) {
	query := r.client.Tag.Query()
	if len(codes) != 0 {
		query.Where(tag.CodeIn(codes...))
	}
	return query.IDs(ctx)
}

func (r *EntgoRepository) CreateOrUpdateTagFromArgs(ctx context.Context, args *domain.CreateTagArgs) error {
	tx, err := r.client.Tx(ctx)
	if err != nil {
		return fmt.Errorf("starting a transaction: %w", err)
	}

	// check if tag exists by code
	tg, err := tx.Tag.Query().Where(tag.CodeEQ(args.Code)).Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return rollback(tx, err)
		}

		// tag not found, create tag
		tg, err = tx.Tag.Create().
			SetCode(args.Code).
			SetType(tag.Type(args.Type)).
			Save(ctx)
		if err != nil {
			return rollback(tx, err)
		}
	} else {
		// update tag
		tg, err = tg.Update().
			SetType(tag.Type(args.Type)).
			Save(ctx)
		if err != nil {
			return rollback(tx, err)
		}
	}

	// delete old translations if exist
	// TODO: maybe change this to bulk upsert
	_, err = tx.TagTranslation.Delete().
		Where(tagtranslation.HasTagWith(tag.IDEQ(tg.ID))).
		Exec(ctx)
	if err != nil {
		return rollback(tx, err)
	}

	// create tag translations
	// this happens only on start time, so time doesn't matter
	// and thus bulk is not used
	for _, t := range args.Translations {
		_, err = tx.TagTranslation.Create().
			SetLocale(tagtranslation.Locale(t.Locale)).
			SetContent(t.Content).
			SetTagID(tg.ID).
			Save(ctx)
		if err != nil {
			return rollback(tx, err)
		}
	}

	return tx.Commit()
}

func entToDomainTag(t *ent.Tag, locale string) *domain.Tag {
	content := "no content for this locale: " + locale

	if t.Edges.Translations != nil {
		if len(t.Edges.Translations) == 1 {
			trans := t.Edges.Translations[0]
			content = trans.Content
		}
	}

	return &domain.Tag{
		ID:      t.ID,
		Code:    t.Code,
		Type:    t.Type.String(),
		Content: content,
	}
}
