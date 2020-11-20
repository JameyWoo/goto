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
	ch := make(chan int, 0)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go f1(ch, wg)
	go f2(ch, wg)
	wg.Wait()
}

func f1(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 27; i += 1 {
		fmt.Print(i)
		//fmt.Print(i + 1)
		<- ch
	}
	fmt.Print(28)
}

func f2(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for c := 'A'; c <= 'Z'; c += 1 {
		ch <- 1
		fmt.Printf("%c", c)
		//fmt.Printf("%c", c + 1)
	}
	ch <- 1
}
