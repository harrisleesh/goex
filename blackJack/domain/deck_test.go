package domain

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDeck(t *testing.T) {
	cards := make([]Card, 0)
	for i := 1; i < 5; i++ {
		for j := 1; j < 14; j++ {
			cards = append(cards, Card{
				Suit:   suit(i),
				number: j,
			})
		}
	}
	tests := []struct {
		name string
		want Deck
	}{

		{
			name: "deck 을 생성한다.",
			want: cards,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDeck(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDeck() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeck_hands(t *testing.T) {
	tests := []struct {
		name    string
		d       Deck
		want    []Card
		remains Deck
	}{
		{
			name: "Hands test",
			d:    NewDeck(),
			want: []Card{
				{
					Suit:   spade,
					number: 1,
				},
				{
					Suit:   spade,
					number: 2,
				},
			},
			remains: NewDeck()[2:],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, remains := tt.d.Hands()
			assert.Equalf(t, tt.want, got, "Hands()")
			assert.Equalf(t, tt.remains, remains, "Hands()")
		})
	}
}
