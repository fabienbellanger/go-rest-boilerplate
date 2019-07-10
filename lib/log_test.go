package lib

import (
	"bytes"
	"strings"
	"testing"
	"time"
)

// TestCustomLog
func TestCustomLog(t *testing.T) {
	b := new(bytes.Buffer)
	DefaultEchoLogWriter = b

	// String
	// ------
	getStr := "ERR  | "
	want := "ERR  | "
	CustomLog(getStr)
	result := b.String()
	if !strings.Contains(result, want) {
		t.Errorf("CustomLog(%q) == %q, want %q", getStr, result, want)
	}
	b.Reset()

	// Number
	// ------
	getFloat := 1542.545
	want = "1542.545"
	CustomLog(getFloat)
	result = b.String()
	if !strings.Contains(result, want) {
		t.Errorf("CustomLog(%f) == %q, want %q", getFloat, result, want)
	}
	b.Reset()

	// Custom type
	// -----------
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

// TestSQLLog
func TestSQLLog(t *testing.T) {
	b := new(bytes.Buffer)
	DefaultEchoLogWriter = b
	Config.Environment = "development"
	Config.SQLLog.Limit = 0.01
	Config.SQLLog.Level = 1
	Config.SQLLog.DisplayOverLimit = true
	query := `
		SELECT *
		FROM Test
		WHERE id = ?`
	query = strings.Join(strings.Fields(query), " ")

	var latency time.Duration
	var want, result string

	// Only time
	// ---------
	latency = time.Second
	want = "|  SEL |\x1b[97;41m " + latency.String() + " \x1b[0m\t|"
	SQLLog(latency, query, nil)
	result = b.String()
	if !strings.Contains(result, want) {
		t.Errorf("SQLLog() == %q, want %q", result, want)
	}
	b.Reset()

	Config.Environment = "production"
	want = "|  SEL | " + latency.String() + " \t|"
	SQLLog(latency, query, nil)
	result = b.String()
	if !strings.Contains(result, want) {
		t.Errorf("SQLLog() == %q, want %q", result, want)
	}
	b.Reset()

	// Time + query
	// ------------
	Config.Environment = "development"
	Config.SQLLog.Level = 2
	want = "|  SEL |\x1b[97;41m " + latency.String() + " \t| " + query
	SQLLog(latency, query, nil)
	result = b.String()
	if !strings.Contains(result, want) {
		t.Errorf("SQLLog() == %q, want %q", result, want)
	}

	b.Reset()
	Config.Environment = "production"
	want = "|  SEL | " + latency.String() + " \t| " + query
	SQLLog(latency, query, nil)
	result = b.String()
	if !strings.Contains(result, want) {
		t.Errorf("SQLLog() == %q, want %q", result, want)
	}
	b.Reset()

	// Time + query
	// ------------
	Config.Environment = "development"
	Config.SQLLog.Level = 3
	want = "|  SEL |\x1b[97;41m " + latency.String() + " \t| " + query + " | [1]"
	SQLLog(latency, query, 1)
	result = b.String()
	if !strings.Contains(result, want) {
		t.Errorf("SQLLog() == %q, want %q", result, want)
	}
	b.Reset()

	Config.Environment = "development"
	want = "|  SEL |\x1b[97;41m " + latency.String() + " \t| " + query + " | [1 test]"
	SQLLog(latency, query, 1, "test")
	result = b.String()
	if !strings.Contains(result, want) {
		t.Errorf("SQLLog() == %q, want %q", result, want)
	}
	b.Reset()
}
