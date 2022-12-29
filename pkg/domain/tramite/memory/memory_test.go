package memory

import (
	"fmt"
	"testing"

	"github.com/francoganga/go_reno2/pkg/aggregate"
	"github.com/francoganga/go_reno2/pkg/domain/candidato"
	"github.com/francoganga/go_reno2/pkg/domain/dependencia"
	"github.com/francoganga/go_reno2/pkg/domain/materia"
	_ "github.com/francoganga/go_reno2/pkg/domain/tramite"
	"github.com/francoganga/go_reno2/pkg/domain/user"
	"github.com/google/uuid"
)

func TestMemoryGetTramite(t *testing.T) {

	type testCase struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}

	c := candidato.New("candidato1", "apellido1", "test@mail.com")

	ms := make([]*materia.Materia, 0)

	a := aggregate.New(&c, 2022, ms, user.New("admin"), &dependencia.Dependencia{Nombre: "ING", AreaSudocu: "aaa", Autorizante: user.New("admin")}, "IEI")
}
