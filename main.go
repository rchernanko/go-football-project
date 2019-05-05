package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"go-football-project/database"
)

var router *gin.Engine

func main() {
	database.InitDb()
	router = gin.Default()
	initialiseRoutes()
	router.Run()
}
