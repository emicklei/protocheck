# protocheck

[![Go](https://github.com/emicklei/protocheck/actions/workflows/go.yml/badge.svg)](https://github.com/emicklei/protocheck/actions/workflows/go.yml)
[![GoDoc](https://pkg.go.dev/badge/github.com/emicklei/protocheck)](https://pkg.go.dev/github.com/emicklei/protocheck)

Lightweight solution to ProtocolBuffers message validation.
`protocheck-proto-gen` is a `protoc` plugin that generates Go code.

## features

- expressions with Google/CEL
- both message and field checks
- nested messages
- repeated, oneof and map fields
- syntax validation of CEL expressions at generation time
- supports proto3 and edition 2023

## install

```bash
  go install github.com/emicklei/protocheck/cmd/protoc-gen-protocheck@latest
```

### generate

```bash
  protoc --go_out=. --go_opt=paths=source_relative --protocheck_out=.
```

## example

```protobuf
package hr;

import "check.proto";

message Person {
  
  // message cross-field checks
  option (check.message) = { 
    cel:"size(this.name + this.surname) > 0" 
    fail:"name and surname cannot be empty" 
    id:"person_invariant" }; // id is optional
  
  // with per field state checks
  string name = 1 [(check.field) = { 
      cel:"size(this.name) > 1"                  
      fail:"name must be longer than 1" }];

  optional string middle_name = 2 [(check.field) = { 
      cel:"size(this.middle_name) > 0"           
      fail:"middle name (if set) cannot be empty" }];

  string surname = 3 [(check.field) = { 
      cel:"size(this.surname) > 1"               
      fail:"surname must be longer than 1" }];

  google.protobuf.Timestamp birth_date  = 4 [(check.field) = { 
      cel:"this.birth_date.getFullYear() > 2000" 
      id:"check_birth_date" }];
```
`getFullYear` is one of the supported [CEL macros](https://github.com/google/cel-spec/blob/master/doc/langdef.md#macros).

See [CEL language definition](https://github.com/google/cel-spec/blob/master/doc/langdef.md) for creating CEL expressions.

### usage

```go
p := &Person{
    Name:      "",
    BirthDate: &timestamppb.Timestamp{Seconds:1}
}
err := p.Validate() // a protocheck.ValidationError
```

### considerations

The `protocheck` package is inspired by `bufbuild/protovalidate` which also uses the powerful CEL expression language. However, it differs by:

  - `this` always refers to the message
  - minimal dependencies (cel and protobuf)
  - syntax is more compact