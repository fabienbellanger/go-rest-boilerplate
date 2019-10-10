package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/fabienbellanger/go-rest-boilerplate/issues"
	"github.com/fabienbellanger/go-rest-boilerplate/models"
	"github.com/fabienbellanger/go-rest-boilerplate/repositories"
	userRepository "github.com/fabienbellanger/go-rest-boilerplate/repositories/user"
)

// BenchmarkHandler
type BenchmarkHandler struct {
	userRepository repositories.UserRepository
}

// NewBenchmarkHandler
func NewBenchmarkHandler() *BenchmarkHandler {
	return &BenchmarkHandler{
		userRepository: userRepository.NewMysqlUserRepository(),
	}
}

// BenchmarkUser large query without map or slice
func (h *BenchmarkHandler) BenchmarkUser(c echo.Context) error {
	rows, _ := h.userRepository.GetAllSqlRows(100000)

	response := c.Response()
	response.Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	response.WriteHeader(http.StatusOK)

	if _, err := io.WriteString(response, "["); err != nil {
		return err
	}

	encoder := json.NewEncoder(response)
	var user models.User

	i := 0
	for ; rows.Next(); i++ {
		if i > 0 {
			if _, err := io.WriteString(response, ","); err != nil {
				return err
			}
		}

		rows.Scan(
			&user.ID,
			&user.Username,
			&user.Password,
			&user.Lastname,
			&user.Firstname,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.DeletedAt)

		if err := encoder.Encode(user); err != nil {
			return err
		}
	}

	if _, err := io.WriteString(response, "]"); err != nil {
		return err
	}

	return nil
}

// BenchmarkLargeMap large query with map
func (h *BenchmarkHandler) BenchmarkLargeMap(c echo.Context) error {
	data := issues.InitData()

	response := c.Response()
	response.Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	response.WriteHeader(http.StatusOK)

	if _, err := io.WriteString(response, "{"); err != nil {
		return err
	}

	encoder := json.NewEncoder(response)
	i := 0
	for applicationID, application := range *data {
		if i > 0 {
			if _, err := io.WriteString(response, ","); err != nil {
				return err
			}
		}

		if _, err := io.WriteString(response, `"`+strconv.Itoa(int(applicationID))+`":`); err != nil {
			return err
		}

		if err := encoder.Encode(application); err != nil {
			return err
		}

		i++
	}

	if _, err := io.WriteString(response, "}"); err != nil {
		return err
	}

	return nil
}
