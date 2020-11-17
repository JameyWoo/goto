package main

import (
	"fmt"
	"math/rand"
	"runtime"
)

func Generater (done chan struct{}) chan int {
	ch := make(chan int)
	go func() {
	Lable:
		for {
			select {
			case ch <- rand.Int():
			case <- done:
				break Lable
			}
		}
		close(ch)
	}()
	return ch
}

func main() {
	done := make(chan struct{})
	ch := Generater(done)

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	close(done)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	println("num of goruntine:",runtime.NumGoroutine())
}