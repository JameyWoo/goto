package strings

import (
	"fmt"
	"strings"
	"testing"
)

func TestStrings(t *testing.T) {
	dots := "a,b,124, g, f"
	res := strings.Split(dots, ",")
	fmt.Print(res)
	for _, r := range res {
		fmt.Printf("-%s-\n", r)
	}
}
