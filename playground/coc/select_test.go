package coc

import (
	"fmt"
	"log"
	"testing"
	"time"
)


func TestSelect(t *testing.T) {
	ti := time.NewTicker(time.Second * 2)

	count := 0

	for {
		select {
		case <-ti.C:
			log.Printf("count: %d", count)

		default:
			count += 1
		}
	}

}

func TestCountOneSecond(t *testing.T) {
	out := false
	count := 0
	go func() {
		<-time.Tick(time.Second * 1)
		out = true
	}()

	for {
		count += 1
		if out == true {
			break
		}
	}
	log.Printf("count: %d\n", count)
}

// 测试select与定时器
func TestSelectTime(t *testing.T) {
	ch := make(chan int)
	tim := time.NewTimer(time.Second * 3)

	for {
		select {
		case i := <-ch:
			fmt.Println("i = " + string(i))
		case _ = <- tim.C:
			tim.Reset(time.Second * 2)
			fmt.Println("time out")
		}
	}
}

// 测试 int 的最大值
func TestIntMax(t *testing.T) {
	var a int
	a = 1
	for {
		a *= 10
		fmt.Printf("a = %d\n", a)
	}
}