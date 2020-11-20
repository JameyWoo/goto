package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"time"
)

func main() {
	runtime.GOMAXPROCS(100)
	f, err := os.Create("./trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	trace.Start(f)

	//a := 0

	for i := 0; i < 10000; i++ {
		go func(k int) {
			k = 0
			for i := 0; i < 1000000; i++ {
				k++
			}
			//time.Sleep(300 * time.Millisecond)
			fmt.Print("")
		}(i)
	}

	time.Sleep(time.Millisecond * 3000)
	trace.Stop()
}