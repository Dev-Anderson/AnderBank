package models

import (
	"fmt"
	"log"

	"github.com/dev-anderson/AnderBank/database"
	"github.com/dev-anderson/AnderBank/schemas"
	"github.com/dev-anderson/AnderBank/services"
)

type Accounts struct {
	ID            int     `json:"id"`
	NumberAccount string  `json:"numberaccount"`
	Balance       float64 `json:"balance"`
	DateCreate    string  `json:"datecreate"`
	DateUpdate    *string `json:"dateupdate"` //o campo pode ser null no banco de dados
	DateDelete    *string `json:"datedelete"` //o capmo pode ser null no banco de dados
	Debit         bool    `json:"debit"`
	Credit        bool    `json:"credit"`
	Active        bool    `json:"active"`
}

func (a *Accounts) TodasContas() ([]*Accounts, error) {
	db, err := database.ConnectDatabase()
	if err != nil {
		log.Fatal("Erro ao conectar com o banco de dados", err.Error())
	}
	defer db.Close()

	query := `select id, "numberAccount", balance, "dateCreate", "dateUpdate", "dateDelete", debit, credit, active from anderbank_account`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	var accounts []*Accounts
	for rows.Next() {
		var account Accounts
		err := rows.Scan(
			&account.ID,
			&account.NumberAccount,
			&account.Balance,
			&account.DateCreate,
			&account.DateUpdate,
			&account.DateDelete,
			&account.Debit,
			&account.Credit,
			&account.Active,
		)
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, &account)
	}
	return accounts, nil
}

func GetAllAccounts() ([]schemas.Account, error) {
	db, err := database.ConnectDatabase()
	if err != nil {
		log.Fatal("Erro ao conectar com o banco de dados", err.Error())
	}
	defer db.Close()

	query := `select * from anderbank_account order by id`

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("Erro ao consultar as contas", err.Error())
	}
	defer rows.Close()

	var accounts []schemas.Account
	for rows.Next() {
		var account schemas.Account
		if err := rows.Scan(&account.ID, &account.NumberAccount, &account.Balance, &account.DateCreate, &account.DateUpdate, &account.DateDelete, &account.Debit, &account.Credit, &account.Active); err != nil {
			fmt.Println(err)
		}
		accounts = append(accounts, account)
	}

	if err := rows.Err(); err != nil {
		fmt.Print(err)
	}

	return accounts, nil
}

type Account struct {
	Balance float64 `json:"balance"`
	Debit   bool    `json:"debit"`
	Credit  bool    `json:"credit"`
}

func CreateAccount(balance float64, debit bool, credit bool) error {
	db, err := database.ConnectDatabase()
	if err != nil {
		log.Fatal("Erro ao conectar com o banco de dados", err.Error())
	}
	defer db.Close()

	numberAccount, err := services.RandomAccount()
	if err != nil {
		log.Fatal("Falha ao gerar o numero da conta", err.Error())
	}

	query := `insert into anderbank_account ("numberAccount" , balance,  "dateCreate", debit, credit, active) values ($1, $2, now(), $3, $4, true)`

	err = db.QueryRow(query, numberAccount, balance, debit, credit).Err()
	return err
}

func BalanceAccount(numberAccount string) (float64, error) {
	db, err := database.ConnectDatabase()
	if err != nil {
		log.Fatal("Erro ao conectar com o banco de dados", err.Error())
	}
	defer db.Close()

	query := `select balance from anderbank_account  where "numberAccount"  = $1`
	var balance float64
	row := db.QueryRow(query, numberAccount)
	err = row.Scan(&balance)
	if err != nil {
		log.Fatal("Falha ao carregar o saldo da conta", err.Error())
	}

	return balance, nil

}

func DeleteAccount(numberAccount string) error {
	db, err := database.ConnectDatabase()
	if err != nil {
		log.Fatal("Erro ao conectar com o banco de dados", err.Error())
	}
	defer db.Close()

	query := `select id from anderbank_account where "numberAccount"= $1`
	var id int

	db.QueryRow(query, numberAccount).Scan(&id)

	queryUpdate := `update anderbank_account set active = false where id = $1`
	err = db.QueryRow(queryUpdate, id).Err()

	return err
}
