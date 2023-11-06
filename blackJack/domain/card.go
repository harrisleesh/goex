package domain

import (
	"strconv"
)

type Card struct {
	Suit   suit
	number int
}

type suit int

const (
	spade suit = iota + 1
	heart
	clubs
	diamond
)

func (s suit) String() string {
	switch s {
	case spade:
		return "스페이드"
	case heart:
		return "하트"
	case clubs:
		return "클로버"
	case diamond:
		return "다이아몬드"
	default:
		return ""
	}
}

func (c Card) ToString() string {
	switch c.number {
	case 1:
		return "A" + c.Suit.String()
	case 10:
		return "J" + c.Suit.String()
	case 11:
		return "Q" + c.Suit.String()
	case 12:
		return "K" + c.Suit.String()
	}
	return strconv.FormatInt(int64(c.number), 10) + c.Suit.String()
}
