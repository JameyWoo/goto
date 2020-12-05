package tests

import (
	"fmt"
	"github.com/JameyWoo/goto/ds/bloom_filter"
	"testing"
)

func TestBloomFilterInit(t *testing.T) {
	bf := bloom_filter.NewBloomFilter()
	//fmt.Println(bf)
	bf.Put("hello")
	bf.Put("world")
	bf.Put("shit")
	bf.Put("abcd")
	bf.Put("event")
	bf.Put("event1")

	fmt.Println(bf.Get("event"))
	fmt.Println(bf.Get("nihao"))
}
