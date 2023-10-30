package calculator

import (
	"strconv"
	"strings"
)

func Add(origin string) int64 {
	f := func(r rune) bool {
		return r == ',' || r == ':'
	}
	splits := strings.FieldsFunc(origin, f)
	res := int64(0)
	for _, str := range splits {
		adder, _ := strconv.ParseInt(str, 10, 0)
		res = res + adder
	}
	return res
}
