package tracking

import (
	"context"
	"fmt"
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

type trackingService struct{}

func NewTrackingService() TrackingService {
	return &trackingService{}
}

func (s *trackingService) ReceiveLiveLocations(ctx context.Context, tripID int, updates <-chan dtos.ReceiveLiveLocationsInput) shared.Error {
	for {
		select {
		case <-ctx.Done():
			return nil
		case loc, ok := <-updates:
			if !ok {
				return nil
			}

			fmt.Printf("[Trip %d] new location received: %.6f, %.6f\n", loc.TripID, loc.Lat, loc.Lng)
		}
	}
}
