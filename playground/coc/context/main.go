package main

import (
	"context"
	"fmt"
	"math/rand"
)

func work(ctx context.Context, ch chan int) chan int {
	for {
		select {
		case <- ctx.Done():
			close(ch)
			return nil
		default:
			ch <- rand.Int()
			//fmt.Println("运行一些东西...")
			//time.Sleep(2 * time.Second)
		}
	}
	return nil
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	ch := make(chan int)
	go work(ctx, ch)
	//time.Sleep(5 * time.Second)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	cancel()
	fmt.Println(<-ch)
}