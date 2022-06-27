package domain

import (
	"fmt"
	"reflect"
)

func (t *CreateTagArgs) ValidateTranslations() error {
	locs := make(map[string]struct{})
	for _, l := range Locales() {
		locs[l] = struct{}{}
	}

	// check translations
	transLocs := make(map[string]struct{})
	for _, trans := range t.Translations {
		transLocs[trans.Locale] = struct{}{}
	}

	if eq := reflect.DeepEqual(locs, transLocs); !eq {
		return fmt.Errorf("tag %s'' translations are not full: expected %v but got %v", t.Code, locs, transLocs)
	}

	return nil
}
