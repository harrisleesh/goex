package main

import (
	"Lotto/lotto/view"
	"fmt"
	"strconv"
)

func main() {
	amount := view.GetPurchaseAmount()
	lottoCount := amount / 1000
	fmt.Println(strconv.Itoa(int(lottoCount)) + "개를 구매했습니다.")
	view.GetWinnerNumbers()
}
