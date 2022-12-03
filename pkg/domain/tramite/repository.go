package tramite

import (
	"errors"

	"github.com/francoganga/go_reno2/pkg/aggregate"
	"github.com/google/uuid"
)

var (
    ErrTramiteNotFound = errors.New("El tramite no fue encontrado en el repositorio")
    ErrFailedToAddTramite = errors.New("Fallo al agregar tramite al repositorio")
    ErrUpdateTramite = errors.New("Fallo el update al customer en el repositorio")
)


type TramiteRepository interface {
    Get(uuid.UUID) (aggregate.Tramite, error)
    Add(aggregate.Tramite) error
    Update(aggregate.Tramite) error
}
