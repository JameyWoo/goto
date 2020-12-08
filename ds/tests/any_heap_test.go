package tests

import (
	"fmt"
	"github.com/JameyWoo/goto/ds/any_heap"
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

type IntHeap []int

func (ih *IntHeap) Push(x interface{}) {
	*ih = append(*ih, x.(int))
}

func (ih *IntHeap) Pop() interface{} {
	top := (*ih)[len(*ih) - 1]
	*ih = (*ih)[:len(*ih) - 1]
	return top
}

func (ih *IntHeap) Top() interface{} {
	return (*ih)[0]
}

func (ih *IntHeap) Swap(i, j int) {
	tmp := (*ih)[i]
	(*ih)[i] = (*ih)[j]
	(*ih)[j] = tmp
}

func (ih *IntHeap) Len() int {
	return len(*ih)
}

func (ih *IntHeap) Less(i, j int) bool {
	return (*ih)[i] < (*ih)[j]
}

func TestAnyHeap(t *testing.T) {
	ah := &any_heap.AnyHeap{Elems: &IntHeap{}}

	nums := []int{3, 5, 1, 8, 4, 2, 9, 7, 3}
	for _, num := range nums {
		ah.Push(num)
		fmt.Println(ah.Top())
	}
	for !ah.Empty() {
		val, err := ah.Pop()
		if err != nil {
			panic(err)
		}
		fmt.Printf("%d ", val)
	}
}
