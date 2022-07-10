// Code generated by ent, DO NOT EDIT.

package testtranslation

import (
	"fmt"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the testtranslation type in the database.
	Label = "test_translation"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldLocale holds the string denoting the locale field in the database.
	FieldLocale = "locale"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldDetails holds the string denoting the details field in the database.
	FieldDetails = "details"
	// FieldInstruction holds the string denoting the instruction field in the database.
	FieldInstruction = "instruction"
	// FieldResultPreambule holds the string denoting the result_preambule field in the database.
	FieldResultPreambule = "result_preambule"
	// EdgeTest holds the string denoting the test edge name in mutations.
	EdgeTest = "test"
	// Table holds the table name of the testtranslation in the database.
	Table = "test_translations"
	// TestTable is the table that holds the test relation/edge.
	TestTable = "test_translations"
	// TestInverseTable is the table name for the Test entity.
	// It exists in this package in order to avoid circular dependency with the "test" package.
	TestInverseTable = "tests"
	// TestColumn is the table column denoting the test relation/edge.
	TestColumn = "test_translations"
)

// Columns holds all SQL columns for testtranslation fields.
var Columns = []string{
	FieldID,
	FieldLocale,
	FieldTitle,
	FieldDescription,
	FieldDetails,
	FieldInstruction,
	FieldResultPreambule,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "test_translations"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"test_translations",
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
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
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
		return fmt.Errorf("testtranslation: invalid enum value for locale field: %q", l)
	}
}
