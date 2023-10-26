package learning

import (
	"fmt"
	"math/rand"
)

func Example_Rand() {
	oddEven := rand.Int() % 2
	fmt.Println(oddEven)
	//output: 0
}
