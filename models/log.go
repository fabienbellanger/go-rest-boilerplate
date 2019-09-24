package models

type LogFile struct {
	Error *LogErrorFile
	Echo  *LogEchoFile
	Sql   *LogSqlFile
}

type LogErrorFile struct {
	Source    string `json:"source"`
	Timestamp string `json:"timestamp"`
	Message   string `json:"message"`
}

type LogEchoFile struct {
	Source    string `json:"source"`
	Timestamp string `json:"timestamp"`
	Code      string `json:"code"`
	Latency   string `json:"latency"`
	Method    string `json:"method"`
	Uri       string `json:"uri"`
}

type LogSqlFile struct {
	Source     string `json:"source"`
	Timestamp  string `json:"timestamp"`
	Request    string `json:"request"`
	Latency    string `json:"latency"`
	Query      string `json:"query"`
	Parameters string `json:"parameters"`
}
