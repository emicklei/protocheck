gen:
	protoc -I=. -I=.. --go_out=.  --go_opt=paths=source_relative check.proto sample.proto

fmt:
	protofmt -w sample.proto
	protofmt -w ../check.proto
