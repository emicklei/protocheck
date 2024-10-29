# go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
# go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
# go install github.com/emicklei/proto-contrib/cmd/protofmt@latest
gen:
	cd ../cmd/protoc-gen-protocheck && go install
	protoc -I=. -I=.. \
		--go_out=. --go_opt=paths=source_relative \
		--protocheck_out=. \
	holiday.proto person.proto

fmt:
	protofmt -w person.proto 
	protofmt -w holiday.proto
