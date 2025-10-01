package application

import (
	"github.com/joaovds/chegae-server/internal/shared/errs"
	"github.com/joaovds/chegae-server/internal/tracking/domain"
	"github.com/joaovds/chegae-server/internal/tracking/dtos"
	"github.com/joaovds/chegae-server/internal/tracking/ports"
)

type StartTripUC struct {
	repo   ports.TripRepository
	stream ports.TripStream
}

func NewStartTripUC(repo ports.TripRepository) *StartTripUC {
	return &StartTripUC{repo: repo}
}

func (uc *StartTripUC) Execute(input *dtos.StartTripInput) errs.Error {
	trip := &domain.Trip{Driver: domain.Driver{ID: input.DriverID}}
	err := uc.repo.Save(trip)
	return err
}
