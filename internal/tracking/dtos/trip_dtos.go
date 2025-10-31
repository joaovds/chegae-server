package dtos

import "time"

type TripBase struct {
	ID        int       `json:"id"`
	StartedAt time.Time `json:"started_at"`
}

type (
	GetTripInput struct {
		ID int `json:"trip_id"`
	}

	GetTripOutput struct {
		TripBase
	}
)

type (
	StartTripOutput struct {
		ID        int       `json:"trip_id"`
		StartedAt time.Time `json:"started_at"`
	}
)
