package user

import (
	"context"
	"errors"
)


var (
    ErrUserNotFound = errors.New("El usuario no fue encontrado en el repositorio")
    ErrFailedToAddUser = errors.New("Fallo al agregar usuario al repositorio")
    ErrUpdateUser = errors.New("Fallo el update al user en el repositorio")
)

type UserRepository interface {
    Get(ctx context.Context, uuid string) (User, error)
    Save(User) error
}
