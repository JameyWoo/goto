package kit

/*
go的莫名其妙的heap
不知道怎么实现. 无语...
不如自己写
 */

/*
这是一个失败的例子
 */

import (
	"container/heap"
)
//以下实现优先级队列
// This example demonstrates a Priority queue built using the heap interface.
// Entry 是 PriorityQueue 中的元素
type Entry struct {
	Key      string
	Priority int
	// Index 是 Entry 在 heap 中的索引号
	// Entry 加入 Priority Queue 后， Priority 会变化时，很有用
	// 如果 Entry.Priority 一直不变的话，可以删除 Index
	Index int
}

// PQ implements heap.Interface and holds entries.
type PQ []*Entry

func (pq PQ) Len() int { return len(pq) }

func (pq PQ) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

// Push 往 pq 中放 Entry
func (pq *PQ) Push(x interface{}) {
	temp := x.(*Entry)
	temp.Index = len(*pq)
	*pq = append(*pq, temp)
}

// Pop 从 pq 中取出最优先的 Entry
func (pq *PQ) Pop() interface{} {
	temp := (*pq)[len(*pq)-1]
	temp.Index = -1 // for safety
	*pq = (*pq)[0 : len(*pq)-1]
	return temp
}

// update modifies the Priority and value of an Entry in the queue.
func (pq *PQ) update(Entry *Entry, value string, Priority int) {
	Entry.Key = value
	Entry.Priority = Priority
	heap.Fix(pq, Entry.Index)
}