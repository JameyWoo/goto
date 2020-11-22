package main

import "fmt"

func main() {
	pase_student()
}

type student struct {
	Name string
	Age  int
}

func (m student) String() string {
	return fmt.Sprintf("{%s: %d}", m.Name, m.Age)
}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for i, stu := range stus {
		fmt.Println(stu)
		//m[stu.Name] = &stu
		m[stu.Name] = &stus[i]
	}
	fmt.Println(m)
}