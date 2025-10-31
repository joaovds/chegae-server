package dtos

import "time"

type (
	StartTripOutput struct {
		ID        int       `json:"trip_id"`
		StartedAt time.Time `json:"started_at"`
	}
)
