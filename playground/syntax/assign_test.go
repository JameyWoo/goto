package syntax

import (
	"fmt"
	"testing"
)

func TestAssign(t *testing.T) {
	i := 0
	fmt.Printf("*i = %p\n", &i)
	if true {
		i, j := 1, 2
		fmt.Printf("*i = %p\n", &i)
		fmt.Println(i, j)
	}
	fmt.Println(i)

	a := 1
	fmt.Printf("%p\n", &a)
	a, b := 1, 2
	fmt.Printf("%p, %p\n", &a, &b)
}

func TestMultiArgs(t *testing.T) {
	argsTest(1, 2, 3)

	x := []int{0, 1, 2, 3, 4}
	fmt.Println(x)
}

func argsTest(args ...int)  {
	fmt.Printf("type: %T\n", args)
}
