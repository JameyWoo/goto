package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	stack := Stack{}
	stack.Pop()
	stack.Top()
	stack.Push(1, 3, 5, 7)
	stack.Push(4)
	stack.print()
	
	stack.Pop()
	stack.Pop()
	fmt.Println(stack.Top())
}
