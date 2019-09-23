package logs

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/fabienbellanger/go-rest-boilerplate/models"
)

// TestParseError
func TestParseError(t *testing.T) {
	log := `ERR  | 2019-09-20 21:51:55 | [template: no template "logs/index.gohtml" associated with template "index.gohtml"]`
	log1 := models.LogErrorFile{
		Source:    "ERR",
		Timestamp: "2019-09-20 21:51:55",
		Message:   `[template: no template "logs/index.gohtml" associated with template "index.gohtml"]`,
	}
	log2, _ := parseError(log)
	assert.Equal(t, log1, log2)
}

// TestParseEcho
func TestEchoError(t *testing.T) {
	log := `ECHO | 2019-09-20 21:51:55 |  500 | 326.921µs	| GET	| /logs`
	log1 := models.LogEchoFile{
		Source:    "ECHO",
		Timestamp: "2019-09-20 21:51:55",
		Code:      "500",
		Latency:   "326.921µs",
		Method:    "GET",
		Uri:       "/logs",
	}
	log2, _ := parseEcho(log)
	assert.Equal(t, log1, log2)
}

// TestSqlError
func TestSqlError(t *testing.T) {
	log := `SQL  | 2019-09-21 21:37:13 |  SEL | 87.148393ms 	| SELECT id, username, lastname, firstname, created_at, deleted_at FROM users WHERE username = ? AND password = ? AND deleted_at IS NULL LIMIT 1 | [[[fabien 62670d1e1eea06b6c975e12bc8a16131b278f6d7bcbe017b65f854c58476baba86c2082b259fd0c1310935b365dc40f609971b6810b065e528b0b60119e69f61]]]`
	log1 := models.LogSqlFile{
		Source:    "SQL",
		Timestamp: "2019-09-21 21:37:13",
		Request:   "SEL",
		Latency:   "87.148393ms",
		Query:     "SELECT id, username, lastname, firstname, created_at, deleted_at FROM users WHERE username = ? AND password = ? AND deleted_at IS NULL LIMIT 1",
		Paramters: "[[[fabien 62670d1e1eea06b6c975e12bc8a16131b278f6d7bcbe017b65f854c58476baba86c2082b259fd0c1310935b365dc40f609971b6810b065e528b0b60119e69f61]]]",
	}
	log2, _ := parseSql(log)
	assert.Equal(t, log1, log2)

	log = `SQL  | 2019-09-21 21:37:13 |  SEL | 87.148393ms 	| SELECT id, username, lastname, firstname, created_at, deleted_at FROM users WHERE username = ? AND password = ? AND deleted_at IS NULL LIMIT 1 |`
	log1 = models.LogSqlFile{
		Source:    "SQL",
		Timestamp: "2019-09-21 21:37:13",
		Request:   "SEL",
		Latency:   "87.148393ms",
		Query:     "SELECT id, username, lastname, firstname, created_at, deleted_at FROM users WHERE username = ? AND password = ? AND deleted_at IS NULL LIMIT 1",
		Paramters: "",
	}
	log2, _ = parseSql(log)
	assert.Equal(t, log1, log2)

	log = `SQL  | 2019-09-21 21:37:13 |  SEL | 87.148393ms 	|`
	log1 = models.LogSqlFile{
		Source:    "SQL",
		Timestamp: "2019-09-21 21:37:13",
		Request:   "SEL",
		Latency:   "87.148393ms",
		Query:     "",
		Paramters: "",
	}
	log2, _ = parseSql(log)
	assert.Equal(t, log1, log2)
}
