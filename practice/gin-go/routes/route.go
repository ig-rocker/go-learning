package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Route() *gin.Engine {
	route := gin.Default()

	route.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Hello")
	})

	return route
}
