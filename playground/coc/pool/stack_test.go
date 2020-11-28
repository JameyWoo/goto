package pool

import (
	"fmt"
	"github.com/Jameywoo/goto/ds/stack"
	"testing"
)

func TestStack(t *testing.T) {
	s := stack.Stack{}
	s.Push(4, 4, 6, 1, 2)
	fmt.Println(s.Top())
	fmt.Println(s.Size())
}
