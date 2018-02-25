package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.Caller(0))

	pc := make([]uintptr, 100)
	n := runtime.Callers(0, pc)
	pc = pc[:n]
	fmt.Printf("PC: %v\n\n", pc)

	frames := runtime.CallersFrames(pc)
	for {
		frame, ok := frames.Next()
		if !ok {
			break
		}

		fmt.Printf("PC %v; Function %v\n", frame.PC, frame.Function)
	}
}
