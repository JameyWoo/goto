package main

import (
	"context"
	"fmt"
	"github.com/Jameywoo/goto/playground/pkg/grpc/myrpc"
	_ "github.com/Jameywoo/goto/playground/pkg/grpc/myrpc"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := myrpc.NewHelloServiceClient(conn)
	obj, err := client.Hello(context.TODO(), &myrpc.String{Value: "Hello, World!"})
	if err != nil {
		panic(err)
	}
	fmt.Println(obj)
}
