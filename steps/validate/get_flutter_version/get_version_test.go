package get_flutter_version

import (
	"testing"
)

func Test_cleanVersion(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"v1.2.3", "1.2.3"},
	}
	for _, tt := range tests {
		got := cleanVersion(tt.input)
		if got != tt.want {
			t.Errorf("cleanVersion(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

func Test_CleanVersion(t *testing.T) {
	output := "Flutter 3.35.7 • channel stable • https://github.com/flutter/flutter.git"
	got, err := CleanVersion(output)
	if err != nil {
		t.Fatalf("CleanVersion() error: %v", err)
	}
	if got != "3.35.7" {
		t.Errorf("CleanVersion() = %q, want %q", got, "3.35.7")
	}

	outputInvalid := "Flutter stable"
	_, err = CleanVersion(outputInvalid)
	if err == nil {
		t.Errorf("CleanVersion() should error for invalid input")
	}
}

func Test_ParseVersion(t *testing.T) {
	valid := "3.35.7"
	ver, err := ParseVersion(valid)
	if err != nil {
		t.Fatalf("ParseVersion() error: %v", err)
	}
	if ver.String() != valid {
		t.Errorf("ParseVersion() = %v, want %v", ver, valid)
	}

	invalid := "not_a_version"
	_, err = ParseVersion(invalid)
	if err == nil {
		t.Errorf("ParseVersion() should error for invalid input")
	}
}

func Test_GetVersion_regex(t *testing.T) {
	output := "Flutter 3.35.7 • channel stable • https://github.com/flutter/flutter.git"
	cleaned, err := CleanVersion(output)
	if err != nil {
		t.Fatalf("CleanVersion() error: %v", err)
	}
	parsed, err := ParseVersion(cleaned)
	if err != nil {
		t.Fatalf("ParseVersion() error: %v", err)
	}
	if parsed.String() != "3.35.7" {
		t.Errorf("parsed version = %v, want 3.35.7", parsed)
	}
}
