package domain

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
