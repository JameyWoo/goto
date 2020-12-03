module github.com/Jameywoo/goto/playground/pkg/grpc/client

go 1.14

require (
	github.com/Jameywoo/goto/playground/pkg/grpc/myrpc v0.0.0
	google.golang.org/grpc v1.33.2
)

replace github.com/Jameywoo/goto/playground/pkg/grpc/myrpc v0.0.0 => ../myrpc/
