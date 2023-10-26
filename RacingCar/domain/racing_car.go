package domain

import "math/rand"

type RacingCar struct {
	position int
}

func (c *RacingCar) move() {
	c.position += 1
}

type RacingCars struct {
	racingCars []*RacingCar
}

func (cars *RacingCars) move() {
	for _, c := range cars.racingCars {
		c.move()
	}
}

func (c *RacingCar) moveOrNot() {
	if rand.Int()%2 == 0 {
		c.move()
	}
}

func (cars *RacingCars) moveOrNot() {
	for _, c := range cars.racingCars {
		c.moveOrNot()
	}
}
