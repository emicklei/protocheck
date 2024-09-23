gen:
	protoc -I=. -I=.. \
		--go_out=. --go_opt=paths=source_relative \
		--protocheck_out=. \
	holiday.proto person.proto

fmt:
	protofmt -w sample.proto 
	protofmt -w holiday.proto
