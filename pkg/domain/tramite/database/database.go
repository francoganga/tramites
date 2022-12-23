package database

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/francoganga/go_reno2/pkg/aggregate"
	"github.com/francoganga/go_reno2/pkg/domain/candidato"
	"github.com/francoganga/go_reno2/pkg/domain/dependencia"
	"github.com/francoganga/go_reno2/pkg/domain/materia"
	"github.com/francoganga/go_reno2/pkg/domain/tramite/database/models"
	"github.com/francoganga/go_reno2/pkg/domain/user"
	"github.com/uptrace/bun"
)

type DatabaseRepository struct {
	Bun *bun.DB
}

func New(bun *bun.DB) DatabaseRepository {
	return DatabaseRepository{Bun: bun}
}

func (r *DatabaseRepository) Get(ctx context.Context, id string) (aggregate.Tramite, error) {

	tramite := new(models.Tramite)

	err := r.Bun.NewSelect().
		Model(tramite).
		Where("id = ?", id).
		Relation("Events").
		Scan(ctx)

	if err != nil {
		return aggregate.Tramite{}, err
	}

	a, err := toAggregate(tramite)

	if err != nil {
		return aggregate.Tramite{}, err
	}

	return a, nil
}

func (r *DatabaseRepository) Save(ctx context.Context, a *aggregate.Tramite) error {

	nt := new(models.Tramite)

	tx, err := r.Bun.BeginTx(ctx, &sql.TxOptions{})

	if err != nil {
		return err
	}

	defer tx.Rollback()

	exists, err := r.Bun.NewSelect().Model((*models.Tramite)(nil)).Where("id = ?", a.GetID()).Exists(ctx)

	if err != nil {
		return err
	}

	if !exists {

		fmt.Println("no existe, creando")
		nt.ID = a.GetID()
		nt.Categoria = string(a.Categoria)
		nt.Link = a.Link
		nt.AnoPresupuestario = a.AnoProsupuestario
		nt.Estado = a.Estado
		nt.CreatedAt = a.CreatedAt
		nt.UpdatedAt = a.UpdatedAt

		obs := make([]*models.Observation, 0)
		evs := make([]*models.Event, 0)

		for _, e := range a.Events {
			switch e := e.(type) {
			case *aggregate.ObservationAdded:
				obs = append(obs, &models.Observation{
					Content:   e.Content,
					TramiteID: a.GetID(),
				})
			}

			j, err := json.Marshal(e)

			if err != nil {
				return err
			}

			evs = append(evs, &models.Event{
				Kind:      e.String(),
				TramiteID: a.GetID(),
				Payload:   j,
			})
		}

		_, err := tx.NewInsert().
			Model(nt).
			Exec(ctx)

		if err != nil {
			return err
		}

		if len(obs) > 0 {

			_, err = tx.NewInsert().Model(&obs).Exec(ctx)

			if err != nil {
				return err
			}
		}

		if len(evs) > 0 {

			_, err = tx.NewInsert().Model(&evs).Exec(ctx)

			if err != nil {
				return err
			}
		}

		if err = tx.Commit(); err != nil {
			return err
		}

		a.Events = nil

		return nil
	}

	fmt.Println("existe, updateando")

	if len(a.Events) == 0 {
		fmt.Println("no events returning early")
		return nil
	}

	err = tx.NewSelect().
		Model(nt).
		Where("id = ?", a.GetID()).
		Scan(ctx)

	if err != nil {
		return err
	}

	fmt.Printf("nt=%v\n", nt)

	obs := make([]*models.Observation, 0)
	evs := make([]*models.Event, 0)

	for _, e := range a.Events {
		switch e := e.(type) {
		case *aggregate.TramiteIniciado:
			nt.Estado = aggregate.EstadoTramiteIniciado
		case *aggregate.ObservationAdded:
			obs = append(obs, &models.Observation{
				Content:   e.Content,
				TramiteID: a.GetID(),
			})
		}

		nt.UpdatedAt = time.Now()

		j, err := json.Marshal(e)

		if err != nil {
			return err
		}

		evs = append(evs, &models.Event{
			Kind:      e.String(),
			TramiteID: a.GetID(),
			Payload:   j,
		})

	}

	_, err = tx.NewUpdate().
		Model(nt).
		WherePK().
		Exec(ctx)

	if err != nil {
		return err
	}

	if len(obs) > 0 {
		_, err = tx.NewInsert().Model(&obs).Exec(ctx)
		if err != nil {
			return err
		}
	}

	_, err = tx.NewInsert().Model(&evs).Exec(ctx)

	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	fmt.Println("setting events to nil")
	a.Events = nil

	return nil
}

func toAggregate(t *models.Tramite) (aggregate.Tramite, error) {

	user := user.New("admin")

	a := aggregate.Tramite{
		Id:                t.ID,
		AnoProsupuestario: t.AnoPresupuestario,
		Observaciones:     make([]string, 0),
		Link:              t.Link,
		CreatedAt:         t.CreatedAt,
		UpdatedAt:         t.UpdatedAt,
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

	for _, e := range t.Events {

		var ev aggregate.Event
		switch e.Kind {
		case eventName(&aggregate.ObservationAdded{}):
			ev = &aggregate.ObservationAdded{}
		case eventName(&aggregate.TramiteIniciado{}):
			ev = &aggregate.TramiteIniciado{}
		}

		err := json.Unmarshal(e.Payload, ev)

		if err != nil {
			return aggregate.Tramite{}, err
		}

		a.On(ev, false)
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
