package shared

import (
	"testing"

	"google.golang.org/protobuf/reflect/protoreflect"
)

func TestIsCELParseable(t *testing.T) {
	tests := []struct {
		expr string
		want bool
	}{
		{"size(this.name) > 3", true},
		{"size(this.name > 3", false},
		{"matches(this.name,'^[0-9]$')", true},
		{"this.name.matches('^[0-9]$')", true},
	}
	for _, tt := range tests {
		err := parseCEL(tt.expr)
		if got := err == nil; got != tt.want {
			t.Errorf("isCELParseable() = %v, want %v, error = %v", got, tt.want, err)
		}
	}
}

func TestDotsAndSlashsToUnderscore(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"", ""},
		{"simple", "simple"},
		{"path/to/file", "path_to_file"},
		{"package.proto", "package_proto"},
		{"path/to/package.proto", "path_to_package_proto"},
		{"./relative.proto", "__relative_proto"},
		{"../parent.proto", "___parent_proto"},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := dotsAndSlashsToUnderscore(tt.input)
			if got != tt.want {
				t.Errorf("dotsAndSlashsToUnderscore(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestMapKindToJavaType(t *testing.T) {
	tests := []struct {
		kind protoreflect.Kind
		want string
	}{
		{protoreflect.BoolKind, "Bool"},
		{protoreflect.StringKind, "java.lang.String"},
		{protoreflect.Int32Kind, "Object"},
		{protoreflect.Int64Kind, "Object"},
		{protoreflect.FloatKind, "Object"},
		{protoreflect.DoubleKind, "Object"},
	}
	for _, tt := range tests {
		t.Run(tt.kind.String(), func(t *testing.T) {
			got := mapKindToJavaType(tt.kind)
			if got != tt.want {
				t.Errorf("mapKindToJavaType(%v) = %q, want %q", tt.kind, got, tt.want)
			}
		})
	}
}

func TestIfEmpty(t *testing.T) {
	tests := []struct {
		content string
		alt     string
		want    string
	}{
		{"", "alternative", "alternative"},
		{"content", "alternative", "content"},
		{"   ", "alternative", "   "},
		{"non-empty", "", "non-empty"},
	}
	for _, tt := range tests {
		t.Run(tt.content+"_"+tt.alt, func(t *testing.T) {
			got := ifEmpty(tt.content, tt.alt)
			if got != tt.want {
				t.Errorf("ifEmpty(%q, %q) = %q, want %q", tt.content, tt.alt, got, tt.want)
			}
		})
	}
}

func TestMessageHasChecks(t *testing.T) {
	tests := []struct {
		name     string
		md       protoreflect.MessageDescriptor
		expected bool
	}{
		{"nil descriptor", nil, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := messageHasChecks(tt.md)
			if got != tt.expected {
				t.Errorf("messageHasChecks() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestFieldHasChecks(t *testing.T) {
	tests := []struct {
		name     string
		fd       protoreflect.FieldDescriptor
		expected bool
	}{
		{"nil descriptor", nil, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fieldHasChecks(tt.fd)
			if got != tt.expected {
				t.Errorf("fieldHasChecks() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestMessageHasMessageChecks(t *testing.T) {
	tests := []struct {
		name     string
		md       protoreflect.MessageDescriptor
		expected bool
	}{
		{"nil descriptor", nil, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := messageHasMessageChecks(tt.md)
			if got != tt.expected {
				t.Errorf("messageHasMessageChecks() = %v, want %v", got, tt.expected)
			}
		})
	}
}
