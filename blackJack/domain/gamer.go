package domain

type Gamer struct {
	Name  string
	Cards []Card
}

func (g *Gamer) Take(c Card) {
	g.Cards = append(g.Cards, c)
}

func (g Gamer) Result() (total int64) {
	hasAce := false
	for _, card := range g.Cards {
		if card.number > 10 {
			total += 10
		} else {
			total = total + int64(card.number)
		}
		if card.number == 1 {
			hasAce = true
		}
	}
	if total <= 11 && hasAce {
		total += 10
	}
	return total
}
