
# protoc-gen-validate: Lightweight Protobuf Validation in Go

For Go developers working with Protocol Buffers, ensuring data integrity is crucial. While Protobufs guarantee type safety, they don't enforce semantic rules out of the box. This is where `protoc-gen-validate` comes in, a lightweight and powerful tool for adding validation logic to your Protobuf messages.

This blog post will explore the design choices behind `protoc-gen-validate`, focusing on its simplicity of declaration and code generation. We'll also compare it to its successor, `buf-validate`, to help you choose the right tool for your project.

## Simplicity of Declaration

One of the standout features of `protoc-gen-validate` is its intuitive and straightforward approach to defining validation rules. The validations are declared directly within the `.proto` file using custom options. This co-location of data structure and validation rules makes the schema easy to read and maintain.

Let's look at an example from the `person.proto` file in the `protoc-gen-validate` repository:

```protobuf
message Person {
  // message cross-field checks
  option (check.message) = {
    cel:"size(this.name + this.surname) > 0"
    fail:"name and surname cannot be empty"
    id:"person_invariant"
  };

  // with per field state checks
  string name = 1 [(check.field) = {
      cel:"size(this.name) > 1"
      fail:"name must be longer than 1"
  }];

  optional string middle_name = 2 [(check.field) = {
      cel:"size(this.middle_name) > 0"
      fail:"middle name (if set) cannot be empty"
  }];
}
```

As you can see, the validation rules are expressed using the Common Expression Language (CEL), a powerful and flexible expression language from Google. The `(check.message)` option is used for message-level checks, while the `(check.field)` option is used for field-level checks.

This approach has several advantages:

*   **Readability:** The validation rules are right next to the fields they apply to, making it easy to understand the expected data format.
*   **Single Source of Truth:** The `.proto` file becomes the single source of truth for both the data structure and its validation rules.
*   **Language Agnostic:** The validation rules are defined in a language-agnostic way, which means they can be used to generate validation code in any supported language.

## Code Generation

The `protoc-gen-validate` tool is a `protoc` plugin that generates Go code for the validation rules defined in the `.proto` file. The generated code is placed in a separate file (e.g., `person.check.go`), which keeps the validation logic separate from the core Protobuf-generated code.

The generated code is efficient and easy to use. For each message with validation rules, a `Validate()` method is generated. This method can be called to check if a message instance is valid.

Here's an example of how to use the generated `Validate()` method:

```go
p := &Person{
    Name:      "",
    BirthDate: &timestamppb.Timestamp{Seconds:1}
}
if err := p.Validate(); err != nil {
  for _ , each := range err {
    log.Println(each)
  }
}
```

The generated code is also highly optimized. It uses a shared CEL environment to avoid the overhead of creating a new environment for each validation check.

## protoc-gen-validate vs. buf-validate

The `buf-validate` package is the successor to `protoc-gen-validate`. It's built on the same principles but offers a more powerful and flexible validation framework.

Here are some of the key differences between the two packages:

*   **Standardized Rules:** `buf-validate` provides a set of standardized validation rules for common use cases, such as string length, number range, and regex matching. This can save you from writing custom CEL expressions for common scenarios.
*   **More Powerful CEL Integration:** `buf-validate` offers a more powerful and flexible integration with CEL. It allows you to define custom functions and macros, which can be used to create more complex validation rules.
*   **Better Error Reporting:** `buf-validate` provides more detailed and structured error messages, which can be helpful for debugging.
*   **Active Development:** `buf-validate` is actively developed and maintained by Buf, while `protoc-gen-validate` is no longer actively developed.

## Conclusion

The `protoc-gen-validate` package is a lightweight and powerful tool for adding validation logic to your Protobuf messages. Its simplicity of declaration and efficient code generation make it a great choice for projects that need a straightforward and easy-to-use validation solution.

However, if you need a more powerful and flexible validation framework, or if you want to take advantage of the latest features and improvements, you should consider using `buf-validate`.

Ultimately, the choice between `protoc-gen-validate` and `buf-validate` depends on the specific needs of your project. If you're looking for a simple and easy-to-use solution, `protoc-gen-validate` is a great choice. If you need a more powerful and flexible framework, `buf-validate` is the way to go.
