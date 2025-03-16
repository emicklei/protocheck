package shared

import "testing"

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
