package main

import "sync"
import atomic "sync/atomic"

var initialized uint32
var mu sync.Mutex

type singleton struct {}

var instance *singleton

func GetInstance() *singleton {

	if atomic.LoadUint32(&initialized) == 1 {  // 原子操作
		return instance
	}

	// 这里有锁, 因此下方就不需要用 atomic 判断 initialized
	mu.Lock()
	defer mu.Unlock()

	if initialized == 0 {
		instance = &singleton{}
		atomic.StoreUint32(&initialized, 1)
	}

	return instance
}

func main() {
	_ = GetInstance()
}