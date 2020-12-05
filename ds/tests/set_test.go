package tests

import (
	"fmt"
	"github.com/JameyWoo/goto/ds/set"
	"log"
	"testing"
)

func TestSet(t *testing.T) {
	// 新建 set
	s := set.NewSet()
	log.Printf("size: %d\n", s.Size())

	s.Insert("Hello")
	s.Insert(1)
	s.Insert([...]int{1, 2, 3, 4, 5})
	log.Printf("size: %d\n", s.Size())
	log.Printf("%t\n", s.Find("Hello"))

	s.Delete([3]int{1, 2, 3})
	s.Delete("Hello")
	s.Insert("world")
	s.Insert(true)

	// 获取所有的key
	iter := s.Iter()
	log.Printf("set size: %d\n", s.Size())
	log.Printf("iter size: %d\n", len(iter))
	for _, x := range iter {
		fmt.Printf("key = %v\n", x)
	}
}

func TestInterfaceSlice(t *testing.T) {
	var inters []interface{}
	inters = append(inters, "hello", "shit", [2]int{3, 1020})

	for x := range inters {
		log.Println(x)
	}
}