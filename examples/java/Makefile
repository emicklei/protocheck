gen: install
	protoc -I=../protos \
		--experimental_allow_proto3_optional \
		--java_out=src/main/java \
		--protocheck_out=src/main/java \
		--protocheck_opt=paths=source_relative,lang=java \
		../protos/person.proto

install:
	cd ../../cmd/protoc-gen-protocheck && go install	