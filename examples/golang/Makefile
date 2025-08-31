install:
	cd ../../cmd/protoc-gen-protocheck && go install

gen: install
	protoc -I=../protos \
		--experimental_allow_proto3_optional \
		--go_out=. \
		--go_opt=paths=source_relative \
		--protocheck_out=. \
		--protocheck_opt=paths=source_relative \
	person.proto toy.proto
	gofmt -w person.check.go

opaque: install
	mkdir -p api_opaque/golang
	protoc -I=../protos \
		--experimental_allow_proto3_optional \
		--go_out=api_opaque/golang \
		--go_opt=paths=source_relative,default_api_level=API_OPAQUE \
		--protocheck_out=api_opaque/golang \
		--protocheck_opt=paths=source_relative,default_api_level=API_OPAQUE \
	person.proto
	gofmt -w api_opaque/golang/person.check.go