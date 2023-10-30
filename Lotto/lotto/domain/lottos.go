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
func (l Lottos) DrawResults(winners []int64) *LottoResult {
	result := &LottoResult{}
	for _, lotto := range l.Lottos {
		count := int64(0)
		for _, winner := range winners {
			if contains(lotto.Numbers, winner) {
				count += 1
			}
		}
		result.AddResult(count)
	}
	result.AddResultRate(int64(len(l.Lottos)))
	return result
}

func contains(s []int64, e int64) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
