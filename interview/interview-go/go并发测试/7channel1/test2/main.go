package main

import "fmt"

func main() {
	var ch chan int
	go func() {
		<- ch
	}()
	ch <- 1

	fmt.Println("hello")
}
