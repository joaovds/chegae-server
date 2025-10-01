package ports

import (
	"context"

	"github.com/joaovds/chegae-server/internal/shared/errs"
	"github.com/joaovds/chegae-server/internal/tracking/domain"
)

type TripStream interface {
	AddFollower(ctx context.Context, tripID string, conn domain.FollowerConn) errs.Error
	RemoveFollower(ctx context.Context, tripID string, conn domain.FollowerConn) errs.Error
	PublishLocation(ctx context.Context, tripID string, location *domain.Location) errs.Error
}
