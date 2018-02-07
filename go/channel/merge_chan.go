package main

import (
	"fmt"
	"reflect"
	"sync"
)

func main() {
	a := asChan(1, 2, 3)
	b := asChan(4, 5, 6)
	for v := range mergeReflect(a, b) {
		fmt.Println(v)
	}
}

func asChan(vs ...int) <-chan int {
	ch := make(chan int)
	go func() {
		for _, v := range vs {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

func merge(chs ...<-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		var wg sync.WaitGroup
		wg.Add(len(chs))
		for _, ch := range chs {
			ch := ch
			go func() {
				for v := range ch {
					out <- v
				}
				wg.Done()
			}()
		}
		wg.Wait()
	}()
	return out
}

func mergeReflect(chs ...<-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		var cases []reflect.SelectCase
		for i := range chs {
			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(chs[i]),
			})
		}
		for len(cases) > 0 {
			chosen, recv, recvOK := reflect.Select(cases)
			if !recvOK {
				cases = append(cases[:chosen], cases[chosen+1:]...)
				continue
			}
			out <- recv.Interface().(int)
		}
	}()
	return out
}
