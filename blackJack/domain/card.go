package domain

type Card struct {
	suit   suit
	number int
}

type suit int

const (
	spade suit = iota + 1
	heart
	clubs
	diamond
)

func (c Card) toString() string {

}
