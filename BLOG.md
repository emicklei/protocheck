
# protocheck: Lightweight Protobuf Validation in Go

For Go developers working with Protocol Buffers, ensuring data integrity is crucial. While Protobufs guarantee type safety, they don't enforce semantic rules out of the box. This is where `protocheck` comes in, a lightweight and powerful tool for adding validation logic to your Protobuf messages.

This blog post will explore the design choices behind `protocheck`, focusing on its simplicity of declaration and code generation. We'll also compare it to `bufbuild/protovalidate` to help you choose the right tool for your project.

## Simplicity of Declaration

One of the standout features of `protocheck` is its intuitive and straightforward approach to defining validation rules. The validations are declared directly within the `.proto` file using custom options. This co-location of data structure and validation rules makes the schema easy to read and maintain.

Let's look at an example from the `person.proto` file in the `protocheck` repository:

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

The `protocheck` tool is a `protoc` plugin that generates code for the validation rules defined in the `.proto` file. The generated code is placed in a separate file (e.g., `person.check.go`), which keeps the validation logic separate from the core Protobuf-generated code. A key design benefit is that the validation definitions in the `.proto` file are language-agnostic. This allows the `protoc-gen-protocheck` plugin to be extended to support multiple programming languages. Currently, it supports both Go and Java.

The generated code is efficient and easy to use. For each message with validation rules, a `Validate()` method is generated. This method can be called to check if a message instance is valid.

Here's an example of how to use the generated `Validate()` method in Go:

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

## protocheck vs. bufbuild/protovalidate

The `protocheck` package is inspired by `bufbuild/protovalidate`, which also uses the powerful CEL expression language for defining validation rules. However, `protocheck` makes deliberate design choices that will appeal to experienced developers who prioritize simplicity, elegance, and minimal dependencies.

While `bufbuild/protovalidate` is a powerful, feature-rich framework, `protocheck` offers a more focused and lightweight approach. Here’s why you might prefer `protocheck`:

*   **Design Simplicity and Consistency:** In `protocheck`, the `this` keyword *always* refers to the message being validated, even in field-level validations. This elegant consistency simplifies the mental model and reduces the cognitive load when writing validation rules. In `bufbuild/protovalidate`, the context of `this` can change, which can lead to more complex and less predictable expressions.
*   **Minimal Dependencies:** `protocheck` keeps your dependency tree lean. It relies only on the core `cel` and `protobuf` libraries for each language. This makes it a lightweight choice that is easy to integrate into existing projects without introducing unnecessary bloat.
*   **Developer Control over Readability:** `protocheck` trusts the developer to write clear and effective CEL expressions. The compact syntax keeps your definitions clean and readable, putting the power of CEL directly in your hands. `protocheck` is designed for developers who prefer the expressiveness of a powerful language like CEL over a large set of pre-defined, and sometimes limiting, rules.

In short, `protocheck` is not just a validation tool; it's a statement of a design philosophy that favors simplicity, developer control, and a lean dependency graph.

## Future Development: Adding Python Support

The modular design of `protocheck`'s code generator makes it straightforward to add support for new languages. Here is a high-level plan for adding Python support:

1.  **Create a Python Generator:** A new generator will be created in the `cmd/protoc-gen-protocheck/lang/` directory, alongside the existing `golang` and `java` implementations.
2.  **Implement the Generator Interface:** This new generator will be responsible for translating the protobuf validation rules into Python-specific code.
3.  **Develop a Python Template:** A template file will be created to define the structure of the generated Python validation methods. This template will use the `cel-python` library to execute the CEL expressions.
4.  **Integrate into the Plugin:** Finally, the `lang=python` option will be added to the `protoc-gen-protocheck` plugin, allowing users to generate Python validation code.

This clear path to extension demonstrates the robustness and forward-thinking design of the `protocheck` package.

## Conclusion

For the discerning developer, `protocheck` offers a compelling solution for Protobuf validation. It's a tool that respects your desire for simplicity, minimal dependencies, and design consistency. While `bufbuild/protovalidate` is a comprehensive framework, `protocheck` provides a more elegant, lightweight, and extensible alternative.

If you value a focused tool that gives you the full power of CEL without the overhead of a heavy framework, `protocheck` is the clear choice. It helps you write clean, maintainable, and efficient validation logic, making it a valuable addition to any developer's toolkit.
