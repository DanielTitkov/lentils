// Code generated by entc, DO NOT EDIT.

package interpretationtranslation

import (
	"fmt"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the interpretationtranslation type in the database.
	Label = "interpretation_translation"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldLocale holds the string denoting the locale field in the database.
	FieldLocale = "locale"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// EdgeInterpretation holds the string denoting the interpretation edge name in mutations.
	EdgeInterpretation = "interpretation"
	// Table holds the table name of the interpretationtranslation in the database.
	Table = "interpretation_translations"
	// InterpretationTable is the table that holds the interpretation relation/edge.
	InterpretationTable = "interpretation_translations"
	// InterpretationInverseTable is the table name for the Interpretation entity.
	// It exists in this package in order to avoid circular dependency with the "interpretation" package.
	InterpretationInverseTable = "interpretations"
	// InterpretationColumn is the table column denoting the interpretation relation/edge.
	InterpretationColumn = "interpretation_translations"
)

// Columns holds all SQL columns for interpretationtranslation fields.
var Columns = []string{
	FieldID,
	FieldLocale,
	FieldContent,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "interpretation_translations"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"interpretation_translations",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// ContentValidator is a validator for the "content" field. It is called by the builders before save.
	ContentValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Locale defines the type for the "locale" enum field.
type Locale string

// Locale values.
const (
	LocaleEn Locale = "en"
	LocaleRu Locale = "ru"
)

func (l Locale) String() string {
	return string(l)
}

// LocaleValidator is a validator for the "locale" field enum values. It is called by the builders before save.
func LocaleValidator(l Locale) error {
	switch l {
	case LocaleEn, LocaleRu:
		return nil
	default:
		return fmt.Errorf("interpretationtranslation: invalid enum value for locale field: %q", l)
	}
}
