package main

import "fmt"

func main() {
	m := make(map[int]int)
	m[2] = 3
	key, val := m[2]
	fmt.Println(key)
	fmt.Println(val)
}
