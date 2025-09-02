tools:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/emicklei/proto-contrib/cmd/protofmt@latest
	
gen:
	protoc -I=. \
		--go_out=. --go_opt=paths=source_relative \
	check.proto

fmt:
	protofmt -w check.proto

install:
	cd cmd/protoc-gen-protocheck && go install

test:
	go test -v -race -cover

test-golang: test install
	protoc -I=examples/protos \
		--experimental_allow_proto3_optional \
		--go_out=examples/golang \
		--go_opt=paths=source_relative \
		--protocheck_out=examples/golang \
		--protocheck_opt=paths=source_relative \
	person.proto toy.proto
	gofmt -w examples/golang/person.check.go
	cd examples/golang && go test

test-golang-opaque: test install
	protoc -I=examples/protos \
		--experimental_allow_proto3_optional \
		--go_out=examples/golang/api_opaque/golang \
		--go_opt=paths=source_relative,default_api_level=API_OPAQUE  \
		--protocheck_out=examples/golang/api_opaque/golang \
		--protocheck_opt=paths=source_relative,default_api_level=API_OPAQUE  \
	person.proto toy.proto
	gofmt -w examples/golang/api_opaque/golang/person.check.go
	cd examples/golang/api_opaque/golang && go test

test-java: test install
