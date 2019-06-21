package routes

import (
	"net/http"
	"strconv"

	"github.com/fabienbellanger/go-rest-boilerplate/database"
	"github.com/fabienbellanger/go-rest-boilerplate/lib"
	"github.com/fabienbellanger/go-rest-boilerplate/orm"
	"github.com/fabienbellanger/go-rest-boilerplate/orm/models"
	"github.com/gin-gonic/gin"
)

func exampleRoutes(group *gin.RouterGroup) {
	// Benchmark large query with Gorm
	group.GET("/benchmark", func(c *gin.Context) {
		users := make([]models.User, 0)
		orm.DB.Limit(1000).Find(&users)

		c.JSON(http.StatusOK, lib.GetHTTPResponse(
			http.StatusOK,
			"Success",
			users,
		))
	})

	// Benchmark large query with pure mysql
	group.GET("/benchmark2", func(c *gin.Context) {
		query := "SELECT * FROM users LIMIT 1000"
		rows, _ := database.Select(query)

		users := make([]models.User, 0)
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

		c.JSON(http.StatusOK, lib.GetHTTPResponse(
			http.StatusOK,
			"Success",
			users,
		))
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
