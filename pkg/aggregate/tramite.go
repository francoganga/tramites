package aggregate

import (
	"fmt"
	"time"

	"github.com/francoganga/go_reno2/pkg/domain/candidato"
	"github.com/francoganga/go_reno2/pkg/domain/dependencia"
	"github.com/francoganga/go_reno2/pkg/domain/materia"
	"github.com/francoganga/go_reno2/pkg/domain/user"
	"github.com/francoganga/go_reno2/internal"
	"github.com/google/uuid"
)

type Event interface {
	isEvent()
}

type Categoria string

// type Contratacion
// type Solicitud
type Tramite struct {
	Id                uuid.UUID
	AnoProsupuestario int
	Observaciones     []string
	Link              string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	candidato         *candidato.Candidato
	events            []Event
	autor             *user.User
	Dependencia       *dependencia.Dependencia
	Materias          []*materia.Materia
	categoria         Categoria
}

func New(
	candidato *candidato.Candidato,
	anoPresupuestario int,
	materias []*materia.Materia,
	autor *user.User,
	dependencia *dependencia.Dependencia,
	categoria Categoria,
) Tramite {

	rs, err := internal.GenerateRandomBytes(18)

	if err != nil {
		panic(err)
	}

	return Tramite{
		Id:                uuid.New(),
		AnoProsupuestario: anoPresupuestario,
		Observaciones:     make([]string, 0),
		Link:              fmt.Sprintf("%x", rs),
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
		candidato:         candidato,
		events:            make([]Event, 0),
		autor:             autor,
		Dependencia:       dependencia,
		Materias:          materias,
		categoria:         categoria,
	}
}

func (t *Tramite) GetID() uuid.UUID {
    return t.Id
}
