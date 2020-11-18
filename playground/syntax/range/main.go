package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	si:=[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := range si {
		println(i)
		wg.Add(1)
		go func() {
			fmt.Println(i)
			wg.Done()
		}()
	}
	wg. Wait()
}