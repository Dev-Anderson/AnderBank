package main

import (
	"fmt"

	"github.com/dev-anderson/AnderBank/database"
	"github.com/dev-anderson/AnderBank/routes"
)

func main() {
	fmt.Println("Teste")

	database.ConnectDatabase()
	routes.Initialize()
}
