package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	fmt.Println(l.Back() == nil)
	fmt.Println(l.Front() == nil)

	fmt.Println()
	for i := 1; i <= 128; i *= 2 {
		l.PushBack(i)
	}
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	fmt.Println()
	l.Remove(l.Front())
	l.Remove(l.Front())
	for e := l.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}

}
