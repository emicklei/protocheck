package golang

import (
	"testing"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func TestIsSetConditionSource(t *testing.T) {
	tests := []struct {
		name     string
		kind     protoreflect.Kind
		expected string
	}{
		{"MessageKind", protoreflect.MessageKind, "typedX.GetField() != nil"},
		{"StringKind", protoreflect.StringKind, "typedX.GetField() != \"\""},
		{"Int32Kind", protoreflect.Int32Kind, "typedX.GetField() != 0"},
		{"Int64Kind", protoreflect.Int64Kind, "typedX.GetField() != 0"},
		{"Uint32Kind", protoreflect.Uint32Kind, "typedX.GetField() != 0"},
		{"Uint64Kind", protoreflect.Uint64Kind, "typedX.GetField() != 0"},
		{"Sint32Kind", protoreflect.Sint32Kind, "typedX.GetField() != 0"},
		{"Sint64Kind", protoreflect.Sint64Kind, "typedX.GetField() != 0"},
		{"Sfixed32Kind", protoreflect.Sfixed32Kind, "typedX.GetField() != 0"},
		{"Sfixed64Kind", protoreflect.Sfixed64Kind, "typedX.GetField() != 0"},
		{"Fixed32Kind", protoreflect.Fixed32Kind, "typedX.GetField() != 0"},
		{"Fixed64Kind", protoreflect.Fixed64Kind, "typedX.GetField() != 0"},
		{"DoubleKind", protoreflect.DoubleKind, "typedX.GetField() != 0.0"},
		{"BytesKind", protoreflect.BytesKind, "len(typedX.GetField()) > 0"},
		{"BoolKind", protoreflect.BoolKind, "typedX.GetField() != false"},
		{"FloatKind", protoreflect.FloatKind, "typedX.GetField() != 0.0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := &protogen.Field{
				Desc:   &mockFieldDescriptor{kind: tt.kind},
				GoName: "Field",
			}
			result := isSetConditionSource(field)
			if result != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, result)
			}
		})
	}
}

type mockFieldDescriptor struct {
	protoreflect.FieldDescriptor
	kind protoreflect.Kind
}

func (m *mockFieldDescriptor) Kind() protoreflect.Kind {
	return m.kind
}
