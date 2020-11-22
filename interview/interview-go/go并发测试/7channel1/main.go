package main
import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var ch chan int
	//ch = make(chan int, 1)
	go func() {
		ch = make(chan int, 1)
		ch <- 1
	}()
	// 注意下面这里是传了参数的! 值传递
	go func(ch chan int) {
		time.Sleep(1 * time.Second)
		<-ch
	}(ch)
	c := time.Tick(1 * time.Second)
	for range c {
		fmt.Printf("#goroutines: %d\n", runtime.NumGoroutine())
	}
}