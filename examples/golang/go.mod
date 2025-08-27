module golang

replace github.com/emicklei/protocheck => ../..

go 1.24

toolchain go1.24.3

require (
	github.com/emicklei/protocheck v0.0.0-20241212220435-83bbc3a88fb5
	github.com/google/cel-go v0.26.0
	google.golang.org/protobuf v1.36.7
)

require (
	cel.dev/expr v0.24.0 // indirect
	github.com/antlr4-go/antlr/v4 v4.13.1 // indirect
	github.com/stoewer/go-strcase v1.3.1 // indirect
	golang.org/x/exp v0.0.0-20250808145144-a408d31f581a // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20250804133106-a7a43d27e69b // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250804133106-a7a43d27e69b // indirect
)
