package domain

type LottoResult struct {
	Three       int64
	Four        int64
	Five        int64
	Six         int64
	WinningRate float64
}

func (l *LottoResult) AddResult(matchCount int64) {
	if matchCount == 3 {
		l.Three += 1
	} else if matchCount == 4 {
		l.Four += 1
	} else if matchCount == 5 {
		l.Five += 1
	} else if matchCount == 6 {
		l.Six += 1
	}
}

func (l *LottoResult) AddResultRate(count int64) {
	l.WinningRate = float64(l.Three*5000+l.Four*50000+l.Five*1500000+l.Six*2000000000) / float64(1000*count)
}
