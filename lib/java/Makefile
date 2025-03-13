gen: gen_test gen_main

gen_main:
	protoc -I=../../ --java_out=src/main/java ../../check.proto

gen_test:
	protoc -I=src/test/resources --java_out=src/test/java src/test/resources/bike.proto