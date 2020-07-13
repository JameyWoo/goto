package pool

import (
	"fmt"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	task := func() error {
		fmt.Println("doing task")
		return nil
	}
	p := NewPool(8)

	go func() {
		for {
			p.Add(task)
		}
	}()
	p.Run()
	time.Sleep(3 * time.Second)
	fmt.Printf("job count: %d\n", p.jobCnt)
}
