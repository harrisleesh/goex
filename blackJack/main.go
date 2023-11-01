package main

import (
	"blackJack/view"
	"fmt"
)

func main() {
	fmt.Println("블랙잭 게임을 시작합니다.")
	names := view.GetParticipateNames()
	fmt.Println(names)
	// hands, hit, stay, burst, soft, hard
}
