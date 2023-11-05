package domain

import (
	"math/rand"
	"time"
)

type Deck []Card

func NewDeck() Deck {
	cards := make([]Card, 0)
	for i := 1; i < 5; i++ {
		for j := 1; j < 14; j++ {
			cards = append(cards, Card{
				Suit:   suit(i),
				number: j,
			})
		}
	}
	Deck(cards).shuffle()
	return cards
}

func (d Deck) shuffle() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < len(d); i++ {
		newPos := r.Intn(52)
		d[i], d[newPos] = d[newPos], d[i]
	}
}

func (d Deck) Hands() ([]Card, Deck) {
	return d[0:2], d[2:]
}
