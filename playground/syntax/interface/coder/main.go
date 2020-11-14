package main

import "fmt"

type coder interface {
	code()
	debug()
}

type Gopher struct {
	language string
	Value int
}

func (p *Gopher) code() {
	p.Value = 123
	fmt.Printf("I am coding %s language\n", p.language)
}

func (p *Gopher) debug() {
	p.Value = 321
	fmt.Printf("I am debuging %s language\n", p.language)
}

func main() {
	gopher := &Gopher{"Go", 0}
	var c coder = gopher
	c.code()
	c.debug()
	gopher.debug()
	fmt.Println(gopher.Value)
}