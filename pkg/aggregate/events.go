package aggregate

import (
	"encoding/json"
	"reflect"
	"time"
)

type Event interface {
	isEvent()
	String() string
	visit(*Tramite)
}

type ObservationAdded struct {
	Content string
}

func (e ObservationAdded) isEvent() {}
func (e ObservationAdded) String() string {
	return reflect.TypeOf(e).Name()
}
func (e ObservationAdded) visit(t *Tramite) {
	t.Observaciones = append(t.Observaciones, e.Content)
}

type TramiteIniciado struct {
	SolicitudContratacionID string
}

func (e TramiteIniciado) isEvent() {}
func (e TramiteIniciado) String() string {
	return reflect.TypeOf(e).Name()
}
func (e TramiteIniciado) visit(t *Tramite) {
	t.SolicitudContratacionID = e.SolicitudContratacionID
	t.Estado = "tramite_iniciado"
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

func (t *Tramite) On(event Event, new bool) {
	event.visit(t)

	t.UpdatedAt = time.Now()

	if !new {
		t.Version++
	}
}
