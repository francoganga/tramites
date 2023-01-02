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

	a := createAggregate()

	a.AddObservation("una observacion")
	a.IniciarTramite(uuid.New().String())
	a.AddObservation("segunda observacion")

	t.Logf("aggregate=%+v\n\n", a)

	t.Logf("events: %s", a.PrintEvents())

	if len(a.Observaciones) == 0 {
		t.Fatal("fallo es 0")
	}

}

func TestCantStartTramiteIfNotBorrador(t *testing.T) {

	a := createAggregate()
	a.Estado = "otro"

	docId := uuid.New()

	err := a.IniciarTramite(docId.String())

	if err == nil {
		t.Fatalf("It should return an error if trying to start tramite while not in %s status", EstadoTramiteBorrador)
	}
}

func TestCanStartTramiteIfBorrador(t *testing.T) {
	a := createAggregate()

	docId := uuid.New()

	err := a.IniciarTramite(docId.String())

	if err != nil {
		t.Fatalf("It should be possible to start tramite from %s state", EstadoTramiteBorrador)
	}
}

func createAggregate() Tramite {

	c := candidato.New("franco", "ganga", "asd@mail.com")

	materias := make([]*materia.Materia, 0)

	dep := dependencia.Dependencia{
		Nombre:      "IEI",
		AreaSudocu:  "2222222",
		Autorizante: user.New("admin"),
	}

	return New(c, 2022, materias, user.New("admin"), &dep, "CJ1")
}
