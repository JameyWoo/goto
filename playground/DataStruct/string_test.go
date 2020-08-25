package DataStruct

import (
	"crypto/sha256"
	"fmt"
	"path/filepath"
	"testing"
	"unsafe"
)

func TestString(t *testing.T) {
	var str string
	str = "hello"
	fmt.Println(str)

	var a int = 12
	fmt.Println(a)

	var b = 3
	println(b)

	println("hello, World")

	c := &b
	fmt.Println(c)
	*c = 111
	fmt.Println(b)

	d := b
	b = 45
	fmt.Println(d)
}

func localPoint() *int {
	var a = 1
	fmt.Println("a =", &a)
	return &a
}

func TestLocalPoint(t *testing.T) {
	var b = localPoint()
	fmt.Println("b =", b)
}

func TestStringModify(t *testing.T) {
	var str = "fuck"
	str2 := str[:2] + str[2:]  // 通过拼接
	//str[2] = 'g'  // 不能修改
	fmt.Println(&str)
	fmt.Println(&str2)
	fmt.Println(len(str))
}

func TestString2(t *testing.T) {
	var str = "abcde"
	var str2 = str[:]
	fmt.Println(&str, &str2)

	str3 := "abc"
	str4 := str3

	fmt.Println(&str3, &str4)
	fmt.Println(str4)
}

func TestPath(t *testing.T) {
	filename := "/var/vals/test.txt"
	file := filepath.Base(filename)
	fmt.Println(file)

	dir, file := filepath.Split(filename)
	fmt.Println(dir)
	fmt.Println(file)
}

func TestArray(t *testing.T) {
	var a [3]int = [3]int{1, 2, 3}
	a = [3]int{4, 5}
	a = [...]int{6, 7, 8}  // 合法, 3个
	//a = [...]int{6}  // 违法, 长度不符合
	fmt.Println(a)
}

func TestArray2(t *testing.T) {
	var a [3][4]int
	fmt.Println(len(a))
	fmt.Println(len(a[0]))
	fmt.Println(a)

	const b = 5
	//var b = 5 // 不行, 数组的大小参数需要是常数
	var c [b]int
	fmt.Println(c, b)
}

func TestArray3(t *testing.T) {
	a := [...]int{1: 6, 8: 9}
	fmt.Println(a)

	c1 := sha256.Sum256([]byte("fuck"))
	fmt.Println(c1)
	fmt.Printf("%x\n", c1)
	fmt.Println(len(c1))
}

func TestArray4(t *testing.T) {
	a := [...]int{1, 2, 3}
	modify(a)
	fmt.Println(a)
}

func modify(a [3]int)  {
	// 传值
	a[2] = 4
}

func TestStringSize(t *testing.T) {
	var s string
	fmt.Println(unsafe.Sizeof(s))

	s2 := "hello, world"
	fmt.Println(unsafe.Sizeof(s2))

	s3 := "1111111111111111111111111111111111111111111111111111111111111"
	fmt.Println(unsafe.Sizeof(s3))
}