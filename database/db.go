package database

import (
	"github.com/jinzhu/gorm"
	"go-football-project/models"
)

var Db *gorm.DB

func InitDb() {

	var err error

	Db, err = gorm.Open("mysql", "root:QcvA8gRn{DpCLJQmr2*taDJwHBaBcd2e@tcp(127.0.0.1:3306)/richard_test?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("Failed to connect to the database")
	}

	Db.AutoMigrate(&models.Player{}, &models.Team{})
}