package lib

import (
	"bytes"
	"strings"
	"testing"
)

// TestCustomLog
func TestCustomLog(t *testing.T) {
	b := new(bytes.Buffer)
	DefaultEchoLogWriter = b

	// String
	getStr := "ERR  | "
	want := "ERR  | "
	CustomLog(getStr)
	result := b.String()
	if !strings.Contains(result, want) {
		t.Errorf("CustomLog(%q) == %q, want %q", getStr, result, want)
	}
	b.Reset()

	// Number
	getFloat := 1542.545
	want = "1542.545"
	CustomLog(getFloat)
	result = b.String()
	if !strings.Contains(result, want) {
		t.Errorf("CustomLog(%f) == %q, want %q", getFloat, result, want)
	}
	b.Reset()

	// Custom type
	type test struct {
		prop1 string
		prop2 int
	}
	getType := test{
		prop1: "str",
		prop2: 23,
	}
	want = "{prop1:str prop2:23}"
	CustomLog(getType)
	result = b.String()
	if !strings.Contains(result, want) {
		t.Errorf("CustomLog(%v) == %q, want %q", getType, result, want)
	}
	b.Reset()
}
