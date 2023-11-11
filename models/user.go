package models

import (
	"fmt"
	"log"

	"github.com/dev-anderson/AnderBank/database"
	"github.com/dev-anderson/AnderBank/schemas"
)

func GetAllUser() ([]schemas.User, error) {
	db, err := database.ConnectDatabase()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	query := "select id, name, email, password from anderbank_user where active = true order by id"

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

	query := `insert into anderbank_user (name, email, password, "dateCreate", active) values ($1, $2, $3, now(), true);`

	err = db.QueryRow(query, name, email, password).Err()

	return err
}

func GetUserID(id int) (User, error) {
	db, err := database.ConnectDatabase()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	query := `select name, email, password from anderbank_user where id = $1`

	var user User

	row := db.QueryRow(query, id)
	err = row.Scan(&user.Name, &user.Email, &user.Password)
	if err != nil {
		log.Fatal("Falha ao carregar o ID do user", err)
	}

	return user, nil
}

func GetAllUserDelete() ([]User, error) {
	db, err := database.ConnectDatabase()
	if err != nil {
		log.Fatal("Erro ao conectar com o banco de dados", err.Error())
	}
	defer db.Close()

	query := `select name, email, password from anderbank_user where active = false`

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("Erro ao buscar User Delete", err.Error())
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Name, &user.Email, &user.Password); err != nil {
			fmt.Println(err.Error())
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
	}

	return users, nil
}

func DeleteIDUser(id int) error {
	db, err := database.ConnectDatabase()
	if err != nil {
		log.Fatal("Falha ao conectar com o banco de dados", err.Error())
	}
	defer db.Close()

	query := `update anderbank_user set active = false where id = $1`

	_, err = db.Query(query, id)
	if err != nil {
		log.Fatal("Falha ao deletar o user", err)
	}

	return nil

}
