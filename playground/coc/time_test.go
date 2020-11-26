package coc

import (
	"fmt"
	"testing"
	"time"
)

func TestAfterFuncA(t *testing.T) {
	time.AfterFunc(3 * time.Second, func() {
		fmt.Println("hello, world")
	})
	time.Sleep(5 * time.Second)
}

func TestAfterFuncB(t *testing.T) {
	ch := make(chan int)
	time.AfterFunc(3 * time.Second, func() {
		fmt.Println("hello, world")
		close(ch)
	})
	for _ = range ch {}
}