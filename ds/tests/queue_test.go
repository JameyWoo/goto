package tests

import (
	"fmt"
	"github.com/JameyWoo/goto/ds/queue"
	"testing"
	"time"
)

func TestQueue(t *testing.T) {
	Q := queue.Queue{}
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
