package routes

import (
	"net/http/pprof"

	"github.com/labstack/echo/v4"
)

// echoPprofRoutes adds several routes from package `net/http/pprof` to *echo.Group object.
func echoPprofRoutes(g *echo.Group) {
	routers := []struct {
		Method  string
		Path    string
		Handler echo.HandlerFunc
	}{
		{"GET", "/debug/pprof/", PprofIndexHandler()},
		{"GET", "/debug/pprof/heap", PprofHeapHandler()},
		{"GET", "/debug/pprof/goroutine", PprofGoroutineHandler()},
		{"GET", "/debug/pprof/block", PprofBlockHandler()},
		{"GET", "/debug/pprof/threadcreate", PprofThreadCreateHandler()},
		{"GET", "/debug/pprof/cmdline", PprofCmdlineHandler()},
		{"GET", "/debug/pprof/profile", PprofProfileHandler()},
		{"GET", "/debug/pprof/symbol", PprofSymbolHandler()},
		{"POST", "/debug/pprof/symbol", PprofSymbolHandler()},
		{"GET", "/debug/pprof/trace", PprofTraceHandler()},
		{"GET", "/debug/pprof/mutex", PprofMutexHandler()},
	}

	for _, r := range routers {
		switch r.Method {
		case "GET":
			g.GET(r.Path, r.Handler)
		case "POST":
			g.POST(r.Path, r.Handler)
		}
	}
}

// PprofIndexHandler will pass the call from /debug/pprof to pprof.
func PprofIndexHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		pprof.Index(ctx.Response().Writer, ctx.Request())
		return nil
	}
}

// PprofHeapHandler will pass the call from /debug/pprof/heap to pprof.
func PprofHeapHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		pprof.Handler("heap").ServeHTTP(ctx.Response(), ctx.Request())
		return nil
	}
}

// PprofGoroutineHandler will pass the call from /debug/pprof/goroutine to pprof.
func PprofGoroutineHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		pprof.Handler("goroutine").ServeHTTP(ctx.Response().Writer, ctx.Request())
		return nil
	}
}

// PprofBlockHandler will pass the call from /debug/pprof/block to pprof.
func PprofBlockHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		pprof.Handler("block").ServeHTTP(ctx.Response().Writer, ctx.Request())
		return nil
	}
}

// PprofThreadCreateHandler will pass the call from /debug/pprof/threadcreate to pprof.
func PprofThreadCreateHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		pprof.Handler("threadcreate").ServeHTTP(ctx.Response().Writer, ctx.Request())
		return nil
	}
}

// PprofCmdlineHandler will pass the call from /debug/pprof/cmdline to pprof.
func PprofCmdlineHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		pprof.Cmdline(ctx.Response().Writer, ctx.Request())
		return nil
	}
}

// PprofProfileHandler will pass the call from /debug/pprof/profile to pprof.
func PprofProfileHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		pprof.Profile(ctx.Response().Writer, ctx.Request())
		return nil
	}
}

// PprofSymbolHandler will pass the call from /debug/pprof/symbol to pprof.
func PprofSymbolHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		pprof.Symbol(ctx.Response().Writer, ctx.Request())
		return nil
	}
}

// PprofTraceHandler will pass the call from /debug/pprof/trace to pprof.
func PprofTraceHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		pprof.Trace(ctx.Response().Writer, ctx.Request())
		return nil
	}
}

// PprofMutexHandler will pass the call from /debug/pprof/mutex to pprof.
func PprofMutexHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		pprof.Handler("mutex").ServeHTTP(ctx.Response().Writer, ctx.Request())
		return nil
	}
}
