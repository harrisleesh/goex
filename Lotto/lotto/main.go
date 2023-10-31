package main

import (
	"Lotto/lotto/domain"
	"Lotto/lotto/view"
	"fmt"
	"strconv"
)

func main() {
	amount := view.GetPurchaseAmount()
	manualCount := view.GetManualCount()
	manualNumbers := view.GetManualLottoNumbers(manualCount)
	autoCount := amount/1000 - manualCount
	fmt.Println("수동으로 "+strconv.Itoa(int(manualCount))+"개, 자동으로", strconv.Itoa(int(autoCount))+"개를 구매했습니다.")
	lottosManual := domain.NewLottosManual(manualNumbers)
	lottosAuto := domain.NewLottos(autoCount)
	lottos := lottosManual.Append(lottosAuto)
	view.PrintLottos(lottos)
	winnerNumbers, bonus := view.GetWinnerNumbers()
	fmt.Println(bonus)
	result := lottos.DrawResults(winnerNumbers, bonus)
	view.PrintResult(result)
}
