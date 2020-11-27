package main

import "fmt"

func test() {
	s := make([]int, 1000, 10000)
	_ = len(s)
	return
}

// 测试 inline
// 有循环就不可以inline了
func testInline() {
	s := make([]int, 1, 1)
	for i := 0; i < len(s); i += 1 {
		s[i] = i
	}
}

func main() {
	f := Fib()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
	x := "shit"
	SelfInterface(x)
}

func Fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a + b
		return a
	}
}

func fmtTest() {
	s := "hello, world"
	fmt.Println(s)
}

// 测试自己的 interface{} 函数
func SelfInterface(i interface{}) {
	y := i
	_ = y
}