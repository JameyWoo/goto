package main

import (
	"fmt"
	"unsafe"
)

type Si struct {
	i int32
	b byte
	o int32
}

type Si2 struct {
	i int32
	o int32
	b byte
}

type Si3 struct {
	i int32
	o int32
	b byte
	c byte
	d byte
}

type Si4 struct {
	i int32
	o int64
	b byte
	c byte
	d byte
}

func main() {
	fmt.Println(unsafe.Sizeof(int8(1)))
	fmt.Println(unsafe.Sizeof(int(1)))
	fmt.Println(unsafe.Sizeof(Si{}))
	fmt.Println(unsafe.Sizeof(Si2{}))
	fmt.Println(unsafe.Sizeof(Si3{}))
	fmt.Println(unsafe.Sizeof(Si4{}))

	ch := make(chan int)
	fmt.Println(unsafe.Sizeof(ch))
	fmt.Println(unsafe.Sizeof(struct{ bool; int16; float64}{}))
	fmt.Println(unsafe.Sizeof(struct{ bool; float64; int16}{}))
	fmt.Println(unsafe.Sizeof(struct{ bool; float32; int16}{}))
}
