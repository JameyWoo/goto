package _testing

import (
	"fmt"
	"testing"
)

func ExampleAdd() {
	a := 1
	b := 2
	fmt.Println(a + b)
	// Output: 3
}

func BenchmarkAdd(b *testing.B) {
	a := 1
	c := 2
	fmt.Println(a + c)
	// Output: 3
}