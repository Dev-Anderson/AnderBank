package models

import (
	"fmt"

	"github.com/dev-anderson/AnderBank/database"
	"github.com/dev-anderson/AnderBank/schemas"
)

func Login(email string) (schemas.Login, error) {
	db, err := database.ConnectDatabase()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	query := `select id, email, password from anderbank_user where email = $1`

	rows, err := db.Query(query, email)
	if err != nil {
		fmt.Println("Erro in get user", err)
	}
	defer rows.Close()

	var login schemas.Login
	for rows.Next() {
		err := rows.Scan(&login.ID, &login.Email, &login.Password)
		if err != nil {
			fmt.Println("Erro in get user", err)
		}
	}

	return login, nil
}
