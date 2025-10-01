package domain

import (
	"context"

	"github.com/joaovds/chegae-server/internal/shared/errs"
)

type Follower struct {
	ID int
}

type FollowerConn interface {
	GetFollower() *Follower
	Send(ctx context.Context, location *Location) errs.Error
	Close() errs.Error
}
