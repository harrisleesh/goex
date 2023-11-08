package domain

type Gamer struct {
	Name  string
	Cards []Card
	Point Point
}
type Point struct {
	winCount  int64
	loseCount int64
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
	if total > 21 {
		return -1
	}
	if total <= 11 && hasAce {
		total += 10
	}
	return total
}

func WinningPoint(dealer Gamer, gamers []Gamer) map[string][]int64 {
	if dealer.Result() < 0 {

	}
	for _, gamer := range gamers {
	}
}
