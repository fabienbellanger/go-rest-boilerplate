package api

import (
	"github.com/fabienbellanger/go-rest-boilerplate/handlers/api"
	"github.com/fabienbellanger/go-rest-boilerplate/routes"

	"github.com/labstack/echo/v4"
)

type apiBenchmarkRoute struct {
	Group *echo.Group
}

// NewApiBenchmarkRoute returns implement of api example routes
func NewApiBenchmarkRoute(g *echo.Group) routes.ApiBenchmarkRoutes {
	return &apiBenchmarkRoute{
		Group: g,
	}
}

// Routes associées au framework Echo
func (r *apiBenchmarkRoute) BenchmarkRoutes() {
	benchmarkHandler := api.NewBenchmarkHandler()

	// Benchmark large query without using an array
	r.Group.GET("/benchmark-users", benchmarkHandler.BenchmarkUser)
	r.Group.GET("/benchmark-map", benchmarkHandler.BenchmarkLargeMap)
}
