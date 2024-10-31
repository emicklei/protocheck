fmt:
	protofmt -w check.proto

install:
	cd cmd/protoc-gen-protocheck && go install