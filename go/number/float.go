package main

import (
	"fmt"
)

func main() {
	nums := []float32{0}
	for _, n := range nums {
		fmt.Printf("%v: %b\n", n, n)
	}
}
