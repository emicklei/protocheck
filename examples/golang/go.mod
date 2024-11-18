module golang

replace github.com/emicklei/protocheck => ../../../protocheck

go 1.23

require (
	github.com/emicklei/protocheck v0.0.0-00010101000000-000000000000
	github.com/google/cel-go v0.22.0
	google.golang.org/protobuf v1.35.1
)

require (
	cel.dev/expr v0.18.0 // indirect
	github.com/antlr4-go/antlr/v4 v4.13.1 // indirect
	github.com/stoewer/go-strcase v1.3.0 // indirect
	golang.org/x/exp v0.0.0-20241108190413-2d47ceb2692f // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20241104194629-dd2ea8efbc28 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241104194629-dd2ea8efbc28 // indirect
)
