package main

import (
	"fmt"
	"strings"
)

func main() {
	var b strings.Builder
	fmt.Fprint(&b, "hello")
	fmt.Fprint(&b, " ")
	fmt.Fprintln(&b, "world")
	fmt.Println(b.String())
}
