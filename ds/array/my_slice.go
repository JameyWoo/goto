package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

/*
使用 unsafe.Pointer 和 uintptr 操作数组的方法

但是不知道怎么动态生成数组.
 */

func main() {
	bools := [...]bool{true, false, false, false, true}
	p := uintptr(unsafe.Pointer(&bools)) + 3 * unsafe.Sizeof(bool(true))
	fmt.Println(reflect.TypeOf(p))
	*(*bool)(unsafe.Pointer(p)) = true
	fmt.Println(bools)

	n := 10
	n ++
}