package main

import (
	"fmt"
	"time"
)

func main() {
	// time.After
	tc := time.After(1 * time.Second)
	<- tc

	fmt.Println("time ok")
}
