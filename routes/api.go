package routes

// ApiAuthRoutes manages authentication routes
type ApiAuthRoutes interface {
	AuthRoutes()
}

// ApiUserRoutes manages user routes
type ApiUserRoutes interface {
	UsersRoutes()
}

// ApiBenchmarkRoutes manages benchmark routes
type ApiBenchmarkRoutes interface {
	BenchmarkRoutes()
}
