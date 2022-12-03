package memory

import (
	"testing"

	"github.com/francoganga/go_reno/aggregate"
	"github.com/francoganga/go_reno/domain/candidato"
	"github.com/francoganga/go_reno/domain/tramite"
	"github.com/google/uuid"
)

func TestMemoryGetTramite(t *testing.T) {

    type testCase struct {
        name string
        id   uuid.UUID
        expectedErr error
    }

    c := candidato.New("candidato1", "apellido1", "test@mail.com")
}
