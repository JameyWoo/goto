package _map

import (
	"fmt"
	"testing"
)

func TestMapLen(t *testing.T) {
	m := map[int]int {
		1: 2,
		2: 3,
	}
	fmt.Println(m)
	fmt.Println(len(m))

	n := make(map[int]int, 10)
	fmt.Println(len(n))
	n[1] = 3
	v, ok := n[2]
	if ok {
		fmt.Println(v)
	}
	fmt.Println(n)
}

func TestDelete(t *testing.T) {
	m := map[int]int {
		1 : 2,
		2 : 4,
	}
	fmt.Println(m)
	delete(m, 3)
	fmt.Print(m)
}