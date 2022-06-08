// Code generated by entc, DO NOT EDIT.

package response

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the response type in the database.
	Label = "response"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldValue holds the string denoting the value field in the database.
	FieldValue = "value"
	// FieldMeta holds the string denoting the meta field in the database.
	FieldMeta = "meta"
	// EdgeItem holds the string denoting the item edge name in mutations.
	EdgeItem = "item"
	// EdgeTake holds the string denoting the take edge name in mutations.
	EdgeTake = "take"
	// Table holds the table name of the response in the database.
	Table = "responses"
	// ItemTable is the table that holds the item relation/edge.
	ItemTable = "responses"
	// ItemInverseTable is the table name for the Item entity.
	// It exists in this package in order to avoid circular dependency with the "item" package.
	ItemInverseTable = "items"
	// ItemColumn is the table column denoting the item relation/edge.
	ItemColumn = "item_responses"
	// TakeTable is the table that holds the take relation/edge.
	TakeTable = "responses"
	// TakeInverseTable is the table name for the Take entity.
	// It exists in this package in order to avoid circular dependency with the "take" package.
	TakeInverseTable = "takes"
	// TakeColumn is the table column denoting the take relation/edge.
	TakeColumn = "take_responses"
)

// Columns holds all SQL columns for response fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldValue,
	FieldMeta,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "responses"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"item_responses",
	"take_responses",
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
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// DefaultValue holds the default value on creation for the "value" field.
	DefaultValue int
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)