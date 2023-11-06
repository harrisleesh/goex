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

func MoreCard(name string) bool {
	fmt.Println(name + "는 한장의 카드를 더 받겠습니까?(예는 y, 아니오는 n)")
	var res string
	_, err := fmt.Scanf("%s", &res)
	if err != nil {
		return false
	}
	if res == "n" {
		return false
	}
	return true
}
