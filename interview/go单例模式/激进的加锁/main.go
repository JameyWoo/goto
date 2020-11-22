package main

import "sync"

var mu sync.Mutex

type singleton struct {}

var instance *singleton

func GetInstance() *singleton {
	mu.Lock()                    // 如果实例存在没有必要加锁
	defer mu.Unlock()

	if instance == nil {
		instance = &singleton{}
	}
	return instance
}

func main() {
	_ = GetInstance()
}
