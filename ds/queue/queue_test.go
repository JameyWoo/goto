package queue

import (
	"fmt"
	"testing"
	"time"
)

func TestQueue(t *testing.T) {
	Q := queue{}
	Q.Push(5, 1, 2, 4)
	fmt.Println(Q.Size())

	for !Q.Empty() {
		time.Sleep(time.Second * 2)
		front := Q.Front()
		Q.Pop()
		fmt.Println(front)
	}
	//Q.Pop()
}
