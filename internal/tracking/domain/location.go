package domain

import "time"

type Location struct {
	Lat       float64
	Lng       float64
	Timestamp time.Time
}
