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
	result := ParseInt(parsed[0])
	for i := 1; i < len(parsed); i += 2 {
		result = operateBinary(result, parsed[i], ParseInt(parsed[i+1]))
	}
	return result
}

func operateBinary(res int64, operator string, temp int64) int64 {
	if operator == "+" {
		return Add(res, temp)
	} else if operator == "-" {
		return Minus(res, temp)
	} else if operator == "*" {
		return Multiply(res, temp)
	} else if operator == "/" {
		return Divide(res, temp)
	} else {
		return 0
	}
}
