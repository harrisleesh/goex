package view

import (
	"Lotto/lotto/domain"
	"fmt"
)

func PrintLottos(lottos domain.Lottos) {
	for _, lotto := range lottos.Lottos {
		fmt.Println(lotto.Numbers)
	}
}

func PrintResult(result *domain.LottoResult) {
	fmt.Println("당첨 통계\n---------")
	fmt.Println("3개 일치 (5000원)-", result.Three, "개")
	fmt.Println("4개 일치 (50000원)-", result.Four, "개")
	fmt.Println("5개 일치 (1500000원)-", result.Five, "개")
	fmt.Println("개 일치, 보너스 볼 일치(30000000원)-", result.FiveWithBonus, "개")
	fmt.Println("6개 일치 (2000000000원)-", result.Six, "개")
	fmt.Println("총 수익률은", result.WinningRate, "입니다.(기준이 1이기 때문에 결과적으로 손해라는 의미임)")
}
