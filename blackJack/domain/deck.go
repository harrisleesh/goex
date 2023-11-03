package domain

type Deck []card

func NewDeck() Deck {
	cards := make([]card, 0)
	for i := 1; i < 5; i++ {
		for j := 1; j < 14; j++ {
			cards = append(cards, card{
				suit:   suit(i),
				number: j,
			})
		}
	}
	return cards
}

func (d Deck) hands() ([]card, Deck) {
	return d[0:2], d[2:]
}
