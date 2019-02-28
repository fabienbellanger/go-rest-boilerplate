package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func exampleRoutes(group *gin.RouterGroup) {
	// This handler will match /user/john but will not match /user/ or /user
	group.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
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
}
