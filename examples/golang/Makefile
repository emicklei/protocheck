install:
	cd ../../cmd/protoc-gen-protocheck && go install

gen: install
	protoc -I=../protos \
		--experimental_allow_proto3_optional \
		--go_out=. \
		--go_opt=paths=source_relative \
		--protocheck_out=. \
		--protocheck_opt=paths=source_relative \
	person.proto
	gofmt -w person.check.go