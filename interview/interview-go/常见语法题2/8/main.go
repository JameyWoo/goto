package main

import (
	"fmt"
	"sync"
)

func main() {
	ua := &UserAges{ages: map[string]int{"age": 1}}
	ua.Add("name", 23)
	ua.Add("hello", 1)
	fmt.Println(ua.Get("shit"))
	fmt.Println(ua.Get("hello"))
}

type UserAges struct {
	ages map[string]int
	sync.Mutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}

func (ua *UserAges) Get(name string) int {
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}