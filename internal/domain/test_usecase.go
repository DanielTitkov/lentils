package domain

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"sort"
)

func (t *Test) OrderQuestions(seed int64) {
	if t.Display.RandomizeOrder {
		rand.Seed(seed)
		rand.Shuffle(len(t.Questions), func(i, j int) {
			t.Questions[i], t.Questions[j] = t.Questions[j], t.Questions[i]
		})
	} else {
		sort.Slice(t.Questions, func(i, j int) bool {
			return t.Questions[i].Order < t.Questions[j].Order
		})
	}
}

func (t *Test) PageCount() int {
	if t.Display.QuestionsPerPage < 1 {
		return 1
	}
	return int(math.Ceil(float64(len(t.Questions)) / float64(t.Display.QuestionsPerPage)))
}

func (t *Test) GetItem(code string) *Item {
	// test wont have too many items
	// so map is not required probably
	// TODO: but it'd better to check
	for _, q := range t.Questions {
		for _, i := range q.Items {
			if i.Code == code {
				return i
			}
		}
	}

	return nil
}

func (t *Test) QuestionsForPage(page int) []*Question {
	if page < 1 {
		return []*Question{}
	}

	if page > t.PageCount() {
		return []*Question{}
	}

	// TODO: maybe pregenerate pages
	page = page - 1 // in the app page count goes from 1
	var pages [][]*Question
	for i := 0; i < t.PageCount(); i++ {
		var pq []*Question
		for j := 0; j < t.Display.QuestionsPerPage; j++ {
			idx := i*t.Display.QuestionsPerPage + j
			if idx >= len(t.Questions) {
				break
			}
			pq = append(pq, t.Questions[idx])
		}
		pages = append(pages, pq)
	}

	return pages[page]
}

func (t *Test) IsPageDone(page int) bool {
	for _, q := range t.QuestionsForPage(page) {
		if !q.IsDone() {
			return false
		}
	}

	return true
}

func (t *Test) IsPageNotDone(page int) bool {
	return !t.IsPageDone(page)
}

func (t *Test) IsDone() bool {
	for _, q := range t.Questions {
		if !q.IsDone() {
			return false
		}
	}

	return true
}

func (t *Test) IsNotDone() bool {
	return !t.IsDone()
}

func (t *Test) CalculateResult() error {
	for _, s := range t.Scales {
		err := s.CalculateResult()
		if err != nil {
			return err
		}
	}

	return nil
}

func (t *CreateTestArgs) ValidateTranslations() error {
	locs := make(map[string]struct{})
	for _, l := range t.AvailableLocales {
		locs[l] = struct{}{}
	}

	// check test translations
	transLocs := make(map[string]struct{})
	for _, trans := range t.Translations {
		transLocs[trans.Locale] = struct{}{}
	}

	if eq := reflect.DeepEqual(locs, transLocs); !eq {
		return fmt.Errorf("test %s'' translations are not full: expected %v but got %v", t.Code, locs, transLocs)
	}

	// check scales
	for _, scale := range t.Scales {
		scaleLocs := make(map[string]struct{})
		for _, trans := range scale.Translations {
			scaleLocs[trans.Locale] = struct{}{}
		}
		if eq := reflect.DeepEqual(locs, scaleLocs); !eq {
			return fmt.Errorf("scale '%s' translations are not full: expected %v but got %v", scale.Code, locs, scaleLocs)
		}

		// check items
		for _, itm := range scale.Items {
			itemLocs := make(map[string]struct{})
			for _, trans := range itm.Translations {
				itemLocs[trans.Locale] = struct{}{}
			}
			if eq := reflect.DeepEqual(locs, itemLocs); !eq {
				return fmt.Errorf("item '%s' translations are not full: expected %v but got %v", itm.Code, locs, itemLocs)
			}
		}
	}

	// check questions
	for _, q := range t.Questions {
		questionLocs := make(map[string]struct{})
		for _, trans := range q.Translations {
			questionLocs[trans.Locale] = struct{}{}
		}
		if eq := reflect.DeepEqual(locs, questionLocs); !eq {
			return fmt.Errorf("question '%s' translations are not full: expected %v but got %v", q.Code, locs, questionLocs)
		}
	}

	return nil
}
