package main

import (
	"fmt"
	"time"
)

func main() {
	n := time.Now()

	time.Sleep(3 * time.Second)

	if time.Since(n) > 2 * time.Second {
		fmt.Println("two Second!")
	}
}
