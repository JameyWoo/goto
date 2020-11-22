package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	a :="aaa"
	ssh := *(*reflect.StringHeader)(unsafe.Pointer(&a))
	b := *(*[]byte)(unsafe.Pointer(&ssh))
	//b[2] = 12
	//b[0] = 'b'
	fmt.Printf("%v\n",b)
	fmt.Println(a)


	fmt.Println(reflect.TypeOf(a))

	fmt.Println(reflect.TypeOf(ssh))

	fmt.Println(reflect.TypeOf(b))

	c := []byte{1, 2, 3}
	c[1] = 4
	fmt.Println(c)
	fmt.Println(reflect.TypeOf(c))

	d := b[:]
	fmt.Println(d)
	d[0] = 11
	fmt.Println(d)
}