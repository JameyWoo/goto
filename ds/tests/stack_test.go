package tests

import (
	"fmt"
	stack2 "github.com/JameyWoo/goto/ds/stack"
	"testing"
)

func TestStack(t *testing.T) {
	stack2 := stack2.Stack{}
	//stack2.Pop()
	//stack2.Top()
	stack2.Push(1, 3, 5, 7)
	stack2.Push(4)
	
	stack2.Pop()
	stack2.Pop()
	fmt.Println(stack2.Top())
}
