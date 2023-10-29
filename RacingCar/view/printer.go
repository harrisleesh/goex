package view

import (
	"RacingCar/domain"
	"fmt"
)

func PrintPositions(cars *domain.RacingCars) {
	for _, car := range cars.Cars {
		fmt.Print(car.Name + " : ")
		for i := 0; i < car.Position; i++ {
			fmt.Print("-")
		}
		fmt.Println()
	}
	fmt.Println()
}

func PrintWinners(winners string) {
	fmt.Println(winners + "가 최종 우승했습니다.")
}
