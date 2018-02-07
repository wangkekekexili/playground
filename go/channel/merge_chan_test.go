package main

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	c := merge(asChan(1, 2, 3), asChan(4, 5, 6), asChan(7, 8, 9))
	seen := make(map[int]bool)
	for v := range c {
		if seen[v] {
			t.Fatalf("%v is seen at least twice", v)
		}
		seen[v] = true
	}
	for i := 1; i != 9; i++ {
		if !seen[i] {
			t.Errorf("%v is missing", i)
		}
	}
}

func BenchmarkMerge(b *testing.B) {
	funcs := []struct {
		name string
		f    func(...<-chan int) <-chan int
	}{
		{"goroutines", merge},
		{"reflection", mergeReflect},
	}
	for _, f := range funcs {
		for n := 1; n <= 1024; n *= 2 {
			chans := make([]<-chan int, n)
			b.Run(fmt.Sprintf("%v-%v", f.name, n), func(b *testing.B) {
				for i := 0; i != b.N; i++ {
					for j := range chans {
						chans[j] = asChan(1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
					}
					c := f.f(chans...)
					for range c {
					}
				}
			})
		}
	}
}
