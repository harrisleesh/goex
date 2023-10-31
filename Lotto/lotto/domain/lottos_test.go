package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLottosCreate(t *testing.T) {
	lottos := NewLottos(2)
	assert.Equal(t, len(lottos.Lottos), 2)
}
