package tests

import (
	"fmt"
	"github.com/JameyWoo/goto/ds/priority_queue"
	"testing"
)

type IntNode int

func (node IntNode) Less(i, j int) bool {
	return false
}

func TestPQFrame(t *testing.T) {
	pq := priority_queue.PriorityQueue{}
	pq.Push(IntNode(2))
	pq.Push(IntNode(3))
	pq.Push(IntNode(1))

	x := pq.Top()
	fmt.Println(x)
}
