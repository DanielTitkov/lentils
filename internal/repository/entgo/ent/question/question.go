// Code generated by entc, DO NOT EDIT.

package question

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the question type in the database.
	Label = "question"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldCode holds the string denoting the code field in the database.
	FieldCode = "code"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// EdgeItems holds the string denoting the items edge name in mutations.
	EdgeItems = "items"
	// EdgeTranslations holds the string denoting the translations edge name in mutations.
	EdgeTranslations = "translations"
	// EdgeTest holds the string denoting the test edge name in mutations.
	EdgeTest = "test"
	// Table holds the table name of the question in the database.
	Table = "questions"
	// ItemsTable is the table that holds the items relation/edge. The primary key declared below.
	ItemsTable = "question_items"
	// ItemsInverseTable is the table name for the Item entity.
	// It exists in this package in order to avoid circular dependency with the "item" package.
	ItemsInverseTable = "items"
	// TranslationsTable is the table that holds the translations relation/edge.
	TranslationsTable = "question_translations"
	// TranslationsInverseTable is the table name for the QuestionTranslation entity.
	// It exists in this package in order to avoid circular dependency with the "questiontranslation" package.
	TranslationsInverseTable = "question_translations"
	// TranslationsColumn is the table column denoting the translations relation/edge.
	TranslationsColumn = "question_translations"
	// TestTable is the table that holds the test relation/edge. The primary key declared below.
	TestTable = "test_questions"
	// TestInverseTable is the table name for the Test entity.
	// It exists in this package in order to avoid circular dependency with the "test" package.
	TestInverseTable = "tests"
)

// Columns holds all SQL columns for question fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldCode,
	FieldType,
}

var (
	// ItemsPrimaryKey and ItemsColumn2 are the table columns denoting the
	// primary key for the items relation (M2M).
	ItemsPrimaryKey = []string{"question_id", "item_id"}
	// TestPrimaryKey and TestColumn2 are the table columns denoting the
	// primary key for the test relation (M2M).
	TestPrimaryKey = []string{"test_id", "question_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// CodeValidator is a validator for the "code" field. It is called by the builders before save.
	CodeValidator func(string) error
	// TypeValidator is a validator for the "type" field. It is called by the builders before save.
	TypeValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)