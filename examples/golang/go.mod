module golang

replace github.com/emicklei/protocheck => ../../../protocheck

go 1.23

require (
	github.com/emicklei/protocheck v0.0.0-20241212220435-83bbc3a88fb5
	github.com/google/cel-go v0.24.1
	google.golang.org/protobuf v1.36.6
)

require (
	cel.dev/expr v0.19.1 // indirect
	github.com/antlr4-go/antlr/v4 v4.13.1 // indirect
	github.com/stoewer/go-strcase v1.3.0 // indirect
	golang.org/x/exp v0.0.0-20250106191152-7588d65b2ba8 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20250106144421-5f5ef82da422 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250106144421-5f5ef82da422 // indirect
)
