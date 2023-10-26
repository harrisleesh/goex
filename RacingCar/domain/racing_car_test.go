package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMoveRacingCar(t *testing.T) {
	racingCar := RacingCar{}
	racingCar.move()
	assert.Equal(t, 1, racingCar.Position)
}

func TestMoveRacingCars(t *testing.T) {
	cars := RacingCars{
		make([]*RacingCar, 3),
	}
	for i := 0; i < 3; i++ {
		cars.Cars[i] = &RacingCar{
			Position: 0,
		}
	}

	cars.move()
	for _, car := range cars.Cars {
		assert.Equal(t, 1, car.Position)
	}
}

func TestRandomMoveRacingCars(t *testing.T) {
	cars := RacingCars{
		make([]*RacingCar, 3),
	}
	for i := 0; i < 3; i++ {
		cars.Cars[i] = &RacingCar{
			Position: 0,
		}
	}
	cars.MoveOrNot()
	cars.MoveOrNot()
	for _, car := range cars.Cars {
		assert.GreaterOrEqual(t, 2, car.Position)
	}
}
