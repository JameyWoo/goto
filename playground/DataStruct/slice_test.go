package DataStruct

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
	"unsafe"
)

func TestSlice(t *testing.T) {
	a := [...]int{1: 3, 6: 7}
	b := a[:] // 是slice了, 修改它会改变原数组, 如何得到一个新的slice呢?
	fmt.Println(b)
	Modify(b)
	fmt.Println(a)
}

func Modify(b []int) {
	b[2] = 2
}

func TestMain0(t *testing.T) {
	type Point struct {
		X, Y int
	}
	type Point2 struct {
		X, Y int
	}

	type Circle struct {
		Point
		R int
	}

	type Test2 struct {
		Point
		Point2
	}

	x := Circle{
		Point{1, 2},
		4,
	}
	fmt.Println(x)
	fmt.Printf("%#v\n", x)

	t2 := Test2{Point{1, 2}, Point2{3, 4}}
	t2.Point.X = 5
	fmt.Printf("%#v\n", t2)
}

// 初始化方式
func TestInit(t *testing.T) {
	a := [...]int{}  // 这个是数组
	fmt.Println(a)
	fmt.Println(reflect.TypeOf(a))

	var b []int // 这个是切片
	fmt.Println(b)
	fmt.Println(reflect.TypeOf(b))

	// 多维切片
	var c [][]int
	c = append(c, []int{1, 2, 3}, []int{4, 5, 6})
	fmt.Println(c)
	fmt.Println(reflect.TypeOf(c))

	d := [][][]int{4:{3:{2, 3}}}
	fmt.Println(d)
	fmt.Println(reflect.TypeOf(d))

	// 切片与底层数组
	e := [6]int{}
	f := e[:3]
	h := e[3:]
	fmt.Println(f)
	fmt.Printf("cas:%d, len:%d\n", cap(f), len(f))
	f = append(f, 4, 5, 6)
	// f = append(f, 4, 5, 6, 7)  // 注意这行代码和上一行代码引起的巨大差异
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(h)

	// 看默认多大
	g := []int{1, 2, 3}
	fmt.Printf("cas:%d, len:%d\n", cap(g), len(g))  // 默认跟初始化的数量一样大
	g = append(g, 4, 5)
	fmt.Printf("cas:%d, len:%d\n", cap(g), len(g))  // 扩容之后, cas * 2
}

func TestSize(t *testing.T) {
	var a int
	fmt.Println(unsafe.Sizeof(a))

	var b []int
	fmt.Println(unsafe.Sizeof(b))
}

func TestMap(t *testing.T) {
	var colors map[string]string
	fmt.Println(unsafe.Sizeof(colors))
	colors = make(map[string]string, 0)
	fmt.Println(unsafe.Sizeof(colors))
	colors["hello"] = "world"
	fmt.Println(unsafe.Sizeof(colors))

	colors["shit"] = "eat"
	colors["a"] = "b"
	colors["aaabaaa"] = "fuck"

	for key, value := range colors {
		fmt.Printf("key: %s, value: %s\n", key, value)
	}

	keys := make([]string, len(colors))
	i := 0
	for key, _ := range colors {
		keys[i] = key
		i += 1
	}
	sort.Strings(keys)
	fmt.Println(keys)

	for _, key := range keys {
		fmt.Printf("k: %s, v: %s\n", key, colors[key])
	}
}