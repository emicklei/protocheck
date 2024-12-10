gen:
	protoc -I=../protos \
		--experimental_allow_proto3_optional \
		--go_out=. --go_opt=paths=source_relative \
		--protocheck_out=. \
	person.proto