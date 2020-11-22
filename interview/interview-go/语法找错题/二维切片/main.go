package main

import "fmt"

func main() {
	s := [][]int{
		[]int{1, 2, 3},
		[]int{1, 3, 5},
	}
	fmt.Println(s)

	x := make([][]int, 10)
	fmt.Println(x)
}
