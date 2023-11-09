package models

import (
	"fmt"

	"github.com/dev-anderson/AnderBank/database"
	"github.com/dev-anderson/AnderBank/schemas"
)

func GetAllUser() ([]schemas.User, error) {
	db, err := database.ConnectDatabase()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	query := "select id, name, email, password from anderbank_user order by id"

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("Erro in get user", err)
	}
	defer rows.Close()

	var users []schemas.User
	for rows.Next() {
		var user schemas.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password); err != nil {
			fmt.Println(err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
	}

	return users, nil
}

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(name, email, password string) error {
	db, err := database.ConnectDatabase()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	query := `insert into anderbank_user (name, email, password, "dateCreate") values ($1, $2, $3, now());`

	err = db.QueryRow(query, name, email, password).Err()

	return err
}
