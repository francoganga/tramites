// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/francoganga/go_reno2/ent/observacion"
)

// Observacion is the model entity for the Observacion schema.
type Observacion struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Observacion) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case observacion.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Observacion", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Observacion fields.
func (o *Observacion) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case observacion.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			o.ID = int(value.Int64)
		}
	}
	return nil
}

// Update returns a builder for updating this Observacion.
// Note that you need to call Observacion.Unwrap() before calling this method if this Observacion
// was returned from a transaction, and the transaction was committed or rolled back.
func (o *Observacion) Update() *ObservacionUpdateOne {
	return (&ObservacionClient{config: o.config}).UpdateOne(o)
}

// Unwrap unwraps the Observacion entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (o *Observacion) Unwrap() *Observacion {
	_tx, ok := o.config.driver.(*txDriver)
	if !ok {
		panic("ent: Observacion is not a transactional entity")
	}
	o.config.driver = _tx.drv
	return o
}

// String implements the fmt.Stringer.
func (o *Observacion) String() string {
	var builder strings.Builder
	builder.WriteString("Observacion(")
	builder.WriteString(fmt.Sprintf("id=%v", o.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Observacions is a parsable slice of Observacion.
type Observacions []*Observacion

func (o Observacions) config(cfg config) {
	for _i := range o {
		o[_i].config = cfg
	}
}
