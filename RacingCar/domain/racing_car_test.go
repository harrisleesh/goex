package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMoveRacingCar(t *testing.T) {
	racingCar := RacingCar{position: 0}
	racingCar.move()
	assert.Equal(t, 1, racingCar.position)
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
		assert.Equal(t, 1, car.position)
	}
}

func TestRandomMoveRacingCars(t *testing.T) {
	cars := RacingCars{
		[]*RacingCar{
			{0},
			{0},
			{0},
		},
	}
	cars.moveOrNot()
	cars.moveOrNot()
	for _, car := range cars.racingCars {
		assert.GreaterOrEqual(t, 2, car.position)
	}
}
