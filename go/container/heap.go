package main

import (
	"container/heap"
	"fmt"
)

type city struct {
	name     string
	distance int
	index    int
}

func (c *city) String() string {
	return c.name
}

type cityPriorityQueue []*city

func (q *cityPriorityQueue) Len() int {
	return len(*q)
}

func (q *cityPriorityQueue) Less(i, j int) bool {
	return (*q)[i].distance < (*q)[j].distance
}

func (q *cityPriorityQueue) Swap(i, j int) {
	(*q)[i], (*q)[j] = (*q)[j], (*q)[i]
	(*q)[i].index = i
	(*q)[j].index = j
}

func (q *cityPriorityQueue) Push(x interface{}) {
	city := x.(*city)
	city.index = len(*q)
	*q = append(*q, city)
}

func (q *cityPriorityQueue) Pop() interface{} {
	n := len(*q)
	city := (*q)[n-1]
	city.index = -1
	*q = (*q)[:n-1]
	return city
}

func main() {
	cityQueue := &cityPriorityQueue{
		{
			name:     "Shanghai",
			distance: 10,
			index:    0,
		},
		{
			name:     "Beijing",
			distance: 5,
			index:    1,
		},
		{
			name:     "Shenzhen",
			distance: 20,
			index:    2,
		},
	}
	fmt.Println(cityQueue)
	heap.Init(cityQueue)
	fmt.Println(cityQueue)

	city := heap.Pop(cityQueue)
	fmt.Println(city)
	fmt.Println(cityQueue)
}
