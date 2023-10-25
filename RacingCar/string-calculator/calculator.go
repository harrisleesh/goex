package string_calculator

func Add(a, b int64) int64 {
	return a + b
}

func Minus(a, b int64) int64 {
	return a - b
}

func Multiply(a, b int64) int64 {
	return a * b
}

func Divide(a, b int64) int64 {
	return a / b
}

func Calculate(origin string) int64 {
	parsed := Parse(origin)
	if parsed[1] == "+" {
		return Add(ParseInt(parsed[0]), ParseInt(parsed[2]))
	} else if parsed[1] == "-" {
		return Minus(ParseInt(parsed[0]), ParseInt(parsed[2]))
	} else if parsed[1] == "*" {
		return Multiply(ParseInt(parsed[0]), ParseInt(parsed[2]))
	} else if parsed[1] == "/" {
		return Divide(ParseInt(parsed[0]), ParseInt(parsed[2]))
	} else {
		return 0
	}
}
