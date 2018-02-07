package main

import (
	"fmt"
	"reflect"
	"sync"
)

func main() {
	a := asChan(1, 2, 3)
	b := asChan(4, 5, 6)
	for v := range mergeRecursive(a, b) {
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

func mergeRecursive(chs ...<-chan int) <-chan int {
	switch len(chs) {
	case 0:
		ch := make(chan int)
		close(ch)
		return ch
	case 1:
		return chs[0]
	case 2:
		return mergeTwo(chs[0], chs[1])
	default:
		m := len(chs) / 2
		return mergeTwo(mergeRecursive(chs[0:m]...), mergeRecursive(chs[m:]...))
	}
}

func mergeTwo(a, b <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for a != nil || b != nil {
			select {
			case v, ok := <-a:
				if ok {
					out <- v
				} else {
					a = nil
				}
			case v, ok := <-b:
				if ok {
					out <- v
				} else {
					b = nil
				}
			}
		}
	}()
	return out
}
