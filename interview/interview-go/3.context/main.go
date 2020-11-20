package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
)

/*
在 golang 协程和channel配合使用
写代码实现两个 goroutine，其中一个产生随机数并写入到 go channel 中，另外一个从 channel 中读取数字并打印到标准输出。最终输出五个随机数。
*/

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	randnum := make(chan int)

	wg := &sync.WaitGroup{}

	wg.Add(2)

	go f1(ctx, randnum, wg)

	go f2(cancel, randnum, wg)

	wg.Wait()
}

func f1(ctx context.Context, randnum chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	count := 0
	for {
		select {
		case num := <- randnum:
			fmt.Println(num)
			count += 1
		case <- ctx.Done():
			return
		}
	}
}

func f2(cancel func(), randnum chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		randnum <- rand.Int()
	}
	cancel()
}