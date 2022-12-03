package memory

import (
	"fmt"
	"sync"

	"github.com/francoganga/go_reno/aggregate"
	"github.com/francoganga/go_reno/domain/tramite"
	"github.com/google/uuid"
)



type MemoryRepository struct {
    tramites map[uuid.UUID]aggregate.Tramite
    sync.Mutex
}

func New() *MemoryRepository {
    return &MemoryRepository{
        tramites: make(map[uuid.UUID]aggregate.Tramite),
    }
}

func (mr *MemoryRepository) Get(id uuid.UUID) (aggregate.Tramite, error) {
    if tramite, ok := mr.tramites[id]; ok {
        return tramite, nil
    }

    return aggregate.Tramite{}, tramite.ErrTramiteNotFound
}

func (mr *MemoryRepository) Add(t aggregate.Tramite) error {
    if mr.tramites == nil {
        mr.Lock()
        mr.tramites = make(map[uuid.UUID]aggregate.Tramite)
        mr.Unlock()
    }

    if _, ok := mr.tramites[t.GetID()]; ok {
        return fmt.Errorf("Tramite ya existe: %w", tramite.ErrFailedToAddTramite)
    }

    mr.Lock()
    mr.tramites[t.GetID()] = t
    mr.Unlock()

    return nil
}

func (mr *MemoryRepository) Update(t aggregate.Tramite) error {
    if _, ok := mr.tramites[t.GetID()]; !ok {
        return fmt.Errorf("tramite no existe: %w", tramite.ErrUpdateTramite)
    }

    mr.Lock()
    mr.tramites[t.GetID()] = t
    mr.Unlock()

    return nil
}
