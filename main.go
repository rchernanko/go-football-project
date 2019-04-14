package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"go-football-project/models/players"
)

var db *gorm.DB
var err error

func main() {
	db, err = gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	db.AutoMigrate(&players.Player{})

	r := gin.Default()
	r.GET("/players", getPlayers)
	r.GET("/players/:id", getPlayer)
	r.POST("/players", createPlayer)
	r.Run()
}

func createPlayer(c *gin.Context) {
	var pl players.Player
	c.BindJSON(&pl)
	db.Create(&pl)
	c.JSON(200, pl)
}

func getPlayers(c *gin.Context) {
	var pl []players.Player

	if err := db.Find(&pl).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, pl)
	}
}

func getPlayer(c *gin.Context) {
	id := c.Params.ByName("id")
	var pl players.Player
	if err := db.Where("id = ?", id).First(&pl).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, pl)
	}
}