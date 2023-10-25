package string_calculator

import (
	"strconv"
	"strings"
)

func Parse(origin string) []string {
	return strings.Split(origin, " ")
}

func ParseInt(origin string) int64 {
	i, _ := strconv.ParseInt(origin, 10, 0)
	return i
}
