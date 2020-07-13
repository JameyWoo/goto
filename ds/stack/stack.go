package stack

import (
	"errors"
	"fmt"
)

type Stack []int

func (stack *Stack) Push(x ...int) {
	*stack = append(*stack, x...)
}

func (stack *Stack) Top() int {
	if len(*stack) > 0 {
		return (*stack)[len(*stack) - 1]
	} else {
		panic(0)
		return 0
	}
}

func (stack *Stack) Pop() error {
	if len(*stack) > 0 {
		*stack = (*stack)[:len(*stack)-1]
		return nil
	} else {
		return errors.New("Error: Stack is empty")
	}
}

func (stack *Stack) IsEmpty() bool {
	return len(*stack) == 0
}

func (stack *Stack) Size() int {
	return len(*stack)
}

// 自低向顶打印栈
func (stack *Stack) print() {
	for _, x := range *stack {
		fmt.Printf("%d ", x)
	}
	fmt.Println()
}