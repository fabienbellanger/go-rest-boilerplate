package echo

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/fabienbellanger/go-rest-boilerplate/lib"
)

// customHTTPErrorHandler manages HTTP errors
func customHTTPErrorHandler(err error, c echo.Context) {
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
		lib.CheckError(err, 0)

		c.JSON(code, map[string]string{"message": "Internal Server Error"})
	default:
		lib.CheckError(err, 0)

		c.JSON(code, "")
	}
}
