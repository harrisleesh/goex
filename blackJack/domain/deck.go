package domain

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
	return cards
}

func (d Deck) Hands() ([]Card, Deck) {
	return d[0:2], d[2:]
}
