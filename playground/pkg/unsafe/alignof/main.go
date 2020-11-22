package main

import "fmt"
import "unsafe"

func main() {
	var b bool
	var i8 int8
	var i16 int16
	var i64 int64

	var f32 float32

	var s string

	var m map[string]string

	var p *int32

	fmt.Println(unsafe.Alignof(b))   // 1
	fmt.Println(unsafe.Alignof(i8))  // 1
	fmt.Println(unsafe.Alignof(i16)) // 2
	fmt.Println(unsafe.Alignof(i64)) // 8
	fmt.Println(unsafe.Alignof(f32)) // 4
	fmt.Println(unsafe.Alignof(s))   // 8
	fmt.Println(unsafe.Alignof(m))   // 8
	fmt.Println(unsafe.Alignof(p))   // 8
}