package domain

import "sync"

type Trip struct {
	ID        string
	Driver    Driver
	Followers map[int]Follower
	mu        sync.Mutex
}

func NewTrip(id string, driver Driver) *Trip {
	return &Trip{
		ID:        id,
		Driver:    driver,
		Followers: make(map[int]Follower),
	}
}

func (t *Trip) AddFollower(f Follower) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.Followers[f.ID] = f
}
