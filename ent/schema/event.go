package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Event holds the schema definition for the Event entity.
type Event struct {
	ent.Schema
}



// Fields of the Event.
func (Event) Fields() []ent.Field {
    return []ent.Field{
        field.String("type"),
        field.JSON("payload", map[string]string{}),
    }
}

// Edges of the Event.
func (Event) Edges() []ent.Edge {
	return nil
}
