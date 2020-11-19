package main

import "fmt"

func main() {

	done := make(chan bool)

	values := []string{"a", "b", "c"}
	x := ""
	cou := 0

	for _, v := range values {
		fmt.Println("--->", v)
		go func(u string) {
			//fmt.Println(u)
			cou += 1
			x = v
			done <- true
		}(v)
	}

	// wait for all goroutines to complete before exiting
	for _ = range values {
		<-done
	}

	fmt.Println(x)
	fmt.Println(cou)

}