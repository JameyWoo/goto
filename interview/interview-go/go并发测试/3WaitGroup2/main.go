package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func f1() {
	time.Sleep(1 * time.Second)
	wg.Done()
}

func main() {
	var i int
	for i=0;i<3;i++ {
		wg.Add(1)
		go f1()
	}
	wg.Wait()
	fmt.Println("end...")
}