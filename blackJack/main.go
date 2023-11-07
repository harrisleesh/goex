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
		hands := make([]domain.Card, 2)
		hands, deck = deck.Hands()
		gamers[i] = domain.Gamer{
			Name:  name,
			Cards: hands,
		}
		view.PrintHands(gamers[i])
	}
	fmt.Println()
	for i, _ := range gamers {
		for more := view.MoreCard(gamers[i].Name); more; more = view.MoreCard(gamers[i].Name) {
			var newCard domain.Card
			newCard, deck = deck.Hit()
			gamers[i].Take(newCard)
			view.PrintHands(gamers[i])
		}
		view.PrintHands(gamers[i])
	}
	fmt.Println()
	view.PrintResult(gamers)
	// hands, hit, stay, burst, soft, hard
}
