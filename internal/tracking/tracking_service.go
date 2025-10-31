package tracking

import (
	"context"
	"log"
	"slices"
	"sync"
	"time"

	"github.com/joaovds/chegae-server/internal/shared"
	"github.com/joaovds/chegae-server/internal/tracking/dtos"
)

type tripService struct {
	repo TripRepository
}

func NewTripService(repo TripRepository) TripService {
	return &tripService{repo}
}

func (s *tripService) StartTrip(ctx context.Context) (*dtos.StartTripOutput, shared.Error) {
	trip := &Trip{
		StartedAt: time.Now(),
	}
	if err := s.repo.Create(ctx, trip); err != nil {
		return nil, err
	}
	return &dtos.StartTripOutput{
		ID:        trip.ID,
		StartedAt: trip.StartedAt,
	}, nil
}

func (s *tripService) GetTrip(ctx context.Context, input *dtos.GetTripInput) (*dtos.GetTripOutput, shared.Error) {
	trip, err := s.repo.FindByID(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	return &dtos.GetTripOutput{
		TripBase: dtos.TripBase{
			ID:        trip.ID,
			StartedAt: trip.StartedAt,
		},
	}, nil
}

// ----- .. -----

type trackingService struct {
	mu          sync.RWMutex
	subscribers map[int][]Client
}

func NewTrackingService() TrackingService {
	return &trackingService{
		subscribers: make(map[int][]Client),
	}
}

func (s *trackingService) StreamLiveLocations(ctx context.Context, tripID int, updates <-chan dtos.LiveLocations) shared.Error {
	for {
		select {
		case <-ctx.Done():
			return nil
		case loc, ok := <-updates:
			if !ok {
				return nil
			}

			s.mu.RLock()
			clients := s.subscribers[tripID]
			s.mu.RUnlock()

			for _, client := range clients {
				if err := client.GetConn().SendLocationUpdate(loc); err != nil {
					log.Printf("failed to send location to clientID %d", client.GetID())
				}
			}
		}
	}
}

func (s *trackingService) TrackLiveLocations(ctx context.Context, tripID int, client Client) shared.Error {
	s.mu.Lock()
	s.subscribers[tripID] = append(s.subscribers[tripID], client)
	s.mu.Unlock()

	go func() {
		<-ctx.Done()
		s.mu.Lock()
		defer s.mu.Unlock()

		clients := s.subscribers[tripID]
		s.subscribers[tripID] = slices.DeleteFunc(clients, func(c Client) bool {
			return c.GetID() == client.GetID()
		})
	}()

	return nil
}
