package aggregate

import (
	"database/sql/driver"
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

type UJson map[string]string

func (j *UJson) Scan(val string) error {

    return json.Unmarshal([]byte(val), &j)
}

func (pc *UJson) Value() (driver.Value, error) {
	return json.Marshal(pc)
}

func NewJson(v map[string]string) UJson {
    return UJson(v)
}


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
	ID      string
	Content string
}

func (e ObservationAdded) isEvent() {}
func (e ObservationAdded) String() string {
	return reflect.TypeOf(e).Name()
}

type TramiteIniciado struct {
	ID                      string
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
	}
}

func (t *Tramite) GetID() uuid.UUID {
	return t.Id
}

func (t *Tramite) AddObservation(content string) error {
	t.raise(&ObservationAdded{
		ID:      uuid.New().String(),
		Content: content,
	})

	return nil
}

func (t *Tramite) IniciarTramite(solicitudID string) error {
	t.raise(&TramiteIniciado{
		ID:                      uuid.New().String(),
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
