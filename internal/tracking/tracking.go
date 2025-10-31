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

	Client interface {
		GetID() int
		GetConn() LiveConnection
	}

	LiveConnection interface {
		SendLocationUpdate(loc dtos.LiveLocations) shared.Error
	}
)

type (
	TripService interface {
		StartTrip(ctx context.Context) (*dtos.StartTripOutput, shared.Error)
		GetTrip(ctx context.Context, input *dtos.GetTripInput) (*dtos.GetTripOutput, shared.Error)
	}

	TripRepository interface {
		Create(ctx context.Context, trip *Trip) shared.Error
		FindByID(ctx context.Context, id int) (*Trip, shared.Error)
	}

	TrackingService interface {
		StreamLiveLocations(ctx context.Context, tripID int, updates <-chan dtos.LiveLocations) shared.Error
		TrackLiveLocations(ctx context.Context, tripID int, client Client) shared.Error
	}
)
