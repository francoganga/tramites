package tramite

import (
	"context"
	"errors"

	"github.com/francoganga/go_reno2/pkg/aggregate"
	"github.com/google/uuid"
)

var (
	ErrTramiteNotFound    = errors.New("El tramite no fue encontrado en el repositorio")
	ErrFailedToAddTramite = errors.New("Fallo al agregar tramite al repositorio")
	ErrUpdateTramite      = errors.New("Fallo el update al customer en el repositorio")
)

type TramiteRepository interface {
	Get(context.Context, uuid.UUID) (aggregate.Tramite, error)
	Save(*aggregate.Tramite) error
}
