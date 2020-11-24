package main

import (
	"fmt"
)

type People interface {
	Hello()
}

type Student struct {}

func (s Student) Hello() {
	fmt.Println("hello")
}

func main() {
	var p People = Student{}
	switchTypeTest(p)
}

func switchTypeTest(p interface{}) {
	switch x := p.(type) {
	case Student:
		fmt.Println("Student")
	case People:
		fmt.Println("People")
	default:
		fmt.Println(x)
	}
}
