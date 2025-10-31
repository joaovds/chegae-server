package dtos

import "time"

type (
	TripBase struct {
		ID        int       `json:"id"`
		StartedAt time.Time `json:"started_at"`
	}

	GetTripInput struct {
		ID int `json:"trip_id"`
	}

	GetTripOutput struct {
		TripBase
	}
)
