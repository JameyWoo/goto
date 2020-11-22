package main

import "sync"

var mu sync.Mutex

func main() {
	Map := make(map[int]int)

	group := &sync.WaitGroup{}

	for i := 0; i < 100000; i++ {
		group.Add(2)
		go writeMap(Map, i, i, group)
		go readMap(Map, i, group)
	}

	group.Wait()
}

func readMap(Map map[int]int, key int, group *sync.WaitGroup) int {
	mu.Lock()
	defer group.Done()
	defer mu.Unlock()
	return Map[key]
}

func writeMap(Map map[int]int, key int, value int, group *sync.WaitGroup) {
	defer group.Done()
	mu.Lock()
	defer mu.Unlock()
	Map[key] = value
}
