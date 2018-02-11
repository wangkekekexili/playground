package main

import (
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	var a, b sync.Mutex

	go func() {
		defer wg.Done()
		a.Lock()
		time.Sleep(time.Second)
		b.Lock()
		b.Unlock()
		a.Unlock()
	}()

	go func() {
		defer wg.Done()
		b.Lock()
		time.Sleep(time.Second)
		a.Lock()
		a.Unlock()
		b.Unlock()
	}()

	wg.Wait()
}
