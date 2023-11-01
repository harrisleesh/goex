package view

import (
	"fmt"
	"strings"
)

func GetParticipateNames() []string {
	fmt.Println("게임에 참여할 사람의 이름을 입력하세요.(쉼표 기준으로 분리)")
	var names string
	_, err := fmt.Scanln(&names)
	if err != nil {
		return nil
	}
	return strings.Split(names, ",")
}
