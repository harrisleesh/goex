package view

import (
	"fmt"
	"strconv"
	"strings"
)

func GetPurchaseAmount() int64 {
	fmt.Println("구입금액을 입력해 주세요.")
	var amount int64
	_, err := fmt.Scanf("%d", &amount)
	if err != nil {
		return 0
	}
	return amount
}

func GetWinnerNumbers() ([]int64, int64) {
	fmt.Println("지난 주 당첨 번호를 입력해 주세요.")
	var s string
	_, err := fmt.Scanln(&s)
	if err != nil {
		return nil, 0
	}
	winnerStrings := strings.Split(s, ",")
	winners := make([]int64, len(winnerStrings))
	for i, winner := range winnerStrings {
		winners[i], _ = strconv.ParseInt(winner, 10, 0)
	}
	fmt.Println("보너스 볼을 입력해 주세요.")
	var bonus int64
	_, err = fmt.Scanf("%d", &bonus)
	if err != nil {
		return nil, 0
	}
	return winners, bonus
}
