package main

/*
只能模拟, 不能
 */

import "fmt"

func main() {
	stu := Student{
		Name: "Tom",
		Male: true,
		S: []*Score{&Score{Score: []int32{1, 2, 3}},
			&Score{Score: []int32{4, 5, 6}}},
	}

	fmt.Println(stu)
}
