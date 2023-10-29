package main

import (
	"RacingCar/domain"
	"RacingCar/view"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("경주할 자동차 이름을 입력하세요(이름은 쉼표(,)를 기준으로 구분).")

	var carNames string
	_, err := fmt.Scanf("%s", &carNames)
	names := strings.Split(carNames, ",")
	if err != nil {
		return
	}
	var tryCount int
	fmt.Println("시도할 횟수는 몇 회인가요?")
	_, err = fmt.Scanf("%d", &tryCount)
	if err != nil {
		return
	}

	cars := domain.RacingCars{Cars: make([]*domain.RacingCar, len(names))}
	for i := 0; i < len(names); i++ {
		cars.Cars[i] = &domain.RacingCar{
			Position: 0,
			Name:     names[i],
		}
	}

	fmt.Println("실행 결과")
	for i := 0; i < tryCount; i++ {
		cars.MoveOrNot()
		view.PrintPositions(&cars)
	}
	view.PrintWinners(cars.GetWinners())
}
