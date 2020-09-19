package conv

import (
	"fmt"
	"testing"
)

func TestCnov(t *testing.T) {
	type Map map[string]string
	a := Map{}
	a["hello"] = "world"

	fmt.Println(a)

	type Map2 Map
	var b Map2 = Map2(a)
	fmt.Println(b)
	println(b)  // 这个方法很奇怪, 输出了地址, 为什么
}

func TestSprintf(t *testing.T) {
	s := fmt.Sprintf("%s 去吃 %s吧\n", "你", "屎")
	fmt.Println(s)
}

func TestScan(t *testing.T) {
	var a int
	fmt.Scan(a)
	fmt.Println(a)
}