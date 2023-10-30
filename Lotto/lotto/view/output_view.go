package view

import (
	"Lotto/lotto/domain"
	"fmt"
)

func PrintLottos(lottos []domain.Lotto) {
	for _, lotto := range lottos {
		fmt.Println(lotto.Numbers)
	}
}
