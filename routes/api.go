package routes

type ApiAuthRoutes interface {
	AuthRoutes()
}

type ApiUserRoutes interface {
	UsersRoutes()
}

type ApiExampleRoutes interface {
	ExampleRoutes()
}
