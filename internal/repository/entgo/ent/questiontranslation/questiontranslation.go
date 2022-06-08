// Code generated by entc, DO NOT EDIT.

package questiontranslation

import (
	"fmt"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the questiontranslation type in the database.
	Label = "question_translation"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldLocale holds the string denoting the locale field in the database.
	FieldLocale = "locale"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldHeaderContent holds the string denoting the header_content field in the database.
	FieldHeaderContent = "header_content"
	// FieldFooterContent holds the string denoting the footer_content field in the database.
	FieldFooterContent = "footer_content"
	// EdgeQuestion holds the string denoting the question edge name in mutations.
	EdgeQuestion = "question"
	// Table holds the table name of the questiontranslation in the database.
	Table = "question_translations"
	// QuestionTable is the table that holds the question relation/edge.
	QuestionTable = "question_translations"
	// QuestionInverseTable is the table name for the Question entity.
	// It exists in this package in order to avoid circular dependency with the "question" package.
	QuestionInverseTable = "questions"
	// QuestionColumn is the table column denoting the question relation/edge.
	QuestionColumn = "question_translations"
)

// Columns holds all SQL columns for questiontranslation fields.
var Columns = []string{
	FieldID,
	FieldLocale,
	FieldContent,
	FieldHeaderContent,
	FieldFooterContent,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "question_translations"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"question_translations",
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
		return fmt.Errorf("questiontranslation: invalid enum value for locale field: %q", l)
	}
}