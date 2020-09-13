package priority_queue

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestMain0(m *testing.T) {
	pq := PriorityQueue{}
	pq.Push(4)
	pq.Push(6)
	pq.Push(1)
	for pq.Len() > 0 {
		fmt.Printf("%d\n", heap.Pop(&pq))
	}
}
