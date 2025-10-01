package application

import (
	"testing"

	"github.com/joaovds/chegae-server/internal/shared/errs"
	"github.com/joaovds/chegae-server/internal/tracking/dtos"
	"github.com/joaovds/chegae-server/internal/tracking/ports/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func makeSUT() (*StartTripUC, *mocks.TripRepositoryMock) {
	mockRepo := new(mocks.TripRepositoryMock)
	sut := NewStartTripUC(mockRepo)
	return sut, mockRepo
}

func TestStartTripUC_Execute(t *testing.T) {
	t.Run("successful trip start", func(t *testing.T) {
		startTripUC, mockRepo := makeSUT()
		input := &dtos.StartTripInput{DriverID: 123}

		mockRepo.On("Save", mock.Anything).Return(nil)

		assert.Nil(t, startTripUC.Execute(input))

		mockRepo.AssertExpectations(t)
	})

	t.Run("error on trip save", func(t *testing.T) {
		startTripUC, mockRepo := makeSUT()
		input := &dtos.StartTripInput{DriverID: 123}

		mockRepo.On("Save", mock.Anything).Return(errs.NewErr("Save error"))

		err := startTripUC.Execute(input)

		assert.NotNil(t, err)
		assert.Equal(t, "Save error", err.GetMessage())

		mockRepo.AssertExpectations(t)
	})
}
