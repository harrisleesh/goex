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
