package view

import (
	"fmt"
	"strings"
)

func PrintHands(names []string) {
	nameRes := strings.Join(names, ", ")
	fmt.Println(nameRes + "에게 2장의 카드를 나누었습니다.")
}
