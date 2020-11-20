package main

import (
	"fmt"
	"sync"
)

// https://github.com/lifei6671/interview-go/blob/master/question/q001.md

/*
交替打印数字和字母
问题描述

使用两个 goroutine 交替打印序列，一个 goroutine 打印数字， 另外一个 goroutine 打印字母， 最终效果如下：

12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728
*/

func main() {
	chars := make(chan int, 0)
	nums := make(chan int, 0)

	wg := &sync.WaitGroup{}
	wg.Add(2)
	go f1(chars, nums, wg)
	go f2(chars, nums, wg)
	nums <- 1
	wg.Wait()
}

func f1(chars, nums chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	i := 1
	for {
		select {
		case <- nums:
			fmt.Print(i)
			i++
			fmt.Print(i)
			i++
			if i == 29 {
				return
			}
			// 注意它的顺序很重要!!! 如果它在return前面, 那么会导致阻塞
			chars <- 1
		default:
			break
		}
	}
}

func f2(chars, nums chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	c := 'A'
	for {
		select {
		case <- chars:
			fmt.Printf("%c", c)
			c++
			fmt.Printf("%c", c)
			c++
			nums <- 1
			if c > 'Z' {
				return
			}
		default:
			break
		}
	}
}
