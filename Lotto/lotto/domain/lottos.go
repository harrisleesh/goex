package domain

type Lottos struct {
	Lottos []Lotto
}

func NewLottos(count int64) Lottos {
	lottos := make([]Lotto, count)
	for i := 0; i < int(count); i++ {
		lottos[i] = newLotto()
	}
	return Lottos{
		Lottos: lottos,
	}
}

func NewLottosManual(numbers [][]int64) Lottos {
	lottos := make([]Lotto, len(numbers))
	for i := 0; i < len(numbers); i++ {
		lottos[i] = newLottoManual(numbers[i])
	}
	return Lottos{
		Lottos: lottos,
	}
}
func (l Lottos) DrawResults(winners []int64, bonus int64) *LottoResult {
	result := &LottoResult{}
	for _, lotto := range l.Lottos {
		count := int64(0)
		for _, winner := range winners {
			if contains(lotto.Numbers, winner) {
				count += 1
			}
		}
		if count == 5 && contains(lotto.Numbers, bonus) {
			result.AddSecond()
		} else {
			result.AddResult(count)
		}
	}
	result.AddResultRate(int64(len(l.Lottos)))
	return result
}

func (l1 Lottos) Append(l2 Lottos) Lottos {
	return Lottos{
		Lottos: append(l1.Lottos, l2.Lottos...),
	}
}

func contains(s []int64, e int64) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
