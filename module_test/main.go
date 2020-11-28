package main

import (
	"count"
	"fmt"
	"github.com/JameyWoo/goto/ds/queue"  // 普通的网络包. 而且是仓库的子目录作为的独立包
	"github.com/JameyWoo/goto/module_test/funcs"  // 自己module内的一部分(非独立包)
	"github.com/JameyWoo/modtest"  // 本地module
	"github.com/JameyWoo/modtest/calcs"  // 本地module的一部分(非独立包)
	hello "github.com/Jameywoo/modtest/helloworld"  // 本地module内的单独module(独立包)
	"helloworld"  // 本地module
)

func main() {
	x := modtest.Add(1, 2)
	fmt.Println(x)
	q := queue.Queue{}
	q.Push(2)
	fmt.Println(q.Front())

	fmt.Println(calcs.Sub(1, 2))
	funcs.Hello()

	fmt.Println(count.One())

	helloworld.HelloWorld()

	hello.HelloWorld()
}
