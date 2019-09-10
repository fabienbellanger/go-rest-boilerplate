package web

import (
	"net/http/pprof"

	"github.com/labstack/echo/v4"

	"github.com/fabienbellanger/go-rest-boilerplate/routes"
)

type webPprofRoute struct {
	Group *echo.Group
}

// NewWebPprofRoute returns implement of api authentication routes
func NewWebPprofRoute(g *echo.Group) routes.WebPprofRoutes {
	return &webPprofRoute{
		Group: g,
	}
}

// PprofRoutes adds several routes from package `net/http/pprof` to *echo.Group object.
func (r *webPprofRoute) PprofRoutes() {
	routers := []struct {
		Method  string
		Path    string
		Handler echo.HandlerFunc
	}{
		{"GET", "/debug/pprof/", pprofIndexHandler()},
		{"GET", "/debug/pprof/allocs", pprofAllocsHandler()},
		{"GET", "/debug/pprof/heap", pprofHeapHandler()},
		{"GET", "/debug/pprof/goroutine", pprofGoroutineHandler()},
		{"GET", "/debug/pprof/block", pprofBlockHandler()},
		{"GET", "/debug/pprof/threadcreate", pprofThreadCreateHandler()},
		{"GET", "/debug/pprof/cmdline", pprofCmdlineHandler()},
		{"GET", "/debug/pprof/profile", pprofProfileHandler()},
		{"GET", "/debug/pprof/symbol", pprofSymbolHandler()},
		{"POST", "/debug/pprof/symbol", pprofSymbolHandler()},
		{"GET", "/debug/pprof/trace", pprofTraceHandler()},
		{"GET", "/debug/pprof/mutex", pprofMutexHandler()},
	}

	for _, router := range routers {
		switch router.Method {
		case "GET":
			r.Group.GET(router.Path, router.Handler)
		case "POST":
			r.Group.POST(router.Path, router.Handler)
		}
	}
}

// pprofIndexHandler will pass the call from /debug/pprof to pprof.
func pprofIndexHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		pprof.Index(ctx.Response().Writer, ctx.Request())
		return nil
	}
}

// pprofAllocsHandler will pass the call from /debug/pprof/allocs to pprof.
func pprofAllocsHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		pprof.Handler("allocs").ServeHTTP(ctx.Response(), ctx.Request())
		return nil
	}
}

// pprofHeapHandler will pass the call from /debug/pprof/heap to pprof.
func pprofHeapHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		pprof.Handler("heap").ServeHTTP(ctx.Response(), ctx.Request())
		return nil
	}
}

// pprofGoroutineHandler will pass the call from /debug/pprof/goroutine to pprof.
func pprofGoroutineHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		pprof.Handler("goroutine").ServeHTTP(ctx.Response().Writer, ctx.Request())
		return nil
	}
}

// pprofBlockHandler will pass the call from /debug/pprof/block to pprof.
func pprofBlockHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		pprof.Handler("block").ServeHTTP(ctx.Response().Writer, ctx.Request())
		return nil
	}
}

// pprofThreadCreateHandler will pass the call from /debug/pprof/threadcreate to pprof.
func pprofThreadCreateHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		pprof.Handler("threadcreate").ServeHTTP(ctx.Response().Writer, ctx.Request())
		return nil
	}
}

// pprofCmdlineHandler will pass the call from /debug/pprof/cmdline to pprof.
func pprofCmdlineHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		pprof.Cmdline(ctx.Response().Writer, ctx.Request())
		return nil
	}
}

// pprofProfileHandler will pass the call from /debug/pprof/profile to pprof.
func pprofProfileHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		pprof.Profile(ctx.Response().Writer, ctx.Request())
		return nil
	}
}

// pprofSymbolHandler will pass the call from /debug/pprof/symbol to pprof.
func pprofSymbolHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		pprof.Symbol(ctx.Response().Writer, ctx.Request())
		return nil
	}
}

// pprofTraceHandler will pass the call from /debug/pprof/trace to pprof.
func pprofTraceHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		pprof.Trace(ctx.Response().Writer, ctx.Request())
		return nil
	}
}

// pprofMutexHandler will pass the call from /debug/pprof/mutex to pprof.
func pprofMutexHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		pprof.Handler("mutex").ServeHTTP(ctx.Response().Writer, ctx.Request())
		return nil
	}
}
