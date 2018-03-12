// https://www.youtube.com/watch?v=f6kdp27TYZs&t=312s

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	hello := fanIn2(sayHelloBy("ke"), sayHelloBy("yiyang"))
	for i := 0; i != 10; i++ {
		fmt.Println(<-hello)
	}
}

func fanIn(left, right <-chan string) <-chan string {
	ch := make(chan string)
	go func() {
		for s := range left {
			ch <- s
		}
	}()
	go func() {
		for s := range right {
			ch <- s
		}
	}()
	return ch
}

func fanIn2(left, right <-chan string) <-chan string {
	ch := make(chan string)
	go func() {
		for {
			select {
			case s := <-left:
				ch <- s
			case s := <-right:
				ch <- s
			}
		}
	}()
	return ch
}

func sayHelloBy(name string) <-chan string {
	ch := make(chan string)
	go func() {
		i := 0
		for {
			ch <- fmt.Sprintf("hello from %s %d", name, i)
			time.Sleep(100 * time.Millisecond * time.Duration(rand.Intn(10)))
			i++
		}
	}()
	return ch
}
