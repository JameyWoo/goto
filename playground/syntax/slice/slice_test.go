package slice

import (
	"fmt"
	"testing"
)

// 测试 随着append, slice容量的变化
func TestSliceInitA(t *testing.T) {
	var a []int
	fmt.Println(len(a), cap(a))

	a = append(a, 1)
	fmt.Println(len(a), cap(a))

	a = append(a, 1)
	fmt.Println(len(a), cap(a))

	a = append(a, 1)
	fmt.Println(len(a), cap(a))

	a = append(a, 1, 2)
	fmt.Println(len(a), cap(a))
}

// 测试 初始化slice得到的cap
// ! 初始化时cap = len, 并不会自动扩展或怎样
func TestSliceInitB(t *testing.T) {
	b := []int{}
	fmt.Println(len(b), cap(b))

	c := []int{1}
	fmt.Println(len(c), cap(c))

	d := []int{1, 2}
	fmt.Println(len(d), cap(d))

	e := []int{1, 2, 3}
	fmt.Println(len(e), cap(e))
}

// 测试 slice 作为函数参数传递, 修改会不会改变原slice的值
// 修改会改变!
func TestSliceFunc(t *testing.T) {
	//s := []int{1, 2, 3}
	s := make([]int, 3, 4)
	fmt.Println(s)
	fmt.Println(len(s), cap(s))

	modify(s)

	fmt.Println(s)
	fmt.Println(len(s), cap(s))

	appendAndModify(s)

	fmt.Println(s)
	fmt.Println(len(s), cap(s))
}

func modify(s []int) {
	s[0] = 10
}

func appendAndModify(s []int) {
	s = append(s, 1)
	s[0] = 21
}

// 测试slice的扩展表达式
// slice[low: high: max]
// max代表对这个slice最大容量的限制, 是用来避免对原数组的后续空间进行修改的
// 当slice进行append且超过了max的大小时, 就会申请新的空间
// 其实原理就是  slice.cap = min{max, slice.old_cap}
func TestSliceExtend(t *testing.T) {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	a := s[0:5:6]
	b := s[0:5]
	fmt.Println(a)
	fmt.Println(len(a), cap(a))

	fmt.Println(b)
	fmt.Println(len(b), cap(b))
}

// append slice
func TestSliceAppend(t *testing.T) {
	a := []int{1, 2, 3}
	a = append(a, []int{4, 5, 6}...)
	fmt.Println(a)
}

// 扩容测试
func TestAppendT(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	a = append(a, 6)
	fmt.Println(len(a), cap(a))

	a = append(a, 7, 8, 9, 10, 11)
	fmt.Println(len(a), cap(a))
}