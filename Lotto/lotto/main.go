package main

import (
	"Lotto/lotto/domain"
	"Lotto/lotto/view"
	"fmt"
	"strconv"
)

func main() {
	amount := view.GetPurchaseAmount()
	lottoCount := amount / 1000
	fmt.Println(strconv.Itoa(int(lottoCount)) + "개를 구매했습니다.")
	lottos := domain.NewLottos(lottoCount)
	view.PrintLottos(lottos)
	view.GetWinnerNumbers()
}
