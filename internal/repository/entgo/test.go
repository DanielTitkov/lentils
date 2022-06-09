package entgo

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/item"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/itemtranslation"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/scale"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/scaletranslation"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/test"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/testtranslation"

	"github.com/DanielTitkov/lentils/internal/domain"
)

func (r *EntgoRepository) GetTests(ctx context.Context, locale string) ([]*domain.Test, error) {
	tests, err := r.client.Test.Query().
		WithTranslations(
			func(q *ent.TestTranslationQuery) {
				q.Where(testtranslation.LocaleEQ(testtranslation.Locale(locale)))
			}).
		All(ctx)
	if err != nil {
		return nil, err
	}

	var res []*domain.Test
	for _, t := range tests {
		res = append(res, entToDomainTest(t, locale))
	}

	return res, nil
}

// TODO: function too long, refactor please
func (r *EntgoRepository) CreateOrUpdateTestFromArgs(ctx context.Context, args *domain.CreateTestArgs) error {
	tx, err := r.client.Tx(ctx)
	if err != nil {
		return fmt.Errorf("starting a transaction: %w", err)
	}

	// check if test exists by code
	tst, err := tx.Test.Query().Where(test.CodeEQ(args.Code)).Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return rollback(tx, err)
		}

		// test not found, create test
		tst, err = tx.Test.Create().
			SetCode(args.Code).
			SetPublished(args.Published).
			Save(ctx)
		if err != nil {
			return rollback(tx, err)
		}
	} else {
		// test exists, update
		tst, err = tst.Update().
			SetPublished(args.Published).
			// clear edges
			ClearScales().
			Save(ctx)
		if err != nil {
			return rollback(tx, err)
		}
	}

	// delete old translations if exist
	// TODO: maybe change this to bulk upsert
	_, err = tx.TestTranslation.Delete().
		Where(testtranslation.HasTestWith(test.IDEQ(tst.ID))).
		Exec(ctx)
	if err != nil {
		return rollback(tx, err)
	}

	// create test translations
	// this happens only on start time, so time doesn't matter
	// and thus bulk is not used
	for _, t := range args.Translations {
		_, err = tx.TestTranslation.Create().
			SetLocale(testtranslation.Locale(t.Locale)).
			SetTitle(t.Title).
			SetDescription(t.Description).
			SetInstruction(t.Instruction).
			SetTestID(tst.ID).
			Save(ctx)
		if err != nil {
			return rollback(tx, err)
		}
	}

	// create or update scales for test
	updateTst := tst.Update()
	for _, sArgs := range args.Scales {
		// check if scale exists by code
		// TODO: what if different tests use one scale?
		scl, err := tx.Scale.Query().Where(scale.CodeEQ(sArgs.Code)).Only(ctx)
		if err != nil {
			if !ent.IsNotFound(err) {
				return rollback(tx, err)
			}

			// scale not found, create scale
			scl, err = tx.Scale.Create().
				SetCode(sArgs.Code).
				SetType(scale.Type(sArgs.Type)).
				Save(ctx)
			if err != nil {
				return rollback(tx, err)
			}
		} else {
			// scale exists, update if allowed
			if scl.Global {
				r.logger.Error("trying to update global scale", errors.New("not allowed to update global scale from test constructor"))
				updateTst.AddScaleIDs(scl.ID)
				continue
			}

			scl, err = scl.Update().
				SetType(scale.Type(sArgs.Type)).
				ClearItems().
				Save(ctx)
			if err != nil {
				return rollback(tx, err)
			}
		}

		// delete old translations if exist
		// TODO: maybe change this to bulk upsert
		_, err = tx.ScaleTranslation.Delete().
			Where(scaletranslation.HasScaleWith(scale.IDEQ(scl.ID))).
			Exec(ctx)
		if err != nil {
			return rollback(tx, err)
		}

		// create scale translations
		// this happens only on start time, so time doesn't matter
		// and thus bulk is not used
		for _, t := range sArgs.Translations {
			_, err = tx.ScaleTranslation.Create().
				SetLocale(scaletranslation.Locale(t.Locale)).
				SetTitle(t.Title).
				SetDescription(t.Description).
				SetScaleID(scl.ID).
				Save(ctx)
			if err != nil {
				return rollback(tx, err)
			}
		}

		// create and add items for scale
		for _, iArgs := range sArgs.Items {
			// check if item exists by code
			itm, err := tx.Item.Query().Where(item.CodeEQ(iArgs.Code)).Only(ctx)
			if err != nil {
				if !ent.IsNotFound(err) {
					return rollback(tx, err)
				}

				// scale not found, create scale
				itm, err = tx.Item.Create().
					SetCode(iArgs.Code).
					SetSteps(iArgs.Steps).
					Save(ctx)
				if err != nil {
					return rollback(tx, err)
				}
			} else {
				// items exists, update if allowed
				itm, err = itm.Update().
					SetSteps(iArgs.Steps).
					Save(ctx)
				if err != nil {
					return rollback(tx, err)
				}
			}

			// delete old translations if exist
			// TODO: maybe change this to bulk upsert
			_, err = tx.ItemTranslation.Delete().
				Where(itemtranslation.HasItemWith(item.IDEQ(itm.ID))).
				Exec(ctx)
			if err != nil {
				return rollback(tx, err)
			}

			// create scale translations
			// this happens only on start time, so time doesn't matter
			// and thus bulk is not used
			for _, t := range iArgs.Translations {
				_, err = tx.ItemTranslation.Create().
					SetLocale(itemtranslation.Locale(t.Locale)).
					SetContent(t.Content).
					SetItemID(itm.ID).
					Save(ctx)
				if err != nil {
					return rollback(tx, err)
				}
			}

			// create item-scale edge
			_, err = tx.ScaleItem.Create().
				SetItemID(itm.ID).
				SetScaleID(scl.ID).
				SetReverse(iArgs.Reverse).
				Save(ctx)
			if err != nil {
				return rollback(tx, err)
			}
			// finished creating an item
		}

		// add scale to test
		updateTst.AddScaleIDs(scl.ID)
		// finished creating a scale
	}

	// save test updates (adds scales with items and stuff)
	_, err = updateTst.Save(ctx)
	if err != nil {
		return rollback(tx, err)
	}

	return tx.Commit()
}

func entToDomainTest(t *ent.Test, locale string) *domain.Test {
	title := "no title for this locale: " + locale
	description := "no description for this locale: " + locale
	instruction := "no instructuin for this locale: " + locale

	if t.Edges.Translations != nil {
		if len(t.Edges.Translations) == 1 {
			trans := t.Edges.Translations[0]
			title = trans.Title
			description = trans.Description
			instruction = trans.Instruction
		} else {
			log.Println("got multiple translations for test, something is wrong") // FIXME
		}
	}

	return &domain.Test{
		ID:          t.ID,
		Code:        t.Code,
		Published:   t.Published,
		Title:       title,
		Description: description,
		Instruction: instruction,
	}
}
