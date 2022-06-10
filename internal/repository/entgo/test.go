package entgo

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/item"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/itemtranslation"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/question"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/questiontranslation"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/scale"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/scaletranslation"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/test"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/testdisplay"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/testtranslation"
	"github.com/DanielTitkov/lentils/internal/util"

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

func (r *EntgoRepository) GetTestByCode(ctx context.Context, code string, locale string) (*domain.Test, error) {
	tst, err := r.client.Test.Query().
		Where(test.CodeEQ(code)).
		WithDisplay().
		WithQuestions(
			func(q *ent.QuestionQuery) {
				q.WithTranslations(
					func(tq *ent.QuestionTranslationQuery) {
						tq.Where(questiontranslation.LocaleEQ(questiontranslation.Locale(locale)))
					},
				).WithItems(
					func(iq *ent.ItemQuery) {
						iq.WithTranslations(
							func(itq *ent.ItemTranslationQuery) {
								itq.Where(itemtranslation.LocaleEQ(itemtranslation.Locale(locale)))
							},
						)
					},
				)
			},
		).
		WithTranslations(
			func(q *ent.TestTranslationQuery) {
				q.Where(testtranslation.LocaleEQ(testtranslation.Locale(locale)))
			},
		).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	// TODO

	return entToDomainTest(tst, locale), nil
}

// TODO: function too long, refactor please
func (r *EntgoRepository) CreateOrUpdateTestFromArgs(ctx context.Context, args *domain.CreateTestArgs) error {
	defer util.DebugExecutionTime(time.Now(), "entgo.CreateOrUpdateTestFromArgs", r.logger)
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
			SetAvailableLocales(args.AvailableLocales).
			Save(ctx)
		if err != nil {
			return rollback(tx, err)
		}
	} else {
		// test exists, update
		tst, err = tst.Update().
			SetPublished(args.Published).
			SetAvailableLocales(args.AvailableLocales).
			// clear edges
			ClearScales().
			ClearQuestions().
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

	// update test display
	// delete old display if exist
	// TODO: maybe change this to upsert
	_, err = tx.TestDisplay.Delete().
		Where(testdisplay.HasTestWith(test.IDEQ(tst.ID))).
		Exec(ctx)
	if err != nil {
		return rollback(tx, err)
	}

	_, err = tx.TestDisplay.Create().
		SetTestID(tst.ID).
		SetQuestionsPerPage(args.Display.QuestionsPerPage).
		SetRandomizeOrder(args.Display.RandomizeOrder).
		Save(ctx)
	if err != nil {
		return rollback(tx, err)
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

				// item not found, create item
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

			// create item translations
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

	// add questions for test
	for _, qArgs := range args.Questions {
		// check if question exists by code
		q, err := tx.Question.Query().Where(question.CodeEQ(qArgs.Code)).Only(ctx)
		if err != nil {
			if !ent.IsNotFound(err) {
				return rollback(tx, err)
			}

			// scale not found, create scale
			q, err = tx.Question.Create().
				SetCode(qArgs.Code).
				SetOrder(qArgs.Order).
				SetType(question.Type(qArgs.Type)).
				Save(ctx)
			if err != nil {
				return rollback(tx, err)
			}
		} else {
			// question exists, update
			q, err = q.Update().
				SetType(question.Type(qArgs.Type)).
				SetOrder(qArgs.Order).
				ClearItems().
				Save(ctx)
			if err != nil {
				return rollback(tx, err)
			}
		}

		// delete old translations if exist
		// TODO: maybe change this to bulk upsert
		_, err = tx.QuestionTranslation.Delete().
			Where(questiontranslation.HasQuestionWith(question.IDEQ(q.ID))).
			Exec(ctx)
		if err != nil {
			return rollback(tx, err)
		}

		// create question translations
		// this happens only on start time, so time doesn't matter
		// and thus bulk is not used
		for _, t := range qArgs.Translations {
			_, err = tx.QuestionTranslation.Create().
				SetLocale(questiontranslation.Locale(t.Locale)).
				SetContent(t.Content).
				SetHeaderContent(t.HeaderContent).
				SetFooterContent(t.FooterConent).
				SetQuestionID(q.ID).
				Save(ctx)
			if err != nil {
				return rollback(tx, err)
			}
		}

		// add items for question
		updateQuestion := q.Update()
		for _, iArgs := range qArgs.Items {
			// item for question must exist.
			// not allowed to create items without scale
			itm, err := tx.Item.Query().Where(item.CodeEQ(iArgs.Code)).Only(ctx)
			if err != nil {
				return rollback(tx, err)
			}

			updateQuestion.AddItemIDs(itm.ID)
			// finished adding an item to question
		}
		q, err = updateQuestion.Save(ctx)
		if err != nil {
			return rollback(tx, err)
		}

		// add scale to test
		updateTst.AddQuestionIDs(q.ID)
		// finished creating a question
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
	instruction := "no instruction for this locale: " + locale

	if t.Edges.Translations != nil {
		if len(t.Edges.Translations) == 1 {
			trans := t.Edges.Translations[0]
			title = trans.Title
			description = trans.Description
			instruction = trans.Instruction
		}
	}

	var display domain.TestDisplay
	if t.Edges.Display != nil {
		display = entToDomainTestDisplay(t.Edges.Display)
	}

	var questions []*domain.Question
	if t.Edges.Questions != nil {
		for _, q := range t.Edges.Questions {
			questions = append(questions, entToDomainQuestion(q, locale))
		}
	}

	return &domain.Test{
		ID:               t.ID,
		Code:             t.Code,
		Published:        t.Published,
		AvailableLocales: t.AvailableLocales,
		Title:            title,
		Description:      description,
		Instruction:      instruction,
		Display:          display,
		Questions:        questions,
	}
}

func entToDomainTestDisplay(d *ent.TestDisplay) domain.TestDisplay {
	return domain.TestDisplay{
		QuestionsPerPage: d.QuestionsPerPage,
		RandomizeOrder:   d.RandomizeOrder,
	}
}
