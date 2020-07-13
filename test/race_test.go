package test

import (
	"fmt"
	"sync"
	"testing"
)

var cnt int
var wg sync.WaitGroup
var lock sync.Mutex

func TestRace(t *testing.T) {
	cnt = 0
	wg = sync.WaitGroup{}
	lock = sync.Mutex{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go inc()
	}
	wg.Wait()
	fmt.Printf("cnt = %d\n", cnt)
}

func inc() {
	for i := 0; i < 10000; i++ {
		lock.Lock()
		cnt += 1
		lock.Unlock()
	}
	wg.Done()
}