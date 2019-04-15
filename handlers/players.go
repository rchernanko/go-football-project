package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-football-project/models/players"
	"net/http"
)

func GetPlayers(c *gin.Context) {
	var pl []players.Player

	if err := db.Find(&pl).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, pl)
	}
}

func GetPlayer(c *gin.Context) {
	id := c.Params.ByName("id")
	var pl players.Player
	if err := db.Where("id = ?", id).First(&pl).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, pl)
	}
}
func CreatePlayer(c *gin.Context) {
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

func UpdatePlayer(c *gin.Context) {
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

func DeletePlayer(c *gin.Context) {
	id := c.Params.ByName("id")
	var pl players.Player
	d := db.Where("id = ?", id).Delete(&pl)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}