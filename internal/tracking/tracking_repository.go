package tracking

import (
	"context"
	"sync"

	"github.com/joaovds/chegae-server/internal/shared"
)

type InMemoryTripRepository struct {
	mu     sync.RWMutex
	trips  map[int]*Trip
	nextID int
}

func NewInMemoryTripRepository() *InMemoryTripRepository {
	return &InMemoryTripRepository{
		trips:  make(map[int]*Trip),
		nextID: 1,
	}
}

func (r *InMemoryTripRepository) Create(ctx context.Context, trip *Trip) shared.Error {
	r.mu.Lock()
	defer r.mu.Unlock()

	trip.ID = r.nextID
	r.nextID++

	r.trips[trip.ID] = trip
	return nil
}

func (r *InMemoryTripRepository) FindByID(ctx context.Context, id int) (*Trip, shared.Error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	trip, ok := r.trips[id]
	if !ok {
		return nil, shared.NewErr("trip not found!").SetCode(404)
	}

	return trip, nil
}
