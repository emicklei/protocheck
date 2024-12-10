# protocheck

Lightweight solution to ProtocolBuffers message validation.
`protocheck-proto-gen` is a `protoc` plugin that generates Go code.

## status

  Work in progress

## example

```go
package hr;

message Person {
  
  // multiple message state checks
  option (check.message) = { cel:"size(this.name + this.description) > 0" fail:"name and description cannot be empty" id:"person_invariant"  };
  
  // with per field state checks
           string                    name        = 1 [(check.field) = { cel:"size(this) > 1"            fail:"name cannot be empty"        }];
  optional string                    description = 2 [(check.field) = { cel:"size(this) > 0"            fail:"description cannot be empty" }];
           google.protobuf.Timestamp birth_date  = 3 [(check.field) = { cel:"this.getFullYear() > 1000" id:"check_birth_date"              }];
}
```

`getFullYear` is one of the supported [CEL macros](https://github.com/google/cel-spec/blob/master/doc/langdef.md#macros).

### usage

```go
p := &Person{
    Name:      "",
    BirthDate: &timestamppb.Timestamp{Seconds:1}
}
err := p.Validate()
```

The `protocheck` package is inspired by `bufbuild/protovalidate` which also uses the powerful CEL expression language.
However, it differs by:

  - `this` always refers to the (root) message
  - no dependency on the buf package
  - syntax is more compact