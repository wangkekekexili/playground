package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	//debug.PrintStack()
	result := add(1, 2)
	fmt.Printf("returned value address: %p\n", &result)
}

func add(a, b int) (c int) {
	debug.PrintStack()
	fmt.Printf("before add is returned: %p\n", &c)
	return a + b
}
