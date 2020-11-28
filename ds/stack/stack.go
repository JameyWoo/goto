package stack

import (
	"errors"
	"fmt"
)

type Stack []int

func (Stack *Stack) Push(x ...int) {
	*Stack = append(*Stack, x...)
}

func (Stack *Stack) Top() int {
	if len(*Stack) > 0 {
		return (*Stack)[len(*Stack) - 1]
	} else {
		panic(0)
		return 0
	}
}

func (Stack *Stack) Pop() error {
	if len(*Stack) > 0 {
		*Stack = (*Stack)[:len(*Stack)-1]
		return nil
	} else {
		return errors.New("Error: Stack is empty")
	}
}

func (Stack *Stack) IsEmpty() bool {
	return len(*Stack) == 0
}

func (Stack *Stack) Size() int {
	return len(*Stack)
}

// 自低向顶打印栈
func (Stack *Stack) print() {
	for _, x := range *Stack {
		fmt.Printf("%d ", x)
	}
	fmt.Println()
}