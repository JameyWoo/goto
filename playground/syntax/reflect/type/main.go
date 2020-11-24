package main

import (
	"fmt"
	"reflect"
)

type People interface {
	Hello()
}

type Student struct {
	Name string
	Age int
}

func (s Student) Hello() {
	fmt.Println("hello")
}

func main() {
	var p People = Student{}
	v := reflect.ValueOf(p)
	fmt.Println(v.Type())
	fmt.Println(v.NumField())

	f0 := v.Type().Field(0)
	fmt.Println(f0.Name)

	fmt.Println(reflect.TypeOf(v.Field(0)))
	fmt.Println(reflect.TypeOf(v.Type().Field(0)))

	//fmt.Println(reflect.TypeOf(reflect.ValueOf(1).Field(0)))
}
