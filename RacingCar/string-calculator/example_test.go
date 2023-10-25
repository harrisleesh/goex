package string_calculator

import (
	"fmt"
)

func Example_OneNumber() {
	fmt.Println("1")

	//Output: 1
}

func ExampleAddTwoNumbers() {
	fmt.Println(Add(1, 2))

	//Output: 3
}

func ExampleParseString() {
	fmt.Println(Parse("1 + 2"))

	//Output: [1 + 2]
}

func ExampleAddString() {
	parses := Parse("1 + 2")
	fmt.Println(Add(ParseInt(parses[0]), ParseInt(parses[2])))

	//Output: 3
}

func ExampleMinusString() {
	parses := Parse("1 - 2")
	fmt.Println(Minus(ParseInt(parses[0]), ParseInt(parses[2])))
	//Output: -1
}
