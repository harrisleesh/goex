package view

import (
	"RacingCar/domain"
	"fmt"
)

func PrintPositions(cars *domain.RacingCars) {
	for _, car := range cars.Cars {
		for i := 0; i < car.Position; i++ {
			fmt.Print("-")
		}
		fmt.Println()
	}
	fmt.Println()
}
