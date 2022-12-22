package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Tramite struct {
	bun.BaseModel `bun:"tramites,alias:a"`

    ID  uuid.UUID `bun:"type:uuid,pk"`
    AnoPresupuestario int
    Link string
    CreatedAt time.Time
    UpdatedAt time.Time
    Categoria string
    Events []*Event `bun:"rel:has-many,join:id=tramite_id"`
    Observations []*Observation `bun:"rel:has-many,join:id=tramite_id"`
    Estado string
    Version int
}

type Event struct {
    bun.BaseModel `bun:"events,alias:e"`

    ID       uuid.UUID `bun:"type:uuid,pk,default:uuid_generate_v4()"`

    Kind string

    Payload json.RawMessage `bun:"type:jsonb"`
    TramiteID uuid.UUID
}

type Observation struct {
    bun.BaseModel `bun:"observations,alias:o"`

    ID int `bun:"id,pk,autoincrement"`

    Content string
    TramiteID uuid.UUID
}
