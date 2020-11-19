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

	for i := 0; i < 100000; i++ {
		go func(k int) {
			k = 0
			for i := 0; i < 1; i++ {
				k++
			}
			time.Sleep(300 * time.Millisecond)
		}(i)
	}

	time.Sleep(time.Millisecond * 30000)

	trace.Stop()
}
