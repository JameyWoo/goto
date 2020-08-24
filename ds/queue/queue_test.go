package queue

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	Q := queue{}
	Q.Push(5, 1, 2, 4)
	fmt.Println(Q.Size())

	for !Q.Empty() {
		front := Q.Front()
		Q.Pop()
		fmt.Println(front)
	}
	//Q.Pop()
}
