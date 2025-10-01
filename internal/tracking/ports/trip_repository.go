package ports

import (
	"github.com/joaovds/chegae-server/internal/shared/errs"
	"github.com/joaovds/chegae-server/internal/tracking/domain"
)

type TripRepository interface {
	Save(trip *domain.Trip) errs.Error
	GetByID(id string) (*domain.Trip, bool)
}
