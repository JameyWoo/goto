module github.com/Jameywoo/goto/playground/pkg/grpc/client

go 1.14

require (
	github.com/golang/protobuf v1.4.3
	google.golang.org/grpc v1.33.2
	github.com/Jameywoo/goto/playground/pkg/grpc/myrpc v0.0.0
)

replace (
	github.com/Jameywoo/goto/playground/pkg/grpc/myrpc v0.0.0 => ../myrpc/
)