package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	for {
		select {
		case <-ch:
			fmt.Println("out")
		case ch <- 1:
			fmt.Print("in")
		default:
			fmt.Print("default\n")
			time.Sleep(time.Second)
		}
	}
}
