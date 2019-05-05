package main

import (
	"go-football-project/handlers"
)

func initialiseRoutes() {
	playerRoutes()
}

func playerRoutes() {
	playerRoutes := router.Group("/players")
	{
		playerRoutes.GET("", handlers.GetPlayers)
		playerRoutes.GET("/:id", handlers.GetPlayer)
		playerRoutes.POST("", handlers.CreatePlayer)
		playerRoutes.PUT("/:id", handlers.UpdatePlayer)
		playerRoutes.DELETE("/:id", handlers.DeletePlayer)
	}
}
