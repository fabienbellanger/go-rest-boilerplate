package routes

type WebPprofRoutes interface {
	PprofRoutes()
}

type WebExampleRoutes interface {
	ExampleRoutes()
}

type WebLogsRoutes interface {
	LogsRoutes()
}
