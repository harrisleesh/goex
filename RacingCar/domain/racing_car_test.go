package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMoveRacingCar(t *testing.T) {
	racingCar := RacingCar{position: 0}
	racingCar.move()
	assert.Equal(t, racingCar.position, 1)
}

func TestMoveRacingCars(t *testing.T) {
	cars := RacingCars{
		[]*RacingCar{
			{0},
			{0},
			{0},
		},
	}
	cars.move()
	for _, car := range cars.racingCars {
		assert.Equal(t, car.position, 1)
	}
}
