package syntax

import (
	"fmt"
	"testing"
	"time"
)

func TestSelectTime(t *testing.T) {
	//ch := make(chan int)
	ti := time.NewTicker(2 * time.Second)
	for {
		select {
		case <- ti.C:
			fmt.Println("ch")
		//case <- ti.C:
		//	fmt.Println("time!")
		//default:
		//	fmt.Println("default")
		//	time.Sleep(1 * time.Second)
		}
		fmt.Println("out")
	}
}


// select既输入也输出
func TestInOutChan(t *testing.T) {
	ch := make(chan int, 1)

	for {
		select {
		case ch <- 1:
			fmt.Println("in")


		case <- ch:
			fmt.Println("out")
		}
	}
}

// 测试空select, 会死锁
func TestEmptySelect(t *testing.T) {
	select {

	}
	// OutPut: h
}