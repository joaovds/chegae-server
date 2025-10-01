package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrip(t *testing.T) {
	t.Run("Test new trip", func(t *testing.T) {
		driver := Driver{ID: 1, Name: "any_driver"}
		new_trip := NewTrip("any_trip_id", driver)
		assert.Equal(t, "any_trip_id", new_trip.ID)
		assert.Equal(t, driver, new_trip.Driver)
	})

	t.Run("Test add follower", func(t *testing.T) {
		driver := Driver{ID: 1, Name: "any_driver"}
		trip := NewTrip("any_trip_id", driver)

		assert.Len(t, trip.Followers, 0)

		f1 := Follower{ID: 1}
		f2 := Follower{ID: 2}

		trip.AddFollower(f1)
		trip.AddFollower(f2)

		assert.Len(t, trip.Followers, 2)
		assert.Equal(t, "any_trip_id", trip.ID)
		assert.Equal(t, driver, trip.Driver)
	})
}
