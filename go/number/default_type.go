package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Default type for untyped int.
	const a = 1
	b := a
	fmt.Println(reflect.TypeOf(b))

	// Default type for untyped float.
	const c = 1.0
	d := c
	fmt.Println(reflect.TypeOf(d))
}
