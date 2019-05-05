package main

import (
	"go-football-project/handlers"
)

func initialiseRoutes() {
	router.GET("/players", handlers.GetPlayers)
	router.GET("/players/:id", handlers.GetPlayer)
	router.POST("/players", handlers.CreatePlayer)
	router.PUT("/players/:id", handlers.UpdatePlayer)
	router.DELETE("/players/:id", handlers.DeletePlayer)
}
