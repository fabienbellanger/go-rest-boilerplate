package routes

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/fabienbellanger/go-rest-boilerplate/database"
	"github.com/fabienbellanger/go-rest-boilerplate/lib"
	"github.com/fabienbellanger/go-rest-boilerplate/orm/models"
	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
)

// Routes associées au framework Gin
func exampleRoutes(group *gin.RouterGroup) {
	// Benchmark large query with pure mysql
	group.GET("/benchmark", func(c *gin.Context) {
		query := "SELECT * FROM users"
		rows, _ := database.Select(query)

		users := make([]models.User, 0, 100000)
		var user models.User

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

			users = append(users, user)
		}

		if len(users) == 0 {
			c.JSON(http.StatusNotFound, users)
		} else {
			res := lib.GetHTTPResponse(
				http.StatusOK,
				"Success",
				&users,
			)

			c.Writer.WriteHeader(200)
			if err := json.NewEncoder(c.Writer).Encode(res); err != nil {
				lib.CheckError(err, 0)
			}
			c.Writer.Flush()
		}
	})

	// This handler will match /user/john but will not match /user/ or /user
	group.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")

		c.HTML(http.StatusOK, "example/index.gohtml", gin.H{
			"title": "Example page",
			"name":  name,
		})
	})

	// However, this one will match /user/john/ and also /user/john/send
	// If no other routers match /user/john, it will redirect to /user/john/
	// :param : Paramètre obligatoire
	// *param : Paramètre optionnel
	group.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action

		c.String(http.StatusOK, message)
	})

	// Test page for websockets
	group.GET("/websockets", func(c *gin.Context) {
		c.HTML(http.StatusOK, "example/websockets.gohtml", gin.H{
			"title":        "Websockets example",
			"webSocketUrl": strconv.Itoa(lib.Config.WebSocketServer.Port),
		})
	})

	// Test page for VueJS
	group.GET("/vuejs", func(c *gin.Context) {
		c.HTML(http.StatusOK, "example/vuejs.gohtml", gin.H{
			"title": "VueJS example",
		})
	})
}

// Routes associées au framework Echo
func echoExampleRoutes(e *echo.Echo, g *echo.Group) {
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

	// Benchmark large query with pure mysql
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

func benchmarkEcho(rows *sql.Rows) []*models.User {
	users := make([]*models.User, 100000)
	var user models.User

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

		users[i] = &user
		// users = append(users, user)

		i++
	}

	return users
}
