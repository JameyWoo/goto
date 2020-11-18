package main

import "fmt"

func f1() (r int) {
	defer func() {
		r++
	}()
	return r
}

func main() {
	fmt.Println(f1())
}
