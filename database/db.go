package database

import (
	"database/sql"
	"fmt"

	"github.com/dev-anderson/AnderBank/config"
	_ "github.com/lib/pq"
)

func ConnectDatabase() (*sql.DB, error) {
	env := config.LoadEnv()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		env.Host, env.Port, env.User, env.Password, env.DBName)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	return db, err
}
