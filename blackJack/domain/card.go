package domain

type card struct {
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
