package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"reflect"
)

func main() {
	stu := &Student{
		Name: "Timi",
		Male: true,
		Scores: []int32{99, 91, 96},
	}
	fmt.Printf("%v\n", stu)

	data, err := proto.Marshal(stu)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(data))

	stu2 := &Student{}
	err = proto.Unmarshal(data, stu2)
	fmt.Println(stu2)
	fmt.Println(reflect.TypeOf(stu2))
}
