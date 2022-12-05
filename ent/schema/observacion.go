package schema

import "entgo.io/ent"

// Observacion holds the schema definition for the Observacion entity.
type Observacion struct {
	ent.Schema
}

// Fields of the Observacion.
func (Observacion) Fields() []ent.Field {
	return nil
}

// Edges of the Observacion.
func (Observacion) Edges() []ent.Edge {
	return nil
}
