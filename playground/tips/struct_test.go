package tips

import (
	"log"
	"reflect"
	"testing"
	"unsafe"
)

type Type1 int
type Type2 int

func TestStructComp(t *testing.T) {
	a := Type1(1)
	b := Type2(1)
	log.Printf("a = %d, b = %d\n", a, b)
	e := Type1(b)
	//log.Println(a == b)
	log.Printf("ta: %s, tb: %s\n", reflect.TypeOf(a), reflect.TypeOf(b))
	log.Printf("ta: %T, tb: %T\n", a, e)
	log.Println(unsafe.Sizeof(a))

	c := Type1(3)
	d := Type1(3)
	log.Println(c == d)

	log.Println(reflect.DeepEqual(a, b))
	f := reflect.ValueOf(a)
	log.Println(f)
}
