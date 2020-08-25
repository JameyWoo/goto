package coc

import (
	"fmt"
	"sync"
	"testing"
)

// 测试等待所有子携程执行完之后, 自己再执行
func TestGoWait(t *testing.T) {
	wg := sync.WaitGroup{}
	loc := sync.Mutex{}
	count := 0

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(idx int) {
			fmt.Printf("id: %d\n", idx)
			loc.Lock()
			count += 1
			loc.Unlock()
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Printf("count: %d\n", count)
}