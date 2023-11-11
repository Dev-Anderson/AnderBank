package main

import (
	"github.com/dev-anderson/AnderBank/database"
	"github.com/dev-anderson/AnderBank/database/table"
	"github.com/dev-anderson/AnderBank/routes"
)

func main() {
	db, _ := database.ConnectDatabase()
	table.TableAccount(db)
	table.TableUser(db)
	routes.Initialize()
}
