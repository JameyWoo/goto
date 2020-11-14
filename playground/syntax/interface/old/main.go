package main

import "fmt"

type Person struct {
	Age int
}

func (p *Person) howOld() int {
	p.Age = 0
	return p.Age
}

func (p *Person) growUp() {
	p.Age += 1
}

func main() {
	// qcrao 是值类型
	qcrao := Person{Age: 18}

	// 值类型 调用接收者也是值类型的方法
	fmt.Println(qcrao.howOld())

	// 值类型 调用接收者是指针类型的方法
	qcrao.growUp()
	fmt.Println(qcrao.Age)

	// ----------------------

	// stefno 是指针类型
	stefno := &Person{Age: 100}

	// 指针类型 调用接收者是值类型的方法
	fmt.Println(stefno.howOld())

	// 指针类型 调用接收者也是指针类型的方法
	stefno.growUp()
	fmt.Println(stefno.Age)
}