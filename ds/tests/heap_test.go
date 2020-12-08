package tests

import (
	"fmt"
	"github.com/JameyWoo/goto/ds/heap"
	"testing"
)

func TestIntHeap(t *testing.T) {
	ip := heap.IntHeap{}
	_, err := ip.Pop()
	if err != nil {
		fmt.Println(err)
	}
	ip.Push(3)
	fmt.Println(ip.Top())
	ip.Push(1)
	fmt.Println(ip.Top())
	ip.Push(2)
	fmt.Println(ip.Top())
	fmt.Println(ip.Pop())
	fmt.Println(ip.Top())

	ip2 := heap.IntHeap{}
	nums := []int{3, 5, 1, 8, 4, 2, 9, 7, 3}
	for _, num := range nums {
		ip2.Push(num)
	}
	for !ip2.Empty() {
		val, err := ip2.Pop()
		if err != nil {
			panic(err)
		}
		fmt.Printf("%d ", val)
	}
}
