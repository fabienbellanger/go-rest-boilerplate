package routes

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/fabienbellanger/go-rest-boilerplate/database"
	"github.com/fabienbellanger/go-rest-boilerplate/lib"
	"github.com/fabienbellanger/go-rest-boilerplate/models"
	"github.com/labstack/echo/v4"
)

// Routes associées au framework Echo
func exampleRoutes(e *echo.Echo, g *echo.Group) {
	// Test page for websockets
	g.GET("/websockets", func(c echo.Context) error {
		return c.Render(http.StatusOK, "example/websockets.gohtml", map[string]interface{}{
			"title":        "Websockets example",
			"webSocketUrl": strconv.Itoa(lib.Config.WebSocketServer.Port),
		})
	})

	// Test page for VueJS
	g.GET("/vuejs", func(c echo.Context) error {
		return c.Render(http.StatusOK, "example/vuejs.gohtml", map[string]interface{}{
			"title": "VueJS example",
		})
	})

	// Benchmark large query with pure mysql
	g.GET("/benchmark", func(c echo.Context) error {
		query := `
			SELECT id, username, password, lastname, firstname, created_at, updated_at, deleted_at
			FROM users
			LIMIT 100000`
		rows, _ := database.Select(query)

		users := benchmarkEcho(rows)
		nbUsers := len(users)

		response := c.Response()
		response.WriteHeader(http.StatusOK)

		if _, err := io.WriteString(response, "["); err != nil {
			return err
		}

		encoder := json.NewEncoder(response)

		for i := 0; i < nbUsers; i++ {
			if i > 0 {
				if _, err := io.WriteString(response, ","); err != nil {
					return err
				}
			}

			if err := encoder.Encode(users[i]); err != nil {
				return err
			}
		}

		if _, err := io.WriteString(response, "]"); err != nil {
			return err
		}

		return nil
	})

	// Benchmark large query without using an array
	g.GET("/benchmark2", func(c echo.Context) error {
		query := `
			SELECT id, username, password, lastname, firstname, created_at, updated_at, deleted_at
			FROM users
			LIMIT 100000`
		rows, _ := database.Select(query)

		response := c.Response()
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
	})
}

// benchmarkEcho executes query and create the array of results
func benchmarkEcho(rows *sql.Rows) []models.User {
	var user models.User

	users := make([]models.User, 100000)
	// users := make([]models.User, 0)
	i := 0
	for rows.Next() {
		rows.Scan(
			&user.ID,
			&user.Username,
			&user.Password,
			&user.Lastname,
			&user.Firstname,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.DeletedAt)

		users[i] = user
		// users = append(users, user)

		i++
	}

	return users
}
