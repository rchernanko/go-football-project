package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"go-football-project/models/players"
	"net/http"
)

var db *gorm.DB
var err error

func main() {

	db, err = gorm.Open("mysql", "root:password@tcp(127.0.0.1:3306)/richard_test?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("Failed to connect to the database")
	}

	defer db.Close()

	//Runs the migrations for the Player table automatically
	db.AutoMigrate(&players.Player{})

	r := gin.Default()
	r.GET("/players", getPlayers)
	r.GET("/players/:id", getPlayer)
	r.POST("/players", createPlayer)
	r.PUT("/players/:id", updatePlayer)
	r.DELETE("/players/:id", deletePlayer)
	r.Run()
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
func createPlayer(c *gin.Context) {
	var pl players.Player
	err := c.BindJSON(&pl)

	if (err != nil) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		db.Create(&pl)
		c.JSON(200, pl)
	}
}

func updatePlayer(c *gin.Context) {
	var pl players.Player
	id := c.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&pl).Error;
	err != nil{
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	err := c.BindJSON(&pl)

	if (err != nil) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		db.Save(&pl)
		c.JSON(200, pl)
	}
}

func deletePlayer(c *gin.Context) {
	id := c.Params.ByName("id")
	var pl players.Player
	d := db.Where("id = ?", id).Delete(&pl)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
