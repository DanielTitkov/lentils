// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/take"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/test"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/user"
	"github.com/google/uuid"
)

// Take is the model entity for the Take schema.
type Take struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Seed holds the value of the "seed" field.
	Seed int64 `json:"seed,omitempty"`
	// Meta holds the value of the "meta" field.
	Meta map[string]interface{} `json:"meta,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TakeQuery when eager-loading is set.
	Edges      TakeEdges `json:"edges"`
	test_takes *uuid.UUID
	user_takes *uuid.UUID
}

// TakeEdges holds the relations/edges for other nodes in the graph.
type TakeEdges struct {
	// Responses holds the value of the responses edge.
	Responses []*Response `json:"responses,omitempty"`
	// Test holds the value of the test edge.
	Test *Test `json:"test,omitempty"`
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// ResponsesOrErr returns the Responses value or an error if the edge
// was not loaded in eager-loading.
func (e TakeEdges) ResponsesOrErr() ([]*Response, error) {
	if e.loadedTypes[0] {
		return e.Responses, nil
	}
	return nil, &NotLoadedError{edge: "responses"}
}

// TestOrErr returns the Test value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TakeEdges) TestOrErr() (*Test, error) {
	if e.loadedTypes[1] {
		if e.Test == nil {
			// The edge test was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: test.Label}
		}
		return e.Test, nil
	}
	return nil, &NotLoadedError{edge: "test"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TakeEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[2] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Take) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case take.FieldMeta:
			values[i] = new([]byte)
		case take.FieldSeed:
			values[i] = new(sql.NullInt64)
		case take.FieldCreateTime, take.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case take.FieldID:
			values[i] = new(uuid.UUID)
		case take.ForeignKeys[0]: // test_takes
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case take.ForeignKeys[1]: // user_takes
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Take", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Take fields.
func (t *Take) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case take.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				t.ID = *value
			}
		case take.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				t.CreateTime = value.Time
			}
		case take.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				t.UpdateTime = value.Time
			}
		case take.FieldSeed:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field seed", values[i])
			} else if value.Valid {
				t.Seed = value.Int64
			}
		case take.FieldMeta:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field meta", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &t.Meta); err != nil {
					return fmt.Errorf("unmarshal field meta: %w", err)
				}
			}
		case take.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field test_takes", values[i])
			} else if value.Valid {
				t.test_takes = new(uuid.UUID)
				*t.test_takes = *value.S.(*uuid.UUID)
			}
		case take.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field user_takes", values[i])
			} else if value.Valid {
				t.user_takes = new(uuid.UUID)
				*t.user_takes = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryResponses queries the "responses" edge of the Take entity.
func (t *Take) QueryResponses() *ResponseQuery {
	return (&TakeClient{config: t.config}).QueryResponses(t)
}

// QueryTest queries the "test" edge of the Take entity.
func (t *Take) QueryTest() *TestQuery {
	return (&TakeClient{config: t.config}).QueryTest(t)
}

// QueryUser queries the "user" edge of the Take entity.
func (t *Take) QueryUser() *UserQuery {
	return (&TakeClient{config: t.config}).QueryUser(t)
}

// Update returns a builder for updating this Take.
// Note that you need to call Take.Unwrap() before calling this method if this Take
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Take) Update() *TakeUpdateOne {
	return (&TakeClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the Take entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Take) Unwrap() *Take {
	tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Take is not a transactional entity")
	}
	t.config.driver = tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Take) String() string {
	var builder strings.Builder
	builder.WriteString("Take(")
	builder.WriteString(fmt.Sprintf("id=%v", t.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(t.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(t.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", seed=")
	builder.WriteString(fmt.Sprintf("%v", t.Seed))
	builder.WriteString(", meta=")
	builder.WriteString(fmt.Sprintf("%v", t.Meta))
	builder.WriteByte(')')
	return builder.String()
}

// Takes is a parsable slice of Take.
type Takes []*Take

func (t Takes) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}