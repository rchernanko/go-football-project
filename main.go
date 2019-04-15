package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"go-football-project/handlers"
	"go-football-project/models/players"
)

var db *gorm.DB
var err error

func main() {


	//db, err = gorm.Open("mysql", "root:password@tcp(127.0.0.1:3306)/richard_test?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("Failed to connect to the database")
	}

	defer db.Close()

	//Runs the migrations for the Player table automatically
	db.AutoMigrate(&players.Player{})

	r := gin.Default()
	r.GET("/players", handlers.GetPlayers)
	r.GET("/players/:id", handlers.GetPlayer)
	r.POST("/players", handlers.CreatePlayer)
	r.PUT("/players/:id", handlers.UpdatePlayer)
	r.DELETE("/players/:id", handlers.DeletePlayer)
	r.Run()
}

