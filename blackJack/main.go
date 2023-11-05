package main

import (
	"fmt"

	"blackJack/domain"
	"blackJack/view"
)

func main() {
	fmt.Println("블랙잭 게임을 시작합니다.")
	names := view.GetParticipateNames()
	gamers := make([]domain.Gamer, len(names))
	view.PrintShare(names)
	deck := domain.NewDeck()
	for i, name := range names {
		var hands []domain.Card
		hands, deck = deck.Hands()
		gamers[i] = domain.Gamer{
			Name:  name,
			Cards: hands,
		}
		view.PrintHands(gamers[i])
	}

	// hands, hit, stay, burst, soft, hard
}
