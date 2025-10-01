package memory

import (
	"sync"

	"github.com/joaovds/chegae-server/internal/shared/errs"
	"github.com/joaovds/chegae-server/internal/tracking/domain"
)

type InMemoryTripRepo struct {
	data map[string]*domain.Trip
	mu   sync.Mutex
}

func NewInMemoryTripRepo() *InMemoryTripRepo {
	return &InMemoryTripRepo{
		data: make(map[string]*domain.Trip),
	}
}

func (r *InMemoryTripRepo) Save(trip *domain.Trip) errs.Error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.data[trip.ID] = trip
	return nil
}

func (r *InMemoryTripRepo) GetByID(id string) (*domain.Trip, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	t, ok := r.data[id]
	return t, ok
}
