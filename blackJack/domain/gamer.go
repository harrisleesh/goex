package domain

type Gamer struct {
	Name  string
	Cards []Card
}

func (g *Gamer) Take(c Card) {
	g.Cards = append(g.Cards, c)
}
