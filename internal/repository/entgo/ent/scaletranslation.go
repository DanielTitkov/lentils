// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/scale"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/scaletranslation"
	"github.com/google/uuid"
)

// ScaleTranslation is the model entity for the ScaleTranslation schema.
type ScaleTranslation struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Locale holds the value of the "locale" field.
	Locale scaletranslation.Locale `json:"locale,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ScaleTranslationQuery when eager-loading is set.
	Edges              ScaleTranslationEdges `json:"edges"`
	scale_translations *uuid.UUID
}

// ScaleTranslationEdges holds the relations/edges for other nodes in the graph.
type ScaleTranslationEdges struct {
	// Scale holds the value of the scale edge.
	Scale *Scale `json:"scale,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ScaleOrErr returns the Scale value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ScaleTranslationEdges) ScaleOrErr() (*Scale, error) {
	if e.loadedTypes[0] {
		if e.Scale == nil {
			// The edge scale was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: scale.Label}
		}
		return e.Scale, nil
	}
	return nil, &NotLoadedError{edge: "scale"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ScaleTranslation) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case scaletranslation.FieldLocale, scaletranslation.FieldTitle, scaletranslation.FieldDescription:
			values[i] = new(sql.NullString)
		case scaletranslation.FieldID:
			values[i] = new(uuid.UUID)
		case scaletranslation.ForeignKeys[0]: // scale_translations
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type ScaleTranslation", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ScaleTranslation fields.
func (st *ScaleTranslation) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case scaletranslation.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				st.ID = *value
			}
		case scaletranslation.FieldLocale:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field locale", values[i])
			} else if value.Valid {
				st.Locale = scaletranslation.Locale(value.String)
			}
		case scaletranslation.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				st.Title = value.String
			}
		case scaletranslation.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				st.Description = value.String
			}
		case scaletranslation.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field scale_translations", values[i])
			} else if value.Valid {
				st.scale_translations = new(uuid.UUID)
				*st.scale_translations = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryScale queries the "scale" edge of the ScaleTranslation entity.
func (st *ScaleTranslation) QueryScale() *ScaleQuery {
	return (&ScaleTranslationClient{config: st.config}).QueryScale(st)
}

// Update returns a builder for updating this ScaleTranslation.
// Note that you need to call ScaleTranslation.Unwrap() before calling this method if this ScaleTranslation
// was returned from a transaction, and the transaction was committed or rolled back.
func (st *ScaleTranslation) Update() *ScaleTranslationUpdateOne {
	return (&ScaleTranslationClient{config: st.config}).UpdateOne(st)
}

// Unwrap unwraps the ScaleTranslation entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (st *ScaleTranslation) Unwrap() *ScaleTranslation {
	tx, ok := st.config.driver.(*txDriver)
	if !ok {
		panic("ent: ScaleTranslation is not a transactional entity")
	}
	st.config.driver = tx.drv
	return st
}

// String implements the fmt.Stringer.
func (st *ScaleTranslation) String() string {
	var builder strings.Builder
	builder.WriteString("ScaleTranslation(")
	builder.WriteString(fmt.Sprintf("id=%v", st.ID))
	builder.WriteString(", locale=")
	builder.WriteString(fmt.Sprintf("%v", st.Locale))
	builder.WriteString(", title=")
	builder.WriteString(st.Title)
	builder.WriteString(", description=")
	builder.WriteString(st.Description)
	builder.WriteByte(')')
	return builder.String()
}

// ScaleTranslations is a parsable slice of ScaleTranslation.
type ScaleTranslations []*ScaleTranslation

func (st ScaleTranslations) config(cfg config) {
	for _i := range st {
		st[_i].config = cfg
	}
}
