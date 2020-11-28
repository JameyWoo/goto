package main

import (
	"fmt"
	"github.com/JameyWoo/goto/ds/queue"
	"github.com/JameyWoo/goto/module/funcs"
	"github.com/JameyWoo/modtest"
	"github.com/JameyWoo/modtest/calcs"
)

func main() {
	x := modtest.Add(1, 2)
	fmt.Println(x)
	q := queue.Queue{}
	q.Push(2)
	fmt.Println(q.Front())

	fmt.Println(calcs.Sub(1, 2))
	funcs.Hello()
}
