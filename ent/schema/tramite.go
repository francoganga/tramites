package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Tramite holds the schema definition for the Tramite entity.
type Tramite struct {
	ent.Schema
}

// Fields of the Tramite.
func (Tramite) Fields() []ent.Field {
    return []ent.Field{
        field.UUID("id", uuid.UUID{}).Default(uuid.New),
        field.Int("anoPresupuestario"),
        field.String("link"),
        field.Time("created_at").Optional(),
        field.Time("updated_at").Optional(),
        field.String("categoria"),
        field.Int("version"),
    }
}

// Edges of the Tramite.
func (Tramite) Edges() []ent.Edge {
	return nil
}
