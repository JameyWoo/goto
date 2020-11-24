package coc

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMain0(t *testing.T) {
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println(123)
	fmt.Printf("%d, %d", 1, 2)

}

func TestChanInterface(t *testing.T) {
	ch := make(chan int)

	var a interface{} = ch
	fmt.Println(a)

	chanInter(ch)
}

func chanInter(i interface{}) {
	fmt.Println(i)
	fmt.Println(reflect.TypeOf(i))

	i = i.(chan int)
}