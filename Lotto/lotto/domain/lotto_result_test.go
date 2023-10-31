package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLottoResult_AddResult(t *testing.T) {

	l := Lottos{
		Lottos: []Lotto{
			{
				Numbers: []int64{1, 2, 3, 4, 5, 6},
			},
			{
				Numbers: []int64{1, 2, 3, 4, 5, 7},
			},
		},
	}
	r := l.DrawResults([]int64{1, 2, 3, 4, 5, 6}, 7)
	assert.Equal(t, r.Five, int64(0))
	assert.Equal(t, r.Six, int64(1))
	assert.Equal(t, r.FiveWithBonus, int64(1))
}
