package main

import "sync"

var mu sync.Mutex

type singleton struct {}

var instance *singleton

func GetInstance() *singleton {
	if instance == nil {     // 不太完美 因为这里不是完全原子的
		mu.Lock()
		defer mu.Unlock()

		if instance == nil {
			instance = &singleton{}
		}
	}
	return instance
}

func main() {
	_ = GetInstance()
}