package handlers

import (
	"github.com/gin-gonic/gin"
)

func GetPlayers(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": c.Param("id"),
	})
}
