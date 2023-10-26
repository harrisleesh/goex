package main

import "fmt"

func main() {
	fmt.Println("자동차 대수는 몇 대인가요?")
	var carCount, tryCount int
	_, err := fmt.Scanf("%d", &carCount)
	if err != nil {
		return
	}
	fmt.Println("시도할 횟수는 몇 회인가요?")
	_, err = fmt.Scanf("%d", &tryCount)
	if err != nil {
		return
	}

	fmt.Println("실행 결과")
}
