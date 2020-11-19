package main

import (
	"os"
	"runtime/trace"
)

func foo(arg_val int, f *os.File)(*int) {
	trace.Start(f)
	var foo_val int = 11;
	return &foo_val;
}

func main() {
	f, err := os.Create("./trace")
	if err != nil {
		panic(err)
	}
	defer f.Close()


	x := foo(3, f)
	*x = 4

	trace.Stop()
}
