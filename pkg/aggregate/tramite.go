package aggregate

import (
	"encoding/json"
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

type Event interface {
	isEvent()
    String() string
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
	version           int
}

type eventID uuid.UUID

type ObservationAdded struct {
    ID string
    Content string
}

func (e ObservationAdded) isEvent() {}
func (e ObservationAdded) String() string {
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

func (t *Tramite) AddObservation(content string) error {
    t.raise(&ObservationAdded{
        ID: uuid.New().String(),
        Content: content,
    })

    return nil
}


func (t *Tramite) On(event Event, new bool) {
    switch e := event.(type) {
    case *ObservationAdded:
        t.Observaciones = append(t.Observaciones, e.Content)
    }

    if !new {
        t.version++
    }
}

func (t *Tramite) raise(event Event) {
    t.events = append(t.events, event)
    t.On(event, true)
}

func (t *Tramite) PrintEvents() string {
    s := "["
    for _, e := range t.events {
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
