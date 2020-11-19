package main

func foo(arg_val int)(*int) {

	var fooVal int = 11;
	return &fooVal;
}

func main() {

	mainVal := foo(666)

	println(*mainVal)
}