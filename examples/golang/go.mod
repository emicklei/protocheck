module golang

replace github.com/emicklei/protocheck => ../../../protocheck

go 1.23

require (
	github.com/emicklei/protocheck v0.0.0-20241212220435-83bbc3a88fb5
	github.com/golang/protobuf v1.5.4
	github.com/google/cel-go v0.22.1
	google.golang.org/protobuf v1.36.1
)

require (
	cel.dev/expr v0.19.1 // indirect
	github.com/antlr4-go/antlr/v4 v4.13.1 // indirect
	github.com/stoewer/go-strcase v1.3.0 // indirect
	golang.org/x/exp v0.0.0-20241217172543-b2144cdd0a67 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20241223144023-3abc09e42ca8 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241223144023-3abc09e42ca8 // indirect
)
