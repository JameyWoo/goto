package tests

import (
	"fmt"
	"github.com/JameyWoo/goto/ds/kit"
	"testing"
)

func TestKit(t *testing.T) {
	x := kit.PQ{}
	x.Push(&kit.Entry{Key: "shit", Priority: 1})
	x.Push(&kit.Entry{Key: "you", Priority: 2})
	x.Push(&kit.Entry{Key: "fuck", Priority: 3})
	x.Push(&kit.Entry{Key: "it", Priority: 0})
	fmt.Println(x.Pop())
}
