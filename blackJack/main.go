package main

import (
	"fmt"
	"strings"

	"blackJack/domain"
	"blackJack/view"
)

func main() {
	fmt.Println("블랙잭 게임을 시작합니다.")
	names := view.GetParticipateNames()
	view.PrintHands(names)
	deck := domain.NewDeck()
	for _, name := range names {
		view.
	}
	deck.hands()
	// hands, hit, stay, burst, soft, hard
}
