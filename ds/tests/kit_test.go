package tests

import (
	"container/heap"
	"fmt"
	"github.com/JameyWoo/goto/ds/kit"
	"testing"
)

func TestKit(t *testing.T) {
	x := &kit.PQ{}
	heap.Push(x, &kit.Entry{Key: "shit", Priority: 1})
	heap.Push(x, &kit.Entry{Key: "you", Priority: 2})
	heap.Push(x, &kit.Entry{Key: "fuck", Priority: 3})
	heap.Push(x, &kit.Entry{Key: "it", Priority: 0})
	fmt.Println(heap.Pop(x))
}
