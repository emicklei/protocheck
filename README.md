# protocheck

Lightweight solution to ProtocolBuffers message validation.
`protocheck-proto-gen` is a `protoc` plugin that generates Go code.

## status

  Work in progress

## example

```go
package hr;

message Person {
  
  // overal message state check
  option (check.message) = { fail:"name and description cannot be empty" id:"person_invariant" cel:"size(this.name + this.description) > 0" };
  
  // per field check
           string                    name        = 1 [(check.field) = { cel :"size(this.name) > 1"         }  ];
  optional string                    description = 2 [(check.field) = { fail:"description cannot be empty" cel:"size(this.description) > 0" }];
           google.protobuf.Timestamp birth_date  = 3 [(check.field) = { id  :"check_birth_date"            cel:"this.birth_date < now"      }];
}
```

### usage

```go
p := &Person{
    Name:      "Joe",
    BirthDate: &timestamppb.Timestamp{Seconds: timestamppb.Now().Seconds + 1}
}
err := p.Validate()
```

## comparing to the envoyproxy/protoc-gen-validate package


## comparing to the bufbuild/protovalidate package

The `protocheck` packages is inspired by this and both packages are using the powerful CEL expression language and implementation.

- `this` refers to the message
- slightly more compact syntax
