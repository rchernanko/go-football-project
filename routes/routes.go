package routes

import (
	"github.com/gin-gonic/gin"
	"go-football-project/handlers"
)

func InitialiseRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/players/:id", handlers.GetPlayers)

	return r
}
