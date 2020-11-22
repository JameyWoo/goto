package main

// https://github.com/lifei6671/interview-go/blob/master/question/q020.md

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	for i := 0; i < 100000000; i++ {
		//a :="aaa"
		//ssh := *(*reflect.StringHeader)(unsafe.Pointer(&a))
		//b := *(*[]byte)(unsafe.Pointer(&ssh))
		//_ = b

		a := "aaa"
		b := []byte(a)
		_ = b
	}
	end := time.Now()

	fmt.Println(end.Sub(start))
}