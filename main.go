package main

import (
	"github.com/dev-anderson/AnderBank/database"
	"github.com/dev-anderson/AnderBank/routes"
)

func main() {
	db, err := database.ConnectDatabase()
	if err != nil {
		panic(err)
	}
	database.TableUser(db)
	routes.Initialize()
}
