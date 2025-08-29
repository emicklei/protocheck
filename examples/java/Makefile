lib: 
	cd ../../lib/java && mvn install

install:
	cd ../../cmd/protoc-gen-protocheck && go install	

gen: install lib
	protoc -I=../protos \
		--experimental_allow_proto3_optional \
		--java_out=src/main/java \
		--protocheck_out=src/main/java \
		--protocheck_opt=paths=source_relative,lang=java \
		../protos/person.proto

test: gen
	mvn test