package main

import (
	"fmt"
)

func main()  {

	变量 := 3
	fmt.Println(变量)

	defer func() {
		// recover 会捕获一个panic
		if err := recover(); err != nil{
			fmt.Println(err)
		}else {
			fmt.Println("fatal")
		}
	}()

	defer func() {
		// 这个defer里的panic覆盖了下面的panic
		panic("defer panic")
	}()

	panic("panic")
}