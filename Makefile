gen:
	protoc --go_out=. --go_opt=paths=source_relative check.proto

fmt:
	protofmt -w check.proto

install:
	cd cmd/protoc-gen-protocheck && go install