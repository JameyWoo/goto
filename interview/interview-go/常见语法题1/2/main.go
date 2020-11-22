package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name string
}

func zhoujielun(v interface{}) {
	t := reflect.TypeOf(v)
	fmt.Println(t.String())
	switch msg := v.(type) {
	case *student:
		fmt.Println(msg.Name)
	}
}

func main() {
	st := &student{
		Name: "wjh",
	}
	zhoujielun(st)
}
