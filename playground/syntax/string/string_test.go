package string

import (
	"fmt"
	"testing"
)

func TestStringBytes(t *testing.T) {
	var b []byte = []byte{50, 48, 49}

	c := string(b)
	fmt.Println(c)
	b[0] = 51
	fmt.Print(c)
}
