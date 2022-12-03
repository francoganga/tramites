// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/francoganga/go_reno2/ent/passwordtoken"
	"github.com/francoganga/go_reno2/ent/user"
)

// PasswordToken is the model entity for the PasswordToken schema.
type PasswordToken struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Hash holds the value of the "hash" field.
	Hash string `json:"-"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PasswordTokenQuery when eager-loading is set.
	Edges               PasswordTokenEdges `json:"edges"`
	password_token_user *int
}

// PasswordTokenEdges holds the relations/edges for other nodes in the graph.
type PasswordTokenEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PasswordTokenEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
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
func (*PasswordToken) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case passwordtoken.FieldID:
			values[i] = new(sql.NullInt64)
		case passwordtoken.FieldHash:
			values[i] = new(sql.NullString)
		case passwordtoken.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case passwordtoken.ForeignKeys[0]: // password_token_user
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type PasswordToken", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PasswordToken fields.
func (pt *PasswordToken) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case passwordtoken.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pt.ID = int(value.Int64)
		case passwordtoken.FieldHash:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field hash", values[i])
			} else if value.Valid {
				pt.Hash = value.String
			}
		case passwordtoken.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pt.CreatedAt = value.Time
			}
		case passwordtoken.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field password_token_user", value)
			} else if value.Valid {
				pt.password_token_user = new(int)
				*pt.password_token_user = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the PasswordToken entity.
func (pt *PasswordToken) QueryUser() *UserQuery {
	return (&PasswordTokenClient{config: pt.config}).QueryUser(pt)
}

// Update returns a builder for updating this PasswordToken.
// Note that you need to call PasswordToken.Unwrap() before calling this method if this PasswordToken
// was returned from a transaction, and the transaction was committed or rolled back.
func (pt *PasswordToken) Update() *PasswordTokenUpdateOne {
	return (&PasswordTokenClient{config: pt.config}).UpdateOne(pt)
}

// Unwrap unwraps the PasswordToken entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pt *PasswordToken) Unwrap() *PasswordToken {
	tx, ok := pt.config.driver.(*txDriver)
	if !ok {
		panic("ent: PasswordToken is not a transactional entity")
	}
	pt.config.driver = tx.drv
	return pt
}

// String implements the fmt.Stringer.
func (pt *PasswordToken) String() string {
	var builder strings.Builder
	builder.WriteString("PasswordToken(")
	builder.WriteString(fmt.Sprintf("id=%v", pt.ID))
	builder.WriteString(", hash=<sensitive>")
	builder.WriteString(", created_at=")
	builder.WriteString(pt.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// PasswordTokens is a parsable slice of PasswordToken.
type PasswordTokens []*PasswordToken

func (pt PasswordTokens) config(cfg config) {
	for _i := range pt {
		pt[_i].config = cfg
	}
}
