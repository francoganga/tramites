package database

import (
	"context"

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
}


func (r *DatabaseRepository) Get(ctx context.Context, id uuid.UUID) (aggregate.Tramite, error) {

    t, err := r.ORM.Tramite.Query().Where(tramite.ID(id)).Only(ctx)

    if err != nil {
        return aggregate.Tramite{}, err
    }

    a, err := toAggregate(t)

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


func toAggregate(t *ent.Tramite) (aggregate.Tramite, error) {

    user := user.New("admin")

    return aggregate.Tramite{
        Id: t.ID,
        AnoProsupuestario: t.AnoPresupuestario,
        Observaciones: make([]string, 0),
        Link: t.Link,
        CreatedAt: t.CreatedAt,
        UpdatedAt: t.UpdatedAt,
        Candidato: candidato.New("test", "test", "test"),
        Events: make([]aggregate.Event, 0),
        Autor: user,
        Dependencia: &dependencia.Dependencia{Nombre: "IEI", AreaSudocu: "22222", Autorizante: user},
        Materias: make([]*materia.Materia, 0),
        Categoria: "ING",
        SolicitudContratacionID: "",                        
        Estado: "borrador",
        Version: t.Version,
    }, nil
}

// func fromAggregate(a *aggregate.Tramite) (t *ent.Tramite, error) {
//     return &ent.Tramite
// }
