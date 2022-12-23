package aggregate

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
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
)


func transitionError(from string, to string) error {

    msg := fmt.Sprintf("Ivalid transition from %+v to %s", from, to)

    return errors.New(msg)
}


type Event interface {
	isEvent()
	String() string
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

type ObservationAdded struct {
	Content string
}

func (e ObservationAdded) isEvent() {}
func (e ObservationAdded) String() string {
	return reflect.TypeOf(e).Name()
}

type TramiteIniciado struct {
	SolicitudContratacionID string
}

func (e TramiteIniciado) isEvent() {}
func (e TramiteIniciado) String() string {
	return reflect.TypeOf(e).Name()
}

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
        Estado: "borrador",
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

    if t.Estado != "borrador" {
        return transitionError(t.Estado, EstadoTramiteIniciado)
    }


	t.raise(&TramiteIniciado{
		SolicitudContratacionID: solicitudID,
	})

	return nil
}

func (t *Tramite) On(event Event, new bool) {
	switch e := event.(type) {
	case *ObservationAdded:
		t.Observaciones = append(t.Observaciones, e.Content)
	case *TramiteIniciado:
		t.SolicitudContratacionID = e.SolicitudContratacionID
		t.Estado = "tramite_iniciado"
	}

    t.UpdatedAt = time.Now()

	if !new {
		t.Version++
	}
}

func (t *Tramite) raise(event Event) {
	t.Events = append(t.Events, event)
	t.On(event, true)
}

func (t *Tramite) PrintEvents() string {
	s := "["
	for _, e := range t.Events {
		j, err := json.Marshal(e)
		if err != nil {
			panic(err)
		}

		s += e.String()

		s += ": "

		s += string(j)
		s += ", "
	}
	s += "]"
	return s
}
