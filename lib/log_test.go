package lib

import (
	"bytes"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

// TestCustomLog
func TestCustomLog(t *testing.T) {
	b := new(bytes.Buffer)
	DefaultEchoLogWriter = b

	var want, result string

	// String
	// ------
	getStr := "ERR  | "
	want = "ERR  | "
	CustomLog(getStr)
	result = b.String()
	assert.Contains(t, result, want, "must constaint "+want)
	b.Reset()

	// Number
	// ------
	getFloat := 1542.545
	want = "1542.545"
	CustomLog(getFloat)
	result = b.String()
	assert.Contains(t, result, want, "must constaint "+want)
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
	assert.Contains(t, result, want, "must constaint "+want)
	b.Reset()
}

// TestSQLLog
func TestSQLLog(t *testing.T) {
	b := new(bytes.Buffer)
	DefaultSqlLogWriter = b
	viper.Set("environment", "development")
	viper.Set("log.sql.limit", 0.01)
	viper.Set("log.sql.level", 1)
	viper.Set("log.sql.displayOverLimit", true)

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
	assert.Contains(t, result, want, "should contains time with color")
	b.Reset()

	viper.Set("environment", "production")
	want = "|  SEL | " + latency.String() + " \t|"
	SQLLog(latency, query, nil)
	result = b.String()
	assert.Contains(t, result, want, "should contains time without color")
	b.Reset()

	// Time + query
	// ------------
	viper.Set("environment", "development")
	viper.Set("log.sql.level", 2)
	want = "|  SEL |\x1b[97;41m " + latency.String() + " \t| " + query
	SQLLog(latency, query, nil)
	result = b.String()
	assert.Contains(t, result, want, "should contains time and query with color")
	b.Reset()

	viper.Set("environment", "production")
	want = "|  SEL | " + latency.String() + " \t| " + query
	SQLLog(latency, query, nil)
	result = b.String()
	assert.Contains(t, result, want, "should contains time and query without color")
	b.Reset()

	// Time + query
	// ------------
	viper.Set("environment", "development")
	viper.Set("log.sql.level", 3)
	want = "|  SEL |\x1b[97;41m " + latency.String() + " \t| " + query + " | [1]"
	SQLLog(latency, query, 1)
	result = b.String()
	assert.Contains(t, result, want, "should contains time, query and arguments with color")
	b.Reset()

	viper.Set("environment", "development")
	want = "|  SEL |\x1b[97;41m " + latency.String() + " \t| " + query + " | [1 test]"
	SQLLog(latency, query, 1, "test")
	result = b.String()
	assert.Contains(t, result, want, "should contains time, query and arguments without color")
	b.Reset()
}

// TestDisplaySuccessMessage
func TestDisplaySuccessMessage(t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	var msg, want, result string

	msg = "Test success message"
	DisplaySuccessMessage(msg)
	w.Close()
	out, _ := ioutil.ReadAll(r)
	result = string(out)
	want = msg + "\n"
	assert.Contains(t, result, want, "result should contain non empty message")

	msg = ""
	DisplaySuccessMessage(msg)
	w.Close()
	out, _ = ioutil.ReadAll(r)
	result = string(out)
	want = ""
	assert.Contains(t, result, want, "result should contain empty message")

	os.Stdout = rescueStdout
}
