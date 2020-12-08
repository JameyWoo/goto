package any_heap

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

type AnyHeap struct {
	Elems Array
}

func (ip *AnyHeap) Push(x interface{}) {
	ip.Elems.Push(x)
	ip.swim(ip.Elems.Len() - 1)
}

func (ip *AnyHeap) Pop() (interface{}, error) {
	if ip.Elems.Len() < 1 {
		return nil, errors.New("no element in the Heap")
	}
	ip.Elems.Swap(0, ip.Elems.Len() - 1)
	ip.sink(0, ip.Elems.Len() - 1)
	return ip.Elems.Pop(), nil
}

func (ip *AnyHeap) Top() interface{} {
	if ip.Elems.Len() < 1 {
		log.Fatal("Error! No element in the Heap!")
	}
	return ip.Elems.Top()
}

func (ip *AnyHeap) Empty() bool {
	return ip.Elems.Len() == 0
}

func (ip *AnyHeap) swim(idx int) {
	for idx > 0 && ip.Elems.Less(idx, (idx-1)/2) {
		ip.Elems.Swap(idx, (idx - 1) / 2)
		idx = (idx - 1) / 2
	}
}

func (ip *AnyHeap) sink(idx, n int) {
	for idx*2+1 < n || (idx+1)*2 < n {
		tmpIdx := idx
		if ip.Elems.Less(idx*2+1, tmpIdx) {
			tmpIdx = idx*2 + 1
		}
		if (idx+1)*2 < n && ip.Elems.Less((idx+1)*2, tmpIdx) {
			tmpIdx = (idx + 1) * 2
		}
		ip.Elems.Swap(idx, tmpIdx)
		if idx == tmpIdx {
			break
		}
		idx = tmpIdx
	}
}
