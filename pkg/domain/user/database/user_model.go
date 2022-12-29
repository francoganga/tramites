package database

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"users,alias:u"`

	ID uuid.UUID `bun:"type:uuid,pk"`

	Username string
}
