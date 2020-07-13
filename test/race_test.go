package test

import (
	"fmt"
	"my/golang-code-master/chapter7/patterns/runner"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

var cnt int32
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

	runner.New(5 * time.Second)
}

func inc() {
	for i := 0; i < 10000; i++ {
		atomic.AddInt32(&cnt, 1)  // 这样效率比加锁快大概5倍, 比没有慢6倍
		//lock.Lock()
		//cnt += 1
		//lock.Unlock()
	}
	wg.Done()
}

func TestMap(t *testing.T) {
	m := sync.Map{}
	fmt.Println(m)
}