package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type strs []string

func main() {
	var s []string
	fmt.Printf("%p %v\n", &s, s == nil)

	var t strs
	fmt.Printf("%p %v\n", &t, s == nil)

	ps := &s
	*ps = append(*ps, "hello")
	fmt.Println(ps)

	pt := &t
	*pt = append(*pt, "world")
	fmt.Println(pt)

	h2 := (*reflect.SliceHeader)(unsafe.Pointer(ps))
	fmt.Printf("%X %d %d\n", h2.Data, h2.Len, h2.Cap)
	for _, d := range []string{"a", "b", "c", "d", "e", "f"} {
		s = append(s, d)
		fmt.Printf("%X %d %d\n", h2.Data, h2.Len, h2.Cap)
	}
}
