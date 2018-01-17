package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.NumCPU())
	cpu := runtime.GOMAXPROCS(0)
	fmt.Println(cpu)
}
