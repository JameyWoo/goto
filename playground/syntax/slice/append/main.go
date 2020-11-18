package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4}
	//b := a[1:4]
	//a[1] = 2
	var b []int = []int{1, 2}
	
	c := copy(b, a[2:4])
	fmt.Println(b)
	fmt.Println(c)
}