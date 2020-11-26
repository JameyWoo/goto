package syntax

import (
	"fmt"
	"testing"
)

func TestAppendAssign(t *testing.T) {
	s := make([]int, 2, 4)
	fmt.Println(s)

	s2 := append(s, 3)
	fmt.Println(s)

	fmt.Printf("%p\n", s)
	fmt.Printf("%p\n", s2)
}
