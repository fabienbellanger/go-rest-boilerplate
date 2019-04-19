package controllers

import (
	jwt "github.com/appleboy/gin-jwt"
	"github.com/fabienbellanger/go-rest-boilerplate/lib"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetUserHandler displays authenticated user information
func GetUserHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)

	c.JSON(http.StatusOK, lib.GetHTTPResponse(
		http.StatusOK,
		"Success",
		gin.H{
			"id":        claims["id"],
			"username":  claims["username"],
			"lastname":  claims["lastname"],
			"firstname": claims["firstname"],
			"fullname":  claims["fullname"],
		}),
	)
}
