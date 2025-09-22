package domain

import "time"

type Location struct {
	Lat       float64
	Lon       float64
	Timestamp time.Time
}
