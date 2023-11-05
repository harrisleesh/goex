package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCard(t *testing.T) {
	c := Card{
		Suit:   clubs,
		number: 1,
	}

	assert.Equal(t, c.number, 1)
	assert.Equal(t, c.Suit, clubs)
}
