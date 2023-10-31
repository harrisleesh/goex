package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLottosCreate(t *testing.T) {
	lottos := NewLottos(2)
	assert.Equal(t, len(lottos.Lottos), 2)
}

func TestLottosCreateManual(t *testing.T) {
	lottos := NewLottosManual([][]int64{
		{1, 2, 3, 4, 5, 6},
		{1, 2, 3, 4, 5, 7},
	})
	assert.Equal(t, len(lottos.Lottos), 2)
}

func TestAppendLottos(t *testing.T) {
	lottos := NewLottos(2).Append(NewLottos(2))
	assert.Equal(t, len(lottos.Lottos), 4)
}
