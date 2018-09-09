package main

import (
	"fmt"
	"strconv"
	"syscall/js"
)

func add(vs []js.Value) {
	id1 := vs[0].String()
	id2 := vs[1].String()
	v1, _ := strconv.Atoi(js.Global().Get("document").Call("getElementById", id1).Get("value").String())
	v2, _ := strconv.Atoi(js.Global().Get("document").Call("getElementById", id2).Get("value").String())

	js.Global().Get("document").Call("getElementById", "result").Set("value", js.ValueOf(v1+v2))
}

func registerCallbacks() {
	js.Global().Set("add", js.NewCallback(add))
}

func main() {
	fmt.Println("hello world")
	ch := make(chan bool)
	registerCallbacks()
	<-ch
}
