package learning

import "fmt"

type Test struct {
	testField string
}

func Example_testStruct() {
	t := Test{
		testField: "hello",
	}

	fmt.Println(t.testField)
	// Output: hello
}
