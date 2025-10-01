package mocks

import (
	"github.com/joaovds/chegae-server/internal/shared/errs"
	"github.com/joaovds/chegae-server/internal/tracking/domain"
	"github.com/stretchr/testify/mock"
)

type TripRepositoryMock struct {
	mock.Mock
}

func (m *TripRepositoryMock) Save(trip *domain.Trip) errs.Error {
	args := m.Called(trip)
	var err errs.Error
	if v := args.Get(0); v != nil {
		err = v.(errs.Error)
	}
	return err
}

func (m *TripRepositoryMock) GetByID(id string) (*domain.Trip, bool) {
	args := m.Called(id)
	return args.Get(0).(*domain.Trip), args.Bool(1)
}
