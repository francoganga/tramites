package user

import "github.com/google/uuid"



type Role string
type User struct {
	Id    uuid.UUID
    Username string
	Roles []Role
}

func New(username string) *User {
    return &User{
        Id: uuid.New(),
        Username: username,
        Roles: make([]Role, 0),
    }
}
