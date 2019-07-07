package gin

import (
	"net/http"
	"strconv"

	"github.com/fabienbellanger/go-rest-boilerplate/lib"
	"github.com/gin-gonic/gin"
)

// exampleRoutes lists routes of Gin
func exampleRoutes(group *gin.RouterGroup) {
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
