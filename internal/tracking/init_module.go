package tracking

import "net/http"

type (
	Module struct {
		services *services
	}

	services struct {
		tripService     TripService
		trackingService TrackingService
	}
)

func NewModule() *Module {
	tripRepo := NewInMemoryTripRepository()
	tripService := NewTripService(tripRepo)
	trackingService := NewTrackingService()

	return &Module{
		services: &services{
			tripService:     tripService,
			trackingService: trackingService,
		},
	}
}

func (m *Module) SetupHandlers(mux *http.ServeMux) {
	moduleMux := http.NewServeMux()
	mux.Handle("/tracking/", http.StripPrefix("/tracking", moduleMux))

	tripHandlers := NewTripHandler(m.services.tripService)
	moduleMux.HandleFunc("POST /trips", tripHandlers.StartTrip)
	moduleMux.HandleFunc("GET /trips/{trip_id}", tripHandlers.GetTrip)

	trackingHandlers := NewTrackingHandler(m.services.trackingService)
	moduleMux.HandleFunc("/ws/{trip_id}", trackingHandlers.ReceiveLiveLocationsWS)
}
