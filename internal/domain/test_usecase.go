package domain

import (
	"fmt"
	"reflect"
	"sort"
)

func (t *Test) OrderQuestions() {
	sort.Slice(t.Questions, func(i, j int) bool {
		return t.Questions[i].Order < t.Questions[j].Order
	})
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
