package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age int
}

func (s *Student) String() string {
	return fmt.Sprintln("hello, world")
}

func main() {
	var i interface{} = new(Student)
	s, ok := i.(*Student)
	if ok {
		fmt.Println(s)
	}
	a := i.(*Student)
	fmt.Println(a)
	switch v := i.(type) {
	case nil:
		fmt.Println(v)
	case Student:
		fmt.Println(v)
	case *Student:
		fmt.Println(v)
		fmt.Println(reflect.TypeOf(v))
	}
}
