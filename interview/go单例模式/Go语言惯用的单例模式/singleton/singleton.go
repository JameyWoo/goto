package singleton

import (
	"fmt"
	"sync"
)

type singleton struct {}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

func (s *singleton) String() string {
	return fmt.Sprintf("I am the only instance: %p", s)
}

func main() {
	_ = GetInstance()
}