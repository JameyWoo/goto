package main

import (
	"fmt"
	"reflect"
)

type Man struct {
	age int
}

func (m Man) test() {
	fmt.Println("test")
}

func main() {
	var it interface{}
	it = &Man{}
	if it == nil {
		fmt.Println("nil")
	}
	t := reflect.TypeOf(it)
	v := reflect.ValueOf(it)
	fmt.Println(v)
	fmt.Println(t.String())
	fmt.Println(t.Name())
	fmt.Println(t.Kind())
	fmt.Println(t.NumMethod())
	fmt.Println(t.PkgPath())

	fmt.Println(reflect.TypeOf(Man{}).NumMethod())

	v2 := v.Elem()
	//v2.SetBool(true)
	fmt.Println(v2.CanSet())
	//fmt.Println(v2)

}
