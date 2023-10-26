package domain

import "math/rand"

type RacingCar struct {
	Position int
}

func (c *RacingCar) move() {
	c.Position += 1
}

type RacingCars struct {
	Cars []*RacingCar
}

func (cars *RacingCars) move() {
	for _, c := range cars.Cars {
		c.move()
	}
}

func (c *RacingCar) moveOrNot() {
	if rand.Int()%2 == 0 {
		c.move()
	}
}

func (cars *RacingCars) MoveOrNot() {
	for _, c := range cars.Cars {
		c.moveOrNot()
	}
}
