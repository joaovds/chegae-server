package tracking

import "net/http"

type (
	Module struct {
		services *services
	}

	services struct {
		tripService TripService
	}
)

func NewModule() *Module {
	tripRepo := NewInMemoryTripRepository()
	tripService := NewTripService(tripRepo)

	return &Module{
		services: &services{
			tripService: tripService,
		},
	}
}

func (m *Module) SetupHandlers(mux *http.ServeMux) {
	moduleMux := http.NewServeMux()
	mux.Handle("/tracking/", http.StripPrefix("/tracking", moduleMux))

	tripHandlers := NewTripHandler(m.services.tripService)

	moduleMux.HandleFunc("POST /trips", tripHandlers.StartTrip)
	moduleMux.HandleFunc("GET /trips", tripHandlers.GetTrip)
}
