package tracking

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
	"github.com/joaovds/chegae-server/internal/tracking/dtos"
)

type TripHandler struct {
	service TripService
}

func NewTripHandler(service TripService) *TripHandler {
	return &TripHandler{service}
}

func (h *TripHandler) StartTrip(w http.ResponseWriter, r *http.Request) {
	output, err := h.service.StartTrip(r.Context())
	if err != nil {
		http.Error(w, err.GetMessage(), err.GetCode())
		return
	}
	json.NewEncoder(w).Encode(output)
}

func (h *TripHandler) GetTrip(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("trip_id")
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

	json.NewEncoder(w).Encode(output)
}

// ----- .. -----

type TrackingHandler struct {
	service TrackingService
}

func NewTrackingHandler(service TrackingService) *TrackingHandler {
	return &TrackingHandler{service}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func (h *TrackingHandler) ReceiveLiveLocationsWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "failed to start WebSocket", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	tripID, err := strconv.Atoi(r.PathValue("trip_id"))
	if err != nil {
		http.Error(w, "invalid trip_id", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	updates := make(chan dtos.ReceiveLiveLocationsInput)

	go func() {
		h.service.ReceiveLiveLocations(ctx, tripID, updates)
	}()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			close(updates)
			return
		}

		var loc dtos.ReceiveLiveLocationsInput
		if err := json.Unmarshal(msg, &loc); err != nil {
			continue
		}

		updates <- loc
	}
}
