package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-football-project/database"
	"go-football-project/models"
	"net/http"
)

func GetPlayers(c *gin.Context) {
	var pl []models.Player

	if err := database.Db.Find(&pl).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, pl)
	}
}

func GetPlayer(c *gin.Context) {
	id := c.Params.ByName("id")
	var pl models.Player
	if err := database.Db.Where("id = ?", id).First(&pl).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, pl)
	}
}
func CreatePlayer(c *gin.Context) {
	var pl models.Player
	err := c.BindJSON(&pl)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		database.Db.Create(&pl)
		c.JSON(200, pl)
	}
}

func UpdatePlayer(c *gin.Context) {
	var pl models.Player
	id := c.Params.ByName("id")
	if err := database.Db.Where("id = ?", id).First(&pl).Error;
		err != nil{
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	err := c.BindJSON(&pl)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		database.Db.Save(&pl)
		c.JSON(200, pl)
	}
}

func DeletePlayer(c *gin.Context) {
	id := c.Params.ByName("id")
	var pl models.Player
	d := database.Db.Where("id = ?", id).Delete(&pl)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
