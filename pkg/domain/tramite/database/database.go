package database

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/francoganga/go_reno2/ent"
	"github.com/francoganga/go_reno2/ent/tramite"
	"github.com/francoganga/go_reno2/pkg/aggregate"
	"github.com/francoganga/go_reno2/pkg/domain/candidato"
	"github.com/francoganga/go_reno2/pkg/domain/dependencia"
	"github.com/francoganga/go_reno2/pkg/domain/materia"
	"github.com/francoganga/go_reno2/pkg/domain/user"
	"github.com/google/uuid"
)

type DatabaseRepository struct {
	ORM *ent.Client
	DB  *sql.DB
}

func (r *DatabaseRepository) Get(ctx context.Context, id string) (aggregate.Tramite, error) {

	uid, err := uuid.Parse(id)

	if err != nil {
		return aggregate.Tramite{}, fmt.Errorf("uuid invalido: %w", err)
	}

	t, err := r.ORM.Tramite.Query().Where(tramite.ID(uid)).Only(ctx)

	events, err := r.DB.QueryContext(ctx, "SELECT type, payload from events where tramite_events = $1", id)

	if err != nil {
		return aggregate.Tramite{}, fmt.Errorf("Failed to get tramite events: %w", err)
	}

	if err != nil {
		return aggregate.Tramite{}, err
	}

	a, err := toAggregate(t, events)

	if err != nil {
		return aggregate.Tramite{}, err
	}

	return a, nil
}

func (r *DatabaseRepository) Add(ctx context.Context, a aggregate.Tramite) error {
	_, err := r.ORM.Tramite.Create().
		SetAnoPresupuestario(a.AnoProsupuestario).
		SetID(a.GetID()).
		SetLink(a.Link).
		SetCategoria(string(a.Categoria)).
		SetVersion(0).Save(ctx)

	if err != nil {
		return err
	}

	return nil
}

func (r *DatabaseRepository) Save(ctx context.Context, a aggregate.Tramite) error {

	exists, err := r.ORM.Tramite.Query().Where(tramite.ID(a.GetID())).Exist(ctx)

	if err != nil {
		return err
	}

	// TODO: not handling possible error
	// if exists, _ := r.ORM.Tramite.Query().Where(tramite.ID(a.GetID())).Exist(ctx); !exists {

    tx, err := r.ORM.Tx(ctx)

	if !exists {
		r.ORM.Tramite.Create().
			SetID(a.GetID()).
			SetLink(a.Link).
			SetAnoPresupuestario(a.AnoProsupuestario).
			SetVersion(0).
			SetCategoria(string(a.Categoria)).
			Save(ctx)
	}

	for _, e := range a.Events {


        switch e := e.(type) {
        case *aggregate.TramiteIniciado:
            _, err := r.DB.Exec("UPDATE tramites set estado = 'tramite_iniciado'")
            fmt.Printf("handling e=%v\n", e)

            if err != nil {
                tx.Rollback()
                return err
            }
        }


		fmt.Printf("processing event=%v\n", e)

		j, err := json.Marshal(e)

		if err != nil {
			fmt.Printf("error: %v\n", err)
			return err
		}

		_, err = r.DB.Exec("INSERT INTO events (type, tramite_events, payload) VALUES($1, $2, $3)", e.String(), a.GetID(), j)

		if err != nil {
            tx.Rollback()
			return err
		}

	}
    err = tx.Commit()

    if err != nil {
        tx.Rollback()
        return err
    }

	return nil
}

func toAggregate(t *ent.Tramite, eventRows *sql.Rows) (aggregate.Tramite, error) {

	defer eventRows.Close()

	fmt.Printf("[toAggregate] ent.Tramite=%v\n", t)

    // TODO hardcoded user
	user := user.New("admin")

	a := aggregate.Tramite{
		Id:                      t.ID,
		AnoProsupuestario:       t.AnoPresupuestario,
		Observaciones:           make([]string, 0),
		Link:                    t.Link,
		CreatedAt:               t.CreatedAt,
		UpdatedAt:               t.UpdatedAt,
        // TODO candidato hardcoded
		Candidato:               candidato.New("test", "test", "test"),
		Events:                  make([]aggregate.Event, 0),
		Autor:                   user,
		Dependencia:             &dependencia.Dependencia{Nombre: "IEI", AreaSudocu: "22222", Autorizante: user},
		Materias:                make([]*materia.Materia, 0),
		Categoria:               "ING",
		SolicitudContratacionID: "",
		Estado:                  "borrador",
		Version:                 t.Version,
	}

	for eventRows.Next() {

        var tipo string
        var b []byte

		err := eventRows.Scan(&tipo, &b)

        var ev aggregate.Event
        switch tipo {
        case eventName(&aggregate.ObservationAdded{}):
            ev = &aggregate.ObservationAdded{}
        case eventName(&aggregate.TramiteIniciado{}):
            ev = &aggregate.TramiteIniciado{}
        }

        err = json.Unmarshal(b, ev)

        a.On(ev, false)

		if err != nil {
			return aggregate.Tramite{}, err
		}
	}

	return a, nil
}

func eventName(event aggregate.Event) string {
	t := reflect.TypeOf(event)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	return t.Name()
}

