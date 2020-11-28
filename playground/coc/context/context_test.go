package main

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 测试 子ctx先调用cancel, 然后 父ctx调用cancel
func TestCtxCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	ch := make(chan int)
	go work2(ctx, ch)
	time.Sleep(5 * time.Second)
	cancel()

	time.Sleep(3 * time.Second)
}

func work2(ctx context.Context, ch chan int) chan int {
	ctx2, can := context.WithCancel(ctx)
	go func(ctx context.Context) {
		for {
			select {
			case <- ctx.Done():
				fmt.Println("child cancel")
				return
			}
		}
	}(ctx2)
	can()
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