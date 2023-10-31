package domain

import (
	"math/rand"
	"time"
)

type Lotto struct {
	Numbers []int64
}

func newLotto() Lotto {
	return Lotto{
		Numbers: lottoNumbers()[0:6],
	}
}

func newLottoManual(numbers []int64) Lotto {
	return Lotto{
		Numbers: numbers,
	}
}
func lottoNumbers() []int64 {
	arr := make([]int64, 45)
	for i := 0; i < 45; i++ {
		arr[i] = int64(i + 1)
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
	return arr
}
