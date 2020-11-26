package main

import (
	"fmt"
	"testing"
	"time"
)

// 测试 stop 效果
func TestTickerStop(t *testing.T) {
	tk := time.NewTicker(1 * time.Second)
	<- tk.C
	fmt.Println("1")
	tk.Stop()
	<- tk.C  // stop之后不会close这个channel, 所以他就阻塞了, 导致error
	fmt.Println("2")
}

// 测试创建很多ticker, 因为似乎会挂载到系统协程上, 导致资源浪费
func TestCreateTicker(t *testing.T) {
	c := 0
	for {
		select {
		case <- time.Tick(1 * time.Millisecond):
			c++
		}
		if c > 1000 {
			break
		}
	}
	fmt.Print("over\n")
	time.Sleep(20 * time.Second)
}