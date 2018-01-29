package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Printf("size of int8: %v\n", unsafe.Sizeof(int8(0)))

	nums := []int8{0, 1, 10, 20, -1, -2, -4, -128}
	for _, n := range nums {
		fmt.Printf("%v: %08b\n", n, uint8(n))
	}
}
