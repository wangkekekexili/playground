package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	fmt.Println(debug.SetGCPercent(10))
}
