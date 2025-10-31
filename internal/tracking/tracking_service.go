package tracking

import (
	"context"
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

func (s *tripService) StartTrip(ctx context.Context) shared.Error {
	trip := &Trip{
		StartedAt: time.Now(),
	}
	if err := s.repo.Create(ctx, trip); err != nil {
		return err
	}
	return nil
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
