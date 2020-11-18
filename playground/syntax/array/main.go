package main

import (
	"fmt"
)

func main() {
	var a [3]int
	a[0] = 1
	fmt.Println(a)

	b := a
	b[2] = 1
	fmt.Println(b)
	fmt.Println(a)
}