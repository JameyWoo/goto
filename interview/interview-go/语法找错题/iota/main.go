package main
import "fmt"
const (
	a = iota
	b = iota
	e = '1'
	f = iota
)
const (
	name = "menglu"
	_
	c    = iota
	d    = iota
)
func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}