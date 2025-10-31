package tracking

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/joaovds/chegae-server/internal/tracking/dtos"
)

type TripHandler struct {
	service TripService
}

func NewTripHandler(service TripService) *TripHandler {
	return &TripHandler{service}
}

func (h *TripHandler) StartTrip(w http.ResponseWriter, r *http.Request) {
	if err := h.service.StartTrip(r.Context()); err != nil {
		http.Error(w, err.GetMessage(), err.GetCode())
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *TripHandler) GetTrip(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("trip_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid trip_id", http.StatusBadRequest)
		return
	}

	output, sErr := h.service.GetTrip(r.Context(), &dtos.GetTripInput{ID: id})
	if sErr != nil {
		http.Error(w, sErr.GetMessage(), sErr.GetCode())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}
