package echo

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/fabienbellanger/go-rest-boilerplate/lib"
)

// customHTTPErrorHandler manages HTTP errors
func customHTTPErrorHandler(err error, c echo.Context) {
	lib.CheckError(err, 0)

	code := http.StatusInternalServerError
	if httpError, ok := err.(*echo.HTTPError); ok {
		code = httpError.Code
	}

	switch code {
	case http.StatusBadRequest:
		// 400
		c.JSON(code, map[string]string{"message": "Bad Request"})
	case http.StatusUnauthorized:
		// 401
		c.JSON(code, map[string]string{"message": "Not Authorized"})
	case http.StatusNotFound:
		// 404
		c.JSON(code, map[string]string{"message": "Resource Not Found"})
	case http.StatusInternalServerError:
		// 500
		c.JSON(code, map[string]string{"message": "Internal Server Error"})
	default:
		c.JSON(code, "")
	}
}
