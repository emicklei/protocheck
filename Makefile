# go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
# go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
# go install github.com/emicklei/proto-contrib/cmd/protofmt@latest
gen:
	protoc -I=. \
		--go_out=. --go_opt=paths=source_relative \
	check.proto

fmt:
	protofmt -w check.proto

install:
	cd cmd/protoc-gen-protocheck && go install