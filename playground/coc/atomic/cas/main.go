package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var i rune = 2
	atomic.CompareAndSwapInt32(&i, 1, 3)
	fmt.Println(i)
}