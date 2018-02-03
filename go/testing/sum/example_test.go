package sum_test

import (
	"fmt"

	"github.com/wangkekekexili/playground/go/testing/sum"
)

func ExampleInts() {
	s := sum.Ints(1, 2, 3, 4, 5)
	fmt.Println(s)
	// Output:
	// 15
}
