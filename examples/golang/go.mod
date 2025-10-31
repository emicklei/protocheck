module golang

replace github.com/emicklei/protocheck => ../..

go 1.24.0

toolchain go1.24.3

require (
	github.com/emicklei/protocheck v0.6.6
	github.com/google/cel-go v0.26.1
	google.golang.org/protobuf v1.36.10
)

require (
	cel.dev/expr v0.25.0 // indirect
	github.com/antlr4-go/antlr/v4 v4.13.1 // indirect
	github.com/stoewer/go-strcase v1.3.1 // indirect
	golang.org/x/exp v0.0.0-20251023183803-a4bb9ffd2546 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20251029180050-ab9386a59fda // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20251029180050-ab9386a59fda // indirect
)
