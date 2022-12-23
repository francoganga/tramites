package database

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/francoganga/go_reno2/pkg/aggregate"
	"github.com/francoganga/go_reno2/pkg/domain/candidato"
	"github.com/francoganga/go_reno2/pkg/domain/dependencia"
	"github.com/francoganga/go_reno2/pkg/domain/materia"
	"github.com/francoganga/go_reno2/pkg/domain/tramite/database/models"
	"github.com/francoganga/go_reno2/pkg/domain/user"
	"github.com/francoganga/go_reno2/pkg/services"
	"github.com/google/uuid"
)

func TestAddTramite(t *testing.T) {

	c := services.NewContainer()

	repo := New(c.Bun)

	err := repo.Save(context.Background(), createTramite())

	if err != nil {
		t.Errorf("Error guardando tramite en la base: %s", err)
	}
}

// Add observation should save the event and add one observation
// Should have the event in db and the observation saved
func TestAddTramiteWithEvents(t *testing.T) {

	c := services.NewContainer()

	repo := New(c.Bun)

	a := createTramite()

	a.AddObservation("example")

	err := repo.Save(context.Background(), a)

	if err != nil {
		t.Errorf("Error saving tramite to db: %s", err)
	}

	tramite := new(models.Tramite)

	err = c.Bun.NewSelect().
		Model(tramite).
		Where("id = ?", a.GetID()).
		Relation("Events").
		Relation("Observations").
		Scan(context.Background())

	if err != nil {
		t.Error("Failed getting new tramite from db")
	}

	if len(tramite.Events) != 1 {
		t.Errorf("Tramite should have 1 events but has %d", len(tramite.Events))
	}

	e := tramite.Events[0]

	if e.Kind != "ObservationAdded" {
		t.Errorf("Wrong event type: event should be of type: ObservationAdded, got=%s", e.Kind)
	}

	if len(tramite.Observations) != 1 {
		t.Errorf("Wrong quantity of observations: expected 1 got=%d", len(tramite.Observations))
	}
}

func TestIniciarTramite(t *testing.T) {

	c := services.NewContainer()

	repo := New(c.Bun)

	a := createTramite()

	documentID := uuid.New().String()

	err := a.IniciarTramite(documentID)

	if err != nil {
		t.Fatalf("Error iniciando tramite: %s", err)
	}

	err = repo.Save(context.Background(), a)

	if err != nil {
		t.Errorf("Error saving tramite to db: %s", err)
	}

	tramite, err := getFromDBFunc(c)(a.GetID())

	if err != nil {
		t.Error("failed getting tramite from db")
	}

	if len(tramite.Events) != 1 {
		t.Errorf("Wrong events count expected to have 1, got=%d", len(tramite.Events))
	}

	e := tramite.Events[0]

	if e.Kind != "TramiteIniciado" {
		t.Errorf("Wrong event type expected=TramiteIniciado got=%s", e.Kind)
	}

	event := &aggregate.TramiteIniciado{}

	err = json.Unmarshal(e.Payload, event)

	if err != nil {
		t.Errorf("Failed to Unmarshal event with payload=%s", string(e.Payload))
	}

	if event.SolicitudContratacionID != documentID {
		t.Errorf("Wrong document ID, expected=%s, got=%s", documentID, event.SolicitudContratacionID)
	}

}

func createTramite() *aggregate.Tramite {

	candidato := candidato.New("pepe", "test", "pepe@mail.com")

	materias := []*materia.Materia{
		{Nombre: "Metodologias I", Instituto: "Ingenieria", Propuesta: "9", MateriaId: 1, PropuestaId: 1, InstitutoId: 1},
		{Nombre: "Base de datos I", Instituto: "Ingenieria", Propuesta: "92", MateriaId: 5, PropuestaId: 8, InstitutoId: 9},
	}

	user := user.New("admin")

	a := aggregate.New(candidato, 2022, materias, user, dependencia.New("IEI", uuid.New().String(), user), "CAY3")

	return &a
}

func getFromDBFunc(c *services.Container) func(id uuid.UUID) (*models.Tramite, error) {

	return func(id uuid.UUID) (*models.Tramite, error) {
		tramite := new(models.Tramite)

		err := c.Bun.NewSelect().
			Model(tramite).
			Where("id = ?", id).
			Relation("Events").
			Relation("Observations").
			Scan(context.Background())

		if err != nil {
			return &models.Tramite{}, err
		}

		return tramite, nil
	}
}
