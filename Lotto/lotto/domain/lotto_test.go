package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLottoCreate(t *testing.T) {
	lotto := Lotto{
		Numbers: []int64{1, 2, 3, 4, 5, 6},
	}

	assert.Equal(t, lotto.Numbers, []int64{1, 2, 3, 4, 5, 6})
	assert.Equal(t, len(newLotto().Numbers), 6)
}
