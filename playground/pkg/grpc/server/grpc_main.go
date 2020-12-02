package main

import (
	"context"
	"github.com/Jameywoo/goto/playground/pkg/grpc/myrpc"
	"google.golang.org/grpc"
	"log"
	"net"
)

type HelloServiceImpl struct {}

func (p *HelloServiceImpl) Hello(ctx context.Context, args *myrpc.String) (*myrpc.String, error) {
	reply := &myrpc.String{Value: "hello:" + args.GetValue()}
	return reply, nil
}

func main() {
	grpcServer := grpc.NewServer()
	myrpc.RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)
}