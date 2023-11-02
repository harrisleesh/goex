package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCard(t *testing.T) {
	c := card{
		suit:   clubs,
		number: 1,
	}

	assert.Equal(t, c.number, 1)
	assert.Equal(t, c.suit, clubs)
}
