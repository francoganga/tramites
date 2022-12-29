package user

import "github.com/google/uuid"

type Role string
type User struct {
	ID       uuid.UUID
	Username string
	Roles    []Role
}

func New(username string) *User {
	return &User{
		ID:       uuid.New(),
		Username: username,
		Roles:    make([]Role, 0),
	}
}
