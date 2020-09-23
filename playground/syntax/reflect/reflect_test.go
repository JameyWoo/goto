package reflect

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func TestTypeOf(t *testing.T) {
	type Map map[string]int
	m := Map{"hello": 1, "world": 2}
	fmt.Println(m)

	ty := reflect.TypeOf(m)
	fmt.Println(ty)
	fmt.Println(ty.Key())
	fmt.Println(reflect.Value{})
	fmt.Println(ty.Size())
	fmt.Println(ty.String())
	fmt.Println(ty.String() == "reflect.Map")
}

type List struct {
	l [6]int
	m string
}

// 如果是普通变量这个string 不能是 *List, 只能是 List
// 如果是指针, 这个只能说 *List不能是List
// 所以这个要注意
func (l *List) String() string {
	out := "["
	for _, x := range l.l {
		out += strconv.Itoa(x) + ","
	}
	out += "]"
	return out
}

func TestPrintString(t *testing.T) {
	l := &List{}
	l.l = [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(l)
}

type Test struct {
	name string
}

func (t Test) String() string {
	return fmt.Sprintf("my name is %s", t.name)
}

func TestPrintString2(t *testing.T) {
	test := Test{"hello"}
	fmt.Println(test)
}