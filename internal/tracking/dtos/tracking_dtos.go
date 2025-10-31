package dtos

import "time"

type (
	ReceiveLiveLocationsInput struct {
		Lat       float64   `json:"lat"`
		Lng       float64   `json:"lgn"`
		TripID    int       `json:"trip_id"`
		Timestamp time.Time `json:"timestamp"`
	}
)
