package heap

import (
	"errors"
	"log"
)

type IntHeap struct {
	array []int
}

func (ip *IntHeap) Push(x int) {
	ip.array = append(ip.array, x)
	ip.swim(len(ip.array) - 1)
}

func (ip *IntHeap) Pop() (int, error) {
	if len(ip.array) < 1 {
		//log.Fatal("Error! No element in the Heap!")
		return 0, errors.New("no element in the Heap")
	}
	top := ip.Top()
	ip.array[0] = ip.array[len(ip.array) - 1]
	ip.array = ip.array[:len(ip.array) - 1]
	ip.sink(0)
	return top, nil
}

func (ip *IntHeap) Top() int {
	if len(ip.array) < 1 {
		log.Fatal("Error! No element in the Heap!")
	}
	return ip.array[0]
}

func (ip *IntHeap) Empty() bool {
	return len(ip.array) == 0
}

func (ip *IntHeap) swim(idx int) {
	for idx > 0 && ip.array[idx] < ip.array[(idx - 1) / 2] {
		tmp := ip.array[idx]
		ip.array[idx] = ip.array[(idx - 1) / 2]
		ip.array[(idx - 1) / 2] = tmp
		idx = (idx - 1) / 2
	}
}

func (ip *IntHeap) sink(idx int) {
	for idx * 2 + 1 < len(ip.array) || (idx + 1) * 2 < len(ip.array) {
		tmpIdx := idx
		if ip.array[tmpIdx] > ip.array[idx*2 + 1] {
			tmpIdx = idx*2 + 1
		}
		if (idx + 1) * 2 < len(ip.array) && ip.array[tmpIdx] > ip.array[(idx + 1) * 2] {
			tmpIdx = (idx + 1) * 2
		}
		tmp := ip.array[idx]
		ip.array[idx] = ip.array[tmpIdx]
		ip.array[tmpIdx] = tmp
		if idx == tmpIdx {
			break
		}
		idx = tmpIdx
	}
}