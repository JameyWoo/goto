package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func work(ctx context.Context, ch chan int) chan int {
	ctx2, _ := context.WithTimeout(ctx, 3 * time.Second)
	go func(ctx context.Context) {
		for {
			select {
			case <- ctx.Done():
				fmt.Println("timeout or cancel")
				return
			}
		}
	}(ctx2)
	for {
		select {
		case <- ctx.Done():
			close(ch)
			fmt.Println("parent cancel")
			return nil
		default:
			chVal := rand.Int()
			//ch <- chVal
			fmt.Println(chVal)
			time.Sleep(1 * time.Second)
		}
	}
	return nil
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	ch := make(chan int)
	go work(ctx, ch)
	time.Sleep(5 * time.Second)
	cancel()

	time.Sleep(3 * time.Second)
}