package aggregate

import (
	"testing"

	"github.com/francoganga/go_reno2/pkg/domain/candidato"
	"github.com/francoganga/go_reno2/pkg/domain/dependencia"
	"github.com/francoganga/go_reno2/pkg/domain/materia"
	"github.com/francoganga/go_reno2/pkg/domain/user"
	"github.com/google/uuid"
)


func TestAddObservation(t *testing.T) {


    c := candidato.New("franco", "ganga", "asd@mail.com")

    materias := make([]*materia.Materia, 0)

    dep := dependencia.Dependencia{
        Nombre: "IEI",
        AreaSudocu: "2222222",
        Autorizante: user.New("admin"),
    }

    a := New(&c, 2022, materias, user.New("admin"), &dep, "CJ1")

    a.AddObservation("una observacion")
    a.IniciarTramite(uuid.New().String())
    a.AddObservation("segunda observacion")

    t.Logf("aggregate=%+v\n\n", a)

    t.Logf("events: %s", a.PrintEvents())

    if len(a.Observaciones) == 0 {
        t.Fatal("fallo es 0")
    }

}
