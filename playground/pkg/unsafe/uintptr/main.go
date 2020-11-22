package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	a := 1
	fmt.Printf("%p\n", &a)

	p := &a
	fmt.Println(reflect.TypeOf(p))
	fmt.Println(reflect.TypeOf(unsafe.Pointer(p)))
	fmt.Println(reflect.TypeOf(uintptr(unsafe.Pointer(p))))

	up := unsafe.Pointer(p)
	fmt.Println(up)


}
