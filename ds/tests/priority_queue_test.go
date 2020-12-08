package tests

import (
	"fmt"
	"github.com/JameyWoo/goto/ds/priority_queue"
	"testing"
)

/*
type Array interface {
	Push(interface{})
	Pop() interface{}
	Top() interface{}
	Swap(i, j int)
	Len() int
	Less(i, j int) bool
}
*/

type IntPQ []int

func (ih *IntPQ) Push(x interface{}) {
	*ih = append(*ih, x.(int))
}

func (ih *IntPQ) Pop() interface{} {
	top := (*ih)[len(*ih) - 1]
	*ih = (*ih)[:len(*ih) - 1]
	return top
}

func (ih *IntPQ) Top() interface{} {
	return (*ih)[0]
}

func (ih *IntPQ) Swap(i, j int) {
	tmp := (*ih)[i]
	(*ih)[i] = (*ih)[j]
	(*ih)[j] = tmp
}

func (ih *IntPQ) Len() int {
	return len(*ih)
}

func (ih *IntPQ) Less(i, j int) bool {
	return (*ih)[i] < (*ih)[j]
}

func TestPQFrame(t *testing.T) {
	pq := priority_queue.PriorityQueue{Elems: &IntPQ{}}
	nums := []int{3, 5, 1, 8, 4, 2, 9, 7, 3}
	for _, num := range nums {
		pq.Push(num)
	}
	for !pq.Empty() {
		val, err := pq.Pop()
		if err != nil {
			panic(err)
		}
		fmt.Printf("%d ", val)
	}
}
