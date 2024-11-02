# npm install typescript --save-dev --global && npm i tsx -g
# npm install ts-protoc-gen -g
run:
	tsx hello.ts

runjs:
	tsc hello.ts
	node hello.js

gen:
	protoc \
		--ts_opt=esModuleInterop=true \
		--ts_out="." \
		person.proto check.proto