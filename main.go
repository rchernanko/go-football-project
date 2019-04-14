package main

import (
	"go-football-project/routes"
)

func main() {

	//TODO setup the db connection

	r := routes.SetupRoutes()
	r.Run()
}
