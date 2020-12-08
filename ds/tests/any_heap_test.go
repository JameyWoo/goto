package tests

import (
	"fmt"
	"github.com/JameyWoo/goto/ds/any_heap"
	"testing"
)

type Int struct {
	x int
}

func (e *Int) Less(elem interface{}) bool {
	switch elem.(type) {
	case Int:
		fmt.Println("Int")
	case *Int:
		fmt.Println("*Int")
	case any_heap.Elem, *any_heap.Elem:
		fmt.Println("Elem")
	default:
		fmt.Println("NO")
	}
	return false
}

func TestAnyHeap(t *testing.T) {
	ah := any_heap.IntHeap{}
	ah.Push(&Int{x: 3})
	ah.Push(&Int{x: 1})
	ah.Push(&Int{x: 2})
}
