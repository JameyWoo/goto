package pool

import (
	"fmt"
	"github.com/Jameywoo/goto/pool"
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

func TestPool2(t *testing.T) {
	p := pool.NewPool(10)
	go func() {
		for {
			p.Add(func() error {
				fmt.Println("running task")
				time.Sleep(500 * time.Microsecond)
				fmt.Println("task over")
				return nil
			})
		}
	}()
	p.Run()
	fmt.Println("over ...")
}
