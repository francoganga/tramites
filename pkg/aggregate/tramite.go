package aggregate

import (
	"errors"
	"fmt"
	"time"

	"github.com/francoganga/go_reno2/pkg/domain/candidato"
	"github.com/francoganga/go_reno2/pkg/domain/dependencia"
	"github.com/francoganga/go_reno2/pkg/domain/materia"
	"github.com/francoganga/go_reno2/pkg/domain/user"
	"github.com/francoganga/go_reno2/pkg/random"
	"github.com/google/uuid"
)

const (
	EstadoTramiteIniciado = "tramite_iniciado"
	EstadoTramiteBorrador = "borrador"
)

func transitionError(from string, to string) error {

	msg := fmt.Sprintf("Ivalid transition from %+v to %s", from, to)

	return errors.New(msg)
}

type Categoria string

// type Contratacion
// type Solicitud
type Tramite struct {
	Id                      uuid.UUID
	AnoProsupuestario       int
	Observaciones           []string
	Link                    string
	CreatedAt               time.Time
	UpdatedAt               time.Time
	Candidato               *candidato.Candidato
	Events                  []Event
	Autor                   *user.User
	Dependencia             *dependencia.Dependencia
	Materias                []*materia.Materia
	Categoria               Categoria
	SolicitudContratacionID string
	Estado                  string
	Version                 int
}

type eventID uuid.UUID

func New(
	candidato *candidato.Candidato,
	anoPresupuestario int,
	materias []*materia.Materia,
	autor *user.User,
	dependencia *dependencia.Dependencia,
	categoria Categoria,
) Tramite {

	rs, err := random.GenerateRandomBytes(18)

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
		Candidato:         candidato,
		Events:            make([]Event, 0),
		Autor:             autor,
		Dependencia:       dependencia,
		Materias:          materias,
		Categoria:         categoria,
		Estado:            EstadoTramiteBorrador,
	}
}

func (t *Tramite) GetID() uuid.UUID {
	return t.Id
}

func (t *Tramite) AddObservation(content string) error {
	t.raise(&ObservationAdded{
		Content: content,
	})

	return nil
}

func (t *Tramite) IniciarTramite(solicitudID string) error {

	if t.Estado != EstadoTramiteBorrador {
		return transitionError(t.Estado, EstadoTramiteIniciado)
	}

	t.raise(&TramiteIniciado{
		SolicitudContratacionID: solicitudID,
	})

	return nil
}
