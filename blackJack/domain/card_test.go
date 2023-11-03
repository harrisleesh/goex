package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCard(t *testing.T) {
	c := Card{
		suit:   clubs,
		number: 1,
	}

	assert.Equal(t, c.number, 1)
	assert.Equal(t, c.suit, clubs)
}
