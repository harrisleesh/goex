package domain

type Gamer struct {
	Name  string
	Cards []Card
	Point *Point
}
type Point struct {
	WinCount  int64
	LoseCount int64
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

func ApplyWinningPoint(dealer Gamer, gamers []Gamer) {
	if dealer.Result() < 0 {
		for _, gamer := range gamers {
			gamer.Point.WinCount += 1
			dealer.Point.LoseCount += 1
		}
		return
	}
	for _, gamer := range gamers {
		if gamer.Result() > dealer.Result() {
			gamer.Point.WinCount += 1
			dealer.Point.LoseCount += 1
		} else {
			gamer.Point.LoseCount += 1
			dealer.Point.WinCount += 1
		}
	}
}
