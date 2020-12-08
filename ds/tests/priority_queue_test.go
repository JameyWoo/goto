package tests

import (
	"fmt"
	"github.com/JameyWoo/goto/ds/priority_queue"
	"testing"
)

type IntNode []int

func (node *IntNode) Less(i, j int) bool {
	return (*node)[i] < (*node)[j]
}

func TestPQFrame(t *testing.T) {
	pq := priority_queue.PriorityQueue{}
	pq.Push(2)
	pq.Push(3)
	pq.Push(1)

	x := pq.Top()
	fmt.Println(x)
	fmt.Println(pq.Pop())
	fmt.Println(pq.Top())
}
