package ws

import (
	"context"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/joaovds/chegae-server/internal/shared/errs"
	"github.com/joaovds/chegae-server/internal/tracking/domain"
)

type WsFollowerConn struct {
	FollowerID int
	Conn       *websocket.Conn
}

func (w *WsFollowerConn) GetFollower() *domain.Follower { return &domain.Follower{ID: w.FollowerID} }

func (w *WsFollowerConn) Send(ctx context.Context, location *domain.Location) errs.Error {
	err := w.Conn.WriteJSON(location)
	if err != nil {
		return errs.NewErr("Err sending JSON:" + err.Error())
	}
	return nil
}

func (w *WsFollowerConn) Close() errs.Error {
	err := w.Conn.Close()
	if err != nil {
		return errs.NewErr("Err close connection:" + err.Error())
	}
	return nil
}

// ----- .. -----

type TripStreamWS struct {
	mu        sync.RWMutex
	followers map[string]map[int]domain.FollowerConn
}

func NewTripStreamWS() *TripStreamWS {
	return &TripStreamWS{
		followers: make(map[string]map[int]domain.FollowerConn),
	}
}

func (s *TripStreamWS) AddFollower(ctx context.Context, tripID string, conn domain.FollowerConn) errs.Error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.followers[tripID] == nil {
		s.followers[tripID] = make(map[int]domain.FollowerConn)
	}

	s.followers[tripID][conn.GetFollower().ID] = conn
	return nil
}

func (s *TripStreamWS) RemoveFollower(ctx context.Context, tripID string, conn domain.FollowerConn) errs.Error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if conns, ok := s.followers[tripID]; ok {
		delete(conns, conn.GetFollower().ID)
	}
	return nil
}

func (s *TripStreamWS) PublishLocation(ctx context.Context, tripID string, location *domain.Location) errs.Error {
	s.mu.RLock()
	defer s.mu.RUnlock()

	conns, ok := s.followers[tripID]
	if !ok {
		return nil
	}

	for id, conn := range conns {
		if err := conn.Send(ctx, location); err != nil {
			conn.Close()
			delete(conns, id)
		}
	}

	return nil
}
