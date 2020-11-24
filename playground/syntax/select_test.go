package syntax

import (
	"fmt"
	"testing"
	"time"
)

func TestSelectTime(t *testing.T) {
	//ch := make(chan int)
	ti := time.NewTicker(2 * time.Second)
	for {
		select {
		case <- ti.C:
			fmt.Println("ch")
		//case <- ti.C:
		//	fmt.Println("time!")
		//default:
		//	fmt.Println("default")
		//	time.Sleep(1 * time.Second)
		}
		fmt.Println("out")
	}
}