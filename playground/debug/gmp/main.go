package main

import (
	"os"
	"runtime/trace"
	"time"
)

func main() {
	f, err := os.Create("./trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	trace.Start(f)

	//a := 0

	for i := 0; i < 10; i++ {
		go func(k int) {
			k = 0
			for i := 0; i < 1000000000; i++ {
				k++
			}
		}(i)
	}

	time.Sleep(time.Millisecond * 3000)

	trace.Stop()
}
