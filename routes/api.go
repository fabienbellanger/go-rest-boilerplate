package routes

type ApiAuthRoutes interface {
	AuthRoutes()
}

type ApiUserRoutes interface {
	UsersRoutes()
}

type ApiBenchmarkRoutes interface {
	BenchmarkRoutes()
}
