package main

import (
	"fmt"
	"github.com/Jameywoo/goto/pool"
	"time"
)

func main() {
	p := pool.NewPool(10)
	go func() {
		for {
			p.Add(func() error {
				fmt.Println("running task")
				time.Sleep(500 * time.Microsecond)
				fmt.Println("task over")
				return nil
			})
		}
	}()
	p.Run()
	fmt.Println("over ...")
}