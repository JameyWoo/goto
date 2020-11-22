package main

import (
	"fmt"
	"github.com/Jameywoo/goto/interview/go单例模式/Go语言惯用的单例模式/singleton"
)

func main() {
	instance := singleton.GetInstance()
	fmt.Println(instance)
	fmt.Printf("%p\n", instance)

	instance2 := singleton.GetInstance()
	fmt.Println(instance2)
	fmt.Printf("%p\n", instance2)

	c := 2
	fmt.Printf("&c = %p\n", &c)
	fmt.Println(&c)
}
