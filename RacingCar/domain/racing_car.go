package domain

import "math/rand"

type RacingCar struct {
	Name     string
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
func (cars *RacingCars) GetWinners() string {
	var maxPosition = cars.getMaxPosition()
	var winners string
	for _, c := range cars.Cars {
		if c.Position == maxPosition {
			if winners == "" {
				winners = c.Name
			} else {
				winners = winners + ", " + c.Name
			}
		}
	}
	return winners
}
func (cars *RacingCars) getMaxPosition() int {
	var maxPosition = 0
	for _, c := range cars.Cars {
		if c.Position > maxPosition {
			maxPosition = c.Position
		}
	}
	return maxPosition
}
