package view

import (
	"fmt"
	"strings"

	"blackJack/domain"
)

func PrintShare(names []string) {
	nameRes := strings.Join(names, ", ")
	fmt.Println(nameRes + "에게 2장의 카드를 나누었습니다.")
}

func PrintHands(gamer domain.Gamer) {
	fmt.Println(gamer.Name + "카드: " + PrintCards(gamer.Cards))
}

func PrintCards(cards []domain.Card) string {
	ss := make([]string, len(cards))
	for i, card := range cards {
		ss[i] = card.ToString()
	}
	return strings.Join(ss, ", ")
}

func PrintResult(gamers []domain.Gamer) {
	for _, gamer := range gamers {
		fmt.Println(gamer.Name+"카드:", PrintCards(gamer.Cards), "- 결과:", gamer.Result())
	}
}
