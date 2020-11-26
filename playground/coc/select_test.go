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

// main 永久阻塞测试
func TestMainBlock(t *testing.T) {
	go func() {
		for {
			fmt.Println("tick")
			time.Sleep(1 * time.Second)
		}
	}()
	select {}
}

// 测试如果有chan可执行以及default, 是否一定执行chan, 或者还是随机
func TestChanDefault(t *testing.T) {
	for {
		ch := make(chan int, 1)
		select {
		case ch <- 1:
			fmt.Println(1)
		default:
			fmt.Println(2)
		}
		time.Sleep(1 * time.Second)
	}
}

// 向 close 的chan中写数据的case 测试
func TestCloseChanWrite(t *testing.T) {
	ch := make(chan int)
	close(ch)
	select {
	case ch <- 1:
		fmt.Println("write")
	default:
		fmt.Println("Default")
	}
}