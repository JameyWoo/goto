package priority_queue

import (
	"errors"
	"log"
)

type Array interface {
	Push(interface{})
	Pop() interface{}
	Top() interface{}
	Swap(i, j int)
	Len() int
	Less(i, j int) bool
}

type PriorityQueue struct {
	Elems Array
}

func (pq *PriorityQueue) Push(x interface{}) {
	pq.Elems.Push(x)
	pq.swim(pq.Elems.Len() - 1)
}

func (pq *PriorityQueue) Pop() (interface{}, error) {
	if pq.Elems.Len() < 1 {
		return nil, errors.New("no element in the Heap")
	}
	pq.Elems.Swap(0, pq.Elems.Len() - 1)
	pq.sink(0, pq.Elems.Len() - 1)
	return pq.Elems.Pop(), nil
}

func (pq *PriorityQueue) Top() interface{} {
	if pq.Elems.Len() < 1 {
		log.Fatal("Error! No element in the Heap!")
	}
	return pq.Elems.Top()
}

func (pq *PriorityQueue) Empty() bool {
	return pq.Elems.Len() == 0
}

func (pq *PriorityQueue) swim(idx int) {
	for idx > 0 && pq.Elems.Less(idx, (idx-1)/2) {
		pq.Elems.Swap(idx, (idx - 1) / 2)
		idx = (idx - 1) / 2
	}
}

func (pq *PriorityQueue) sink(idx, n int) {
	for idx*2+1 < n || (idx+1)*2 < n {
		tmpIdx := idx
		if pq.Elems.Less(idx*2+1, tmpIdx) {
			tmpIdx = idx*2 + 1
		}
		if (idx+1)*2 < n && pq.Elems.Less((idx+1)*2, tmpIdx) {
			tmpIdx = (idx + 1) * 2
		}
		pq.Elems.Swap(idx, tmpIdx)
		if idx == tmpIdx {
			break
		}
		idx = tmpIdx
	}
}
