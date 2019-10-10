package models

// LogFile type
type LogFile struct {
	Error *LogErrorFile
	Echo  *LogEchoFile
	Sql   *LogSqlFile
}

// LogErrorFile type
type LogErrorFile struct {
	Source    string `json:"source"`
	Timestamp string `json:"timestamp"`
	Message   string `json:"message"`
}

// LogEchoFile type
type LogEchoFile struct {
	Source    string `json:"source"`
	Timestamp string `json:"timestamp"`
	Code      int    `json:"code"`
	Latency   string `json:"latency"`
	Method    string `json:"method"`
	Uri       string `json:"uri"`
}

// LogSqlFile type
type LogSqlFile struct {
	Source     string `json:"source"`
	Timestamp  string `json:"timestamp"`
	Request    string `json:"request"`
	Latency    string `json:"latency"`
	Query      string `json:"query"`
	Parameters string `json:"parameters"`
}
