package routes

// WebPprofRoutes manages pprof routes
type WebPprofRoutes interface {
	PprofRoutes()
}

// WebExampleRoutes manages example routes
type WebExampleRoutes interface {
	ExampleRoutes()
}

// WebLogsRoutes manages logs routes
type WebLogsRoutes interface {
	LogsRoutes()
}
