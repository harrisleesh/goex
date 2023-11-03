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
	fmt.Println(gamer.Name + "카드" + PrintCards(gamer.Cards))
}

func PrintCards(cards []domain.Card) {
}
