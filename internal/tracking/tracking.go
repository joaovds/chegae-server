package tracking

import (
	"context"
	"time"

	"github.com/joaovds/chegae-server/internal/shared"
	"github.com/joaovds/chegae-server/internal/tracking/dtos"
)

type (
	Trip struct {
		ID        int
		StartedAt time.Time
	}

	Location struct {
		Lat       float64
		Lng       float64
		TripID    int
		Timestamp time.Time
	}
)

type (
	TripService interface {
		StartTrip(ctx context.Context) shared.Error
		GetTrip(ctx context.Context, input *dtos.GetTripInput) (*dtos.GetTripOutput, shared.Error)
	}

	TripRepository interface {
		Create(ctx context.Context, trip *Trip) shared.Error
		FindByID(ctx context.Context, id int) (*Trip, shared.Error)
	}
)
