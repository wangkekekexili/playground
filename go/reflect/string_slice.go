package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	str := "hello"
	strHeader := (*reflect.StringHeader)(unsafe.Pointer(&str))
	fmt.Printf("%p\n", strHeader)
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(strHeader))
	fmt.Printf("%p\n", sliceHeader)
	fmt.Printf("%X %X\n", strHeader.Data, sliceHeader.Data)
	fmt.Printf("%d %d\n", sliceHeader.Len, sliceHeader.Cap)

	bytes := []byte(str)
	fmt.Printf("%p\n", &bytes)
	bytesHeader := (*reflect.SliceHeader)(unsafe.Pointer(&bytes))
	fmt.Printf("%X\n", bytesHeader.Data)
}
