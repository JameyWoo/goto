package _interface

import (
	"fmt"
	"testing"
)

type Map map[string]int

// ? 为什么不能是 (m *Map)
func (m Map) Iter() error {
	for key, value := range m {
		fmt.Printf("key: %s, value: %d\n", key, value)
	}
	return nil
}

type Iter interface {
	Iter () error
}

func TestInter(t *testing.T) {
	var I Iter

	m := Map{}
	m["hello"] = 1
	m["world"] = 2
	m["aaa"] = 3

	I = m
	_ = I.Iter()
}