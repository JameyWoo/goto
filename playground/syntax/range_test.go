package syntax

import (
	"fmt"
	"testing"
	"time"
)

// 测试range nil的channel
// 结果: 会阻塞, 不会panic
func TestRangeNilChan(t *testing.T) {
	var ch chan int
	go func(c *chan int) {
		for {
		time.Sleep(time.Second)
		//close(*c)
		}
	}(&ch)
	for c := range ch {
		fmt.Println(c)
	}
}

// 测试 range 对于close了的chan的反应
func TestCloseChanRange(t *testing.T) {
	ch := make(chan int)
	close(ch)
	for _ = range ch {
		fmt.Println("hello")
	}
	fmt.Println("over")
}


// 定时器 Timer 与 range
func TestRangeTimer(t *testing.T) {
	tr := time.NewTimer(2 * time.Second)
	// 输出一次后会阻塞, 因为 这个 tr.C 没有被close
	for c := range tr.C {
		fmt.Println(c)
		tr.Reset(1 * time.Second)
	}
}