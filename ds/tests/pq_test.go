package tests

import (
	"container/heap"
	"fmt"
	"github.com/JameyWoo/goto/ds/priority_queue"
	"testing"
)

func TestMain0(m *testing.T) {
	pq := &priority_queue.PriorityQueue{}
	pq.Push(4)
	pq.Push(6)
	pq.Push(1)
	for pq.Len() > 0 {
		fmt.Printf("%d\n", heap.Pop(pq))
	}
}
