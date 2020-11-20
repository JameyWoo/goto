package main

import "fmt"

func returnButDefer() (it int) {  //t初始化0， 并且作用域为该函数全域

	defer func() {
		it = it + 71
	}()

	return 891
}

func main() {
	a := 321
	a = returnButDefer()
	fmt.Println(a)
}