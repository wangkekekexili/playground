package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type str string

func main() {
	s1 := "hello world"
	h1 := (*reflect.StringHeader)(unsafe.Pointer(&s1))
	fmt.Printf("%X %X %d\n", &s1, h1.Data, h1.Len)

	s2 := "hello world"
	h2 := (*reflect.StringHeader)(unsafe.Pointer(&s2))
	fmt.Printf("%X %X %d\n", &s2, h2.Data, h2.Len)

	s3 := s1
	h3 := (*reflect.StringHeader)(unsafe.Pointer(&s3))
	fmt.Printf("%X %X %d\n", &s3, h3.Data, h3.Len)

	s4 := str("hello world")
	h4 := (*reflect.StringHeader)(unsafe.Pointer(&s4))
	fmt.Printf("%X %X %d\n", &s4, h4.Data, h4.Len)
}
