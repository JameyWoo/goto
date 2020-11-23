package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	ch := make(chan int, 1000)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	go func() {
		for {
			a, ok := <-ch
			if !ok {
				fmt.Println("close")
				return
			}
			fmt.Println("a: ", a)
		}
	}()
	runtime.Gosched()
	close(ch)
	fmt.Println("ok")
	time.Sleep(time.Second * 3)
	tick := time.NewTicker(1 * time.Second)
	time.Tick(1 * time.Second)
	<- tick.C

	time.After(1 * time.Second)
}