package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	return r
}
